// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/spec/v1/orders.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Order struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	FullfillTime         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=fullfill_time,json=fullfillTime,proto3" json:"fullfill_time,omitempty"`
	Items                []*OrderItem         `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_009697da2aa4fe60, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Order) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Order) GetFullfillTime() *timestamp.Timestamp {
	if m != nil {
		return m.FullfillTime
	}
	return nil
}

func (m *Order) GetItems() []*OrderItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type OrderItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Item                 *Item    `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	Pieces               int32    `protobuf:"varint,4,opt,name=pieces,proto3" json:"pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_009697da2aa4fe60, []int{1}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderItem) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *OrderItem) GetPieces() int32 {
	if m != nil {
		return m.Pieces
	}
	return 0
}

type Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_009697da2aa4fe60, []int{2}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "cqrs.orders.Order")
	proto.RegisterType((*OrderItem)(nil), "cqrs.orders.OrderItem")
	proto.RegisterType((*Item)(nil), "cqrs.orders.Item")
}

func init() { proto.RegisterFile("pkg/spec/v1/orders.proto", fileDescriptor_009697da2aa4fe60) }

var fileDescriptor_009697da2aa4fe60 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x5d, 0xd7, 0x0e, 0xf6, 0x32, 0x0f, 0x06, 0x19, 0x65, 0x17, 0x67, 0x41, 0xd8, 0x41,
	0x52, 0x36, 0x0f, 0x1e, 0x3c, 0x88, 0x7a, 0xf2, 0xa2, 0x50, 0x3d, 0x79, 0x70, 0x74, 0x5d, 0x36,
	0x82, 0xa9, 0x8d, 0x49, 0x3a, 0xf0, 0xee, 0x1f, 0x2e, 0x79, 0x69, 0x87, 0xe2, 0x64, 0xb7, 0xe6,
	0x7b, 0xdf, 0xef, 0xe3, 0xbd, 0xaf, 0x10, 0xab, 0xb7, 0x75, 0x6a, 0x14, 0x2f, 0xd2, 0xcd, 0x34,
	0xad, 0xf4, 0x92, 0x6b, 0xc3, 0x94, 0xae, 0x6c, 0x45, 0x49, 0xf1, 0xa1, 0x0d, 0xf3, 0xd2, 0xe8,
	0x64, 0x5d, 0x55, 0x6b, 0xc9, 0x53, 0x1c, 0x2d, 0xea, 0x55, 0x6a, 0x45, 0xc9, 0x8d, 0xcd, 0x4b,
	0xe5, 0xdd, 0xc9, 0x57, 0x00, 0xd1, 0xa3, 0xf3, 0x52, 0x0a, 0xe1, 0x7b, 0x5e, 0xf2, 0xb8, 0x33,
	0xee, 0x4c, 0xfa, 0x19, 0x7e, 0xd3, 0x2b, 0x20, 0x85, 0xe6, 0xb9, 0xe5, 0x73, 0xc7, 0xc5, 0xc1,
	0xb8, 0x33, 0x21, 0xb3, 0x11, 0xf3, 0xa1, 0xac, 0x0d, 0x65, 0xcf, 0x6d, 0x68, 0x06, 0xde, 0xee,
	0x04, 0x07, 0xd7, 0x6a, 0xb9, 0x85, 0xbb, 0xfb, 0x61, 0x6f, 0x47, 0xf8, 0x1a, 0x0e, 0x57, 0xb5,
	0x94, 0x2b, 0x21, 0xa5, 0xc7, 0xc3, 0xbd, 0xf8, 0xa0, 0x05, 0x30, 0xe0, 0x1c, 0x22, 0x61, 0x79,
	0x69, 0xe2, 0x68, 0xdc, 0x9d, 0x90, 0xd9, 0x90, 0xfd, 0xa8, 0x85, 0xe1, 0xc5, 0xf7, 0x96, 0x97,
	0x99, 0x37, 0x25, 0xaf, 0xd0, 0xdf, 0x6a, 0x3b, 0x9b, 0x38, 0x83, 0xd0, 0x39, 0x9b, 0x2b, 0x8e,
	0x7e, 0xa5, 0x61, 0x10, 0x8e, 0xe9, 0x10, 0x7a, 0x4a, 0xf0, 0x82, 0x1b, 0xdc, 0x37, 0xca, 0x9a,
	0x57, 0xf2, 0x04, 0xe1, 0xbf, 0xd1, 0xa7, 0x30, 0x58, 0x0a, 0xa3, 0x64, 0xfe, 0x39, 0xc7, 0x59,
	0x80, 0x33, 0xd2, 0x68, 0x0f, 0xce, 0x72, 0x0c, 0x91, 0xd2, 0xa2, 0xf0, 0x25, 0x06, 0x99, 0x7f,
	0xcc, 0x6e, 0xa0, 0x87, 0x4b, 0x1b, 0x7a, 0x09, 0xe4, 0x0e, 0x8b, 0x6f, 0x7e, 0xe5, 0xdf, 0x63,
	0x47, 0x3b, 0xb4, 0xe4, 0xe0, 0x36, 0x7c, 0x09, 0x36, 0xd3, 0x45, 0x0f, 0xdb, 0xbc, 0xf8, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x8f, 0xb7, 0x09, 0x55, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrdersClient is the client API for Orders service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrdersClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
}

type ordersClient struct {
	cc *grpc.ClientConn
}

func NewOrdersClient(cc *grpc.ClientConn) OrdersClient {
	return &ordersClient{cc}
}

func (c *ordersClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/cqrs.orders.Orders/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServer is the server API for Orders service.
type OrdersServer interface {
	CreateOrder(context.Context, *Order) (*Order, error)
}

// UnimplementedOrdersServer can be embedded to have forward compatible implementations.
type UnimplementedOrdersServer struct {
}

func (*UnimplementedOrdersServer) CreateOrder(ctx context.Context, req *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func RegisterOrdersServer(s *grpc.Server, srv OrdersServer) {
	s.RegisterService(&_Orders_serviceDesc, srv)
}

func _Orders_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cqrs.orders.Orders/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orders_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cqrs.orders.Orders",
	HandlerType: (*OrdersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Orders_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/spec/v1/orders.proto",
}
