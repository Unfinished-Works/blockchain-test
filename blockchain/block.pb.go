// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: blockchain/block.proto

package blockchain

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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref        []byte `protobuf:"bytes,1,opt,name=Ref,proto3" json:"Ref,omitempty"`
	ID         uint64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	PrevHash   []byte `protobuf:"bytes,3,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Difficulty uint64 `protobuf:"varint,4,opt,name=Difficulty,proto3" json:"Difficulty,omitempty"`
	Timestamp  uint64 `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_blockchain_block_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetRef() []byte {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *Header) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Header) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *Header) GetDifficulty() uint64 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *Header) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce      uint64         `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Header     *Header        `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	HeaderHash []byte         `protobuf:"bytes,3,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	Tx         []*Transaction `protobuf:"bytes,5,rep,name=tx,proto3" json:"tx,omitempty"`
	Data       []byte         `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blockchain_block_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Block) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetHeaderHash() []byte {
	if x != nil {
		return x.HeaderHash
	}
	return nil
}

func (x *Block) GetTx() []*Transaction {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *Block) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNonce uint64 `protobuf:"varint,1,opt,name=AccountNonce,proto3" json:"AccountNonce,omitempty"`
	Amount       uint64 `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Fee          uint64 `protobuf:"varint,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Recipient    string `protobuf:"bytes,4,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Payload      []byte `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	PublicKey    string `protobuf:"bytes,6,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	R            string `protobuf:"bytes,7,opt,name=R,proto3" json:"R,omitempty"`
	S            string `protobuf:"bytes,8,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_blockchain_block_proto_rawDescGZIP(), []int{2}
}

func (x *Transaction) GetAccountNonce() uint64 {
	if x != nil {
		return x.AccountNonce
	}
	return 0
}

func (x *Transaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Transaction) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Transaction) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Transaction) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Transaction) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *Transaction) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

var File_blockchain_block_proto protoreflect.FileDescriptor

var file_blockchain_block_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x52, 0x65,
	0x66, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x50, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a,
	0x0a, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa6, 0x01, 0x0a, 0x05,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x74, 0x78,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x46,
	0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x52, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x52, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_block_proto_rawDescOnce sync.Once
	file_blockchain_block_proto_rawDescData = file_blockchain_block_proto_rawDesc
)

func file_blockchain_block_proto_rawDescGZIP() []byte {
	file_blockchain_block_proto_rawDescOnce.Do(func() {
		file_blockchain_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_block_proto_rawDescData)
	})
	return file_blockchain_block_proto_rawDescData
}

var file_blockchain_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchain_block_proto_goTypes = []interface{}{
	(*Header)(nil),      // 0: blockchain.Header
	(*Block)(nil),       // 1: blockchain.Block
	(*Transaction)(nil), // 2: blockchain.Transaction
}
var file_blockchain_block_proto_depIdxs = []int32{
	0, // 0: blockchain.Block.Header:type_name -> blockchain.Header
	2, // 1: blockchain.Block.tx:type_name -> blockchain.Transaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blockchain_block_proto_init() }
func file_blockchain_block_proto_init() {
	if File_blockchain_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_blockchain_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_blockchain_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_blockchain_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_block_proto_goTypes,
		DependencyIndexes: file_blockchain_block_proto_depIdxs,
		MessageInfos:      file_blockchain_block_proto_msgTypes,
	}.Build()
	File_blockchain_block_proto = out.File
	file_blockchain_block_proto_rawDesc = nil
	file_blockchain_block_proto_goTypes = nil
	file_blockchain_block_proto_depIdxs = nil
}
