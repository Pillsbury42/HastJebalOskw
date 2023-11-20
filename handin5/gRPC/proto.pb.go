// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: handin5/gRPC/proto.proto

package gRPC

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidderID int64 `protobuf:"varint,1,opt,name=bidderID,proto3" json:"bidderID,omitempty"`
	Amount   int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidMessage) Reset() {
	*x = BidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidMessage) ProtoMessage() {}

func (x *BidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidMessage.ProtoReflect.Descriptor instead.
func (*BidMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{0}
}

func (x *BidMessage) GetBidderID() int64 {
	if x != nil {
		return x.BidderID
	}
	return 0
}

func (x *BidMessage) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BidReplyMessage) Reset() {
	*x = BidReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidReplyMessage) ProtoMessage() {}

func (x *BidReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidReplyMessage.ProtoReflect.Descriptor instead.
func (*BidReplyMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{1}
}

func (x *BidReplyMessage) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ResultReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Over     bool  `protobuf:"varint,1,opt,name=over,proto3" json:"over,omitempty"`         //is auction over?
	WinnerID int64 `protobuf:"varint,2,opt,name=winnerID,proto3" json:"winnerID,omitempty"` //id of the winner
	Highest  int64 `protobuf:"varint,3,opt,name=highest,proto3" json:"highest,omitempty"`   //the highest bid
}

func (x *ResultReplyMessage) Reset() {
	*x = ResultReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReplyMessage) ProtoMessage() {}

func (x *ResultReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReplyMessage.ProtoReflect.Descriptor instead.
func (*ResultReplyMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{2}
}

func (x *ResultReplyMessage) GetOver() bool {
	if x != nil {
		return x.Over
	}
	return false
}

func (x *ResultReplyMessage) GetWinnerID() int64 {
	if x != nil {
		return x.WinnerID
	}
	return 0
}

func (x *ResultReplyMessage) GetHighest() int64 {
	if x != nil {
		return x.Highest
	}
	return 0
}

type ElectionReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyID int64 `protobuf:"varint,1,opt,name=replyID,proto3" json:"replyID,omitempty"`
}

func (x *ElectionReplyMessage) Reset() {
	*x = ElectionReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionReplyMessage) ProtoMessage() {}

func (x *ElectionReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionReplyMessage.ProtoReflect.Descriptor instead.
func (*ElectionReplyMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{3}
}

func (x *ElectionReplyMessage) GetReplyID() int64 {
	if x != nil {
		return x.ReplyID
	}
	return 0
}

type CoordinatorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoordID int64 `protobuf:"varint,1,opt,name=coordID,proto3" json:"coordID,omitempty"`
}

func (x *CoordinatorMessage) Reset() {
	*x = CoordinatorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinatorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinatorMessage) ProtoMessage() {}

func (x *CoordinatorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinatorMessage.ProtoReflect.Descriptor instead.
func (*CoordinatorMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{4}
}

func (x *CoordinatorMessage) GetCoordID() int64 {
	if x != nil {
		return x.CoordID
	}
	return 0
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handin5_gRPC_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_handin5_gRPC_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_handin5_gRPC_proto_proto_rawDescGZIP(), []int{5}
}

var File_handin5_gRPC_proto_proto protoreflect.FileDescriptor

var file_handin5_gRPC_proto_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x35, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x52, 0x50, 0x43,
	0x22, 0x40, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x42, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x5e, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x77, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x22,
	0x30, 0x0a, 0x14, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x49,
	0x44, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x49,
	0x44, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x9d, 0x02, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x42, 0x69, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x42, 0x69,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a,
	0x09, 0x62, 0x69, 0x64, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x52, 0x50,
	0x43, 0x2e, 0x42, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x67,
	0x52, 0x50, 0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x52, 0x50,
	0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18,
	0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e,
	0x67, 0x52, 0x50, 0x43, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x50, 0x69, 0x6c, 0x6c, 0x73, 0x62, 0x75, 0x72, 0x79, 0x34, 0x32, 0x2f, 0x48, 0x61, 0x73, 0x74,
	0x4a, 0x65, 0x62, 0x61, 0x6c, 0x4f, 0x73, 0x6b, 0x77, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x35, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handin5_gRPC_proto_proto_rawDescOnce sync.Once
	file_handin5_gRPC_proto_proto_rawDescData = file_handin5_gRPC_proto_proto_rawDesc
)

func file_handin5_gRPC_proto_proto_rawDescGZIP() []byte {
	file_handin5_gRPC_proto_proto_rawDescOnce.Do(func() {
		file_handin5_gRPC_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_handin5_gRPC_proto_proto_rawDescData)
	})
	return file_handin5_gRPC_proto_proto_rawDescData
}

var file_handin5_gRPC_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_handin5_gRPC_proto_proto_goTypes = []interface{}{
	(*BidMessage)(nil),           // 0: gRPC.BidMessage
	(*BidReplyMessage)(nil),      // 1: gRPC.BidReplyMessage
	(*ResultReplyMessage)(nil),   // 2: gRPC.ResultReplyMessage
	(*ElectionReplyMessage)(nil), // 3: gRPC.ElectionReplyMessage
	(*CoordinatorMessage)(nil),   // 4: gRPC.CoordinatorMessage
	(*EmptyMessage)(nil),         // 5: gRPC.EmptyMessage
}
var file_handin5_gRPC_proto_proto_depIdxs = []int32{
	0, // 0: gRPC.Auction.bid:input_type -> gRPC.BidMessage
	0, // 1: gRPC.Auction.bidupdate:input_type -> gRPC.BidMessage
	5, // 2: gRPC.Auction.result:input_type -> gRPC.EmptyMessage
	5, // 3: gRPC.Auction.election:input_type -> gRPC.EmptyMessage
	4, // 4: gRPC.Auction.coordinator:input_type -> gRPC.CoordinatorMessage
	1, // 5: gRPC.Auction.bid:output_type -> gRPC.BidReplyMessage
	5, // 6: gRPC.Auction.bidupdate:output_type -> gRPC.EmptyMessage
	2, // 7: gRPC.Auction.result:output_type -> gRPC.ResultReplyMessage
	3, // 8: gRPC.Auction.election:output_type -> gRPC.ElectionReplyMessage
	5, // 9: gRPC.Auction.coordinator:output_type -> gRPC.EmptyMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_handin5_gRPC_proto_proto_init() }
func file_handin5_gRPC_proto_proto_init() {
	if File_handin5_gRPC_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handin5_gRPC_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handin5_gRPC_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidReplyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handin5_gRPC_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReplyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handin5_gRPC_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionReplyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handin5_gRPC_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinatorMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handin5_gRPC_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_handin5_gRPC_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handin5_gRPC_proto_proto_goTypes,
		DependencyIndexes: file_handin5_gRPC_proto_proto_depIdxs,
		MessageInfos:      file_handin5_gRPC_proto_proto_msgTypes,
	}.Build()
	File_handin5_gRPC_proto_proto = out.File
	file_handin5_gRPC_proto_proto_rawDesc = nil
	file_handin5_gRPC_proto_proto_goTypes = nil
	file_handin5_gRPC_proto_proto_depIdxs = nil
}
