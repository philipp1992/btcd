// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Common.proto

package lightwalletrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type BlockHeader struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Time                 uint64   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	PrevBlockHash        string   `protobuf:"bytes,4,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{1}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockHeader) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *BlockHeader) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

type Outpoint struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outpoint) Reset()         { *m = Outpoint{} }
func (m *Outpoint) String() string { return proto.CompactTextString(m) }
func (*Outpoint) ProtoMessage()    {}
func (*Outpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{2}
}

func (m *Outpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outpoint.Unmarshal(m, b)
}
func (m *Outpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outpoint.Marshal(b, m, deterministic)
}
func (m *Outpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outpoint.Merge(m, src)
}
func (m *Outpoint) XXX_Size() int {
	return xxx_messageInfo_Outpoint.Size(m)
}
func (m *Outpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Outpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Outpoint proto.InternalMessageInfo

func (m *Outpoint) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Outpoint) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type TxOut struct {
	ScriptPubkey         string   `protobuf:"bytes,1,opt,name=scriptPubkey,proto3" json:"scriptPubkey,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxOut) Reset()         { *m = TxOut{} }
func (m *TxOut) String() string { return proto.CompactTextString(m) }
func (*TxOut) ProtoMessage()    {}
func (*TxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{3}
}

func (m *TxOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxOut.Unmarshal(m, b)
}
func (m *TxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxOut.Marshal(b, m, deterministic)
}
func (m *TxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOut.Merge(m, src)
}
func (m *TxOut) XXX_Size() int {
	return xxx_messageInfo_TxOut.Size(m)
}
func (m *TxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOut.DiscardUnknown(m)
}

var xxx_messageInfo_TxOut proto.InternalMessageInfo

func (m *TxOut) GetScriptPubkey() string {
	if m != nil {
		return m.ScriptPubkey
	}
	return ""
}

