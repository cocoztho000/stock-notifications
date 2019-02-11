// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stock_data.proto

package stockdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Stock struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stock) Reset()         { *m = Stock{} }
func (m *Stock) String() string { return proto.CompactTextString(m) }
func (*Stock) ProtoMessage()    {}
func (*Stock) Descriptor() ([]byte, []int) {
	return fileDescriptor_stock_data_ae3ee931603ed480, []int{0}
}
func (m *Stock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stock.Unmarshal(m, b)
}
func (m *Stock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stock.Marshal(b, m, deterministic)
}
func (dst *Stock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stock.Merge(dst, src)
}
func (m *Stock) XXX_Size() int {
	return xxx_messageInfo_Stock.Size(m)
}
func (m *Stock) XXX_DiscardUnknown() {
	xxx_messageInfo_Stock.DiscardUnknown(m)
}

var xxx_messageInfo_Stock proto.InternalMessageInfo

func (m *Stock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stock) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type StockSummary struct {
	StocksReceived       int32    `protobuf:"varint,1,opt,name=stocks_received,json=stocksReceived,proto3" json:"stocks_received,omitempty"`
	ElapsedTime          int32    `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockSummary) Reset()         { *m = StockSummary{} }
func (m *StockSummary) String() string { return proto.CompactTextString(m) }
func (*StockSummary) ProtoMessage()    {}
func (*StockSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_stock_data_ae3ee931603ed480, []int{1}
}
func (m *StockSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockSummary.Unmarshal(m, b)
}
func (m *StockSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockSummary.Marshal(b, m, deterministic)
}
func (dst *StockSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockSummary.Merge(dst, src)
}
func (m *StockSummary) XXX_Size() int {
	return xxx_messageInfo_StockSummary.Size(m)
}
func (m *StockSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_StockSummary.DiscardUnknown(m)
}

var xxx_messageInfo_StockSummary proto.InternalMessageInfo

func (m *StockSummary) GetStocksReceived() int32 {
	if m != nil {
		return m.StocksReceived
	}
	return 0
}

func (m *StockSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Stock)(nil), "stockdata.Stock")
	proto.RegisterType((*StockSummary)(nil), "stockdata.StockSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockDataClient is the client API for StockData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockDataClient interface {
	// The Aggregator uses this to stream data to the server
	//
	// StockRecorder accepts a stream of stocks to save and send to the engine
	StockRecorder(ctx context.Context, opts ...grpc.CallOption) (StockData_StockRecorderClient, error)
}

type stockDataClient struct {
	cc *grpc.ClientConn
}

func NewStockDataClient(cc *grpc.ClientConn) StockDataClient {
	return &stockDataClient{cc}
}

func (c *stockDataClient) StockRecorder(ctx context.Context, opts ...grpc.CallOption) (StockData_StockRecorderClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StockData_serviceDesc.Streams[0], "/stockdata.StockData/StockRecorder", opts...)
	if err != nil {
		return nil, err
	}
	x := &stockDataStockRecorderClient{stream}
	return x, nil
}

type StockData_StockRecorderClient interface {
	Send(*Stock) error
	CloseAndRecv() (*StockSummary, error)
	grpc.ClientStream
}

type stockDataStockRecorderClient struct {
	grpc.ClientStream
}

func (x *stockDataStockRecorderClient) Send(m *Stock) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stockDataStockRecorderClient) CloseAndRecv() (*StockSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StockSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StockDataServer is the server API for StockData service.
type StockDataServer interface {
	// The Aggregator uses this to stream data to the server
	//
	// StockRecorder accepts a stream of stocks to save and send to the engine
	StockRecorder(StockData_StockRecorderServer) error
}

func RegisterStockDataServer(s *grpc.Server, srv StockDataServer) {
	s.RegisterService(&_StockData_serviceDesc, srv)
}

func _StockData_StockRecorder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StockDataServer).StockRecorder(&stockDataStockRecorderServer{stream})
}

type StockData_StockRecorderServer interface {
	SendAndClose(*StockSummary) error
	Recv() (*Stock, error)
	grpc.ServerStream
}

type stockDataStockRecorderServer struct {
	grpc.ServerStream
}

func (x *stockDataStockRecorderServer) SendAndClose(m *StockSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stockDataStockRecorderServer) Recv() (*Stock, error) {
	m := new(Stock)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StockData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stockdata.StockData",
	HandlerType: (*StockDataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StockRecorder",
			Handler:       _StockData_StockRecorder_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "stock_data.proto",
}

func init() { proto.RegisterFile("stock_data.proto", fileDescriptor_stock_data_ae3ee931603ed480) }

var fileDescriptor_stock_data_ae3ee931603ed480 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0x84, 0x91, 0x72, 0x04, 0x88, 0x4e, 0x48, 0x44, 0x4c, 0x21, 0x0b, 0x99, 0x22,
	0x01, 0x3b, 0x13, 0x1b, 0x9b, 0xc3, 0xc4, 0x12, 0xb9, 0xf6, 0x0d, 0x56, 0xeb, 0x3a, 0x72, 0xdc,
	0x4a, 0xfd, 0xf7, 0x55, 0x2e, 0x51, 0x87, 0x6e, 0xf7, 0xbe, 0xd3, 0x7d, 0x7a, 0x07, 0xe5, 0x94,
	0x82, 0xd9, 0x0e, 0x56, 0x27, 0xdd, 0x8d, 0x31, 0xa4, 0x80, 0x39, 0x93, 0x19, 0x34, 0x1f, 0x20,
	0xfb, 0x39, 0x20, 0xc2, 0xed, 0x5e, 0x7b, 0xaa, 0x44, 0x2d, 0xda, 0x5c, 0xf1, 0x8c, 0xcf, 0x20,
	0xc7, 0xe8, 0x0c, 0x55, 0x59, 0x2d, 0xda, 0x4c, 0x2d, 0xa1, 0xf9, 0x87, 0x82, 0x4f, 0xfa, 0x83,
	0xf7, 0x3a, 0x9e, 0xf0, 0x1d, 0x9e, 0xd8, 0x37, 0x0d, 0x91, 0x0c, 0xb9, 0x23, 0x59, 0x96, 0x48,
	0xf5, 0xb8, 0x60, 0xb5, 0x52, 0x7c, 0x83, 0x82, 0x76, 0x7a, 0x9c, 0xc8, 0x0e, 0xc9, 0xf9, 0xc5,
	0x2a, 0xd5, 0xfd, 0xca, 0xfe, 0x9c, 0xa7, 0xcf, 0x5f, 0xc8, 0xd9, 0xfd, 0xa3, 0x93, 0xc6, 0x6f,
	0x78, 0xe0, 0xa0, 0xc8, 0x84, 0x68, 0x29, 0x62, 0xd9, 0x5d, 0x8a, 0x77, 0xbc, 0x79, 0x7d, 0xb9,
	0x26, 0x6b, 0xa9, 0xe6, 0xa6, 0x15, 0x9b, 0x3b, 0xfe, 0xf6, 0xeb, 0x1c, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0xe3, 0xa0, 0xce, 0x01, 0x01, 0x00, 0x00,
}
