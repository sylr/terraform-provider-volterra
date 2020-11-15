// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/public_customapi.proto

package site

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/gogo/googleapis/google/api"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	strings "strings"

	reflect "reflect"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Set state request
//
// x-displayName: "Set Status Request"
// Set status of the site
type SetStateReq struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Site namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-required
	// x-example: "ce398"
	// Site name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// State
	//
	// x-displayName: "State"
	// x-required
	// x-example: 3
	// Desired (target) state for site (3 = STANDBY)
	State SiteState `protobuf:"varint,3,opt,name=state,proto3,enum=ves.io.schema.site.SiteState" json:"state,omitempty"`
}

func (m *SetStateReq) Reset()                    { *m = SetStateReq{} }
func (*SetStateReq) ProtoMessage()               {}
func (*SetStateReq) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{0} }

func (m *SetStateReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SetStateReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetStateReq) GetState() SiteState {
	if m != nil {
		return m.State
	}
	return ONLINE
}

// Set state responde
//
// x-displayName: "Set Status Response"
// Response for set state request, empty because the only resturned information
// is currently error message
type SetStateResp struct {
}

func (m *SetStateResp) Reset()                    { *m = SetStateResp{} }
func (*SetStateResp) ProtoMessage()               {}
func (*SetStateResp) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{1} }

