// +build linux,btrfs_noversion

package btrfs

// TODO (vbatts) remove this work-around once supported linux distros are on id:209 gh:210
// btrfs utilities of >= 3.16.1

func btrfsBuildVersion() string {
	return "-"
}

func btrfsLibVersion() int {
	return -1
}
