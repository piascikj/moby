// Code generated by protoc-gen-gogo.
// source: github.com/docker/swarmkit/protobuf/plugin/plugin.proto
// DO NOT EDIT!

/*
	Package plugin is a generated protocol buffer package.

	It is generated from these files:
		github.com/docker/swarmkit/protobuf/plugin/plugin.proto

	It has these top-level messages:
		WatchSelectors
		StoreObject
		TLSAuthorization
*/
package plugin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

import github_com_docker_swarmkit_api_deepcopy "github.com/docker/swarmkit/api/deepcopy"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type WatchSelectors struct {
	// supported by all object types
	ID           *bool `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	IDPrefix     *bool `protobuf:"varint,2,opt,name=id_prefix,json=idPrefix" json:"id_prefix,omitempty"`
	Name         *bool `protobuf:"varint,3,opt,name=name" json:"name,omitempty"`
	NamePrefix   *bool `protobuf:"varint,4,opt,name=name_prefix,json=namePrefix" json:"name_prefix,omitempty"`
	Custom       *bool `protobuf:"varint,5,opt,name=custom" json:"custom,omitempty"`
	CustomPrefix *bool `protobuf:"varint,6,opt,name=custom_prefix,json=customPrefix" json:"custom_prefix,omitempty"`
	// supported by tasks only
	ServiceID    *bool `protobuf:"varint,7,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	NodeID       *bool `protobuf:"varint,8,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Slot         *bool `protobuf:"varint,9,opt,name=slot" json:"slot,omitempty"`
	DesiredState *bool `protobuf:"varint,10,opt,name=desired_state,json=desiredState" json:"desired_state,omitempty"`
	// supported by nodes only
	Role       *bool `protobuf:"varint,11,opt,name=role" json:"role,omitempty"`
	Membership *bool `protobuf:"varint,12,opt,name=membership" json:"membership,omitempty"`
	// supported by: resource
	Kind             *bool  `protobuf:"varint,13,opt,name=kind" json:"kind,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WatchSelectors) Reset()                    { *m = WatchSelectors{} }
func (*WatchSelectors) ProtoMessage()               {}
func (*WatchSelectors) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{0} }

type StoreObject struct {
	WatchSelectors   *WatchSelectors `protobuf:"bytes,1,req,name=watch_selectors,json=watchSelectors" json:"watch_selectors,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *StoreObject) Reset()                    { *m = StoreObject{} }
func (*StoreObject) ProtoMessage()               {}
func (*StoreObject) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{1} }

type TLSAuthorization struct {
	// Roles contains the acceptable TLS OU roles for the handler.
	Roles []string `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	// Insecure is set to true if this method does not require
	// authorization. NOTE: Specifying both "insecure" and a nonempty id:453 gh:454
	// list of roles is invalid. This would fail at codegen time.
	Insecure         *bool  `protobuf:"varint,2,opt,name=insecure" json:"insecure,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TLSAuthorization) Reset()                    { *m = TLSAuthorization{} }
func (*TLSAuthorization) ProtoMessage()               {}
func (*TLSAuthorization) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{2} }

var E_Deepcopy = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         70000,
	Name:          "docker.protobuf.plugin.deepcopy",
	Tag:           "varint,70000,opt,name=deepcopy,def=1",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

var E_StoreObject = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*StoreObject)(nil),
	Field:         70001,
	Name:          "docker.protobuf.plugin.store_object",
	Tag:           "bytes,70001,opt,name=store_object,json=storeObject",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

var E_TlsAuthorization = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*TLSAuthorization)(nil),
	Field:         73626345,
	Name:          "docker.protobuf.plugin.tls_authorization",
	Tag:           "bytes,73626345,opt,name=tls_authorization,json=tlsAuthorization",
	Filename:      "github.com/docker/swarmkit/protobuf/plugin/plugin.proto",
}

func init() {
	proto.RegisterType((*WatchSelectors)(nil), "docker.protobuf.plugin.WatchSelectors")
	proto.RegisterType((*StoreObject)(nil), "docker.protobuf.plugin.StoreObject")
	proto.RegisterType((*TLSAuthorization)(nil), "docker.protobuf.plugin.TLSAuthorization")
	proto.RegisterExtension(E_Deepcopy)
	proto.RegisterExtension(E_StoreObject)
	proto.RegisterExtension(E_TlsAuthorization)
}

func (m *WatchSelectors) Copy() *WatchSelectors {
	if m == nil {
		return nil
	}
	o := &WatchSelectors{}
	o.CopyFrom(m)
	return o
}

