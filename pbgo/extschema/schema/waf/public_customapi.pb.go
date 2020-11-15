// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/waf/public_customapi.proto

package waf

import (
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"

	fmt "fmt"

	math "math"

	_ "github.com/gogo/protobuf/gogoproto"

	_ "github.com/gogo/googleapis/google/api"

	_ "github.com/gogo/protobuf/types"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema4 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"

	ves_io_schema_waf_rules2 "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rules"

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

// VirtualHostWafStatusReq
//
// x-displayName: "Virtual Host WAF Status Request"
// Request to get the detailed WAF configuration defined for all waf instances for a given virtual_host
// object identified by (Namespace, Name)
type VirtualHostWafStatusReq struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "blogging-app-namespace-1"
	// Namespace of the virtual host
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Virtual Host Name"
	// x-example: "greatblogs-vhost"
	// Name of the virtual host for which waf status is requested
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *VirtualHostWafStatusReq) Reset()      { *m = VirtualHostWafStatusReq{} }
func (*VirtualHostWafStatusReq) ProtoMessage() {}
func (*VirtualHostWafStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{0}
}

func (m *VirtualHostWafStatusReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *VirtualHostWafStatusReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// WAF Status
//
// x-displayName: "WAF Status"
// A list of detailed rule configurations currently enabled for a WAF instance.
type WafStatus struct {
	// WAF Instance
	//
	// x-displayName: "WAF Instance"
	// WAF instance associated with this virtual host.
	WafRef []*ves_io_schema4.ObjectRefType `protobuf:"bytes,1,rep,name=waf_ref,json=wafRef" json:"waf_ref,omitempty"`
	// WAF Rules Instance
	//
	// x-displayName: "WAF Rules Instance"
	// WAF Rules instance associated with this virtual host.
	WafRulesRef []*ves_io_schema4.ObjectRefType `protobuf:"bytes,2,rep,name=waf_rules_ref,json=wafRulesRef" json:"waf_rules_ref,omitempty"`
	// WAF Rules Status
	//
	// x-displayName: "WAF Rules Status"
	// WAF rules configured for this waf instance
	WafRulesStatus []*ves_io_schema_waf_rules2.WafRulesStatus `protobuf:"bytes,3,rep,name=waf_rules_status,json=wafRulesStatus" json:"waf_rules_status,omitempty"`
}

func (m *WafStatus) Reset()                    { *m = WafStatus{} }
func (*WafStatus) ProtoMessage()               {}
func (*WafStatus) Descriptor() ([]byte, []int) { return fileDescriptorPublicCustomapi, []int{1} }

func (m *WafStatus) GetWafRef() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.WafRef
	}
	return nil
}

func (m *WafStatus) GetWafRulesRef() []*ves_io_schema4.ObjectRefType {
	if m != nil {
		return m.WafRulesRef
	}
	return nil
}

func (m *WafStatus) GetWafRulesStatus() []*ves_io_schema_waf_rules2.WafRulesStatus {
	if m != nil {
		return m.WafRulesStatus
	}
	return nil
}

// Virtual Host WAF Status Response
//
// x-displayName: "Virtual Host WAF Rules Status Response"
// Response is a list of detailed rule configurations currently enabled for the given virtual_host.
type VirtualHostWafStatusRsp struct {
	// Virtual Host WAF Status
	//
	// x-displayName: "Virtual Host WAF Status"
	// Detailed configuration of all WAF instances under this virtual host
	WafStatus []*WafStatus `protobuf:"bytes,1,rep,name=waf_status,json=wafStatus" json:"waf_status,omitempty"`
}

func (m *VirtualHostWafStatusRsp) Reset()      { *m = VirtualHostWafStatusRsp{} }
func (*VirtualHostWafStatusRsp) ProtoMessage() {}
func (*VirtualHostWafStatusRsp) Descriptor() ([]byte, []int) {
	return fileDescriptorPublicCustomapi, []int{2}
}