func (m *TxOut) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TxID struct {
	Txid                 string   `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxID) Reset()         { *m = TxID{} }
func (m *TxID) String() string { return proto.CompactTextString(m) }
func (*TxID) ProtoMessage()    {}
func (*TxID) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{4}
}

func (m *TxID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxID.Unmarshal(m, b)
}
func (m *TxID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxID.Marshal(b, m, deterministic)
}
func (m *TxID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxID.Merge(m, src)
}
func (m *TxID) XXX_Size() int {
	return xxx_messageInfo_TxID.Size(m)
}
func (m *TxID) XXX_DiscardUnknown() {
	xxx_messageInfo_TxID.DiscardUnknown(m)
}

var xxx_messageInfo_TxID proto.InternalMessageInfo

func (m *TxID) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

type BlockHash struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHash) Reset()         { *m = BlockHash{} }
func (m *BlockHash) String() string { return proto.CompactTextString(m) }
func (*BlockHash) ProtoMessage()    {}
func (*BlockHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{5}
}

func (m *BlockHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHash.Unmarshal(m, b)
}
func (m *BlockHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHash.Marshal(b, m, deterministic)
}
func (m *BlockHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHash.Merge(m, src)
}
func (m *BlockHash) XXX_Size() int {
	return xxx_messageInfo_BlockHash.Size(m)
}
func (m *BlockHash) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHash.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHash proto.InternalMessageInfo

func (m *BlockHash) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type HexEncoded struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HexEncoded) Reset()         { *m = HexEncoded{} }
func (m *HexEncoded) String() string { return proto.CompactTextString(m) }
func (*HexEncoded) ProtoMessage()    {}
func (*HexEncoded) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{6}
}

func (m *HexEncoded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HexEncoded.Unmarshal(m, b)
}
func (m *HexEncoded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HexEncoded.Marshal(b, m, deterministic)
}
func (m *HexEncoded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HexEncoded.Merge(m, src)
}
func (m *HexEncoded) XXX_Size() int {
	return xxx_messageInfo_HexEncoded.Size(m)
}
func (m *HexEncoded) XXX_DiscardUnknown() {
	xxx_messageInfo_HexEncoded.DiscardUnknown(m)
}

var xxx_messageInfo_HexEncoded proto.InternalMessageInfo

func (m *HexEncoded) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type Height struct {
	Height               int32    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Height) Reset()         { *m = Height{} }
func (m *Height) String() string { return proto.CompactTextString(m) }
func (*Height) ProtoMessage()    {}
func (*Height) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee72d9a89737215c, []int{7}
}

func (m *Height) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Height.Unmarshal(m, b)
}
func (m *Height) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Height.Marshal(b, m, deterministic)
}
func (m *Height) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Height.Merge(m, src)
}
func (m *Height) XXX_Size() int {
	return xxx_messageInfo_Height.Size(m)
}
func (m *Height) XXX_DiscardUnknown() {
	xxx_messageInfo_Height.DiscardUnknown(m)
}

var xxx_messageInfo_Height proto.InternalMessageInfo

func (m *Height) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "lightwalletrpc.Empty")
	proto.RegisterType((*BlockHeader)(nil), "lightwalletrpc.BlockHeader")
	proto.RegisterType((*Outpoint)(nil), "lightwalletrpc.Outpoint")
	proto.RegisterType((*TxOut)(nil), "lightwalletrpc.TxOut")
	proto.RegisterType((*TxID)(nil), "lightwalletrpc.TxID")
	proto.RegisterType((*BlockHash)(nil), "lightwalletrpc.BlockHash")
	proto.RegisterType((*HexEncoded)(nil), "lightwalletrpc.HexEncoded")
	proto.RegisterType((*Height)(nil), "lightwalletrpc.Height")
}

func init() { proto.RegisterFile("Common.proto", fileDescriptor_ee72d9a89737215c) }

var fileDescriptor_ee72d9a89737215c = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x68, 0x12, 0xe8, 0xd2, 0x72, 0xb0, 0x10, 0x8a, 0xb8, 0x10, 0x59, 0x1c, 0x72, 0xe2,
	0x02, 0x3f, 0xc0, 0xa3, 0x52, 0x38, 0x15, 0x59, 0xfd, 0x81, 0x34, 0x5e, 0x11, 0xab, 0x4e, 0x6c,
	0x39, 0x4e, 0x71, 0xff, 0x1e, 0xd9, 0x29, 0x8f, 0x4a, 0xed, 0x6d, 0x66, 0xb5, 0x3b, 0x33, 0x9a,
	0x85, 0xd9, 0xab, 0x6a, 0x5b, 0xd5, 0x3d, 0x68, 0xa3, 0xac, 0x22, 0x57, 0x52, 0x7c, 0x36, 0xf6,
	0xab, 0x92, 0x12, 0xad, 0xd1, 0x35, 0x3d, 0x87, 0x64, 0xd1, 0x6a, 0xbb, 0xa3, 0x3d, 0x5c, 0xbe,
	0x48, 0x55, 0x6f, 0x4a, 0xac, 0x38, 0x1a, 0x42, 0x20, 0x6e, 0xaa, 0xbe, 0xc9, 0xa2, 0x3c, 0x2a,
	0xa6, 0x2c, 0x60, 0x72, 0x03, 0x69, 0x83, 0xfe, 0x3c, 0x3b, 0xcb, 0xa3, 0x22, 0x61, 0x7b, 0xe6,
	0x77, 0xad, 0x68, 0x31, 0x9b, 0xe4, 0x51, 0x11, 0xb3, 0x80, 0xc9, 0x3d, 0xcc, 0xb5, 0xc1, 0xed,
	0x28, 0xe9, 0x85, 0xe2, 0x20, 0x74, 0x38, 0xa4, 0x4f, 0x70, 0xb1, 0x1c, 0xac, 0x56, 0xa2, 0xb3,
	0x47, 0x1d, 0xaf, 0x21, 0x11, 0x1d, 0x47, 0x17, 0x0c, 0xe7, 0x6c, 0x24, 0xf4, 0x19, 0x92, 0x95,
	0x5b, 0x0e, 0x96, 0x50, 0x98, 0xf5, 0xb5, 0x11, 0xda, 0x7e, 0x0c, 0xeb, 0x0d, 0xee, 0xf6, 0xa7,
	0x07, 0x33, 0x2f, 0xb1, 0xad, 0xe4, 0x80, 0x41, 0x62, 0xc2, 0x46, 0x42, 0x6f, 0x21, 0x5e, 0xb9,
	0xf7, 0xb7, 0x10, 0xdd, 0x09, 0xfe, 0x63, 0xea, 0x31, 0xbd, 0x83, 0xe9, 0x6f, 0xc2, 0x63, 0xa9,
	0x68, 0x0e, 0x50, 0xa2, 0x5b, 0x74, 0xb5, 0xe2, 0xc8, 0x4f, 0x6c, 0xa4, 0xe5, 0xd8, 0xcd, 0x5f,
	0x67, 0xd1, 0xff, 0xce, 0xd6, 0x69, 0x78, 0xc7, 0xe3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3,
	0xa5, 0x03, 0xbb, 0x9e, 0x01, 0x00, 0x00,
}
