// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/spaceenter/exchange/matching_engine/matcher/proto/matcher.proto

/*
Package matcher is a generated protocol buffer package.

It is generated from these files:
	github.com/spaceenter/exchange/matching_engine/matcher/proto/matcher.proto

It has these top-level messages:
	Order
	TradeEvent
	OrderBookEvent
	OrderItem
	OrderTree
	OrderBook
*/
package matcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type OrderBookEvent_Type int32

const (
	OrderBookEvent_TYPE_UNKNOWN OrderBookEvent_Type = 0
	OrderBookEvent_ADD          OrderBookEvent_Type = 1
	OrderBookEvent_REMOVE       OrderBookEvent_Type = 2
)

var OrderBookEvent_Type_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "ADD",
	2: "REMOVE",
}
var OrderBookEvent_Type_value = map[string]int32{
	"TYPE_UNKNOWN": 0,
	"ADD":          1,
	"REMOVE":       2,
}

func (x OrderBookEvent_Type) String() string {
	return proto.EnumName(OrderBookEvent_Type_name, int32(x))
}
func (OrderBookEvent_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Order struct {
	OrderId   string                     `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=order_time,json=orderTime" json:"order_time,omitempty"`
	Type      Order_Type                 `protobuf:"varint,3,opt,name=type,enum=matcher.Order_Type" json:"type,omitempty"`
	IsSell    bool                       `protobuf:"varint,4,opt,name=is_sell,json=isSell" json:"is_sell,omitempty"`
	Price     float64                    `protobuf:"fixed64,5,opt,name=price" json:"price,omitempty"`
	Volume    float64                    `protobuf:"fixed64,6,opt,name=volume" json:"volume,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Order) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Order) GetOrderTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.OrderTime
	}
	return nil
}

func (m *Order) GetType() Order_Type {
	if m != nil {
		return m.Type
	}
	return Order_TYPE_UNKNOWN
}

func (m *Order) GetIsSell() bool {
	if m != nil {
		return m.IsSell
	}
	return false
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

type TradeEvent struct {
	OrderId       string                     `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	Timestamp     *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	IsTaker       bool                       `protobuf:"varint,3,opt,name=is_taker,json=isTaker" json:"is_taker,omitempty"`
	IsSell        bool                       `protobuf:"varint,4,opt,name=is_sell,json=isSell" json:"is_sell,omitempty"`
	Price         float64                    `protobuf:"fixed64,5,opt,name=price" json:"price,omitempty"`
	MatchedVolume float64                    `protobuf:"fixed64,6,opt,name=matched_volume,json=matchedVolume" json:"matched_volume,omitempty"`
	LeftVolume    float64                    `protobuf:"fixed64,7,opt,name=left_volume,json=leftVolume" json:"left_volume,omitempty"`
}

func (m *TradeEvent) Reset()                    { *m = TradeEvent{} }
func (m *TradeEvent) String() string            { return proto.CompactTextString(m) }
func (*TradeEvent) ProtoMessage()               {}
func (*TradeEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TradeEvent) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *TradeEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TradeEvent) GetIsTaker() bool {
	if m != nil {
		return m.IsTaker
	}
	return false
}

func (m *TradeEvent) GetIsSell() bool {
	if m != nil {
		return m.IsSell
	}
	return false
}

func (m *TradeEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TradeEvent) GetMatchedVolume() float64 {
	if m != nil {
		return m.MatchedVolume
	}
	return 0
}

func (m *TradeEvent) GetLeftVolume() float64 {
	if m != nil {
		return m.LeftVolume
	}
	return 0
}

