// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: network.proto

package types

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

// StorageItem is the item for the storage data
type StorageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skey   string `protobuf:"bytes,1,opt,name=skey,proto3" json:"skey,omitempty"`
	Svalue string `protobuf:"bytes,2,opt,name=svalue,proto3" json:"svalue,omitempty"`
}

func (x *StorageItem) Reset() {
	*x = StorageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageItem) ProtoMessage() {}

func (x *StorageItem) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageItem.ProtoReflect.Descriptor instead.
func (*StorageItem) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0}
}

func (x *StorageItem) GetSkey() string {
	if x != nil {
		return x.Skey
	}
	return ""
}

func (x *StorageItem) GetSvalue() string {
	if x != nil {
		return x.Svalue
	}
	return ""
}

type StorageItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*StorageItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StorageItemList) Reset() {
	*x = StorageItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageItemList) ProtoMessage() {}

func (x *StorageItemList) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageItemList.ProtoReflect.Descriptor instead.
func (*StorageItemList) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{1}
}

func (x *StorageItemList) GetItems() []*StorageItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// DeltaItem is the item for the state change
type DeltaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*DeltaItem_StringValue
	//	*DeltaItem_StorageValue
	Value isDeltaItem_Value `protobuf_oneof:"value"`
}

func (x *DeltaItem) Reset() {
	*x = DeltaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaItem) ProtoMessage() {}

func (x *DeltaItem) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaItem.ProtoReflect.Descriptor instead.
func (*DeltaItem) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DeltaItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *DeltaItem) GetValue() isDeltaItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *DeltaItem) GetStringValue() string {
	if x, ok := x.GetValue().(*DeltaItem_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *DeltaItem) GetStorageValue() *StorageItemList {
	if x, ok := x.GetValue().(*DeltaItem_StorageValue); ok {
		return x.StorageValue
	}
	return nil
}

type isDeltaItem_Value interface {
	isDeltaItem_Value()
}

type DeltaItem_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type DeltaItem_StorageValue struct {
	StorageValue *StorageItemList `protobuf:"bytes,4,opt,name=storage_value,json=storageValue,proto3,oneof"`
}

func (*DeltaItem_StringValue) isDeltaItem_Value() {}

func (*DeltaItem_StorageValue) isDeltaItem_Value() {}

// BlockDelta is the delta data for the given block of the specific chain
type BlockDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64       `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	StateRoot   string       `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Chain       string       `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Delta       []*DeltaItem `protobuf:"bytes,4,rep,name=delta,proto3" json:"delta,omitempty"`
	DeltaHash   string       `protobuf:"bytes,5,opt,name=delta_hash,json=deltaHash,proto3" json:"delta_hash,omitempty"`
}

func (x *BlockDelta) Reset() {
	*x = BlockDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockDelta) ProtoMessage() {}

func (x *BlockDelta) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockDelta.ProtoReflect.Descriptor instead.
func (*BlockDelta) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{3}
}

func (x *BlockDelta) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockDelta) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *BlockDelta) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *BlockDelta) GetDelta() []*DeltaItem {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *BlockDelta) GetDeltaHash() string {
	if x != nil {
		return x.DeltaHash
	}
	return ""
}

// BlockHeader is the block header structure
type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber       uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	ParentHash        string `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockHash         string `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	ProposerPubKey    string `protobuf:"bytes,4,opt,name=proposer_pub_key,json=proposerPubKey,proto3" json:"proposer_pub_key,omitempty"`
	ProposerSignature string `protobuf:"bytes,5,opt,name=proposer_signature,json=proposerSignature,proto3" json:"proposer_signature,omitempty"` // (block_number, parent_hash, delta_hash, proof)
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{4}
}

func (x *BlockHeader) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockHeader) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockHeader) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *BlockHeader) GetProposerPubKey() string {
	if x != nil {
		return x.ProposerPubKey
	}
	return ""
}

func (x *BlockHeader) GetProposerSignature() string {
	if x != nil {
		return x.ProposerSignature
	}
	return ""
}

// Block is the block body structure for zk proofs of the block delta
type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delta     *BlockDelta  `protobuf:"bytes,2,opt,name=delta,proto3" json:"delta,omitempty"`
	Proof     string       `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	Signature string       `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"` // aggregated signature of all the validators
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[5]
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
	return file_network_proto_rawDescGZIP(), []int{5}
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetDelta() *BlockDelta {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *Block) GetProof() string {
	if x != nil {
		return x.Proof
	}
	return ""
}

func (x *Block) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// JoinNetworkRequest is the request to join the network
type JoinNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey    string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	StakeAddress string `protobuf:"bytes,2,opt,name=stake_address,json=stakeAddress,proto3" json:"stake_address,omitempty"`
	Signature    string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *JoinNetworkRequest) Reset() {
	*x = JoinNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkRequest) ProtoMessage() {}

