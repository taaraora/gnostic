// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/plugin.proto

package gnostic_plugin_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Message_Level int32

const (
	Message_UNKNOWN Message_Level = 0
	Message_INFO    Message_Level = 1
	Message_WARNING Message_Level = 2
	Message_ERROR   Message_Level = 3
	Message_FATAL   Message_Level = 4
)

var Message_Level_name = map[int32]string{
	0: "UNKNOWN",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "FATAL",
}

var Message_Level_value = map[string]int32{
	"UNKNOWN": 0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
	"FATAL":   4,
}

func (x Message_Level) String() string {
	return proto.EnumName(Message_Level_name, int32(x))
}

func (Message_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{3, 0}
}

// The version number of gnostic.
type Version struct {
	Major int32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch int32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	// A suffix for alpha, beta or rc release, e.g., "alpha-1", "rc2". It should
	// be empty for mainline stable releases.
	Suffix               string   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

// A parameter passed to the plugin from (or through) gnostic.
type Parameter struct {
	// The name of the parameter as specified in the option string
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The parameter value as specified in the option string
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{1}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// An encoded Request is written to the plugin's stdin.
type Request struct {
	// filename or URL of the original source document
	SourceName string `protobuf:"bytes,1,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	// Output path specified in the plugin invocation.
	OutputPath string `protobuf:"bytes,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	// Plugin parameters parsed from the invocation string.
	Parameters []*Parameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	// The version number of gnostic.
	CompilerVersion *Version `protobuf:"bytes,4,opt,name=compiler_version,json=compilerVersion,proto3" json:"compiler_version,omitempty"`
	// API models
	Models               []*any.Any `protobuf:"bytes,5,rep,name=models,proto3" json:"models,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{2}
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

func (m *Request) GetSourceName() string {
	if m != nil {
		return m.SourceName
	}
	return ""
}

func (m *Request) GetOutputPath() string {
	if m != nil {
		return m.OutputPath
	}
	return ""
}

func (m *Request) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Request) GetCompilerVersion() *Version {
	if m != nil {
		return m.CompilerVersion
	}
	return nil
}

func (m *Request) GetModels() []*any.Any {
	if m != nil {
		return m.Models
	}
	return nil
}

// Plugins can return messages to be collated and reported by gnostic.
type Message struct {
	// message severity
	Level Message_Level `protobuf:"varint,1,opt,name=level,proto3,enum=gnostic.plugin.v1.Message_Level" json:"level,omitempty"`
	// a unique message identifier
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// message text
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// an associated key path in an API description
	Keys                 []string `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetLevel() Message_Level {
	if m != nil {
		return m.Level
	}
	return Message_UNKNOWN
}

func (m *Message) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type Messages struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Messages) Reset()         { *m = Messages{} }
func (m *Messages) String() string { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()    {}
func (*Messages) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{4}
}

func (m *Messages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messages.Unmarshal(m, b)
}
func (m *Messages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messages.Marshal(b, m, deterministic)
}
func (m *Messages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messages.Merge(m, src)
}
func (m *Messages) XXX_Size() int {
	return xxx_messageInfo_Messages.Size(m)
}
func (m *Messages) XXX_DiscardUnknown() {
	xxx_messageInfo_Messages.DiscardUnknown(m)
}

var xxx_messageInfo_Messages proto.InternalMessageInfo

func (m *Messages) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// The plugin writes an encoded Response to stdout.
type Response struct {
	// Error message.  If non-empty, the plugin failed.
	// The plugin process should exit with status code zero
	// even if it reports an error in this way.
	//
	// This should be used to indicate errors which prevent the plugin from
	// operating as intended.  Errors which indicate a problem in gnostic
	// itself -- such as the input Document being unparseable -- should be
	// reported by writing a message to stderr and exiting with a non-zero
	// status code.
	Errors []string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	// file output, each file will be written by gnostic to an appropriate location.
	Files []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	// informational messages to be collected and reported by gnostic.
	Messages             []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{5}
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

func (m *Response) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *Response) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// File describes a file generated by a plugin.
type File struct {
	// name of the file
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// data to be written to the file
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be178a99b712ecb, []int{6}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("gnostic.plugin.v1.Message_Level", Message_Level_name, Message_Level_value)
	proto.RegisterType((*Version)(nil), "gnostic.plugin.v1.Version")
	proto.RegisterType((*Parameter)(nil), "gnostic.plugin.v1.Parameter")
	proto.RegisterType((*Request)(nil), "gnostic.plugin.v1.Request")
	proto.RegisterType((*Message)(nil), "gnostic.plugin.v1.Message")
	proto.RegisterType((*Messages)(nil), "gnostic.plugin.v1.Messages")
	proto.RegisterType((*Response)(nil), "gnostic.plugin.v1.Response")
	proto.RegisterType((*File)(nil), "gnostic.plugin.v1.File")
}

func init() { proto.RegisterFile("plugins/plugin.proto", fileDescriptor_5be178a99b712ecb) }

var fileDescriptor_5be178a99b712ecb = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xb1, 0xdd, 0xc4, 0x13, 0x28, 0x66, 0x55, 0x81, 0xa9, 0x90, 0x1a, 0xf9, 0x42, 0x0e,
	0xe0, 0xa8, 0x41, 0xf4, 0xc4, 0x25, 0x91, 0x9a, 0xa8, 0xa2, 0x38, 0xd1, 0x0a, 0xe8, 0x31, 0xda,
	0x3a, 0x1b, 0xc7, 0x60, 0x7b, 0xcd, 0xee, 0x3a, 0x6a, 0x3e, 0x81, 0xdf, 0xe0, 0x4b, 0xf8, 0x32,
	0x84, 0x76, 0xd7, 0x89, 0x8a, 0x08, 0x07, 0x4e, 0x7e, 0xf3, 0xf4, 0xfc, 0x66, 0xe6, 0xed, 0xc0,
	0x49, 0x95, 0xd7, 0x69, 0x56, 0x8a, 0x81, 0xf9, 0x46, 0x15, 0x67, 0x92, 0xa1, 0x27, 0x69, 0xc9,
	0x84, 0xcc, 0x92, 0xa8, 0x61, 0x37, 0xe7, 0xa7, 0xcf, 0x53, 0xc6, 0xd2, 0x9c, 0x0e, 0xb4, 0xe0,
	0xb6, 0x5e, 0x0d, 0x48, 0xb9, 0x35, 0xea, 0x30, 0x81, 0xf6, 0x67, 0xca, 0x45, 0xc6, 0x4a, 0x74,
	0x02, 0x6e, 0x41, 0xbe, 0x30, 0x1e, 0x58, 0x3d, 0xab, 0xef, 0x62, 0x53, 0x68, 0x36, 0x2b, 0x19,
	0x0f, 0x5a, 0x0d, 0xab, 0x0a, 0xc5, 0x56, 0x44, 0x26, 0xeb, 0xc0, 0x36, 0xac, 0x2e, 0xd0, 0x53,
	0x38, 0x12, 0xf5, 0x6a, 0x95, 0xdd, 0x05, 0x4e, 0xcf, 0xea, 0x7b, 0xb8, 0xa9, 0xc2, 0xb7, 0xe0,
	0xcd, 0x09, 0x27, 0x05, 0x95, 0x94, 0x23, 0x04, 0x4e, 0x49, 0x0a, 0xaa, 0xbb, 0x78, 0x58, 0x63,
	0x65, 0xb7, 0x21, 0x79, 0x4d, 0x75, 0x13, 0x0f, 0x9b, 0x22, 0xfc, 0x65, 0x41, 0x1b, 0xd3, 0x6f,
	0x35, 0x15, 0x12, 0x9d, 0x41, 0x57, 0xb0, 0x9a, 0x27, 0x74, 0x71, 0xef, 0x67, 0x30, 0x54, 0xac,
	0x2c, 0xce, 0xa0, 0xcb, 0x6a, 0x59, 0xd5, 0x72, 0x51, 0x11, 0xb9, 0x6e, 0x8c, 0xc0, 0x50, 0x73,
	0x22, 0xd7, 0xe8, 0x1d, 0x40, 0xb5, 0x1b, 0x42, 0x04, 0x76, 0xcf, 0xee, 0x77, 0x87, 0x2f, 0xa2,
	0xbf, 0xc2, 0x8a, 0xf6, 0x93, 0xe2, 0x7b, 0x7a, 0x74, 0x09, 0x7e, 0xc2, 0x8a, 0x2a, 0xcb, 0x29,
	0x5f, 0x6c, 0x4c, 0x60, 0x7a, 0xc9, 0xee, 0xf0, 0xf4, 0x80, 0x47, 0x13, 0x29, 0x7e, 0xbc, 0xfb,
	0x67, 0x97, 0xf1, 0x2b, 0x38, 0x2a, 0xd8, 0x92, 0xe6, 0x22, 0x70, 0xf5, 0x00, 0x27, 0x91, 0x79,
	0x9a, 0x68, 0xf7, 0x34, 0xd1, 0xa8, 0xdc, 0xe2, 0x46, 0x13, 0xfe, 0xb4, 0xa0, 0xfd, 0x81, 0x0a,
	0x41, 0x52, 0x8a, 0x2e, 0xc0, 0xcd, 0xe9, 0x86, 0xe6, 0x7a, 0xf5, 0xe3, 0x61, 0xef, 0x40, 0xd7,
	0x46, 0x1a, 0x5d, 0x2b, 0x1d, 0x36, 0x72, 0x15, 0x77, 0xc2, 0x96, 0xbb, 0x64, 0x35, 0x56, 0x9c,
	0xa4, 0x77, 0x52, 0x3f, 0x9e, 0x87, 0x35, 0x56, 0xdc, 0x57, 0xba, 0x15, 0x81, 0xd3, 0xb3, 0x15,
	0xa7, 0x70, 0x38, 0x02, 0x57, 0x7b, 0xa1, 0x2e, 0xb4, 0x3f, 0xc5, 0xef, 0xe3, 0xd9, 0x4d, 0xec,
	0x3f, 0x40, 0x1d, 0x70, 0xae, 0xe2, 0xc9, 0xcc, 0xb7, 0x14, 0x7d, 0x33, 0xc2, 0xf1, 0x55, 0x3c,
	0xf5, 0x5b, 0xc8, 0x03, 0xf7, 0x12, 0xe3, 0x19, 0xf6, 0x6d, 0x05, 0x27, 0xa3, 0x8f, 0xa3, 0x6b,
	0xdf, 0x09, 0xc7, 0xd0, 0x69, 0xc6, 0x12, 0xe8, 0x02, 0x3a, 0x45, 0x83, 0x03, 0x4b, 0xaf, 0x7f,
	0xfa, 0xef, 0x2d, 0xf0, 0x5e, 0x1b, 0x7e, 0xb7, 0xa0, 0x83, 0xa9, 0xa8, 0x58, 0x29, 0xa8, 0xba,
	0x31, 0xca, 0x39, 0xe3, 0xc6, 0xc2, 0xc3, 0x4d, 0x85, 0x5e, 0x83, 0xbb, 0xca, 0x72, 0x2a, 0x82,
	0x96, 0x76, 0x7e, 0x76, 0xc0, 0x79, 0x92, 0xe5, 0x14, 0x1b, 0xd5, 0x1f, 0xb3, 0xd8, 0xff, 0x31,
	0x4b, 0x04, 0x8e, 0xb2, 0x39, 0x78, 0xc5, 0x08, 0x9c, 0x25, 0x91, 0x44, 0x47, 0xfd, 0x10, 0x6b,
	0x3c, 0x7e, 0x09, 0xc7, 0x8c, 0xa7, 0x7b, 0xeb, 0xcd, 0xf9, 0xf8, 0xd1, 0xd4, 0xe0, 0xb9, 0xee,
	0x32, 0xb7, 0x7e, 0xb4, 0xec, 0x69, 0x3c, 0xbb, 0x3d, 0xd2, 0x17, 0xf0, 0xe6, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0xa0, 0x65, 0xe7, 0xd5, 0x03, 0x00, 0x00,
}
