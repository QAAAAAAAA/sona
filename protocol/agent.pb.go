// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package protocol

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

type MsgTypeId int32

const (
	MsgTypeId_GetConfigReqId    MsgTypeId = 1
	MsgTypeId_RemoveConfigReqId MsgTypeId = 2
	MsgTypeId_PullConfigReqId   MsgTypeId = 3
	MsgTypeId_PullConfigRspId   MsgTypeId = 4
	MsgTypeId_PushConfigReqId   MsgTypeId = 5
	MsgTypeId_AddNewConfigReqId MsgTypeId = 6
	MsgTypeId_UpdateConfigReqId MsgTypeId = 7
	MsgTypeId_DeleteConfigReqId MsgTypeId = 8
	MsgTypeId_ManageConfigRspId MsgTypeId = 9
)

var MsgTypeId_name = map[int32]string{
	1: "GetConfigReqId",
	2: "RemoveConfigReqId",
	3: "PullConfigReqId",
	4: "PullConfigRspId",
	5: "PushConfigReqId",
	6: "AddNewConfigReqId",
	7: "UpdateConfigReqId",
	8: "DeleteConfigReqId",
	9: "ManageConfigRspId",
}
var MsgTypeId_value = map[string]int32{
	"GetConfigReqId":    1,
	"RemoveConfigReqId": 2,
	"PullConfigReqId":   3,
	"PullConfigRspId":   4,
	"PushConfigReqId":   5,
	"AddNewConfigReqId": 6,
	"UpdateConfigReqId": 7,
	"DeleteConfigReqId": 8,
	"ManageConfigRspId": 9,
}

func (x MsgTypeId) Enum() *MsgTypeId {
	p := new(MsgTypeId)
	*p = x
	return p
}
func (x MsgTypeId) String() string {
	return proto.EnumName(MsgTypeId_name, int32(x))
}
func (x *MsgTypeId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgTypeId_value, data, "MsgTypeId")
	if err != nil {
		return err
	}
	*x = MsgTypeId(value)
	return nil
}
func (MsgTypeId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{0}
}

type GetConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigReq) Reset()         { *m = GetConfigReq{} }
func (m *GetConfigReq) String() string { return proto.CompactTextString(m) }
func (*GetConfigReq) ProtoMessage()    {}
func (*GetConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{0}
}
func (m *GetConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigReq.Unmarshal(m, b)
}
func (m *GetConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigReq.Marshal(b, m, deterministic)
}
func (dst *GetConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigReq.Merge(dst, src)
}
func (m *GetConfigReq) XXX_Size() int {
	return xxx_messageInfo_GetConfigReq.Size(m)
}
func (m *GetConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigReq proto.InternalMessageInfo

func (m *GetConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type GetConfigRsp struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRsp) Reset()         { *m = GetConfigRsp{} }
func (m *GetConfigRsp) String() string { return proto.CompactTextString(m) }
func (*GetConfigRsp) ProtoMessage()    {}
func (*GetConfigRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{1}
}
func (m *GetConfigRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRsp.Unmarshal(m, b)
}
func (m *GetConfigRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRsp.Marshal(b, m, deterministic)
}
func (dst *GetConfigRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRsp.Merge(dst, src)
}
func (m *GetConfigRsp) XXX_Size() int {
	return xxx_messageInfo_GetConfigRsp.Size(m)
}
func (m *GetConfigRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRsp proto.InternalMessageInfo

func (m *GetConfigRsp) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *GetConfigRsp) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type PushConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushConfigReq) Reset()         { *m = PushConfigReq{} }
func (m *PushConfigReq) String() string { return proto.CompactTextString(m) }
func (*PushConfigReq) ProtoMessage()    {}
func (*PushConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{2}
}
func (m *PushConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushConfigReq.Unmarshal(m, b)
}
func (m *PushConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushConfigReq.Marshal(b, m, deterministic)
}
func (dst *PushConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushConfigReq.Merge(dst, src)
}
func (m *PushConfigReq) XXX_Size() int {
	return xxx_messageInfo_PushConfigReq.Size(m)
}
func (m *PushConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushConfigReq proto.InternalMessageInfo

func (m *PushConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PushConfigReq) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type RemoveConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveConfigReq) Reset()         { *m = RemoveConfigReq{} }
func (m *RemoveConfigReq) String() string { return proto.CompactTextString(m) }
func (*RemoveConfigReq) ProtoMessage()    {}
func (*RemoveConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{3}
}
func (m *RemoveConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveConfigReq.Unmarshal(m, b)
}
func (m *RemoveConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveConfigReq.Marshal(b, m, deterministic)
}
func (dst *RemoveConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveConfigReq.Merge(dst, src)
}
func (m *RemoveConfigReq) XXX_Size() int {
	return xxx_messageInfo_RemoveConfigReq.Size(m)
}
func (m *RemoveConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveConfigReq proto.InternalMessageInfo

func (m *RemoveConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type PullConfigReq struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullConfigReq) Reset()         { *m = PullConfigReq{} }
func (m *PullConfigReq) String() string { return proto.CompactTextString(m) }
func (*PullConfigReq) ProtoMessage()    {}
func (*PullConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{4}
}
func (m *PullConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullConfigReq.Unmarshal(m, b)
}
func (m *PullConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullConfigReq.Marshal(b, m, deterministic)
}
func (dst *PullConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullConfigReq.Merge(dst, src)
}
func (m *PullConfigReq) XXX_Size() int {
	return xxx_messageInfo_PullConfigReq.Size(m)
}
func (m *PullConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PullConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_PullConfigReq proto.InternalMessageInfo

func (m *PullConfigReq) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PullConfigReq) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type PullConfigRsp struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullConfigRsp) Reset()         { *m = PullConfigRsp{} }
func (m *PullConfigRsp) String() string { return proto.CompactTextString(m) }
func (*PullConfigRsp) ProtoMessage()    {}
func (*PullConfigRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{5}
}
func (m *PullConfigRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullConfigRsp.Unmarshal(m, b)
}
func (m *PullConfigRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullConfigRsp.Marshal(b, m, deterministic)
}
func (dst *PullConfigRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullConfigRsp.Merge(dst, src)
}
func (m *PullConfigRsp) XXX_Size() int {
	return xxx_messageInfo_PullConfigRsp.Size(m)
}
func (m *PullConfigRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PullConfigRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PullConfigRsp proto.InternalMessageInfo

func (m *PullConfigRsp) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PullConfigRsp) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type AddNewConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNewConfigReq) Reset()         { *m = AddNewConfigReq{} }
func (m *AddNewConfigReq) String() string { return proto.CompactTextString(m) }
func (*AddNewConfigReq) ProtoMessage()    {}
func (*AddNewConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{6}
}
func (m *AddNewConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNewConfigReq.Unmarshal(m, b)
}
func (m *AddNewConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNewConfigReq.Marshal(b, m, deterministic)
}
func (dst *AddNewConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNewConfigReq.Merge(dst, src)
}
func (m *AddNewConfigReq) XXX_Size() int {
	return xxx_messageInfo_AddNewConfigReq.Size(m)
}
func (m *AddNewConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNewConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddNewConfigReq proto.InternalMessageInfo

func (m *AddNewConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *AddNewConfigReq) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type UpdateConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	OldValue             *string  `protobuf:"bytes,2,req,name=oldValue" json:"oldValue,omitempty"`
	NewValue             *string  `protobuf:"bytes,3,req,name=newValue" json:"newValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigReq) Reset()         { *m = UpdateConfigReq{} }
func (m *UpdateConfigReq) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigReq) ProtoMessage()    {}
func (*UpdateConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{7}
}
func (m *UpdateConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigReq.Unmarshal(m, b)
}
func (m *UpdateConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigReq.Marshal(b, m, deterministic)
}
func (dst *UpdateConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigReq.Merge(dst, src)
}
func (m *UpdateConfigReq) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigReq.Size(m)
}
func (m *UpdateConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigReq proto.InternalMessageInfo

func (m *UpdateConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *UpdateConfigReq) GetOldValue() string {
	if m != nil && m.OldValue != nil {
		return *m.OldValue
	}
	return ""
}

func (m *UpdateConfigReq) GetNewValue() string {
	if m != nil && m.NewValue != nil {
		return *m.NewValue
	}
	return ""
}

type DeleteConfigReq struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	OldValue             *string  `protobuf:"bytes,2,req,name=oldValue" json:"oldValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConfigReq) Reset()         { *m = DeleteConfigReq{} }
func (m *DeleteConfigReq) String() string { return proto.CompactTextString(m) }
func (*DeleteConfigReq) ProtoMessage()    {}
func (*DeleteConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{8}
}
func (m *DeleteConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConfigReq.Unmarshal(m, b)
}
func (m *DeleteConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConfigReq.Marshal(b, m, deterministic)
}
func (dst *DeleteConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConfigReq.Merge(dst, src)
}
func (m *DeleteConfigReq) XXX_Size() int {
	return xxx_messageInfo_DeleteConfigReq.Size(m)
}
func (m *DeleteConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConfigReq proto.InternalMessageInfo

func (m *DeleteConfigReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *DeleteConfigReq) GetOldValue() string {
	if m != nil && m.OldValue != nil {
		return *m.OldValue
	}
	return ""
}

type ManageConfigRsp struct {
	Retcode              *int32   `protobuf:"varint,1,req,name=retcode" json:"retcode,omitempty"`
	Errmsg               *string  `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManageConfigRsp) Reset()         { *m = ManageConfigRsp{} }
func (m *ManageConfigRsp) String() string { return proto.CompactTextString(m) }
func (*ManageConfigRsp) ProtoMessage()    {}
func (*ManageConfigRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_007ec2824bab5881, []int{9}
}
func (m *ManageConfigRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManageConfigRsp.Unmarshal(m, b)
}
func (m *ManageConfigRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManageConfigRsp.Marshal(b, m, deterministic)
}
func (dst *ManageConfigRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManageConfigRsp.Merge(dst, src)
}
func (m *ManageConfigRsp) XXX_Size() int {
	return xxx_messageInfo_ManageConfigRsp.Size(m)
}
func (m *ManageConfigRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ManageConfigRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ManageConfigRsp proto.InternalMessageInfo

func (m *ManageConfigRsp) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *ManageConfigRsp) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetConfigReq)(nil), "protocol.GetConfigReq")
	proto.RegisterType((*GetConfigRsp)(nil), "protocol.GetConfigRsp")
	proto.RegisterType((*PushConfigReq)(nil), "protocol.PushConfigReq")
	proto.RegisterType((*RemoveConfigReq)(nil), "protocol.RemoveConfigReq")
	proto.RegisterType((*PullConfigReq)(nil), "protocol.PullConfigReq")
	proto.RegisterType((*PullConfigRsp)(nil), "protocol.PullConfigRsp")
	proto.RegisterType((*AddNewConfigReq)(nil), "protocol.AddNewConfigReq")
	proto.RegisterType((*UpdateConfigReq)(nil), "protocol.UpdateConfigReq")
	proto.RegisterType((*DeleteConfigReq)(nil), "protocol.DeleteConfigReq")
	proto.RegisterType((*ManageConfigRsp)(nil), "protocol.ManageConfigRsp")
	proto.RegisterEnum("protocol.MsgTypeId", MsgTypeId_name, MsgTypeId_value)
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_agent_007ec2824bab5881) }

var fileDescriptor_agent_007ec2824bab5881 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xdb, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xd2, 0x53, 0xa6, 0x87, 0x5d, 0x57, 0x84, 0x80, 0x20, 0xa5, 0x57, 0xa5, 0xa2,
	0x88, 0x6f, 0xe0, 0x01, 0x64, 0x2f, 0x2a, 0x52, 0xd4, 0xfb, 0xd0, 0x19, 0xa3, 0x74, 0x9b, 0x5d,
	0xb3, 0x69, 0x4b, 0x1f, 0xd3, 0x37, 0x92, 0x4d, 0x2f, 0xda, 0x2e, 0x16, 0x7b, 0x15, 0xf2, 0xcd,
	0x3f, 0xff, 0xce, 0xfc, 0x03, 0xed, 0x34, 0xa3, 0xbc, 0xbc, 0x36, 0x85, 0x2e, 0xb5, 0x68, 0x55,
	0x9f, 0xa9, 0x56, 0x83, 0x73, 0xe8, 0x3c, 0x51, 0xf9, 0xa0, 0xf3, 0x8f, 0xaf, 0x6c, 0x42, 0xdf,
	0xa2, 0x0d, 0xd1, 0x8c, 0xd6, 0x49, 0xd0, 0x0f, 0x87, 0xf1, 0x60, 0xb4, 0x5b, 0xb4, 0x66, 0xaf,
	0x28, 0xba, 0x50, 0x5f, 0xa6, 0x6a, 0x41, 0x49, 0x58, 0x69, 0x2f, 0xa1, 0xfb, 0xb2, 0xb0, 0x9f,
	0x7f, 0x3b, 0xf9, 0xe2, 0x0b, 0x60, 0x13, 0x9a, 0xeb, 0x25, 0x1d, 0x78, 0xf8, 0xca, 0x99, 0x29,
	0xb5, 0xad, 0x76, 0xa0, 0x36, 0xa3, 0xb5, 0x4d, 0x82, 0x7e, 0x34, 0x8c, 0x45, 0x0f, 0x1a, 0x95,
	0x9b, 0x4d, 0x42, 0xf7, 0xef, 0xc9, 0xad, 0xf9, 0x57, 0xce, 0xee, 0x10, 0x9f, 0x69, 0x75, 0xdc,
	0xb0, 0xf7, 0xc0, 0xde, 0x0c, 0xa6, 0xe5, 0x81, 0x61, 0x05, 0x87, 0x96, 0x56, 0xf8, 0xbe, 0xed,
	0x70, 0x24, 0xa7, 0xd5, 0x86, 0x44, 0x95, 0xc7, 0x0d, 0xb0, 0x47, 0x52, 0x74, 0xbc, 0xc7, 0xe0,
	0x16, 0xd8, 0x38, 0xcd, 0xd3, 0x8c, 0xb6, 0x5b, 0x31, 0x68, 0x16, 0x54, 0x4e, 0x35, 0x52, 0xd5,
	0x55, 0x77, 0x8b, 0x51, 0x51, 0xcc, 0x6d, 0x96, 0x84, 0xfd, 0x60, 0x18, 0x8f, 0x7e, 0x02, 0x88,
	0xc7, 0x36, 0x7b, 0x5d, 0x1b, 0x92, 0x28, 0x04, 0xf4, 0x76, 0x4f, 0x2b, 0x91, 0x07, 0xe2, 0x0c,
	0x4e, 0xbc, 0xe0, 0x25, 0xf2, 0x50, 0x9c, 0x02, 0xdb, 0xcb, 0x5b, 0x22, 0x8f, 0x3c, 0x68, 0x8d,
	0x44, 0x5e, 0xdb, 0xc0, 0x9d, 0x33, 0x4b, 0xe4, 0x75, 0xe7, 0xea, 0x05, 0x2a, 0x91, 0x37, 0x1c,
	0xf6, 0x82, 0x93, 0xc8, 0x9b, 0x0e, 0x7b, 0x59, 0x48, 0xe4, 0x2d, 0x87, 0xbd, 0x85, 0x25, 0xf2,
	0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x84, 0x4b, 0x6d, 0x0d, 0xb8, 0x02, 0x00, 0x00,
}
