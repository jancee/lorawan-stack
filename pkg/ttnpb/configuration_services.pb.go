// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/configuration_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListFrequencyPlansRequest struct {
	// Optional base frequency in MHz for hardware support (433, 470, 868 or 915)
	BaseFrequency        uint32   `protobuf:"varint,1,opt,name=base_frequency,json=baseFrequency,proto3" json:"base_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFrequencyPlansRequest) Reset()      { *m = ListFrequencyPlansRequest{} }
func (*ListFrequencyPlansRequest) ProtoMessage() {}
func (*ListFrequencyPlansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed64f51a0283877, []int{0}
}
func (m *ListFrequencyPlansRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListFrequencyPlansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListFrequencyPlansRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListFrequencyPlansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFrequencyPlansRequest.Merge(m, src)
}
func (m *ListFrequencyPlansRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListFrequencyPlansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFrequencyPlansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFrequencyPlansRequest proto.InternalMessageInfo

func (m *ListFrequencyPlansRequest) GetBaseFrequency() uint32 {
	if m != nil {
		return m.BaseFrequency
	}
	return 0
}

type FrequencyPlanDescription struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the frequency that the current frequency plan is based on.
	BaseID string `protobuf:"bytes,2,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Base frequency in MHz for hardware support (433, 470, 868 or 915)
	BaseFrequency        uint32   `protobuf:"varint,4,opt,name=base_frequency,json=baseFrequency,proto3" json:"base_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyPlanDescription) Reset()      { *m = FrequencyPlanDescription{} }
func (*FrequencyPlanDescription) ProtoMessage() {}
func (*FrequencyPlanDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed64f51a0283877, []int{1}
}
func (m *FrequencyPlanDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FrequencyPlanDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FrequencyPlanDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FrequencyPlanDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyPlanDescription.Merge(m, src)
}
func (m *FrequencyPlanDescription) XXX_Size() int {
	return m.Size()
}
func (m *FrequencyPlanDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyPlanDescription.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyPlanDescription proto.InternalMessageInfo

func (m *FrequencyPlanDescription) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *FrequencyPlanDescription) GetBaseID() string {
	if m != nil {
		return m.BaseID
	}
	return ""
}

func (m *FrequencyPlanDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FrequencyPlanDescription) GetBaseFrequency() uint32 {
	if m != nil {
		return m.BaseFrequency
	}
	return 0
}

type ListFrequencyPlansResponse struct {
	FrequencyPlans       []*FrequencyPlanDescription `protobuf:"bytes,1,rep,name=frequency_plans,json=frequencyPlans,proto3" json:"frequency_plans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListFrequencyPlansResponse) Reset()      { *m = ListFrequencyPlansResponse{} }
func (*ListFrequencyPlansResponse) ProtoMessage() {}
func (*ListFrequencyPlansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed64f51a0283877, []int{2}
}
func (m *ListFrequencyPlansResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListFrequencyPlansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListFrequencyPlansResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListFrequencyPlansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFrequencyPlansResponse.Merge(m, src)
}
func (m *ListFrequencyPlansResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListFrequencyPlansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFrequencyPlansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFrequencyPlansResponse proto.InternalMessageInfo

func (m *ListFrequencyPlansResponse) GetFrequencyPlans() []*FrequencyPlanDescription {
	if m != nil {
		return m.FrequencyPlans
	}
	return nil
}

func init() {
	proto.RegisterType((*ListFrequencyPlansRequest)(nil), "ttn.lorawan.v3.ListFrequencyPlansRequest")
	golang_proto.RegisterType((*ListFrequencyPlansRequest)(nil), "ttn.lorawan.v3.ListFrequencyPlansRequest")
	proto.RegisterType((*FrequencyPlanDescription)(nil), "ttn.lorawan.v3.FrequencyPlanDescription")
	golang_proto.RegisterType((*FrequencyPlanDescription)(nil), "ttn.lorawan.v3.FrequencyPlanDescription")
	proto.RegisterType((*ListFrequencyPlansResponse)(nil), "ttn.lorawan.v3.ListFrequencyPlansResponse")
	golang_proto.RegisterType((*ListFrequencyPlansResponse)(nil), "ttn.lorawan.v3.ListFrequencyPlansResponse")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/configuration_services.proto", fileDescriptor_2ed64f51a0283877)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/configuration_services.proto", fileDescriptor_2ed64f51a0283877)
}