func (m *WatchSelectors) CopyFrom(src interface{}) {

	o := src.(*WatchSelectors)
	*m = *o
}

func (m *StoreObject) Copy() *StoreObject {
	if m == nil {
		return nil
	}
	o := &StoreObject{}
	o.CopyFrom(m)
	return o
}

func (m *StoreObject) CopyFrom(src interface{}) {

	o := src.(*StoreObject)
	*m = *o
	if o.WatchSelectors != nil {
		m.WatchSelectors = &WatchSelectors{}
		github_com_docker_swarmkit_api_deepcopy.Copy(m.WatchSelectors, o.WatchSelectors)
	}
}

func (m *TLSAuthorization) Copy() *TLSAuthorization {
	if m == nil {
		return nil
	}
	o := &TLSAuthorization{}
	o.CopyFrom(m)
	return o
}

func (m *TLSAuthorization) CopyFrom(src interface{}) {

	o := src.(*TLSAuthorization)
	*m = *o
	if o.Roles != nil {
		m.Roles = make([]string, len(o.Roles))
		copy(m.Roles, o.Roles)
	}

}

func (m *WatchSelectors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WatchSelectors) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != nil {
		dAtA[i] = 0x8
		i++
		if *m.ID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IDPrefix != nil {
		dAtA[i] = 0x10
		i++
		if *m.IDPrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Name != nil {
		dAtA[i] = 0x18
		i++
		if *m.Name {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NamePrefix != nil {
		dAtA[i] = 0x20
		i++
		if *m.NamePrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Custom != nil {
		dAtA[i] = 0x28
		i++
		if *m.Custom {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CustomPrefix != nil {
		dAtA[i] = 0x30
		i++
		if *m.CustomPrefix {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ServiceID != nil {
		dAtA[i] = 0x38
		i++
		if *m.ServiceID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NodeID != nil {
		dAtA[i] = 0x40
		i++
		if *m.NodeID {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Slot != nil {
		dAtA[i] = 0x48
		i++
		if *m.Slot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DesiredState != nil {
		dAtA[i] = 0x50
		i++
		if *m.DesiredState {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Role != nil {
		dAtA[i] = 0x58
		i++
		if *m.Role {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Membership != nil {
		dAtA[i] = 0x60
		i++
		if *m.Membership {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Kind != nil {
		dAtA[i] = 0x68
		i++
		if *m.Kind {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StoreObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WatchSelectors == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("watch_selectors")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(m.WatchSelectors.Size()))
		n1, err := m.WatchSelectors.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TLSAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSAuthorization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Insecure != nil {
		dAtA[i] = 0x10
		i++
		if *m.Insecure {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Plugin(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Plugin(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func (m *WatchSelectors) Size() (n int) {
	var l int
	_ = l
	if m.ID != nil {
		n += 2
	}
	if m.IDPrefix != nil {
		n += 2
	}
	if m.Name != nil {
		n += 2
	}
	if m.NamePrefix != nil {
		n += 2
	}
	if m.Custom != nil {
		n += 2
	}
	if m.CustomPrefix != nil {
		n += 2
	}
	if m.ServiceID != nil {
		n += 2
	}
	if m.NodeID != nil {
		n += 2
	}
	if m.Slot != nil {
		n += 2
	}
	if m.DesiredState != nil {
		n += 2
	}
	if m.Role != nil {
		n += 2
	}
	if m.Membership != nil {
		n += 2
	}
	if m.Kind != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StoreObject) Size() (n int) {
	var l int
	_ = l
	if m.WatchSelectors != nil {
		l = m.WatchSelectors.Size()
		n += 1 + l + sovPlugin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TLSAuthorization) Size() (n int) {
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	if m.Insecure != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WatchSelectors) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WatchSelectors{`,
		`ID:` + valueToStringPlugin(this.ID) + `,`,
		`IDPrefix:` + valueToStringPlugin(this.IDPrefix) + `,`,
		`Name:` + valueToStringPlugin(this.Name) + `,`,
		`NamePrefix:` + valueToStringPlugin(this.NamePrefix) + `,`,
		`Custom:` + valueToStringPlugin(this.Custom) + `,`,
		`CustomPrefix:` + valueToStringPlugin(this.CustomPrefix) + `,`,
		`ServiceID:` + valueToStringPlugin(this.ServiceID) + `,`,
		`NodeID:` + valueToStringPlugin(this.NodeID) + `,`,
		`Slot:` + valueToStringPlugin(this.Slot) + `,`,
		`DesiredState:` + valueToStringPlugin(this.DesiredState) + `,`,
		`Role:` + valueToStringPlugin(this.Role) + `,`,
		`Membership:` + valueToStringPlugin(this.Membership) + `,`,
		`Kind:` + valueToStringPlugin(this.Kind) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StoreObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreObject{`,
		`WatchSelectors:` + strings.Replace(fmt.Sprintf("%v", this.WatchSelectors), "WatchSelectors", "WatchSelectors", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TLSAuthorization) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TLSAuthorization{`,
		`Roles:` + fmt.Sprintf("%v", this.Roles) + `,`,
		`Insecure:` + valueToStringPlugin(this.Insecure) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPlugin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WatchSelectors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchSelectors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchSelectors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.ID = &b
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDPrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.IDPrefix = &b
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Name = &b
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamePrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.NamePrefix = &b
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Custom", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Custom = &b
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomPrefix", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.CustomPrefix = &b
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.ServiceID = &b
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.NodeID = &b
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Slot = &b
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredState", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.DesiredState = &b
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Role = &b
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Membership = &b
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Kind = &b
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreObject) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchSelectors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WatchSelectors == nil {
				m.WatchSelectors = &WatchSelectors{}
			}
			if err := m.WatchSelectors.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("watch_selectors")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TLSAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TLSAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TLSAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Insecure", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Insecure = &b
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPlugin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlugin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPlugin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPlugin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/docker/swarmkit/protobuf/plugin/plugin.proto", fileDescriptorPlugin)
}

var fileDescriptorPlugin = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xae, 0xd3, 0x36, 0x4d, 0x26, 0x69, 0xff, 0xfe, 0x2b, 0x54, 0xad, 0x7a, 0x70, 0xaa, 0x46,
	0x42, 0x41, 0x42, 0x8e, 0xd4, 0x0b, 0x52, 0x6e, 0x94, 0x5c, 0x22, 0x01, 0x45, 0x0e, 0x12, 0x37,
	0x22, 0xc7, 0x3b, 0x4d, 0x96, 0x3a, 0x5e, 0x6b, 0x77, 0x4d, 0x0a, 0x27, 0x5e, 0x80, 0x07, 0xe0,
	0xca, 0xd3, 0xf4, 0xc8, 0x91, 0x53, 0x44, 0x2d, 0x71, 0xe0, 0x06, 0x6f, 0x80, 0x76, 0xd7, 0x69,
	0x08, 0x6a, 0xc5, 0xc9, 0x33, 0xdf, 0x7c, 0xdf, 0xcc, 0x7c, 0x3b, 0x86, 0x47, 0x13, 0xae, 0xa7,
	0xf9, 0x38, 0x88, 0xc5, 0xac, 0xcb, 0x44, 0x7c, 0x81, 0xb2, 0xab, 0xe6, 0x91, 0x9c, 0x5d, 0x70,
	0xdd, 0xcd, 0xa4, 0xd0, 0x62, 0x9c, 0x9f, 0x77, 0xb3, 0x24, 0x9f, 0xf0, 0xb4, 0xfc, 0x04, 0x16,
	0x26, 0x07, 0x8e, 0x1d, 0x2c, 0x49, 0x81, 0xab, 0x1e, 0x1e, 0x4d, 0x84, 0x98, 0x24, 0xb8, 0x12,
	0x33, 0x54, 0xb1, 0xe4, 0x99, 0x16, 0x25, 0xf7, 0xf8, 0xd3, 0x26, 0xec, 0xbd, 0x8a, 0x74, 0x3c,
	0x1d, 0x62, 0x82, 0xb1, 0x16, 0x52, 0x91, 0x03, 0xa8, 0x70, 0x46, 0xbd, 0x23, 0xaf, 0x53, 0x3b,
	0xad, 0x16, 0x8b, 0x56, 0x65, 0xd0, 0x0f, 0x2b, 0x9c, 0x91, 0x07, 0x50, 0xe7, 0x6c, 0x94, 0x49,
	0x3c, 0xe7, 0x97, 0xb4, 0x62, 0xcb, 0xcd, 0x62, 0xd1, 0xaa, 0x0d, 0xfa, 0x2f, 0x2c, 0x16, 0xd6,
	0x38, 0x73, 0x11, 0x21, 0xb0, 0x95, 0x46, 0x33, 0xa4, 0x9b, 0x86, 0x15, 0xda, 0x98, 0xb4, 0xa0,
	0x61, 0xbe, 0xcb, 0x06, 0x5b, 0xb6, 0x04, 0x06, 0x2a, 0x45, 0x07, 0x50, 0x8d, 0x73, 0xa5, 0xc5,
	0x8c, 0x6e, 0xdb, 0x5a, 0x99, 0x91, 0x36, 0xec, 0xba, 0x68, 0x29, 0xad, 0xda, 0x72, 0xd3, 0x81,
	0xa5, 0xf8, 0x21, 0x80, 0x42, 0xf9, 0x96, 0xc7, 0x38, 0xe2, 0x8c, 0xee, 0xd8, 0xed, 0x76, 0x8b,
	0x45, 0xab, 0x3e, 0x74, 0xe8, 0xa0, 0x1f, 0xd6, 0x4b, 0xc2, 0x80, 0x91, 0x36, 0xec, 0xa4, 0x82,
	0x59, 0x6a, 0xcd, 0x52, 0xa1, 0x58, 0xb4, 0xaa, 0xcf, 0x05, 0x33, 0xbc, 0xaa, 0x29, 0x0d, 0x98,
	0x31, 0xa1, 0x12, 0xa1, 0x69, 0xdd, 0x99, 0x30, 0xb1, 0xd9, 0x85, 0xa1, 0xe2, 0x12, 0xd9, 0x48,
	0xe9, 0x48, 0x23, 0x05, 0xb7, 0x4b, 0x09, 0x0e, 0x0d, 0x66, 0x84, 0x52, 0x24, 0x48, 0x1b, 0x4e,
	0x68, 0x62, 0xe2, 0x03, 0xcc, 0x70, 0x36, 0x46, 0xa9, 0xa6, 0x3c, 0xa3, 0x4d, 0x67, 0x7e, 0x85,
	0x18, 0xcd, 0x05, 0x4f, 0x19, 0xdd, 0x75, 0x1a, 0x13, 0x1f, 0xbf, 0x86, 0xc6, 0x50, 0x0b, 0x89,
	0x67, 0xe3, 0x37, 0x18, 0x6b, 0x72, 0x06, 0xff, 0xcd, 0xcd, 0xa5, 0x46, 0x6a, 0x79, 0x2a, 0xea,
	0x1d, 0x55, 0x3a, 0x8d, 0x93, 0xfb, 0xc1, 0xed, 0xe7, 0x0f, 0xd6, 0x0f, 0x1b, 0xee, 0xcd, 0xd7,
	0xf2, 0xe3, 0x3e, 0xec, 0xbf, 0x7c, 0x3a, 0x7c, 0x9c, 0xeb, 0xa9, 0x90, 0xfc, 0x7d, 0xa4, 0xb9,
	0x48, 0xc9, 0x3d, 0xd8, 0x36, 0xfb, 0x9a, 0xd6, 0x9b, 0x9d, 0x7a, 0xe8, 0x12, 0x72, 0x08, 0x35,
	0x9e, 0x2a, 0x8c, 0x73, 0x89, 0xee, 0xf2, 0xe1, 0x4d, 0xde, 0x7b, 0x02, 0x35, 0x86, 0x98, 0xc5,
	0x22, 0x7b, 0x47, 0x5a, 0x81, 0xfb, 0xe1, 0x56, 0x9b, 0x3c, 0x43, 0xa5, 0xa2, 0x09, 0x9e, 0x65,
	0xa6, 0xbb, 0xa2, 0x3f, 0x3f, 0xdb, 0xbb, 0xf7, 0xb6, 0xb4, 0xcc, 0x31, 0xbc, 0x11, 0xf6, 0x38,
	0x34, 0x95, 0xb1, 0x3a, 0x12, 0xce, 0xeb, 0x3f, 0x1b, 0xfd, 0xb2, 0x8d, 0x1a, 0x27, 0xed, 0xbb,
	0xbc, 0xff, 0xf1, 0x72, 0x61, 0x43, 0xad, 0x92, 0xde, 0x25, 0xfc, 0xaf, 0x13, 0x35, 0x8a, 0xd6,
	0x6c, 0xfb, 0xb7, 0xcc, 0xd3, 0x53, 0xc1, 0x96, 0xe3, 0x7e, 0x7c, 0xff, 0xd8, 0xb6, 0xf3, 0x3a,
	0x77, 0xcd, 0xfb, 0xfb, 0x25, 0xc3, 0x7d, 0x9d, 0xa8, 0x35, 0xe4, 0x94, 0x5e, 0x5d, 0xfb, 0x1b,
	0x5f, 0xaf, 0xfd, 0x8d, 0x0f, 0x85, 0xef, 0x5d, 0x15, 0xbe, 0xf7, 0xa5, 0xf0, 0xbd, 0x6f, 0x85,
	0xef, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x99, 0x7d, 0xfb, 0xf9, 0x03, 0x00, 0x00,
}
