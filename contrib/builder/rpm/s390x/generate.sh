#!/usr/bin/env bash
set -e

# This file is used to auto-generate Dockerfiles for making rpms via 'make rpm'
#
# usage: ./generate.sh [versions]
#    ie: ./generate.sh
#        to update all Dockerfiles in this directory
#    or: ./generate.sh centos-7
#        to only update centos-7/Dockerfile
#    or: ./generate.sh fedora-newversion
#        to create a new folder and a Dockerfile within it

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

versions=( "$@" )
if [ ${#versions[@]} -eq 0 ]; then
	versions=( */ )
fi
versions=( "${versions[@]%/}" )

for version in "${versions[@]}"; do
	echo "${versions[@]}"
	distro="${version%-*}"
	suite="${version##*-}"
	case "$distro" in
		*opensuse*)
		from="opensuse/s390x:tumbleweed"
		;;
	*clefos*)
		from="sinenomine/${distro}"
		;;
	*)
		echo No appropriate or supported image available.
		exit 1
		;;
    esac
	installer=yum

	mkdir -p "$version"
	echo "$version -> FROM $from"
	cat > "$version/Dockerfile" <<-EOF
		#
		# THIS FILE IS AUTOGENERATED; SEE "contrib/builder/rpm/s390x/generate.sh"!
		#

		FROM $from

	EOF

	echo >> "$version/Dockerfile"

	extraBuildTags=''
	runcBuildTags=

	case "$from" in
		*clefos*)
			# Fix for RHBZ #1213602 + get "Development Tools" packages dependencies
			echo 'RUN touch /var/lib/rpm/* && yum groupinstall -y "Development Tools"' >> "$version/Dockerfile"
			;;
		*opensuse*)
			echo "RUN zypper ar https://download.opensuse.org/ports/zsystems/tumbleweed/repo/oss/ tumbleweed" >> "$version/Dockerfile"
			# get rpm-build and curl packages and dependencies
			echo 'RUN zypper --non-interactive install ca-certificates* curl gzip rpm-build' >> "$version/Dockerfile"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	packages=(
		btrfs-progs-devel # for "btrfs/ioctl.h" (and "version.h" if possible)
		device-mapper-devel # for "libdevmapper.h"
		glibc-static
		libseccomp-devel # for "seccomp.h" & "libseccomp.so"
		libselinux-devel # for "libselinux.so"
		pkgconfig # for the pkg-config command
		selinux-policy
		selinux-policy-devel
		sqlite-devel # for "sqlite3.h"
		systemd-devel # for "sd-journal.h" and libraries
		tar # older versions of dev-tools do not have tar
		git # required for containerd and runc clone
		cmake # tini build
		vim-common # tini build
	)

	case "$from" in
		*clefos*)
			extraBuildTags+=' seccomp'
			runcBuildTags="seccomp selinux"
			;;
		*opensuse*)
			packages=( "${packages[@]/libseccomp-devel}" )
			runcBuildTags="selinux"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	case "$from" in
		*clefos*)
			# Same RHBZ fix needed on all yum lines
			echo "RUN touch /var/lib/rpm/* && ${installer} install -y ${packages[*]}" >> "$version/Dockerfile"
			;;
		*opensuse*)
			packages=( "${packages[@]/btrfs-progs-devel/libbtrfs-devel}" )
			packages=( "${packages[@]/pkgconfig/pkg-config}" )
			packages=( "${packages[@]/vim-common/vim}" )

			packages+=( systemd-rpm-macros ) # for use of >= opensuse:13.*

			# use zypper
			echo "RUN zypper --non-interactive install ${packages[*]}" >> "$version/Dockerfile"
			;;
		*)
			echo No appropriate or supported image available.
			exit 1
			;;
	esac

	echo >> "$version/Dockerfile"

	awk '$1 == "ENV" && $2 == "GO_VERSION" { print; exit }' ../../../../Dockerfile.s390x >> "$version/Dockerfile"
	echo 'RUN curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-s390x.tar.gz" | tar xzC /usr/local' >> "$version/Dockerfile"
	echo 'ENV PATH $PATH:/usr/local/go/bin' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	echo 'ENV AUTO_GOPATH 1' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	# print build tags in alphabetical order
	buildTags=$( echo "selinux $extraBuildTags" | xargs -n1 | sort -n | tr '\n' ' ' | sed -e 's/[[:space:]]*$//' )

	echo "ENV DOCKER_BUILDTAGS $buildTags" >> "$version/Dockerfile"
	echo "ENV RUNC_BUILDTAGS $runcBuildTags" >> "$version/Dockerfile"
	# TODO: Investigate why "s390x-linux-gnu-gcc" is required id:29 gh:30
	echo "RUN ln -s /usr/bin/gcc /usr/bin/s390x-linux-gnu-gcc" >> "$version/Dockerfile"
done