var fileDescriptor_2ed64f51a0283877 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x9f, 0x97, 0x96, 0x88, 0x23, 0x89, 0x30, 0x07, 0x89, 0x41, 0x5e, 0x43, 0x44, 0x89, 0x62,
	0x76, 0xa1, 0xfd, 0x06, 0x31, 0x08, 0x01, 0x0f, 0xba, 0x47, 0x2f, 0x61, 0xb3, 0x99, 0x6c, 0x86,
	0xa4, 0x33, 0xeb, 0xce, 0xa4, 0xc5, 0x5b, 0xf1, 0xd4, 0xa3, 0xd8, 0x8b, 0x47, 0x11, 0x84, 0x1e,
	0x7b, 0xec, 0xb1, 0xc7, 0x1e, 0x0b, 0x5e, 0x7a, 0x2a, 0xdd, 0x59, 0x0f, 0x3d, 0xf6, 0xd8, 0xa3,
	0xec, 0xd4, 0x06, 0xa3, 0x29, 0x78, 0x9b, 0xb7, 0xef, 0xfd, 0xfe, 0xed, 0xbc, 0xa1, 0xde, 0x54,
	0xa5, 0xe1, 0x76, 0x28, 0xdb, 0xda, 0x84, 0xd1, 0xc4, 0x0f, 0x13, 0xe1, 0x47, 0x4a, 0x8e, 0x44,
	0x3c, 0x4b, 0x43, 0x23, 0x94, 0xec, 0x6b, 0x9e, 0x6e, 0x89, 0x88, 0x6b, 0x2f, 0x49, 0x95, 0x51,
	0xac, 0x6a, 0x8c, 0xbc, 0xc1, 0x78, 0x5b, 0x1b, 0xf5, 0x76, 0x2c, 0xcc, 0x78, 0x36, 0xf0, 0x22,
	0xb5, 0xe9, 0xc7, 0x2a, 0x56, 0xbe, 0x1b, 0x1b, 0xcc, 0x46, 0xae, 0x72, 0x85, 0x3b, 0x5d, 0xc3,
	0xeb, 0x8f, 0x62, 0xa5, 0xe2, 0x29, 0x77, 0x3a, 0xa1, 0x94, 0xca, 0x38, 0x91, 0xdf, 0xe4, 0xcd,
	0x0e, 0x7d, 0xf8, 0x5a, 0x68, 0xf3, 0x2a, 0xe5, 0xef, 0x67, 0x5c, 0x46, 0x1f, 0xde, 0x4c, 0x43,
	0xa9, 0x83, 0xa2, 0xd0, 0x86, 0x3d, 0xa1, 0xd5, 0x41, 0xa8, 0x79, 0x7f, 0x74, 0xd3, 0xad, 0x41,
	0x03, 0x5a, 0x95, 0xa0, 0x52, 0x7c, 0x9d, 0x43, 0x9a, 0x9f, 0x81, 0xd6, 0x16, 0x08, 0xba, 0x5c,
	0x47, 0xa9, 0x48, 0x0a, 0x1d, 0xf6, 0x80, 0x96, 0xc4, 0xd0, 0xe1, 0xee, 0x76, 0xca, 0xf6, 0x6c,
	0xad, 0xd4, 0xeb, 0x06, 0x25, 0x31, 0x64, 0x8f, 0xe9, 0x1d, 0xc7, 0x2d, 0x86, 0xb5, 0x92, 0x6b,
	0x52, 0x7b, 0xb6, 0x56, 0xee, 0x84, 0x9a, 0xf7, 0xba, 0x41, 0xb9, 0x68, 0xf5, 0x86, 0x8c, 0xd1,
	0x55, 0x19, 0x6e, 0xf2, 0xda, 0x4a, 0x31, 0x11, 0xb8, 0xf3, 0x12, 0x53, 0xab, 0xcb, 0x4c, 0x29,
	0x5a, 0x5f, 0x16, 0x4c, 0x27, 0x4a, 0x6a, 0xce, 0xde, 0xd2, 0xfb, 0x73, 0x7c, 0x3f, 0x29, 0x5a,
	0x35, 0x68, 0xac, 0xb4, 0xee, 0xad, 0xb7, 0xbc, 0xc5, 0xbf, 0xed, 0xdd, 0x16, 0x2c, 0xa8, 0x8e,
	0x16, 0xa8, 0xd7, 0xbf, 0x03, 0xad, 0xbc, 0xfc, 0xf3, 0x1e, 0xd9, 0x1e, 0x50, 0xf6, 0xaf, 0x07,
	0xf6, 0xec, 0x6f, 0x89, 0x5b, 0x2f, 0xa0, 0xfe, 0xfc, 0x7f, 0x46, 0xaf, 0x23, 0x35, 0x9f, 0x7e,
	0xfc, 0xf1, 0x73, 0xaf, 0xd4, 0x60, 0xb8, 0xb8, 0x4d, 0xfe, 0xdc, 0x66, 0xdb, 0xe5, 0xec, 0x7c,
	0x83, 0xe3, 0x0c, 0xe1, 0x24, 0x43, 0x38, 0xcd, 0x90, 0x9c, 0x67, 0x48, 0x2e, 0x32, 0x24, 0x97,
	0x19, 0x92, 0xab, 0x0c, 0x61, 0xc7, 0x22, 0xec, 0x5a, 0x24, 0xfb, 0x16, 0xe1, 0xc0, 0x22, 0x39,
	0xb4, 0x48, 0x8e, 0x2c, 0x92, 0x63, 0x8b, 0x70, 0x62, 0x11, 0x4e, 0x2d, 0x92, 0x73, 0x8b, 0x70,
	0x61, 0x91, 0x5c, 0x5a, 0x84, 0x2b, 0x8b, 0x64, 0x27, 0x47, 0xb2, 0x9b, 0x23, 0x7c, 0xca, 0x91,
	0x7c, 0xc9, 0x11, 0xbe, 0xe6, 0x48, 0xf6, 0x73, 0x24, 0x07, 0x39, 0xc2, 0x61, 0x8e, 0x70, 0x94,
	0x23, 0xbc, 0x7b, 0x11, 0x2b, 0xcf, 0x8c, 0xb9, 0x19, 0x0b, 0x19, 0x6b, 0x4f, 0x72, 0xb3, 0xad,
	0xd2, 0x89, 0xbf, 0xf8, 0x14, 0x92, 0x49, 0xec, 0x1b, 0x23, 0x93, 0xc1, 0xa0, 0xec, 0xb6, 0x73,
	0xe3, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x0a, 0x80, 0x1a, 0x2c, 0x03, 0x00, 0x00,
}