func init() {
	proto.RegisterType((*SetStateReq)(nil), "ves.io.schema.site.SetStateReq")
	golang_proto.RegisterType((*SetStateReq)(nil), "ves.io.schema.site.SetStateReq")
	proto.RegisterType((*SetStateResp)(nil), "ves.io.schema.site.SetStateResp")
	golang_proto.RegisterType((*SetStateResp)(nil), "ves.io.schema.site.SetStateResp")
}
func (this *SetStateReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetStateReq)
	if !ok {
		that2, ok := that.(SetStateReq)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *SetStateResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetStateResp)
	if !ok {
		that2, ok := that.(SetStateResp)
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
func (this *SetStateReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&site.SetStateReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetStateResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&site.SetStateResp{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CustomStateAPI service

type CustomStateAPIClient interface {
	// Set site state
	//
	// x-displayName: "Set site state"
	// Request changing site state but this request goes through validation as some
	// trainsitions are not allowed.
	// It can be used to decomission site by sending state DECOMISSIONING. Example of
	// forbidden state is PROVISIONING and UPGRADING.
	SetState(ctx context.Context, in *SetStateReq, opts ...grpc.CallOption) (*SetStateResp, error)
}

type customStateAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomStateAPIClient(cc *grpc.ClientConn) CustomStateAPIClient {
	return &customStateAPIClient{cc}
}

func (c *customStateAPIClient) SetState(ctx context.Context, in *SetStateReq, opts ...grpc.CallOption) (*SetStateResp, error) {
	out := new(SetStateResp)
	err := grpc.Invoke(ctx, "/ves.io.schema.site.CustomStateAPI/SetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomStateAPI service

type CustomStateAPIServer interface {
	// Set site state
	//
	// x-displayName: "Set site state"
	// Request changing site state but this request goes through validation as some
	// trainsitions are not allowed.
	// It can be used to decomission site by sending state DECOMISSIONING. Example of
	// forbidden state is PROVISIONING and UPGRADING.
	SetState(context.Context, *SetStateReq) (*SetStateResp, error)
}

func RegisterCustomStateAPIServer(s *grpc.Server, srv CustomStateAPIServer) {
	s.RegisterService(&_CustomStateAPI_serviceDesc, srv)
}

func _CustomStateAPI_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomStateAPIServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.CustomStateAPI/SetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomStateAPIServer).SetState(ctx, req.(*SetStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomStateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.CustomStateAPI",
	HandlerType: (*CustomStateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetState",
			Handler:    _CustomStateAPI_SetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/public_customapi.proto",
}

func (m *SetStateReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetStateReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.State != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func (m *SetStateResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetStateResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SetStateReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovPublicCustomapi(uint64(m.State))
	}
	return n
}

func (m *SetStateResp) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SetStateReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetStateReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetStateResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetStateResp{`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SetStateReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: SetStateReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetStateReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (SiteState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *SetStateResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: SetStateResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetStateResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
				return 0, ErrInvalidLengthPublicCustomapi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPublicCustomapi
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
				next, err := skipPublicCustomapi(dAtA[start:])
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
	ErrInvalidLengthPublicCustomapi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ves.io/schema/site/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x3d, 0x6f, 0x14, 0x31,
	0x10, 0xdd, 0x49, 0x00, 0x11, 0x83, 0xae, 0x70, 0x75, 0x59, 0x82, 0x39, 0xb6, 0x0a, 0x48, 0x67,
	0xa3, 0x44, 0x34, 0x50, 0x01, 0x0d, 0x54, 0xa0, 0x5c, 0x47, 0x83, 0xbc, 0x8b, 0xe3, 0x18, 0x6e,
	0x77, 0xcc, 0xda, 0x7b, 0xe2, 0x43, 0x91, 0x10, 0x25, 0x15, 0x12, 0x7f, 0x82, 0xff, 0x90, 0xe6,
	0x4a, 0xca, 0x08, 0x1a, 0x4a, 0x6e, 0x43, 0x41, 0x99, 0x9f, 0x80, 0xce, 0x9b, 0xcd, 0xe6, 0xc4,
	0x29, 0xdd, 0x9b, 0x79, 0xcf, 0x33, 0xcf, 0xf3, 0xc8, 0xad, 0x89, 0x72, 0xdc, 0xa0, 0x70, 0xd9,
	0x9e, 0xca, 0xa5, 0x70, 0xc6, 0x2b, 0x61, 0xab, 0x74, 0x6c, 0xb2, 0x17, 0x59, 0xe5, 0x3c, 0xe6,
	0xd2, 0x1a, 0x6e, 0x4b, 0xf4, 0x48, 0x69, 0x23, 0xe5, 0x8d, 0x94, 0xcf, 0xa5, 0xf1, 0x50, 0x1b,
	0xbf, 0x57, 0xa5, 0x3c, 0xc3, 0x5c, 0x68, 0xd4, 0x28, 0x82, 0x34, 0xad, 0x76, 0x43, 0x15, 0x8a,
	0x80, 0x9a, 0x11, 0xf1, 0x86, 0x46, 0xd4, 0x63, 0x25, 0xa4, 0x35, 0x42, 0x16, 0x05, 0x7a, 0xe9,
	0x0d, 0x16, 0xee, 0x84, 0xbd, 0xb6, 0xe8, 0x05, 0xed, 0x59, 0x92, 0x2d, 0x31, 0xea, 0xdf, 0x59,
	0xd5, 0xf2, 0xeb, 0x8b, 0xfc, 0x19, 0x2a, 0xf1, 0xe4, 0xca, 0x48, 0xf9, 0x91, 0x97, 0x5e, 0xed,
	0xa8, 0x37, 0x74, 0x83, 0xac, 0x15, 0x32, 0x57, 0xce, 0xca, 0x4c, 0xf5, 0x61, 0x00, 0x9b, 0x6b,
	0x3b, 0x5d, 0x83, 0x52, 0x72, 0x61, 0x5e, 0xf4, 0x57, 0x02, 0x11, 0x30, 0xdd, 0x26, 0x17, 0xdd,
	0xfc, 0x75, 0x7f, 0x75, 0x00, 0x9b, 0xbd, 0xad, 0xeb, 0xfc, 0xff, 0x4b, 0xf0, 0x91, 0xf1, 0xaa,
	0x59, 0xd1, 0x68, 0x93, 0x1e, 0xb9, 0xda, 0x6d, 0x75, 0x76, 0x6b, 0x0a, 0xa4, 0xf7, 0x28, 0x9c,
	0x34, 0xf4, 0x1e, 0x3c, 0x7b, 0x42, 0x3f, 0x03, 0xb9, 0xdc, 0x6a, 0xe8, 0x8d, 0xa5, 0x53, 0x3b,
	0xdf, 0xf1, 0xe0, 0x7c, 0x81, 0xb3, 0xc9, 0xfd, 0x4f, 0x3f, 0xff, 0x7c, 0x5d, 0xb9, 0x9b, 0xdc,
	0x39, 0x89, 0x50, 0x9c, 0x7e, 0xcb, 0x89, 0x0f, 0xa7, 0x78, 0xbf, 0x39, 0x5e, 0x68, 0xec, 0x8b,
	0x60, 0xf6, 0x1e, 0xdc, 0x8e, 0x6f, 0x4e, 0x0f, 0x60, 0xf5, 0xc7, 0x01, 0xac, 0x2f, 0xd9, 0xf2,
	0x34, 0x7d, 0xa5, 0x32, 0xff, 0xf0, 0xfd, 0xe1, 0x8c, 0x45, 0xbf, 0x66, 0x2c, 0x3a, 0x9e, 0x31,
	0xf8, 0x58, 0x33, 0xf8, 0x56, 0x33, 0xf8, 0x5e, 0x33, 0x38, 0xac, 0x19, 0xfc, 0xae, 0x19, 0xfc,
	0xad, 0x59, 0x74, 0x5c, 0x33, 0xf8, 0x72, 0xc4, 0xa2, 0xe9, 0x11, 0x83, 0xe7, 0x8f, 0x35, 0xda,
	0xd7, 0x9a, 0x4f, 0x70, 0xec, 0x55, 0x59, 0x4a, 0x5e, 0x39, 0x11, 0xc0, 0x2e, 0x96, 0xf9, 0xd0,
	0x96, 0x38, 0x31, 0x2f, 0x55, 0x39, 0x6c, 0x69, 0x61, 0x53, 0x8d, 0x42, 0xbd, 0xf5, 0x6d, 0xc8,
	0x5d, 0xd6, 0xe9, 0xa5, 0x90, 0xe5, 0xf6, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xc2, 0xd9,
	0x2f, 0xb1, 0x02, 0x00, 0x00,
}