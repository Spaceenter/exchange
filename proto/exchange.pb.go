// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	OrderInfo
	Order
	MatchedOrder
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return proto1.EnumName(TradingPair_name, int32(x))
}
func (TradingPair) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OrderType int32

const (
	OrderType_TYPE_UNKNOWN OrderType = 0
	OrderType_MARKET       OrderType = 1
	OrderType_LIMIT        OrderType = 2
	OrderType_CANCEL       OrderType = 3
)

var OrderType_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "MARKET",
	2: "LIMIT",
	3: "CANCEL",
}
var OrderType_value = map[string]int32{
	"TYPE_UNKNOWN": 0,
	"MARKET":       1,
	"LIMIT":        2,
	"CANCEL":       3,
}

func (x OrderType) String() string {
	return proto1.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OrderInfo struct {
	OrderId   string                     `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Type      OrderType                  `protobuf:"varint,4,opt,name=type,enum=proto.OrderType" json:"type,omitempty"`
	IsAsk     bool                       `protobuf:"varint,5,opt,name=is_ask,json=isAsk" json:"is_ask,omitempty"`
}

func (m *OrderInfo) Reset()                    { *m = OrderInfo{} }
func (m *OrderInfo) String() string            { return proto1.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()               {}
func (*OrderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OrderInfo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderInfo) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OrderInfo) GetType() OrderType {
	if m != nil {
		return m.Type
	}
	return OrderType_TYPE_UNKNOWN
}

func (m *OrderInfo) GetIsAsk() bool {
	if m != nil {
		return m.IsAsk
	}
	return false
}

type Order struct {
	OrderInfo *OrderInfo `protobuf:"bytes,1,opt,name=order_info,json=orderInfo" json:"order_info,omitempty"`
	Price     float64    `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	Volume    float64    `protobuf:"fixed64,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto1.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Order) GetOrderInfo() *OrderInfo {
	if m != nil {
		return m.OrderInfo
	}
	return nil
}

func (m *Order) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type MatchedOrder struct {
	OrderInfo     *OrderInfo `protobuf:"bytes,1,opt,name=order_info,json=orderInfo" json:"order_info,omitempty"`
	MatchedPrice  float64    `protobuf:"fixed64,2,opt,name=matched_price,json=matchedPrice" json:"matched_price,omitempty"`
	MatchedVolume float64    `protobuf:"fixed64,3,opt,name=matched_volume,json=matchedVolume" json:"matched_volume,omitempty"`
	LeftVolume    float64    `protobuf:"fixed64,4,opt,name=left_volume,json=leftVolume" json:"left_volume,omitempty"`
}

func (m *MatchedOrder) Reset()                    { *m = MatchedOrder{} }
func (m *MatchedOrder) String() string            { return proto1.CompactTextString(m) }
func (*MatchedOrder) ProtoMessage()               {}
func (*MatchedOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MatchedOrder) GetOrderInfo() *OrderInfo {
	if m != nil {
		return m.OrderInfo
	}
	return nil
}

func (m *MatchedOrder) GetMatchedPrice() float64 {
	if m != nil {
		return m.MatchedPrice
	}
	return 0
}

func (m *MatchedOrder) GetMatchedVolume() float64 {
	if m != nil {
		return m.MatchedVolume
	}
	return 0
}

func (m *MatchedOrder) GetLeftVolume() float64 {
	if m != nil {
		return m.LeftVolume
	}
	return 0
}

func init() {
	proto1.RegisterType((*OrderInfo)(nil), "proto.OrderInfo")
	proto1.RegisterType((*Order)(nil), "proto.Order")
	proto1.RegisterType((*MatchedOrder)(nil), "proto.MatchedOrder")
	proto1.RegisterEnum("proto.TradingPair", TradingPair_name, TradingPair_value)
	proto1.RegisterEnum("proto.OrderType", OrderType_name, OrderType_value)
}

func init() { proto1.RegisterFile("exchange.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x6f, 0x9b, 0x30,
	0x1c, 0xc5, 0xe7, 0x34, 0xd0, 0xf2, 0x27, 0x8b, 0x2c, 0xab, 0x9b, 0x58, 0x2f, 0x45, 0xd9, 0x26,
	0xa1, 0x1e, 0x88, 0xd4, 0x5d, 0x76, 0x1c, 0x4d, 0xa3, 0x0d, 0xb5, 0xa1, 0xc8, 0x73, 0x37, 0xed,
	0x84, 0x48, 0x30, 0x04, 0x25, 0x60, 0x04, 0x64, 0x5a, 0x3e, 0xcc, 0xbe, 0xc1, 0x3e, 0xe4, 0x84,
	0x81, 0x2c, 0xb9, 0xee, 0x84, 0xde, 0xe3, 0xf7, 0xff, 0xbf, 0x67, 0x1b, 0xc6, 0xfc, 0xd7, 0x6a,
	0x1d, 0xe6, 0x09, 0xb7, 0x8b, 0x52, 0xd4, 0x82, 0x28, 0xf2, 0x73, 0x75, 0x9d, 0x08, 0x91, 0x6c,
	0xf9, 0x54, 0xaa, 0xe5, 0x2e, 0x9e, 0xd6, 0x69, 0xc6, 0xab, 0x3a, 0xcc, 0x8a, 0x96, 0x9b, 0xfc,
	0x46, 0xa0, 0x3d, 0x95, 0x11, 0x2f, 0xdd, 0x3c, 0x16, 0xe4, 0x0d, 0x5c, 0x88, 0x46, 0x04, 0x69,
	0x64, 0x20, 0x13, 0x59, 0x1a, 0x3d, 0x97, 0xda, 0x8d, 0xc8, 0x47, 0xd0, 0x0e, 0xb3, 0xc6, 0xc0,
	0x44, 0x96, 0x7e, 0x7b, 0x65, 0xb7, 0xdb, 0xed, 0x7e, 0xbb, 0xcd, 0x7a, 0x82, 0xfe, 0x83, 0xc9,
	0x3b, 0x18, 0xd6, 0xfb, 0x82, 0x1b, 0x43, 0x13, 0x59, 0xe3, 0x5b, 0xdc, 0xd2, 0xb6, 0x0c, 0x65,
	0xfb, 0x82, 0x53, 0xf9, 0x97, 0xbc, 0x02, 0x35, 0xad, 0x82, 0xb0, 0xda, 0x18, 0x8a, 0x89, 0xac,
	0x0b, 0xaa, 0xa4, 0x95, 0x53, 0x6d, 0x26, 0x31, 0x28, 0x92, 0x24, 0x53, 0x80, 0xae, 0x5a, 0x1e,
	0x0b, 0x59, 0x4e, 0x3f, 0xdd, 0xd5, 0x1c, 0x80, 0x6a, 0xe2, 0x70, 0x96, 0x4b, 0x50, 0x8a, 0x32,
	0x5d, 0x71, 0x59, 0x16, 0xd1, 0x56, 0x90, 0xd7, 0xa0, 0xfe, 0x14, 0xdb, 0x5d, 0xc6, 0x8d, 0x33,
	0x69, 0x77, 0x6a, 0xf2, 0x07, 0xc1, 0x68, 0x11, 0xd6, 0xab, 0x35, 0x8f, 0xfe, 0x33, 0xef, 0x2d,
	0xbc, 0xcc, 0xda, 0x05, 0xc1, 0x71, 0xee, 0xa8, 0x33, 0x7d, 0x19, 0xff, 0x1e, 0xc6, 0x3d, 0x74,
	0x52, 0xa3, 0x1f, 0xfd, 0x26, 0x4d, 0x72, 0x0d, 0xfa, 0x96, 0xc7, 0x75, 0xcf, 0x0c, 0x25, 0x03,
	0x8d, 0xd5, 0x02, 0x37, 0x0e, 0xe8, 0xac, 0x0c, 0xa3, 0x34, 0x4f, 0xfc, 0x30, 0x2d, 0x89, 0x01,
	0x97, 0x8c, 0x3a, 0xf7, 0xae, 0xf7, 0x39, 0xf0, 0x1d, 0x97, 0x06, 0xcf, 0xde, 0x83, 0xf7, 0xf4,
	0xdd, 0xc3, 0x2f, 0x88, 0x0e, 0xe7, 0x77, 0x6c, 0x16, 0x3c, 0x7f, 0xbd, 0xc7, 0xa8, 0x17, 0xec,
	0xcb, 0x1d, 0x1e, 0xdc, 0x7c, 0xea, 0x1e, 0xbe, 0x79, 0x03, 0x82, 0x61, 0xc4, 0x7e, 0xf8, 0xf3,
	0xa3, 0x41, 0x00, 0x75, 0xe1, 0xd0, 0x87, 0x39, 0xc3, 0x88, 0x68, 0xa0, 0x3c, 0xba, 0x0b, 0x97,
	0xe1, 0x41, 0x63, 0xcf, 0x1c, 0x6f, 0x36, 0x7f, 0xc4, 0x67, 0x4b, 0x55, 0xde, 0xc6, 0x87, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x71, 0x2c, 0x06, 0x7c, 0x02, 0x00, 0x00,
}