func (m *VirtualHostWafStatusRsp) GetWafStatus() []*WafStatus {
	if m != nil {
		return m.WafStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualHostWafStatusReq)(nil), "ves.io.schema.waf.VirtualHostWafStatusReq")
	golang_proto.RegisterType((*VirtualHostWafStatusReq)(nil), "ves.io.schema.waf.VirtualHostWafStatusReq")
	proto.RegisterType((*WafStatus)(nil), "ves.io.schema.waf.WafStatus")
	golang_proto.RegisterType((*WafStatus)(nil), "ves.io.schema.waf.WafStatus")
	proto.RegisterType((*VirtualHostWafStatusRsp)(nil), "ves.io.schema.waf.VirtualHostWafStatusRsp")
	golang_proto.RegisterType((*VirtualHostWafStatusRsp)(nil), "ves.io.schema.waf.VirtualHostWafStatusRsp")
}
func (this *VirtualHostWafStatusReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostWafStatusReq)
	if !ok {
		that2, ok := that.(VirtualHostWafStatusReq)
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
	return true
}
func (this *WafStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WafStatus)
	if !ok {
		that2, ok := that.(WafStatus)
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
	if len(this.WafRef) != len(that1.WafRef) {
		return false
	}
	for i := range this.WafRef {
		if !this.WafRef[i].Equal(that1.WafRef[i]) {
			return false
		}
	}
	if len(this.WafRulesRef) != len(that1.WafRulesRef) {
		return false
	}
	for i := range this.WafRulesRef {
		if !this.WafRulesRef[i].Equal(that1.WafRulesRef[i]) {
			return false
		}
	}
	if len(this.WafRulesStatus) != len(that1.WafRulesStatus) {
		return false
	}
	for i := range this.WafRulesStatus {
		if !this.WafRulesStatus[i].Equal(that1.WafRulesStatus[i]) {
			return false
		}
	}
	return true
}
func (this *VirtualHostWafStatusRsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostWafStatusRsp)
	if !ok {
		that2, ok := that.(VirtualHostWafStatusRsp)
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
	if len(this.WafStatus) != len(that1.WafStatus) {
		return false
	}
	for i := range this.WafStatus {
		if !this.WafStatus[i].Equal(that1.WafStatus[i]) {
			return false
		}
	}
	return true
}
func (this *VirtualHostWafStatusReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&waf.VirtualHostWafStatusReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WafStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&waf.WafStatus{")
	if this.WafRef != nil {
		s = append(s, "WafRef: "+fmt.Sprintf("%#v", this.WafRef)+",\n")
	}
	if this.WafRulesRef != nil {
		s = append(s, "WafRulesRef: "+fmt.Sprintf("%#v", this.WafRulesRef)+",\n")
	}
	if this.WafRulesStatus != nil {
		s = append(s, "WafRulesStatus: "+fmt.Sprintf("%#v", this.WafRulesStatus)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VirtualHostWafStatusRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&waf.VirtualHostWafStatusRsp{")
	if this.WafStatus != nil {
		s = append(s, "WafStatus: "+fmt.Sprintf("%#v", this.WafStatus)+",\n")
	}
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

// Client API for CustomAPI service

type CustomAPIClient interface {
	// Virtual Host WAF Status
	//
	// x-displayName: "Virtual Host WAF Status"
	// Virtual Host WAF Status API is used to get information about the exact configuration, including
	// a list of waf instances and their respective waf-rules that are currently configured for various
	// routes on a given virtual_host identified by (Namespace, Name).
	VirtualHostWafStatus(ctx context.Context, in *VirtualHostWafStatusReq, opts ...grpc.CallOption) (*VirtualHostWafStatusRsp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) VirtualHostWafStatus(ctx context.Context, in *VirtualHostWafStatusReq, opts ...grpc.CallOption) (*VirtualHostWafStatusRsp, error) {
	out := new(VirtualHostWafStatusRsp)
	err := grpc.Invoke(ctx, "/ves.io.schema.waf.CustomAPI/VirtualHostWafStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomAPI service

type CustomAPIServer interface {
	// Virtual Host WAF Status
	//
	// x-displayName: "Virtual Host WAF Status"
	// Virtual Host WAF Status API is used to get information about the exact configuration, including
	// a list of waf instances and their respective waf-rules that are currently configured for various
	// routes on a given virtual_host identified by (Namespace, Name).
	VirtualHostWafStatus(context.Context, *VirtualHostWafStatusReq) (*VirtualHostWafStatusRsp, error)
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_VirtualHostWafStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHostWafStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).VirtualHostWafStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.waf.CustomAPI/VirtualHostWafStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).VirtualHostWafStatus(ctx, req.(*VirtualHostWafStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.waf.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VirtualHostWafStatus",
			Handler:    _CustomAPI_VirtualHostWafStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/waf/public_customapi.proto",
}

func (m *VirtualHostWafStatusReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualHostWafStatusReq) MarshalTo(dAtA []byte) (int, error) {
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
	return i, nil
}

func (m *WafStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WafStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.WafRef) > 0 {
		for _, msg := range m.WafRef {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.WafRulesRef) > 0 {
		for _, msg := range m.WafRulesRef {
			dAtA[i] = 0x12
			i++
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.WafRulesStatus) > 0 {
		for _, msg := range m.WafRulesStatus {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *VirtualHostWafStatusRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VirtualHostWafStatusRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.WafStatus) > 0 {
		for _, msg := range m.WafStatus {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
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
func (m *VirtualHostWafStatusReq) Size() (n int) {
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
	return n
}

func (m *WafStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.WafRef) > 0 {
		for _, e := range m.WafRef {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	if len(m.WafRulesRef) > 0 {
		for _, e := range m.WafRulesRef {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	if len(m.WafRulesStatus) > 0 {
		for _, e := range m.WafRulesStatus {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func (m *VirtualHostWafStatusRsp) Size() (n int) {
	var l int
	_ = l
	if len(m.WafStatus) > 0 {
		for _, e := range m.WafStatus {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
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
func (this *VirtualHostWafStatusReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VirtualHostWafStatusReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WafStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WafStatus{`,
		`WafRef:` + strings.Replace(fmt.Sprintf("%v", this.WafRef), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`WafRulesRef:` + strings.Replace(fmt.Sprintf("%v", this.WafRulesRef), "ObjectRefType", "ves_io_schema4.ObjectRefType", 1) + `,`,
		`WafRulesStatus:` + strings.Replace(fmt.Sprintf("%v", this.WafRulesStatus), "WafRulesStatus", "ves_io_schema_waf_rules2.WafRulesStatus", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VirtualHostWafStatusRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VirtualHostWafStatusRsp{`,
		`WafStatus:` + strings.Replace(fmt.Sprintf("%v", this.WafStatus), "WafStatus", "WafStatus", 1) + `,`,
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
func (m *VirtualHostWafStatusReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VirtualHostWafStatusReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualHostWafStatusReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *WafStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WafStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WafStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WafRef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WafRef = append(m.WafRef, &ves_io_schema4.ObjectRefType{})
			if err := m.WafRef[len(m.WafRef)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WafRulesRef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WafRulesRef = append(m.WafRulesRef, &ves_io_schema4.ObjectRefType{})
			if err := m.WafRulesRef[len(m.WafRulesRef)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WafRulesStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WafRulesStatus = append(m.WafRulesStatus, &ves_io_schema_waf_rules2.WafRulesStatus{})
			if err := m.WafRulesStatus[len(m.WafRulesStatus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *VirtualHostWafStatusRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VirtualHostWafStatusRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VirtualHostWafStatusRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WafStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WafStatus = append(m.WafStatus, &WafStatus{})
			if err := m.WafStatus[len(m.WafStatus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
	proto.RegisterFile("ves.io/schema/waf/public_customapi.proto", fileDescriptorPublicCustomapi)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/waf/public_customapi.proto", fileDescriptorPublicCustomapi)
}

var fileDescriptorPublicCustomapi = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x49, 0x89, 0x64, 0xaa, 0xa2, 0x83, 0x60, 0x8c, 0x71, 0x28, 0xb9, 0x18, 0x84,
	0xcc, 0x40, 0x3d, 0x4a, 0x11, 0x2b, 0xa2, 0x22, 0xf8, 0x23, 0x4a, 0x0b, 0x5e, 0xc2, 0xec, 0x76,
	0x66, 0xb3, 0xba, 0x9b, 0x19, 0x77, 0x66, 0x13, 0x83, 0x04, 0xc4, 0xbf, 0x40, 0xf0, 0x9f, 0x90,
	0xde, 0x3d, 0x79, 0xe9, 0x51, 0x6f, 0x05, 0x2f, 0x1e, 0xcd, 0xd6, 0x83, 0x17, 0xa1, 0x7f, 0x42,
	0xc9, 0xec, 0x66, 0xd3, 0x90, 0x04, 0x7a, 0x59, 0xde, 0xdb, 0xef, 0x7b, 0x9f, 0x79, 0xbf, 0x60,
	0xb3, 0xcf, 0x35, 0xf1, 0x25, 0xd5, 0x6e, 0x97, 0x87, 0x8c, 0x0e, 0x98, 0xa0, 0x2a, 0x76, 0x02,
	0xdf, 0xed, 0xb8, 0xb1, 0x36, 0x32, 0x64, 0xca, 0x27, 0x2a, 0x92, 0x46, 0xa2, 0xcb, 0x69, 0x24,
	0x49, 0x23, 0xc9, 0x80, 0x89, 0x5a, 0xcb, 0xf3, 0x4d, 0x37, 0x76, 0x88, 0x2b, 0x43, 0xea, 0x49,
	0x4f, 0x52, 0x1b, 0xe9, 0xc4, 0xc2, 0x7a, 0xd6, 0xb1, 0x56, 0x4a, 0xa8, 0xd5, 0x3d, 0x29, 0xbd,
	0x80, 0x53, 0xa6, 0x7c, 0xca, 0x7a, 0x3d, 0x69, 0x98, 0xf1, 0x65, 0x4f, 0x67, 0xea, 0xf5, 0x4c,
	0xcd, 0x19, 0x3c, 0x54, 0x66, 0x38, 0x15, 0xe7, 0xcb, 0x94, 0xea, 0x74, 0xe6, 0xb5, 0x79, 0xd1,
	0x0c, 0x15, 0x9f, 0x4a, 0xf5, 0x79, 0xa9, 0xcf, 0x02, 0x7f, 0x8f, 0x19, 0x9e, 0xa9, 0x78, 0xb1,
	0x79, 0xe9, 0xbc, 0xe1, 0xae, 0xc9, 0xf4, 0x1b, 0x8b, 0xfa, 0x69, 0x38, 0x59, 0x90, 0x3b, 0x51,
	0x1c, 0x70, 0xbd, 0x62, 0x82, 0x8d, 0x27, 0xf0, 0xea, 0x8e, 0x1f, 0x99, 0x98, 0x05, 0x8f, 0xa4,
	0x36, 0xbb, 0x4c, 0xbc, 0x34, 0xcc, 0xc4, 0xba, 0xcd, 0xdf, 0xa1, 0x3a, 0xac, 0xf4, 0x58, 0xc8,
	0xb5, 0x62, 0x2e, 0xaf, 0x82, 0x0d, 0xd0, 0xac, 0xb4, 0x67, 0x3f, 0x10, 0x82, 0x6b, 0x13, 0xa7,
	0x5a, 0xb4, 0x82, 0xb5, 0x1b, 0xff, 0x01, 0xac, 0xe4, 0x08, 0xb4, 0x05, 0xcf, 0xd9, 0xe7, 0xb9,
	0xa8, 0x82, 0x8d, 0x52, 0x73, 0x7d, 0xb3, 0x4e, 0xe6, 0xd7, 0xf5, 0xcc, 0xf6, 0xd5, 0xe6, 0xe2,
	0xd5, 0x50, 0xf1, 0xed, 0xf2, 0xfe, 0xa8, 0x34, 0x60, 0xa2, 0x5d, 0x9e, 0x7c, 0xb8, 0x40, 0x4f,
	0xe1, 0x85, 0xbc, 0x7a, 0x0b, 0x29, 0x9e, 0x01, 0x72, 0x7e, 0x7f, 0x54, 0xc9, 0xb3, 0xda, 0xeb,
	0x13, 0x94, 0xb5, 0xb8, 0x40, 0x2f, 0xe0, 0xa5, 0x19, 0x4f, 0xdb, 0x12, 0xab, 0x25, 0x8b, 0xbc,
	0x49, 0x16, 0xce, 0x28, 0x0d, 0x23, 0xbb, 0x59, 0x7e, 0x36, 0x94, 0x8b, 0x83, 0x39, 0xbf, 0xb1,
	0xb3, 0x62, 0x78, 0x5a, 0xa1, 0x3b, 0x10, 0x4e, 0x30, 0xd9, 0x3b, 0xcb, 0xfb, 0x1f, 0x30, 0x41,
	0x66, 0x49, 0x93, 0xba, 0x53, 0x73, 0xf3, 0x27, 0x80, 0x95, 0xfb, 0x76, 0x51, 0xf7, 0x9e, 0x3f,
	0x46, 0xdf, 0x00, 0xbc, 0xb2, 0xec, 0x19, 0x74, 0x6b, 0x09, 0x6f, 0xc5, 0x32, 0x6b, 0x67, 0x8e,
	0xd5, 0xaa, 0xf1, 0xe0, 0xd3, 0xaf, 0xbf, 0x5f, 0x8a, 0x77, 0xd1, 0x56, 0x76, 0x34, 0x34, 0xdf,
	0xba, 0xa6, 0x1f, 0x72, 0x7b, 0x64, 0x6f, 0xaf, 0x9f, 0x32, 0x3a, 0x5d, 0xa9, 0x0d, 0x4d, 0xbb,
	0x4d, 0x63, 0x46, 0xb5, 0xb5, 0x83, 0xef, 0xa0, 0xb4, 0x3d, 0x3c, 0x1c, 0xe3, 0xc2, 0xef, 0x31,
	0x2e, 0x1c, 0x8f, 0x31, 0xf8, 0x98, 0x60, 0xf0, 0x35, 0xc1, 0xe0, 0x47, 0x82, 0xc1, 0x61, 0x82,
	0xc1, 0x9f, 0x04, 0x83, 0x7f, 0x09, 0x2e, 0x1c, 0x27, 0x18, 0x7c, 0x3e, 0xc2, 0x85, 0x83, 0x23,
	0x0c, 0x5e, 0x3f, 0xf4, 0xa4, 0x7a, 0xeb, 0x91, 0xbe, 0x0c, 0x0c, 0x8f, 0x22, 0x46, 0x62, 0x4d,
	0xad, 0x21, 0x64, 0x14, 0xb6, 0x54, 0x24, 0xfb, 0xfe, 0x1e, 0x8f, 0x5a, 0x53, 0x99, 0x2a, 0xc7,
	0x93, 0x94, 0xbf, 0x37, 0xd9, 0xc5, 0xcf, 0x0e, 0xdf, 0x29, 0xdb, 0x13, 0xbf, 0x7d, 0x12, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xd1, 0x6f, 0x7a, 0x50, 0x04, 0x00, 0x00,
}