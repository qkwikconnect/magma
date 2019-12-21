// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package graphgrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Tenant struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (m *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(m, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TenantList struct {
	Tenants              []*Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TenantList) Reset()         { *m = TenantList{} }
func (m *TenantList) String() string { return proto.CompactTextString(m) }
func (*TenantList) ProtoMessage()    {}
func (*TenantList) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *TenantList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantList.Unmarshal(m, b)
}
func (m *TenantList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantList.Marshal(b, m, deterministic)
}
func (m *TenantList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantList.Merge(m, src)
}
func (m *TenantList) XXX_Size() int {
	return xxx_messageInfo_TenantList.Size(m)
}
func (m *TenantList) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantList.DiscardUnknown(m)
}

var xxx_messageInfo_TenantList proto.InternalMessageInfo

func (m *TenantList) GetTenants() []*Tenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

type AlertPayload struct {
	TenantID             string            `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Alertname            string            `protobuf:"bytes,2,opt,name=alertname,proto3" json:"alertname,omitempty"`
	NetworkID            string            `protobuf:"bytes,3,opt,name=networkID,proto3" json:"networkID,omitempty"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AlertPayload) Reset()         { *m = AlertPayload{} }
func (m *AlertPayload) String() string { return proto.CompactTextString(m) }
func (*AlertPayload) ProtoMessage()    {}
func (*AlertPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *AlertPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPayload.Unmarshal(m, b)
}
func (m *AlertPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPayload.Marshal(b, m, deterministic)
}
func (m *AlertPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPayload.Merge(m, src)
}
func (m *AlertPayload) XXX_Size() int {
	return xxx_messageInfo_AlertPayload.Size(m)
}
func (m *AlertPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPayload.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPayload proto.InternalMessageInfo

func (m *AlertPayload) GetTenantID() string {
	if m != nil {
		return m.TenantID
	}
	return ""
}

func (m *AlertPayload) GetAlertname() string {
	if m != nil {
		return m.Alertname
	}
	return ""
}

func (m *AlertPayload) GetNetworkID() string {
	if m != nil {
		return m.NetworkID
	}
	return ""
}

func (m *AlertPayload) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Tenant)(nil), "graph.Tenant")
	proto.RegisterType((*TenantList)(nil), "graph.TenantList")
	proto.RegisterType((*AlertPayload)(nil), "graph.AlertPayload")
	proto.RegisterMapType((map[string]string)(nil), "graph.AlertPayload.LabelsEntry")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0x49, 0x37, 0xbb, 0xbd, 0xeb, 0x8a, 0xce, 0x8a, 0x84, 0x58, 0xd6, 0xa5, 0x2f, 0xee,
	0x83, 0x24, 0x18, 0x59, 0xbf, 0x1e, 0xc4, 0x6a, 0x8b, 0x14, 0x2a, 0x48, 0x5b, 0x7c, 0xf0, 0x6d,
	0x92, 0xde, 0xa6, 0xa1, 0xe9, 0xcc, 0x30, 0x99, 0xb6, 0xe4, 0x57, 0xf9, 0x7f, 0xfc, 0x35, 0x92,
	0x99, 0xa4, 0xb6, 0x6a, 0x11, 0x7d, 0x09, 0xf7, 0x9e, 0x73, 0x6e, 0xee, 0x99, 0x73, 0xa1, 0x2d,
	0x45, 0x12, 0x08, 0xc9, 0x15, 0x27, 0x27, 0xa9, 0xa4, 0x62, 0xe1, 0x3f, 0x4a, 0x39, 0x4f, 0x73,
	0x0c, 0x35, 0x18, 0xaf, 0xe7, 0x21, 0xae, 0x84, 0x2a, 0x8d, 0xc6, 0xbf, 0xfa, 0x95, 0xdc, 0x4a,
	0x2a, 0x04, 0xca, 0xc2, 0xf0, 0xdd, 0xa7, 0xe0, 0x4e, 0x91, 0x51, 0xa6, 0xc8, 0x5d, 0xb0, 0xb3,
	0x99, 0x67, 0x5d, 0x5b, 0x37, 0xed, 0xb1, 0x9d, 0xcd, 0x08, 0x81, 0x16, 0xa3, 0x2b, 0xf4, 0x6c,
	0x8d, 0xe8, 0xba, 0x7b, 0x0b, 0x60, 0xd4, 0xa3, 0xac, 0x50, 0xe4, 0x09, 0x9c, 0x2a, 0xdd, 0x15,
	0x9e, 0x75, 0xed, 0xdc, 0x9c, 0x47, 0x17, 0x81, 0x76, 0x14, 0x18, 0xcd, 0xb8, 0x61, 0xbb, 0xdf,
	0x2d, 0xb8, 0xd3, 0xcb, 0x51, 0xaa, 0xcf, 0xb4, 0xcc, 0x39, 0x9d, 0x11, 0x1f, 0xce, 0x0c, 0x37,
	0xec, 0xd7, 0x1b, 0x77, 0x3d, 0xe9, 0x40, 0x9b, 0x56, 0xda, 0xbd, 0xe5, 0x3f, 0x81, 0x8a, 0x65,
	0xa8, 0xb6, 0x5c, 0x2e, 0x87, 0x7d, 0xcf, 0x31, 0xec, 0x0e, 0x20, 0x2f, 0xc1, 0xcd, 0x69, 0x8c,
	0x79, 0xe1, 0xb5, 0xb4, 0xa1, 0xc7, 0xb5, 0xa1, 0xfd, 0xe5, 0xc1, 0x48, 0x2b, 0x06, 0x4c, 0xc9,
	0x72, 0x5c, 0xcb, 0xfd, 0xd7, 0x70, 0xbe, 0x07, 0x93, 0x7b, 0xe0, 0x2c, 0xb1, 0xac, 0xad, 0x55,
	0x25, 0x79, 0x00, 0x27, 0x1b, 0x9a, 0xaf, 0x1b, 0x47, 0xa6, 0x79, 0x63, 0xbf, 0xb2, 0xa2, 0x6f,
	0x36, 0x5c, 0x98, 0x07, 0x4f, 0x50, 0x6e, 0xb2, 0x04, 0xc9, 0x2d, 0xb8, 0x1f, 0x24, 0x52, 0x85,
	0xa4, 0x13, 0x98, 0xf8, 0x83, 0x26, 0xfe, 0x60, 0xa2, 0x64, 0xc6, 0xd2, 0x2f, 0xd5, 0xb4, 0x7f,
	0x18, 0x17, 0x79, 0x06, 0x2d, 0x1d, 0xeb, 0xc3, 0xdf, 0x86, 0x06, 0xd5, 0x41, 0xfd, 0xfb, 0x07,
	0x72, 0x2d, 0x8d, 0xc0, 0xf9, 0x88, 0xea, 0xdf, 0xd6, 0xbc, 0x83, 0xb3, 0xa9, 0x5c, 0xb3, 0xe4,
	0xef, 0xfe, 0x8e, 0x18, 0x21, 0x6f, 0xc1, 0xed, 0x63, 0x8e, 0xff, 0x3b, 0x1f, 0x7d, 0x82, 0xcb,
	0x5e, 0xa2, 0x32, 0xce, 0x0a, 0x7d, 0x97, 0x26, 0xb6, 0x17, 0x70, 0x3a, 0x95, 0x59, 0x9a, 0xa2,
	0x24, 0x97, 0x7f, 0xb8, 0xdb, 0xb1, 0xdf, 0xbd, 0xbf, 0xfa, 0xda, 0x99, 0xc7, 0x49, 0x58, 0x94,
	0x2b, 0xb1, 0xe0, 0xac, 0x0c, 0xf5, 0xa8, 0xf9, 0xa6, 0x52, 0x24, 0xb1, 0xab, 0xf5, 0xcf, 0x7f,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x07, 0xf4, 0x72, 0x3a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error)
	Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error)
	Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tenantServiceClient struct {
	cc *grpc.ClientConn
}

