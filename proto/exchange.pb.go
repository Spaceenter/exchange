// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

/*
Package exchange is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	Order
*/
package exchange

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TradingPair int32

const (
	TradingPair_TRADING_PAIR_UNKNOWN TradingPair = 0
	TradingPair_BTC_USD              TradingPair = 1
	TradingPair_BTC_THB              TradingPair = 2
)

var TradingPair_name = map[int32]string{
	0: "TRADING_PAIR_UNKNOWN",
	1: "BTC_USD",
	2: "BTC_THB",
}
var TradingPair_value = map[string]int32{
	"TRADING_PAIR_UNKNOWN": 0,
	"BTC_USD":              1,
	"BTC_THB":              2,
}

func (x TradingPair) String() string {
	return proto.EnumName(TradingPair_name, int32(x))
}
func (TradingPair) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Order_Type int32

const (
	Order_TYPE_UNKNOWN Order_Type = 0
	Order_MARKET       Order_Type = 1
	Order_LIMIT        Order_Type = 2
	Order_CANCEL       Order_Type = 3
)

var Order_Type_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "MARKET",
	2: "LIMIT",
	3: "CANCEL",
}
var Order_Type_value = map[string]int32{
	"TYPE_UNKNOWN": 0,
	"MARKET":       1,
	"LIMIT":        2,
	"CANCEL":       3,
}

func (x Order_Type) String() string {
	return proto.EnumName(Order_Type_name, int32(x))
}
func (Order_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Order_Position int32

const (
	Order_POSITION_UNKNOWN Order_Position = 0
	Order_BID              Order_Position = 1
	Order_ASK              Order_Position = 2
)

var Order_Position_name = map[int32]string{
	0: "POSITION_UNKNOWN",
	1: "BID",
	2: "ASK",
}
var Order_Position_value = map[string]int32{
	"POSITION_UNKNOWN": 0,
	"BID":              1,
	"ASK":              2,
}

func (x Order_Position) String() string {
	return proto.EnumName(Order_Position_name, int32(x))
}
func (Order_Position) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Order struct {
	Type     Order_Type     `protobuf:"varint,1,opt,name=type,enum=exchange.Order_Type" json:"type,omitempty"`
	Position Order_Position `protobuf:"varint,2,opt,name=position,enum=exchange.Order_Position" json:"position,omitempty"`
	Value    float64        `protobuf:"fixed64,3,opt,name=value" json:"value,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Order) GetType() Order_Type {
	if m != nil {
		return m.Type
	}
	return Order_TYPE_UNKNOWN
}

func (m *Order) GetPosition() Order_Position {
	if m != nil {
		return m.Position
	}
	return Order_POSITION_UNKNOWN
}

func (m *Order) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "exchange.Order")
	proto.RegisterEnum("exchange.TradingPair", TradingPair_name, TradingPair_value)
	proto.RegisterEnum("exchange.Order_Type", Order_Type_name, Order_Type_value)
	proto.RegisterEnum("exchange.Order_Position", Order_Position_name, Order_Position_value)
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4f, 0x83, 0x30,
	0x18, 0x86, 0x2d, 0x8c, 0x0d, 0xbf, 0x99, 0xa5, 0x69, 0x38, 0x70, 0x5c, 0x38, 0x11, 0x0f, 0x1c,
	0xa6, 0x37, 0x4f, 0x85, 0x11, 0x6d, 0xd8, 0x4a, 0x53, 0xba, 0x18, 0x4f, 0x04, 0x5d, 0x33, 0x49,
	0x0c, 0x10, 0x44, 0xe3, 0xfe, 0xb4, 0xbf, 0xc1, 0x80, 0x73, 0x2e, 0xbb, 0xf5, 0xed, 0xfb, 0xbc,
	0xcf, 0xe1, 0x83, 0x99, 0xfe, 0x7a, 0x79, 0x2d, 0xaa, 0x9d, 0x0e, 0x9a, 0xb6, 0xee, 0x6a, 0x62,
	0xff, 0x65, 0xef, 0x1b, 0x81, 0x95, 0xb6, 0x5b, 0xdd, 0x12, 0x1f, 0x46, 0xdd, 0xbe, 0xd1, 0x2e,
	0x9a, 0x23, 0x7f, 0xb6, 0x70, 0x82, 0xe3, 0x64, 0xa8, 0x03, 0xb5, 0x6f, 0xb4, 0x1c, 0x08, 0x72,
	0x0b, 0x76, 0x53, 0xbf, 0x97, 0x5d, 0x59, 0x57, 0xae, 0x31, 0xd0, 0xee, 0x39, 0x2d, 0x0e, 0xbd,
	0x3c, 0x92, 0xc4, 0x01, 0xeb, 0xb3, 0x78, 0xfb, 0xd0, 0xae, 0x39, 0x47, 0x3e, 0x92, 0xbf, 0xc1,
	0xbb, 0x83, 0x51, 0x6f, 0x26, 0x18, 0xae, 0xd4, 0x93, 0x88, 0xf3, 0x0d, 0x4f, 0x78, 0xfa, 0xc8,
	0xf1, 0x05, 0x01, 0x18, 0xaf, 0xa9, 0x4c, 0x62, 0x85, 0x11, 0xb9, 0x04, 0x6b, 0xc5, 0xd6, 0x4c,
	0x61, 0xa3, 0xff, 0x8e, 0x28, 0x8f, 0xe2, 0x15, 0x36, 0xbd, 0x05, 0xd8, 0xe2, 0x5f, 0x8f, 0x45,
	0x9a, 0x31, 0xc5, 0x52, 0x7e, 0x22, 0x99, 0x80, 0x19, 0xb2, 0x25, 0x46, 0xfd, 0x83, 0x66, 0x09,
	0x36, 0xae, 0x29, 0x4c, 0x55, 0x5b, 0x6c, 0xcb, 0x6a, 0x27, 0x8a, 0xb2, 0x25, 0x2e, 0x38, 0x4a,
	0xd2, 0x25, 0xe3, 0xf7, 0xb9, 0xa0, 0x4c, 0x9e, 0x4c, 0xa7, 0x30, 0x09, 0x55, 0x94, 0x6f, 0xb2,
	0x7e, 0x7e, 0x08, 0xea, 0x21, 0xc4, 0xc6, 0xf3, 0x78, 0x38, 0xe2, 0xcd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0xf7, 0x22, 0xca, 0x56, 0x01, 0x00, 0x00,
}