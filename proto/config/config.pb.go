// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config

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

type App struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{0}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *App) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *App) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *App) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Apps struct {
	Apps                 []*App   `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Apps) Reset()         { *m = Apps{} }
func (m *Apps) String() string { return proto.CompactTextString(m) }
func (*Apps) ProtoMessage()    {}
func (*Apps) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{1}
}

func (m *Apps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Apps.Unmarshal(m, b)
}
func (m *Apps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Apps.Marshal(b, m, deterministic)
}
func (m *Apps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Apps.Merge(m, src)
}
func (m *Apps) XXX_Size() int {
	return xxx_messageInfo_Apps.Size(m)
}
func (m *Apps) XXX_DiscardUnknown() {
	xxx_messageInfo_Apps.DiscardUnknown(m)
}

var xxx_messageInfo_Apps proto.InternalMessageInfo

func (m *Apps) GetApps() []*App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type Cluster struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName,omitempty"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{2}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cluster) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Cluster) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Cluster) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Cluster) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Cluster) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Clusters struct {
	Clusters             []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{3}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type Namespace struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName,omitempty"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName,omitempty"`
	Value                string   `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{4}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Namespace) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Namespace) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Namespace) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Namespace) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Namespace) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Namespace) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Namespace) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Namespaces struct {
	Namespaces           []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Namespaces) Reset()         { *m = Namespaces{} }
func (m *Namespaces) String() string { return proto.CompactTextString(m) }
func (*Namespaces) ProtoMessage()    {}
func (*Namespaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{5}
}

func (m *Namespaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespaces.Unmarshal(m, b)
}
func (m *Namespaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespaces.Marshal(b, m, deterministic)
}
func (m *Namespaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespaces.Merge(m, src)
}
func (m *Namespaces) XXX_Size() int {
	return xxx_messageInfo_Namespaces.Size(m)
}
func (m *Namespaces) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespaces.DiscardUnknown(m)
}

var xxx_messageInfo_Namespaces proto.InternalMessageInfo

func (m *Namespaces) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type Release struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName,omitempty"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName,omitempty"`
	Comment              string   `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{6}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Release) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Release) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Release) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Release) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Release) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{7}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{8}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*App)(nil), "App")
	proto.RegisterType((*Apps)(nil), "Apps")
	proto.RegisterType((*Cluster)(nil), "Cluster")
	proto.RegisterType((*Clusters)(nil), "Clusters")
	proto.RegisterType((*Namespace)(nil), "Namespace")
	proto.RegisterType((*Namespaces)(nil), "Namespaces")
	proto.RegisterType((*Release)(nil), "Release")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_cc332a44e926b360) }

var fileDescriptor_cc332a44e926b360 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0xb5, 0x63, 0x27, 0xb6, 0x5f, 0x26, 0x33, 0x52, 0xc3, 0xc2, 0x44, 0x80, 0x3c, 0xad, 0x61,
	0x08, 0x83, 0x64, 0x46, 0x61, 0xc3, 0x36, 0x0a, 0x4b, 0xc4, 0xc2, 0x12, 0x62, 0x6d, 0xec, 0x06,
	0x2c, 0x25, 0x76, 0x93, 0x76, 0xb8, 0x05, 0xa7, 0xe1, 0x24, 0xdc, 0x04, 0x6e, 0x80, 0xdc, 0x1f,
	0xa7, 0x9d, 0xc0, 0x16, 0x31, 0x9b, 0xc4, 0xfd, 0x5e, 0xb9, 0x5e, 0xbd, 0xea, 0x2a, 0xe3, 0x5e,
	0xd1, 0xd4, 0x1f, 0xab, 0x4f, 0x2f, 0xd4, 0x5f, 0xca, 0x77, 0x4d, 0xdb, 0xd0, 0x6f, 0x2e, 0xbc,
	0x15, 0xe7, 0xe4, 0x1c, 0xa3, 0xaa, 0x8c, 0xdd, 0xc4, 0x5d, 0x78, 0xd9, 0xa8, 0x2a, 0xc9, 0x43,
	0x44, 0xc5, 0x8e, 0xe5, 0x2d, 0x2b, 0x57, 0x6d, 0x3c, 0x92, 0xf0, 0x01, 0xe8, 0xd8, 0x3d, 0x2f,
	0x35, 0xeb, 0x29, 0xb6, 0x07, 0x48, 0x8c, 0x20, 0xe7, 0xfc, 0x6d, 0xbe, 0x65, 0xb1, 0x9f, 0xb8,
	0x8b, 0x28, 0x33, 0x47, 0x92, 0x60, 0x5a, 0x32, 0x51, 0xec, 0x2a, 0xde, 0x56, 0x4d, 0x1d, 0x8f,
	0x25, 0x6b, 0x43, 0x34, 0x81, 0xbf, 0xe2, 0x5c, 0x90, 0x18, 0x7e, 0xce, 0xb9, 0x88, 0xdd, 0xc4,
	0x5b, 0x4c, 0x97, 0x7e, 0xba, 0xe2, 0x3c, 0x93, 0x08, 0xfd, 0xee, 0x22, 0x58, 0x6f, 0xf6, 0xa2,
	0x65, 0xbb, 0x7f, 0x57, 0x75, 0xa1, 0x04, 0x25, 0xab, 0xab, 0xb6, 0xa0, 0x63, 0x5f, 0x93, 0x53,
	0x5f, 0xb7, 0x08, 0x75, 0xd1, 0x82, 0x5c, 0x21, 0xd4, 0x2f, 0x1b, 0x7f, 0x61, 0xaa, 0xc9, 0xac,
	0x67, 0xe8, 0x2f, 0x17, 0x51, 0x97, 0x5c, 0xf0, 0xbc, 0x60, 0xff, 0x8d, 0xd3, 0x2b, 0xcc, 0x6a,
	0x53, 0x94, 0x8c, 0x51, 0x5e, 0x87, 0x20, 0xb9, 0x8f, 0xf1, 0xd7, 0x7c, 0xb3, 0x67, 0x71, 0x20,
	0x59, 0x75, 0x38, 0xee, 0x52, 0x78, 0xda, 0xa5, 0x57, 0x40, 0x6f, 0x59, 0x90, 0x1b, 0xa0, 0x4f,
	0x6b, 0x3a, 0x85, 0xb4, 0x0f, 0xc8, 0x2c, 0x96, 0xfe, 0x70, 0x11, 0x64, 0x6c, 0xc3, 0x72, 0x71,
	0xd7, 0x7a, 0x15, 0x23, 0x28, 0x9a, 0xed, 0x96, 0xd5, 0xad, 0xee, 0x96, 0x39, 0xd2, 0xa8, 0xb3,
	0xf4, 0x65, 0xcf, 0x44, 0x4b, 0x81, 0x30, 0x63, 0x82, 0x37, 0xb5, 0x60, 0xcb, 0x9f, 0x1e, 0x26,
	0x6b, 0xb9, 0xc3, 0xe4, 0x01, 0xa2, 0xb5, 0x34, 0xd2, 0xad, 0xb0, 0x5c, 0x92, 0xb9, 0xfc, 0xa5,
	0x0e, 0x79, 0x8c, 0xe8, 0x35, 0xdb, 0x30, 0x9b, 0x8a, 0x52, 0x93, 0x83, 0x3a, 0xe4, 0x11, 0xc2,
	0x37, 0x95, 0x68, 0xe5, 0xb2, 0x85, 0xa9, 0xd6, 0x99, 0x8f, 0xbb, 0x40, 0x41, 0x1d, 0xf2, 0x04,
	0x33, 0x95, 0xd9, 0xac, 0x5a, 0x3f, 0xa2, 0xf3, 0xfe, 0x89, 0x3a, 0xe4, 0x1a, 0x33, 0xa5, 0x72,
	0x1a, 0x36, 0x50, 0xbb, 0xc4, 0x59, 0xa7, 0xd6, 0xaf, 0x80, 0x29, 0xc8, 0x00, 0xd4, 0x21, 0xcf,
	0x71, 0xa1, 0x14, 0x0f, 0x43, 0x6f, 0x5d, 0xf6, 0xdc, 0x7a, 0xa6, 0x0e, 0xb9, 0xc1, 0x85, 0xd2,
	0xfd, 0x73, 0xf0, 0x40, 0xfb, 0x19, 0xce, 0x3b, 0x6d, 0x6b, 0xb0, 0x0e, 0x45, 0x4e, 0x0f, 0x2f,
	0x75, 0x35, 0x3c, 0xc5, 0xd9, 0x3b, 0x79, 0xf5, 0xba, 0xbf, 0x7f, 0xcd, 0x79, 0x8d, 0x99, 0x9e,
	0x36, 0x1d, 0xd9, 0xb5, 0x50, 0x9e, 0x87, 0x71, 0x14, 0x7e, 0xc6, 0xf2, 0x92, 0xd8, 0x3a, 0xc7,
	0xa2, 0x97, 0x18, 0xbf, 0xcf, 0xdb, 0xe2, 0xb3, 0x75, 0x0d, 0x03, 0xb3, 0xb7, 0xee, 0x87, 0x89,
	0xfc, 0x58, 0xbf, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xee, 0x83, 0xfe, 0xc3, 0x05, 0x00,
	0x00,
}
