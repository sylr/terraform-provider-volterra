// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/fast_acl_set/object.proto

/*
	Package fast_acl_set is a generated protocol buffer package.

	It is generated from these files:
		ves.io/schema/fast_acl_set/object.proto
		ves.io/schema/fast_acl_set/public_crudapi.proto
		ves.io/schema/fast_acl_set/types.proto

	It has these top-level messages:
		Object
		SpecType
		StatusObject
		CreateRequest
		CreateResponse
		ReplaceRequest
		ReplaceResponse
		GetRequest
		GetResponse
		ListRequest
		ListResponseItem
		ListResponse
		DeleteRequest
		GlobalSpecType
		CreateSpecType
		ReplaceSpecType
		GetSpecType
*/
package fast_acl_set

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import ves_io_schema4 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Fast ACL Set
//
// x-displayName: "Fast ACL Set"
// Fast ACL Set object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *ves_io_schema4.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *ves_io_schema4.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the desired behavior of the Fast ACL set
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec" json:"spec,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{0} }

func (m *Object) GetMetadata() *ves_io_schema4.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *ves_io_schema4.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Fast ACL set specification
//
// x-displayName: "Specification"
// Shape of the Fast ACL Set specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	GcSpec *GlobalSpecType `protobuf:"bytes,2,opt,name=gc_spec,json=gcSpec" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()                    { *m = SpecType{} }
func (*SpecType) ProtoMessage()               {}
func (*SpecType) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{1} }

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

// Status of Fast ACL set
//
// x-displayName: "Status"
// Most recently observed status of object
type StatusObject struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard status's metadata
	Metadata *ves_io_schema4.StatusMetaType `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// object_refs
	//
	// x-displayName: "Config Object"
	// Refernce to object for current status
	ObjectRefs []*ves_io_schema4.ObjectRefType `protobuf:"bytes,2,rep,name=object_refs,json=objectRefs" json:"object_refs,omitempty"`
	// conditions
	//
	// x-displayName: "Conditions"
	// Conditions reported by various component of the system
	Conditions []*ves_io_schema4.ConditionType `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *StatusObject) Reset()                    { *m = StatusObject{} }
func (*StatusObject) ProtoMessage()               {}
func (*StatusObject) Descriptor() ([]byte, []int) { return fileDescriptorObject, []int{2} }

func (m *StatusObject) GetMetadata() *ves_io_schema4.StatusMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatusObject) GetObjectRefs() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.ObjectRefs
	}
	return nil
}