type OrderBookEvent struct {
	OrderId   string                     `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Type      OrderBookEvent_Type        `protobuf:"varint,3,opt,name=type,enum=matcher.OrderBookEvent_Type" json:"type,omitempty"`
	IsSell    bool                       `protobuf:"varint,4,opt,name=is_sell,json=isSell" json:"is_sell,omitempty"`
	Price     float64                    `protobuf:"fixed64,5,opt,name=price" json:"price,omitempty"`
	Volume    float64                    `protobuf:"fixed64,6,opt,name=volume" json:"volume,omitempty"`
}

func (m *OrderBookEvent) Reset()                    { *m = OrderBookEvent{} }
func (m *OrderBookEvent) String() string            { return proto.CompactTextString(m) }
func (*OrderBookEvent) ProtoMessage()               {}
func (*OrderBookEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OrderBookEvent) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderBookEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OrderBookEvent) GetType() OrderBookEvent_Type {
	if m != nil {
		return m.Type
	}
	return OrderBookEvent_TYPE_UNKNOWN
}

func (m *OrderBookEvent) GetIsSell() bool {
	if m != nil {
		return m.IsSell
	}
	return false
}

func (m *OrderBookEvent) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderBookEvent) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type OrderItem struct {
	OrderId   string                     `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	OrderTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=order_time,json=orderTime" json:"order_time,omitempty"`
	Price     float64                    `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
	Volume    float64                    `protobuf:"fixed64,4,opt,name=volume" json:"volume,omitempty"`
}

func (m *OrderItem) Reset()                    { *m = OrderItem{} }
func (m *OrderItem) String() string            { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()               {}
func (*OrderItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OrderItem) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderItem) GetOrderTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.OrderTime
	}
	return nil
}

func (m *OrderItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderItem) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type OrderTree struct {
	IsSell bool         `protobuf:"varint,1,opt,name=is_sell,json=isSell" json:"is_sell,omitempty"`
	Items  []*OrderItem `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *OrderTree) Reset()                    { *m = OrderTree{} }
func (m *OrderTree) String() string            { return proto.CompactTextString(m) }
func (*OrderTree) ProtoMessage()               {}
func (*OrderTree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OrderTree) GetIsSell() bool {
	if m != nil {
		return m.IsSell
	}
	return false
}

func (m *OrderTree) GetItems() []*OrderItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type OrderBook struct {
	Pair         TradingPair                `protobuf:"varint,1,opt,name=pair,enum=matcher.TradingPair" json:"pair,omitempty"`
	SnapshotTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=snapshot_time,json=snapshotTime" json:"snapshot_time,omitempty"`
	Trees        []*OrderTree               `protobuf:"bytes,3,rep,name=trees" json:"trees,omitempty"`
}

func (m *OrderBook) Reset()                    { *m = OrderBook{} }
func (m *OrderBook) String() string            { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()               {}
func (*OrderBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *OrderBook) GetPair() TradingPair {
	if m != nil {
		return m.Pair
	}
	return TradingPair_TRADING_PAIR_UNKNOWN
}

func (m *OrderBook) GetSnapshotTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.SnapshotTime
	}
	return nil
}

func (m *OrderBook) GetTrees() []*OrderTree {
	if m != nil {
		return m.Trees
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "matcher.Order")
	proto.RegisterType((*TradeEvent)(nil), "matcher.TradeEvent")
	proto.RegisterType((*OrderBookEvent)(nil), "matcher.OrderBookEvent")
	proto.RegisterType((*OrderItem)(nil), "matcher.OrderItem")
	proto.RegisterType((*OrderTree)(nil), "matcher.OrderTree")
	proto.RegisterType((*OrderBook)(nil), "matcher.OrderBook")
	proto.RegisterEnum("matcher.TradingPair", TradingPair_name, TradingPair_value)
	proto.RegisterEnum("matcher.Order_Type", Order_Type_name, Order_Type_value)
	proto.RegisterEnum("matcher.OrderBookEvent_Type", OrderBookEvent_Type_name, OrderBookEvent_Type_value)
}