func NewTenantServiceClient(cc *grpc.ClientConn) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) Create(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/graph.TenantService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Get(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Truncate(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) Delete(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.TenantService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	Create(context.Context, *wrappers.StringValue) (*Tenant, error)
	List(context.Context, *empty.Empty) (*TenantList, error)
	Get(context.Context, *wrappers.StringValue) (*Tenant, error)
	Truncate(context.Context, *wrappers.StringValue) (*empty.Empty, error)
	Delete(context.Context, *wrappers.StringValue) (*empty.Empty, error)
}

// UnimplementedTenantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (*UnimplementedTenantServiceServer) Create(ctx context.Context, req *wrappers.StringValue) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTenantServiceServer) List(ctx context.Context, req *empty.Empty) (*TenantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTenantServiceServer) Get(ctx context.Context, req *wrappers.StringValue) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTenantServiceServer) Truncate(ctx context.Context, req *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (*UnimplementedTenantServiceServer) Delete(ctx context.Context, req *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Create(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Get(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Truncate(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.TenantService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).Delete(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TenantService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TenantService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TenantService_Get_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _TenantService_Truncate_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TenantService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// ActionsAlertServiceClient is the client API for ActionsAlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionsAlertServiceClient interface {
	Trigger(ctx context.Context, in *AlertPayload, opts ...grpc.CallOption) (*empty.Empty, error)
}

type actionsAlertServiceClient struct {
	cc *grpc.ClientConn
}

func NewActionsAlertServiceClient(cc *grpc.ClientConn) ActionsAlertServiceClient {
	return &actionsAlertServiceClient{cc}
}

func (c *actionsAlertServiceClient) Trigger(ctx context.Context, in *AlertPayload, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/graph.ActionsAlertService/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsAlertServiceServer is the server API for ActionsAlertService service.
type ActionsAlertServiceServer interface {
	Trigger(context.Context, *AlertPayload) (*empty.Empty, error)
}

// UnimplementedActionsAlertServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActionsAlertServiceServer struct {
}

func (*UnimplementedActionsAlertServiceServer) Trigger(ctx context.Context, req *AlertPayload) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}

func RegisterActionsAlertServiceServer(s *grpc.Server, srv ActionsAlertServiceServer) {
	s.RegisterService(&_ActionsAlertService_serviceDesc, srv)
}

func _ActionsAlertService_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAlertServiceServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graph.ActionsAlertService/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAlertServiceServer).Trigger(ctx, req.(*AlertPayload))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionsAlertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "graph.ActionsAlertService",
	HandlerType: (*ActionsAlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _ActionsAlertService_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}