func (m *StatusObject) GetConditions() []*ves_io_schema4.ConditionType {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.fast_acl_set.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.fast_acl_set.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.fast_acl_set.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.fast_acl_set.SpecType")
	proto.RegisterType((*StatusObject)(nil), "ves.io.schema.fast_acl_set.StatusObject")
	golang_proto.RegisterType((*StatusObject)(nil), "ves.io.schema.fast_acl_set.StatusObject")
}
func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *StatusObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusObject)
	if !ok {
		that2, ok := that.(StatusObject)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.ObjectRefs) != len(that1.ObjectRefs) {
		return false
	}
	for i := range this.ObjectRefs {
		if !this.ObjectRefs[i].Equal(that1.ObjectRefs[i]) {
			return false
		}
	}
	if len(this.Conditions) != len(that1.Conditions) {
		return false
	}
	for i := range this.Conditions {
		if !this.Conditions[i].Equal(that1.Conditions[i]) {
			return false
		}
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&fast_acl_set.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&fast_acl_set.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&fast_acl_set.StatusObject{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.ObjectRefs != nil {
		s = append(s, "ObjectRefs: "+fmt.Sprintf("%#v", this.ObjectRefs)+",\n")
	}
	if this.Conditions != nil {
		s = append(s, "Conditions: "+fmt.Sprintf("%#v", this.Conditions)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.SystemMetadata != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.SystemMetadata.Size()))
		n2, err := m.SystemMetadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Spec != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Spec.Size()))
		n3, err := m.Spec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.GcSpec.Size()))
		n4, err := m.GcSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *StatusObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintObject(dAtA, i, uint64(m.Metadata.Size()))
		n5, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.ObjectRefs) > 0 {
		for _, msg := range m.ObjectRefs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Conditions) > 0 {
		for _, msg := range m.Conditions {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintObject(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Object) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *StatusObject) Size() (n int) {
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if len(m.ObjectRefs) > 0 {
		for _, e := range m.ObjectRefs {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	return n
}

func sovObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "ves_io_schema4.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "ves_io_schema4.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(fmt.Sprintf("%v", this.Spec), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusObject{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "StatusMetaType", "ves_io_schema4.StatusMetaType", 1) + `,`,
		`ObjectRefs:` + strings.Replace(fmt.Sprintf("%v", this.ObjectRefs), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`Conditions:` + strings.Replace(fmt.Sprintf("%v", this.Conditions), "ConditionType", "ves_io_schema4.ConditionType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &ves_io_schema4.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
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
			return fmt.Errorf("proto: StatusObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ves_io_schema4.StatusMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectRefs = append(m.ObjectRefs, &ves_io_schema4.ObjectRefType{})
			if err := m.ObjectRefs[len(m.ObjectRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
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
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &ves_io_schema4.ConditionType{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
					return 0, ErrIntOverflowObject
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
				return 0, ErrInvalidLengthObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowObject
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
				next, err := skipObject(dAtA[start:])
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
	ErrInvalidLengthObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/fast_acl_set/object.proto", fileDescriptorObject) }
func init() {
	golang_proto.RegisterFile("ves.io/schema/fast_acl_set/object.proto", fileDescriptorObject)
}

var fileDescriptorObject = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x09, 0x0a, 0xd5, 0x25, 0x6a, 0x2b, 0x23, 0x24, 0x13, 0xe0, 0x54, 0x05, 0x04,
	0x11, 0x92, 0xcf, 0x52, 0x59, 0xa0, 0x42, 0x45, 0x6a, 0x07, 0x16, 0x4a, 0x51, 0xc2, 0xc4, 0x62,
	0x9d, 0x2f, 0xcf, 0xae, 0x21, 0xce, 0x59, 0x77, 0x67, 0x8b, 0x0c, 0x48, 0x88, 0x8d, 0x8d, 0x8f,
	0x81, 0xf8, 0x14, 0x2d, 0x13, 0x62, 0xaa, 0x98, 0x3a, 0x12, 0x77, 0x61, 0x60, 0xe8, 0xc2, 0x8e,
	0x7a, 0xae, 0xa3, 0xb8, 0x0a, 0x55, 0xb7, 0x77, 0x7a, 0xbf, 0xff, 0xff, 0xbd, 0x7b, 0xf7, 0x0e,
	0xdf, 0xcf, 0x40, 0xd1, 0x48, 0xb8, 0x8a, 0xef, 0x41, 0xcc, 0xdc, 0x80, 0x29, 0xed, 0x31, 0x3e,
	0xf2, 0x14, 0x68, 0x57, 0xf8, 0x6f, 0x80, 0x6b, 0x9a, 0x48, 0xa1, 0x85, 0xd5, 0x29, 0x40, 0x5a,
	0x80, 0x74, 0x1e, 0xec, 0x38, 0x61, 0xa4, 0xf7, 0x52, 0x9f, 0x72, 0x11, 0xbb, 0xa1, 0x08, 0x85,
	0x6b, 0x24, 0x7e, 0x1a, 0x98, 0x93, 0x39, 0x98, 0xa8, 0xb0, 0xea, 0xdc, 0xbb, 0xa0, 0xa6, 0x9e,
	0x24, 0xa0, 0xce, 0xb8, 0x9b, 0x55, 0x4e, 0x24, 0x3a, 0x12, 0xe3, 0x32, 0x79, 0xa3, 0x9a, 0x9c,
	0xd7, 0x75, 0xab, 0xa9, 0x0c, 0x14, 0x8c, 0xb3, 0xaa, 0xbc, 0xfb, 0x17, 0xe1, 0xe6, 0xae, 0xb9,
	0x9f, 0xf5, 0x18, 0x2f, 0xc5, 0xa0, 0xd9, 0x90, 0x69, 0x66, 0xa3, 0x35, 0xd4, 0x6b, 0xad, 0xdf,
	0xa6, 0xd5, 0xcb, 0x16, 0xe0, 0x0e, 0x68, 0xf6, 0x6a, 0x92, 0x40, 0x7f, 0x86, 0x5b, 0xcf, 0xf1,
	0x8a, 0x9a, 0x28, 0x0d, 0xb1, 0x37, 0x73, 0xa8, 0x1b, 0x87, 0x3b, 0xe7, 0x1c, 0x06, 0x86, 0x3a,
	0xe7, 0xb3, 0x5c, 0x68, 0x77, 0x4a, 0xb7, 0x47, 0xf8, 0x8a, 0x4a, 0x80, 0xdb, 0x0d, 0x63, 0x71,
	0x97, 0xfe, 0x7f, 0xe2, 0x74, 0x90, 0x00, 0x37, 0x1e, 0x46, 0xb1, 0xd1, 0xfb, 0xf8, 0xb4, 0x59,
	0x98, 0x7d, 0x3b, 0xb0, 0xeb, 0xab, 0xe8, 0xc7, 0x81, 0x7d, 0x3d, 0x03, 0xe5, 0x44, 0xc2, 0x89,
	0xc6, 0x81, 0x64, 0x4a, 0xcb, 0x94, 0xeb, 0x54, 0x42, 0x77, 0x17, 0x2f, 0x95, 0x5a, 0x6b, 0x1b,
	0x5f, 0x0d, 0xb9, 0x67, 0x4a, 0x16, 0x5d, 0x3f, 0xb8, 0xa8, 0xe4, 0xb3, 0x91, 0xf0, 0xd9, 0x68,
	0x56, 0xb8, 0x19, 0xf2, 0xd3, 0xb8, 0xfb, 0x07, 0xe1, 0xf6, 0x40, 0x33, 0x9d, 0xaa, 0x4b, 0x8f,
	0xb3, 0xc0, 0x17, 0x8c, 0xf3, 0x05, 0x6e, 0x15, 0x3b, 0xe7, 0x49, 0x08, 0x94, 0x5d, 0x5f, 0x6b,
	0xf4, 0x5a, 0xeb, 0xb7, 0x16, 0x3e, 0x46, 0x1f, 0x82, 0x53, 0xf1, 0xd6, 0xca, 0xd7, 0xf7, 0xed,
	0xf9, 0x36, 0xfb, 0x58, 0x94, 0x79, 0x65, 0x3d, 0xc1, 0x98, 0x8b, 0xf1, 0x30, 0x32, 0x0f, 0x6f,
	0x37, 0x16, 0xda, 0x6d, 0x97, 0x80, 0xe9, 0x65, 0x8e, 0xdf, 0xb8, 0xf6, 0x73, 0x73, 0x15, 0x2f,
	0xe3, 0x76, 0xd9, 0x1e, 0x4d, 0xa3, 0xe1, 0xd6, 0x27, 0xb4, 0xbf, 0x59, 0x3b, 0x9c, 0x92, 0xda,
	0xd1, 0x94, 0xd4, 0x4e, 0xa6, 0x04, 0x7d, 0xc8, 0x09, 0xfa, 0x92, 0x13, 0xf4, 0x3d, 0x27, 0xe8,
	0x30, 0x27, 0xe8, 0x28, 0x27, 0xe8, 0x57, 0x4e, 0xd0, 0xef, 0x9c, 0xd4, 0x4e, 0x72, 0x82, 0x3e,
	0x1f, 0x93, 0xda, 0xfe, 0x31, 0x41, 0xaf, 0x5f, 0x86, 0x22, 0x79, 0x1b, 0xd2, 0x4c, 0x8c, 0x34,
	0x48, 0xc9, 0x68, 0xaa, 0x5c, 0x13, 0x04, 0x42, 0xc6, 0x4e, 0x22, 0x45, 0x16, 0x0d, 0x41, 0x3a,
	0x65, 0xda, 0x4d, 0xfc, 0x50, 0xb8, 0xf0, 0x4e, 0x9f, 0x2d, 0xf1, 0x82, 0xbf, 0xe2, 0x37, 0xcd,
	0x2a, 0x3f, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x55, 0xa1, 0xb3, 0xe7, 0xc4, 0x03, 0x00, 0x00,
}