func init() {
	proto.RegisterFile("github.com/spaceenter/exchange/matching_engine/matcher/proto/matcher.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0xfd, 0x56, 0xfe, 0x51, 0x3c, 0x4e, 0x8c, 0xd8, 0x2f, 0xb4, 0x6a, 0x28, 0xc4, 0x18, 0x4a,
	0x45, 0xa1, 0x72, 0x71, 0x6f, 0x5a, 0x7a, 0x51, 0x9c, 0xc4, 0xb4, 0x6e, 0x12, 0x27, 0x6c, 0x94,
	0x94, 0x5e, 0x09, 0xc5, 0x9e, 0xc8, 0x4b, 0xf4, 0xc7, 0x6a, 0x13, 0x9a, 0x57, 0x28, 0x94, 0xbe,
	0x42, 0xdf, 0xb1, 0x2f, 0x50, 0x76, 0x25, 0x99, 0x08, 0xd2, 0x92, 0x5e, 0xe4, 0x4e, 0x67, 0x74,
	0xf6, 0xec, 0x9c, 0x33, 0xb3, 0xf0, 0x29, 0xe4, 0x72, 0x79, 0x75, 0xee, 0xce, 0xd3, 0x78, 0x98,
	0x67, 0xc1, 0x1c, 0x31, 0x91, 0x28, 0x86, 0xf8, 0x75, 0xbe, 0x0c, 0x92, 0x10, 0x87, 0x71, 0x20,
	0xe7, 0x4b, 0x9e, 0x84, 0x3e, 0x26, 0x21, 0x4f, 0x4a, 0x8c, 0x62, 0x98, 0x89, 0x54, 0xa6, 0x15,
	0x72, 0x35, 0xa2, 0x66, 0x09, 0xb7, 0xb6, 0xc3, 0x34, 0x0d, 0x23, 0x2c, 0x48, 0xe7, 0x57, 0x17,
	0x43, 0xc9, 0x63, 0xcc, 0x65, 0x10, 0x67, 0x05, 0x73, 0xf0, 0xcd, 0x80, 0xd6, 0x91, 0x58, 0xa0,
	0xa0, 0x4f, 0x60, 0x2d, 0x55, 0x1f, 0x3e, 0x5f, 0xd8, 0xa4, 0x4f, 0x9c, 0x0e, 0x33, 0x35, 0x9e,
	0x2e, 0xe8, 0x5b, 0x80, 0xe2, 0x97, 0x3a, 0x6d, 0x1b, 0x7d, 0xe2, 0x74, 0x47, 0x5b, 0x6e, 0x21,
	0xed, 0x56, 0xd2, 0xae, 0x57, 0x49, 0xb3, 0x8e, 0x66, 0x2b, 0x4c, 0x9f, 0x43, 0x53, 0xde, 0x64,
	0x68, 0x37, 0xfa, 0xc4, 0xe9, 0x8d, 0xfe, 0x77, 0xab, 0x3e, 0xf5, 0x9d, 0xae, 0x77, 0x93, 0x21,
	0xd3, 0x04, 0xfa, 0x18, 0x4c, 0x9e, 0xfb, 0x39, 0x46, 0x91, 0xdd, 0xec, 0x13, 0x67, 0x8d, 0xb5,
	0x79, 0x7e, 0x82, 0x51, 0x44, 0x37, 0xa1, 0x95, 0x09, 0x3e, 0x47, 0xbb, 0xd5, 0x27, 0x0e, 0x61,
	0x05, 0xa0, 0x8f, 0xa0, 0x7d, 0x9d, 0x46, 0x57, 0x31, 0xda, 0x6d, 0x5d, 0x2e, 0xd1, 0xe0, 0x1d,
	0x34, 0x95, 0x28, 0xb5, 0x60, 0xdd, 0xfb, 0x72, 0x3c, 0xf1, 0x4f, 0x67, 0xfb, 0xb3, 0xa3, 0xcf,
	0x33, 0xeb, 0x3f, 0x0a, 0xd0, 0x3e, 0x1c, 0xb3, 0xfd, 0x89, 0x67, 0x11, 0xda, 0x81, 0xd6, 0xc1,
	0xf4, 0x70, 0xea, 0x59, 0x86, 0x2a, 0xef, 0x8e, 0x67, 0xbb, 0x93, 0x03, 0xab, 0x31, 0xf8, 0x45,
	0x00, 0x3c, 0x11, 0x2c, 0x70, 0x72, 0x8d, 0x89, 0xfc, 0x5b, 0x22, 0x6f, 0xa0, 0xb3, 0x4a, 0xf2,
	0x3e, 0x81, 0xac, 0xc8, 0x4a, 0x94, 0xe7, 0xbe, 0x0c, 0x2e, 0x51, 0xe8, 0x50, 0xd6, 0x98, 0xc9,
	0x73, 0x4f, 0xc1, 0x7f, 0x8d, 0xe0, 0x19, 0xf4, 0x8a, 0x34, 0x17, 0x7e, 0x2d, 0x8a, 0x8d, 0xb2,
	0x7a, 0xa6, 0x8b, 0x74, 0x1b, 0xba, 0x11, 0x5e, 0xc8, 0x8a, 0x63, 0x6a, 0x0e, 0xa8, 0x52, 0x41,
	0x18, 0x7c, 0x37, 0xa0, 0xa7, 0xc7, 0xb1, 0x93, 0xa6, 0x97, 0x0f, 0xe8, 0xfc, 0x55, 0x6d, 0x15,
	0x9e, 0xd6, 0x57, 0x61, 0x75, 0xf7, 0x03, 0xec, 0xc4, 0xcb, 0x3f, 0xee, 0x84, 0x09, 0x8d, 0xf1,
	0xde, 0x9e, 0x45, 0xd4, 0x16, 0xb0, 0xc9, 0xe1, 0xd1, 0xd9, 0xc4, 0x32, 0x06, 0x3f, 0x08, 0x74,
	0x74, 0x4f, 0x53, 0x89, 0xf1, 0x03, 0x3d, 0x8b, 0x95, 0x81, 0xc6, 0xdd, 0x06, 0x9a, 0x35, 0x03,
	0xb3, 0xb2, 0x21, 0x4f, 0x60, 0x2d, 0x14, 0x52, 0x0b, 0xc5, 0x81, 0x16, 0x97, 0x18, 0xe7, 0xb6,
	0xd1, 0x6f, 0x38, 0xdd, 0x11, 0xad, 0x07, 0xac, 0xcc, 0xb0, 0x82, 0x30, 0xf8, 0x59, 0x39, 0x54,
	0xa9, 0x53, 0x07, 0x9a, 0x59, 0xc0, 0x85, 0x56, 0xeb, 0x8d, 0x36, 0x57, 0xc7, 0xd4, 0x4b, 0xe0,
	0x49, 0x78, 0x1c, 0x70, 0xc1, 0x34, 0x83, 0xbe, 0x87, 0x8d, 0x3c, 0x09, 0xb2, 0x7c, 0x99, 0xca,
	0xfb, 0x7a, 0x5e, 0xaf, 0x0e, 0x68, 0xdb, 0x0e, 0xb4, 0xa4, 0x40, 0xcc, 0xed, 0xc6, 0x5d, 0x2d,
	0x2a, 0x7b, 0xac, 0x20, 0xbc, 0x18, 0x43, 0xf7, 0xd6, 0xfd, 0xd4, 0x86, 0x4d, 0x8f, 0x8d, 0xf7,
	0xa6, 0xb3, 0x0f, 0xfe, 0xf1, 0x78, 0xca, 0x6e, 0x8d, 0xb0, 0x0b, 0xe6, 0x8e, 0xb7, 0xeb, 0x9f,
	0x9e, 0xa8, 0x31, 0x96, 0xc0, 0xfb, 0xb8, 0x63, 0x19, 0xe7, 0x6d, 0xdd, 0xce, 0xeb, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x08, 0x95, 0x56, 0x3c, 0x59, 0x05, 0x00, 0x00,
}