func (x *JoinNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkRequest.ProtoReflect.Descriptor instead.
func (*JoinNetworkRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{6}
}

func (x *JoinNetworkRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *JoinNetworkRequest) GetStakeAddress() string {
	if x != nil {
		return x.StakeAddress
	}
	return ""
}

func (x *JoinNetworkRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// JoinNetworkResponse is the response for joining the network
type JoinNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JoinNetworkResponse) Reset() {
	*x = JoinNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinNetworkResponse) ProtoMessage() {}

func (x *JoinNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinNetworkResponse.ProtoReflect.Descriptor instead.
func (*JoinNetworkResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{7}
}

func (x *JoinNetworkResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *JoinNetworkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// GetBlockRequest is the request to get the given block
type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlockRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

// GetBlockResponse is the response for getting the given block
type GetBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetBlockResponse) Reset() {
	*x = GetBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockResponse) ProtoMessage() {}

func (x *GetBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockResponse.ProtoReflect.Descriptor instead.
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{9}
}

func (x *GetBlockResponse) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

// CommitBlockRequest is the request to commit the signature
type CommitBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Signature   string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *CommitBlockRequest) Reset() {
	*x = CommitBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitBlockRequest) ProtoMessage() {}

func (x *CommitBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitBlockRequest.ProtoReflect.Descriptor instead.
func (*CommitBlockRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{10}
}

func (x *CommitBlockRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *CommitBlockRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

// CommitBlockResponse is the response for uploading the signature
type CommitBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommitBlockResponse) Reset() {
	*x = CommitBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitBlockResponse) ProtoMessage() {}

func (x *CommitBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitBlockResponse.ProtoReflect.Descriptor instead.
func (*CommitBlockResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{11}
}

func (x *CommitBlockResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CommitBlockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_network_proto protoreflect.FileDescriptor

var file_network_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x65, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x75,
	0x62, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2f, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x76, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x55, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xf7, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x67,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x4c, 0x61, 0x67, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x2d, 0x4e, 0x6f, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_proto_rawDescOnce sync.Once
	file_network_proto_rawDescData = file_network_proto_rawDesc
)

func file_network_proto_rawDescGZIP() []byte {
	file_network_proto_rawDescOnce.Do(func() {
		file_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_proto_rawDescData)
	})
	return file_network_proto_rawDescData
}

var file_network_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_network_proto_goTypes = []interface{}{
	(*StorageItem)(nil),         // 0: network.v1.StorageItem
	(*StorageItemList)(nil),     // 1: network.v1.StorageItemList
	(*DeltaItem)(nil),           // 2: network.v1.DeltaItem
	(*BlockDelta)(nil),          // 3: network.v1.BlockDelta
	(*BlockHeader)(nil),         // 4: network.v1.BlockHeader
	(*Block)(nil),               // 5: network.v1.Block
	(*JoinNetworkRequest)(nil),  // 6: network.v1.JoinNetworkRequest
	(*JoinNetworkResponse)(nil), // 7: network.v1.JoinNetworkResponse
	(*GetBlockRequest)(nil),     // 8: network.v1.GetBlockRequest
	(*GetBlockResponse)(nil),    // 9: network.v1.GetBlockResponse
	(*CommitBlockRequest)(nil),  // 10: network.v1.CommitBlockRequest
	(*CommitBlockResponse)(nil), // 11: network.v1.CommitBlockResponse
}
var file_network_proto_depIdxs = []int32{
	0,  // 0: network.v1.StorageItemList.items:type_name -> network.v1.StorageItem
	1,  // 1: network.v1.DeltaItem.storage_value:type_name -> network.v1.StorageItemList
	2,  // 2: network.v1.BlockDelta.delta:type_name -> network.v1.DeltaItem
	4,  // 3: network.v1.Block.header:type_name -> network.v1.BlockHeader
	3,  // 4: network.v1.Block.delta:type_name -> network.v1.BlockDelta
	5,  // 5: network.v1.GetBlockResponse.block:type_name -> network.v1.Block
	6,  // 6: network.v1.NetworkService.JoinNetwork:input_type -> network.v1.JoinNetworkRequest
	8,  // 7: network.v1.NetworkService.GetBlock:input_type -> network.v1.GetBlockRequest
	10, // 8: network.v1.NetworkService.CommitBlock:input_type -> network.v1.CommitBlockRequest
	7,  // 9: network.v1.NetworkService.JoinNetwork:output_type -> network.v1.JoinNetworkResponse
	9,  // 10: network.v1.NetworkService.GetBlock:output_type -> network.v1.GetBlockResponse
	11, // 11: network.v1.NetworkService.CommitBlock:output_type -> network.v1.CommitBlockResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_network_proto_init() }
func file_network_proto_init() {
	if File_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageItem); i {
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
		file_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageItemList); i {
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
		file_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaItem); i {
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
		file_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockDelta); i {
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
		file_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkRequest); i {
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
		file_network_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinNetworkResponse); i {
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
		file_network_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
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
		file_network_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockResponse); i {
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
		file_network_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitBlockRequest); i {
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
		file_network_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitBlockResponse); i {
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
	file_network_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DeltaItem_StringValue)(nil),
		(*DeltaItem_StorageValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_proto_goTypes,
		DependencyIndexes: file_network_proto_depIdxs,
		MessageInfos:      file_network_proto_msgTypes,
	}.Build()
	File_network_proto = out.File
	file_network_proto_rawDesc = nil
	file_network_proto_goTypes = nil
	file_network_proto_depIdxs = nil
}