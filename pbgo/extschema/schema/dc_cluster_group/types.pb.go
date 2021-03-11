// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/dc_cluster_group/types.proto

package dc_cluster_group

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
import _ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Global Specification
//
// x-displayName: "Global Specification"
// DC Cluster Group specification
type GlobalSpecType struct {
}

func (m *GlobalSpecType) Reset()                    { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage()               {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

// Create DC Cluster group
//
// x-displayName: "Create DC Cluster Group"
// Create DC Cluster group in given namespace
type CreateSpecType struct {
}

func (m *CreateSpecType) Reset()                    { *m = CreateSpecType{} }
func (*CreateSpecType) ProtoMessage()               {}
func (*CreateSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

// Replace DC Cluster Group
//
// x-displayName: "Replace DC Cluster Group"
// Replace given DC Cluster Group in given namespace
type ReplaceSpecType struct {
}

func (m *ReplaceSpecType) Reset()                    { *m = ReplaceSpecType{} }
func (*ReplaceSpecType) ProtoMessage()               {}
func (*ReplaceSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

// Get DC Cluster Group
//
// x-displayName: "Get DC Cluster Group"
// Gets DC Cluster Group in given namespace
type GetSpecType struct {
}

func (m *GetSpecType) Reset()                    { *m = GetSpecType{} }
func (*GetSpecType) ProtoMessage()               {}
func (*GetSpecType) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.dc_cluster_group.GlobalSpecType")
	golang_proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.dc_cluster_group.GlobalSpecType")
	proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.dc_cluster_group.CreateSpecType")
	golang_proto.RegisterType((*CreateSpecType)(nil), "ves.io.schema.dc_cluster_group.CreateSpecType")
	proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.dc_cluster_group.ReplaceSpecType")
	golang_proto.RegisterType((*ReplaceSpecType)(nil), "ves.io.schema.dc_cluster_group.ReplaceSpecType")
	proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.dc_cluster_group.GetSpecType")
	golang_proto.RegisterType((*GetSpecType)(nil), "ves.io.schema.dc_cluster_group.GetSpecType")
}
func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	return true
}
func (this *CreateSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateSpecType)
	if !ok {
		that2, ok := that.(CreateSpecType)
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
	return true
}
func (this *ReplaceSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReplaceSpecType)
	if !ok {
		that2, ok := that.(ReplaceSpecType)
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
	return true
}
func (this *GetSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSpecType)
	if !ok {
		that2, ok := that.(GetSpecType)
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
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.GlobalSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.CreateSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReplaceSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.ReplaceSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&dc_cluster_group.GetSpecType{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CreateSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ReplaceSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSpecType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GlobalSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *CreateSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ReplaceSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetSpecType) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *CreateSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *ReplaceSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReplaceSpecType{`,
		`}`,
	}, "")
	return s
}
func (this *GetSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSpecType{`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func (m *CreateSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: CreateSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func (m *ReplaceSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ReplaceSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func (m *GetSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GetSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTypes
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
				next, err := skipTypes(dAtA[start:])
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
	ErrInvalidLengthTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/dc_cluster_group/types.proto", fileDescriptorTypes) }
func init() {
	golang_proto.RegisterFile("ves.io/schema/dc_cluster_group/types.proto", fileDescriptorTypes)
}

var fileDescriptorTypes = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0x87, 0xe3, 0xe5, 0x1d, 0xf2, 0x4a, 0xe5, 0xcf, 0x82, 0x28, 0xd2, 0x09, 0x15, 0x26, 0xa4,
	0xda, 0x03, 0x1b, 0x03, 0x03, 0x0c, 0xdd, 0xa1, 0x13, 0x4b, 0x95, 0xb8, 0x57, 0x37, 0x22, 0xe9,
	0x59, 0x8e, 0x13, 0xd1, 0x8d, 0x99, 0x89, 0x8f, 0xc1, 0xc7, 0x60, 0x44, 0x4c, 0x1d, 0x3b, 0x12,
	0x77, 0x61, 0xec, 0x47, 0x40, 0x72, 0x52, 0x89, 0x54, 0xa2, 0xdb, 0x9d, 0x9f, 0xc7, 0xf6, 0xcf,
	0xbe, 0xf0, 0xa2, 0xc4, 0x9c, 0x27, 0x24, 0x72, 0x39, 0xc5, 0x2c, 0x12, 0x63, 0x39, 0x92, 0x69,
	0x91, 0x5b, 0x34, 0x23, 0x65, 0xa8, 0xd0, 0xc2, 0xce, 0x35, 0xe6, 0x5c, 0x1b, 0xb2, 0x74, 0x08,
	0xb5, 0xcb, 0x6b, 0x97, 0x6f, 0xbb, 0xdd, 0xbe, 0x4a, 0xec, 0xb4, 0x88, 0xb9, 0xa4, 0x4c, 0x28,
	0x52, 0x24, 0xfc, 0xb6, 0xb8, 0x98, 0xf8, 0xce, 0x37, 0xbe, 0xaa, 0x8f, 0xeb, 0x1e, 0xb5, 0xaf,
	0x9e, 0xa1, 0x6d, 0xc0, 0x49, 0x1b, 0x90, 0xb6, 0x09, 0xcd, 0x9a, 0x10, 0xdd, 0xe3, 0x36, 0xfc,
	0x95, 0xaf, 0xb7, 0x1f, 0x76, 0x06, 0x29, 0xc5, 0x51, 0x7a, 0xaf, 0x51, 0x0e, 0xe7, 0x1a, 0x7b,
	0x67, 0x61, 0xe7, 0xd6, 0x60, 0x64, 0x71, 0xb3, 0x72, 0x75, 0xf0, 0x79, 0xbd, 0x2d, 0x9d, 0x87,
	0x7b, 0x77, 0xa8, 0xd3, 0x48, 0xee, 0xb4, 0x4e, 0xc3, 0xff, 0x03, 0xb4, 0x3b, 0x8c, 0x9b, 0x17,
	0xb6, 0xa8, 0x20, 0x58, 0x56, 0x10, 0xac, 0x2b, 0x60, 0xcf, 0x0e, 0xd8, 0x9b, 0x03, 0xf6, 0xe1,
	0x80, 0x2d, 0x1c, 0xb0, 0xa5, 0x03, 0xf6, 0xe5, 0x80, 0x7d, 0x3b, 0x08, 0xd6, 0x0e, 0xd8, 0xeb,
	0x0a, 0x82, 0xf7, 0x15, 0xb0, 0x87, 0xa1, 0x22, 0xfd, 0xa8, 0x78, 0x49, 0xa9, 0x45, 0x63, 0x22,
	0x5e, 0xe4, 0xc2, 0x17, 0x13, 0x32, 0x59, 0x5f, 0x1b, 0x2a, 0x93, 0x31, 0x9a, 0xfe, 0x06, 0x0b,
	0x1d, 0x2b, 0x12, 0xf8, 0x64, 0x9b, 0xf7, 0xff, 0x31, 0xb7, 0xf8, 0x9f, 0xff, 0x92, 0xcb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xfa, 0x35, 0x6c, 0xe0, 0x01, 0x00, 0x00,
}