func (this *ListFrequencyPlansRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListFrequencyPlansRequest)
	if !ok {
		that2, ok := that.(ListFrequencyPlansRequest)
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
	if this.BaseFrequency != that1.BaseFrequency {
		return false
	}
	return true
}
func (this *FrequencyPlanDescription) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FrequencyPlanDescription)
	if !ok {
		that2, ok := that.(FrequencyPlanDescription)
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
	if this.ID != that1.ID {
		return false
	}
	if this.BaseID != that1.BaseID {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.BaseFrequency != that1.BaseFrequency {
		return false
	}
	return true
}
func (this *ListFrequencyPlansResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListFrequencyPlansResponse)
	if !ok {
		that2, ok := that.(ListFrequencyPlansResponse)
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
	if len(this.FrequencyPlans) != len(that1.FrequencyPlans) {
		return false
	}
	for i := range this.FrequencyPlans {
		if !this.FrequencyPlans[i].Equal(that1.FrequencyPlans[i]) {
			return false
		}
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	ListFrequencyPlans(ctx context.Context, in *ListFrequencyPlansRequest, opts ...grpc.CallOption) (*ListFrequencyPlansResponse, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) ListFrequencyPlans(ctx context.Context, in *ListFrequencyPlansRequest, opts ...grpc.CallOption) (*ListFrequencyPlansResponse, error) {
	out := new(ListFrequencyPlansResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Configuration/ListFrequencyPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	ListFrequencyPlans(context.Context, *ListFrequencyPlansRequest) (*ListFrequencyPlansResponse, error)
}

// UnimplementedConfigurationServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (*UnimplementedConfigurationServer) ListFrequencyPlans(ctx context.Context, req *ListFrequencyPlansRequest) (*ListFrequencyPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFrequencyPlans not implemented")
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_ListFrequencyPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFrequencyPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).ListFrequencyPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Configuration/ListFrequencyPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).ListFrequencyPlans(ctx, req.(*ListFrequencyPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFrequencyPlans",
			Handler:    _Configuration_ListFrequencyPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/configuration_services.proto",
}

func (m *ListFrequencyPlansRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListFrequencyPlansRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListFrequencyPlansRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseFrequency != 0 {
		i = encodeVarintConfigurationServices(dAtA, i, uint64(m.BaseFrequency))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FrequencyPlanDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FrequencyPlanDescription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FrequencyPlanDescription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseFrequency != 0 {
		i = encodeVarintConfigurationServices(dAtA, i, uint64(m.BaseFrequency))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConfigurationServices(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseID) > 0 {
		i -= len(m.BaseID)
		copy(dAtA[i:], m.BaseID)
		i = encodeVarintConfigurationServices(dAtA, i, uint64(len(m.BaseID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintConfigurationServices(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListFrequencyPlansResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListFrequencyPlansResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListFrequencyPlansResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FrequencyPlans) > 0 {
		for iNdEx := len(m.FrequencyPlans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FrequencyPlans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConfigurationServices(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfigurationServices(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfigurationServices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedListFrequencyPlansRequest(r randyConfigurationServices, easy bool) *ListFrequencyPlansRequest {
	this := &ListFrequencyPlansRequest{}
	this.BaseFrequency = r.Uint32()
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedFrequencyPlanDescription(r randyConfigurationServices, easy bool) *FrequencyPlanDescription {
	this := &FrequencyPlanDescription{}
	this.ID = randStringConfigurationServices(r)
	this.BaseID = randStringConfigurationServices(r)
	this.Name = randStringConfigurationServices(r)
	this.BaseFrequency = r.Uint32()
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedListFrequencyPlansResponse(r randyConfigurationServices, easy bool) *ListFrequencyPlansResponse {
	this := &ListFrequencyPlansResponse{}
	if r.Intn(5) != 0 {
		v1 := r.Intn(5)
		this.FrequencyPlans = make([]*FrequencyPlanDescription, v1)
		for i := 0; i < v1; i++ {
			this.FrequencyPlans[i] = NewPopulatedFrequencyPlanDescription(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyConfigurationServices interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneConfigurationServices(r randyConfigurationServices) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringConfigurationServices(r randyConfigurationServices) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneConfigurationServices(r)
	}
	return string(tmps)
}
func randUnrecognizedConfigurationServices(r randyConfigurationServices, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldConfigurationServices(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldConfigurationServices(dAtA []byte, r randyConfigurationServices, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateConfigurationServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateConfigurationServices(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ListFrequencyPlansRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseFrequency != 0 {
		n += 1 + sovConfigurationServices(uint64(m.BaseFrequency))
	}
	return n
}

func (m *FrequencyPlanDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovConfigurationServices(uint64(l))
	}
	l = len(m.BaseID)
	if l > 0 {
		n += 1 + l + sovConfigurationServices(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfigurationServices(uint64(l))
	}
	if m.BaseFrequency != 0 {
		n += 1 + sovConfigurationServices(uint64(m.BaseFrequency))
	}
	return n
}

func (m *ListFrequencyPlansResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FrequencyPlans) > 0 {
		for _, e := range m.FrequencyPlans {
			l = e.Size()
			n += 1 + l + sovConfigurationServices(uint64(l))
		}
	}
	return n
}

func sovConfigurationServices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfigurationServices(x uint64) (n int) {
	return sovConfigurationServices((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ListFrequencyPlansRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListFrequencyPlansRequest{`,
		`BaseFrequency:` + fmt.Sprintf("%v", this.BaseFrequency) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FrequencyPlanDescription) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FrequencyPlanDescription{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`BaseID:` + fmt.Sprintf("%v", this.BaseID) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`BaseFrequency:` + fmt.Sprintf("%v", this.BaseFrequency) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListFrequencyPlansResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForFrequencyPlans := "[]*FrequencyPlanDescription{"
	for _, f := range this.FrequencyPlans {
		repeatedStringForFrequencyPlans += strings.Replace(f.String(), "FrequencyPlanDescription", "FrequencyPlanDescription", 1) + ","
	}
	repeatedStringForFrequencyPlans += "}"
	s := strings.Join([]string{`&ListFrequencyPlansResponse{`,
		`FrequencyPlans:` + repeatedStringForFrequencyPlans + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfigurationServices(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ListFrequencyPlansRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigurationServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListFrequencyPlansRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListFrequencyPlansRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFrequency", wireType)
			}
			m.BaseFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseFrequency |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfigurationServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfigurationServices
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
func (m *FrequencyPlanDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigurationServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FrequencyPlanDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FrequencyPlanDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFrequency", wireType)
			}
			m.BaseFrequency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseFrequency |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfigurationServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfigurationServices
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
func (m *ListFrequencyPlansResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigurationServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListFrequencyPlansResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListFrequencyPlansResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrequencyPlans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FrequencyPlans = append(m.FrequencyPlans, &FrequencyPlanDescription{})
			if err := m.FrequencyPlans[len(m.FrequencyPlans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfigurationServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigurationServices
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfigurationServices
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
func skipConfigurationServices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfigurationServices
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
					return 0, ErrIntOverflowConfigurationServices
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfigurationServices
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
			if length < 0 {
				return 0, ErrInvalidLengthConfigurationServices
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfigurationServices
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfigurationServices
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfigurationServices        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfigurationServices          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfigurationServices = fmt.Errorf("proto: unexpected end of group")
)
