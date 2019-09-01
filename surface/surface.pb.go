// Code generated by protoc-gen-go. DO NOT EDIT.
// source: surface.proto

package surface_v1

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

type FieldKind int32

const (
	FieldKind_SCALAR    FieldKind = 0
	FieldKind_MAP       FieldKind = 1
	FieldKind_ARRAY     FieldKind = 2
	FieldKind_REFERENCE FieldKind = 3
	FieldKind_ANY       FieldKind = 4
)

var FieldKind_name = map[int32]string{
	0: "SCALAR",
	1: "MAP",
	2: "ARRAY",
	3: "REFERENCE",
	4: "ANY",
}

var FieldKind_value = map[string]int32{
	"SCALAR":    0,
	"MAP":       1,
	"ARRAY":     2,
	"REFERENCE": 3,
	"ANY":       4,
}

func (x FieldKind) String() string {
	return proto.EnumName(FieldKind_name, int32(x))
}

func (FieldKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{0}
}

type TypeKind int32

const (
	TypeKind_STRUCT TypeKind = 0
	TypeKind_OBJECT TypeKind = 1
)

var TypeKind_name = map[int32]string{
	0: "STRUCT",
	1: "OBJECT",
}

var TypeKind_value = map[string]int32{
	"STRUCT": 0,
	"OBJECT": 1,
}

func (x TypeKind) String() string {
	return proto.EnumName(TypeKind_name, int32(x))
}

func (TypeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{1}
}

type Position int32

const (
	Position_BODY     Position = 0
	Position_HEADER   Position = 1
	Position_FORMDATA Position = 2
	Position_QUERY    Position = 3
	Position_PATH     Position = 4
)

var Position_name = map[int32]string{
	0: "BODY",
	1: "HEADER",
	2: "FORMDATA",
	3: "QUERY",
	4: "PATH",
}

var Position_value = map[string]int32{
	"BODY":     0,
	"HEADER":   1,
	"FORMDATA": 2,
	"QUERY":    3,
	"PATH":     4,
}

func (x Position) String() string {
	return proto.EnumName(Position_name, int32(x))
}

func (Position) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{2}
}

// Field is a field in a definition and can be associated with
// a position in a request structure.
type Field struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Kind                 FieldKind `protobuf:"varint,3,opt,name=kind,proto3,enum=surface.v1.FieldKind" json:"kind,omitempty"`
	Format               string    `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Position             Position  `protobuf:"varint,5,opt,name=position,proto3,enum=surface.v1.Position" json:"position,omitempty"`
	NativeType           string    `protobuf:"bytes,6,opt,name=nativeType,proto3" json:"nativeType,omitempty"`
	FieldName            string    `protobuf:"bytes,7,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	ParameterName        string    `protobuf:"bytes,8,opt,name=parameterName,proto3" json:"parameterName,omitempty"`
	Serialize            bool      `protobuf:"varint,9,opt,name=serialize,proto3" json:"serialize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetKind() FieldKind {
	if m != nil {
		return m.Kind
	}
	return FieldKind_SCALAR
}

func (m *Field) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Field) GetPosition() Position {
	if m != nil {
		return m.Position
	}
	return Position_BODY
}

func (m *Field) GetNativeType() string {
	if m != nil {
		return m.NativeType
	}
	return ""
}

func (m *Field) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Field) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

func (m *Field) GetSerialize() bool {
	if m != nil {
		return m.Serialize
	}
	return false
}

// Type typically corresponds to a definition, parameter, or response
// in an API and is represented by a type in generated code.
type Type struct {
	Name                   string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                   TypeKind    `protobuf:"varint,2,opt,name=kind,proto3,enum=surface.v1.TypeKind" json:"kind,omitempty"`
	Description            string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ContentType            string      `protobuf:"bytes,4,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Fields                 []*Field    `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	TypeName               string      `protobuf:"bytes,6,opt,name=typeName,proto3" json:"typeName,omitempty"`
	SpecificationExtension []*NamedAny `protobuf:"bytes,7,rep,name=specification_extension,json=specificationExtension,proto3" json:"specification_extension,omitempty"`
	TypeSource             string      `protobuf:"bytes,100,opt,name=typeSource,proto3" json:"typeSource,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}    `json:"-"`
	XXX_unrecognized       []byte      `json:"-"`
	XXX_sizecache          int32       `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{1}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetKind() TypeKind {
	if m != nil {
		return m.Kind
	}
	return TypeKind_STRUCT
}

func (m *Type) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Type) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *Type) GetSpecificationExtension() []*NamedAny {
	if m != nil {
		return m.SpecificationExtension
	}
	return nil
}

func (m *Type) GetTypeSource() string {
	if m != nil {
		return m.TypeSource
	}
	return ""
}

// Method is an operation of an API and typically has associated client and server code.
type Method struct {
	Operation              string      `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Path                   string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Method                 string      `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Description            string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Name                   string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	HandlerName            string      `protobuf:"bytes,6,opt,name=handlerName,proto3" json:"handlerName,omitempty"`
	ProcessorName          string      `protobuf:"bytes,7,opt,name=processorName,proto3" json:"processorName,omitempty"`
	ClientName             string      `protobuf:"bytes,8,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ParametersTypeName     string      `protobuf:"bytes,9,opt,name=parametersTypeName,proto3" json:"parametersTypeName,omitempty"`
	ResponsesTypeName      string      `protobuf:"bytes,10,opt,name=responsesTypeName,proto3" json:"responsesTypeName,omitempty"`
	SpecificationExtension []*NamedAny `protobuf:"bytes,100,rep,name=specification_extension,json=specificationExtension,proto3" json:"specification_extension,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}    `json:"-"`
	XXX_unrecognized       []byte      `json:"-"`
	XXX_sizecache          int32       `json:"-"`
}

func (m *Method) Reset()         { *m = Method{} }
func (m *Method) String() string { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()    {}
func (*Method) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{2}
}

func (m *Method) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method.Unmarshal(m, b)
}
func (m *Method) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method.Marshal(b, m, deterministic)
}
func (m *Method) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method.Merge(m, src)
}
func (m *Method) XXX_Size() int {
	return xxx_messageInfo_Method.Size(m)
}
func (m *Method) XXX_DiscardUnknown() {
	xxx_messageInfo_Method.DiscardUnknown(m)
}

var xxx_messageInfo_Method proto.InternalMessageInfo

func (m *Method) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Method) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Method) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Method) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Method) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Method) GetHandlerName() string {
	if m != nil {
		return m.HandlerName
	}
	return ""
}

func (m *Method) GetProcessorName() string {
	if m != nil {
		return m.ProcessorName
	}
	return ""
}

func (m *Method) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Method) GetParametersTypeName() string {
	if m != nil {
		return m.ParametersTypeName
	}
	return ""
}

func (m *Method) GetResponsesTypeName() string {
	if m != nil {
		return m.ResponsesTypeName
	}
	return ""
}

func (m *Method) GetSpecificationExtension() []*NamedAny {
	if m != nil {
		return m.SpecificationExtension
	}
	return nil
}

// Model represents an API for code generation.
type Model struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Types                []*Type   `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	Methods              []*Method `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	CommandStreamTypes   string    `protobuf:"bytes,100,opt,name=command_stream_types,json=commandStreamTypes,proto3" json:"command_stream_types,omitempty"`
	FactsStreamTypes     string    `protobuf:"bytes,101,opt,name=facts_stream_types,json=factsStreamTypes,proto3" json:"facts_stream_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{3}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetTypes() []*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Model) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Model) GetCommandStreamTypes() string {
	if m != nil {
		return m.CommandStreamTypes
	}
	return ""
}

func (m *Model) GetFactsStreamTypes() string {
	if m != nil {
		return m.FactsStreamTypes
	}
	return ""
}

// Automatically-generated message used to represent maps of Any as ordered (name,value) pairs.
type NamedAny struct {
	// Map key
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Mapped value
	Value                *Any     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedAny) Reset()         { *m = NamedAny{} }
func (m *NamedAny) String() string { return proto.CompactTextString(m) }
func (*NamedAny) ProtoMessage()    {}
func (*NamedAny) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{4}
}

func (m *NamedAny) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedAny.Unmarshal(m, b)
}
func (m *NamedAny) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedAny.Marshal(b, m, deterministic)
}
func (m *NamedAny) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedAny.Merge(m, src)
}
func (m *NamedAny) XXX_Size() int {
	return xxx_messageInfo_NamedAny.Size(m)
}
func (m *NamedAny) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedAny.DiscardUnknown(m)
}

var xxx_messageInfo_NamedAny proto.InternalMessageInfo

func (m *NamedAny) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedAny) GetValue() *Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type Any struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Yaml                 string   `protobuf:"bytes,2,opt,name=yaml,proto3" json:"yaml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
func (*Any) Descriptor() ([]byte, []int) {
	return fileDescriptor_3306d3fe7fd759d5, []int{5}
}

func (m *Any) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Any.Unmarshal(m, b)
}
func (m *Any) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Any.Marshal(b, m, deterministic)
}
func (m *Any) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Any.Merge(m, src)
}
func (m *Any) XXX_Size() int {
	return xxx_messageInfo_Any.Size(m)
}
func (m *Any) XXX_DiscardUnknown() {
	xxx_messageInfo_Any.DiscardUnknown(m)
}

var xxx_messageInfo_Any proto.InternalMessageInfo

func (m *Any) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Any) GetYaml() string {
	if m != nil {
		return m.Yaml
	}
	return ""
}

func init() {
	proto.RegisterEnum("surface.v1.FieldKind", FieldKind_name, FieldKind_value)
	proto.RegisterEnum("surface.v1.TypeKind", TypeKind_name, TypeKind_value)
	proto.RegisterEnum("surface.v1.Position", Position_name, Position_value)
	proto.RegisterType((*Field)(nil), "surface.v1.Field")
	proto.RegisterType((*Type)(nil), "surface.v1.Type")
	proto.RegisterType((*Method)(nil), "surface.v1.Method")
	proto.RegisterType((*Model)(nil), "surface.v1.Model")
	proto.RegisterType((*NamedAny)(nil), "surface.v1.NamedAny")
	proto.RegisterType((*Any)(nil), "surface.v1.Any")
}

func init() { proto.RegisterFile("surface.proto", fileDescriptor_3306d3fe7fd759d5) }

var fileDescriptor_3306d3fe7fd759d5 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0xc5, 0xb1, 0x9d, 0x38, 0x93, 0xcb, 0xbd, 0x66, 0xc5, 0xe5, 0xfa, 0xa2, 0xaa, 0x8a, 0xa2,
	0xb6, 0x0a, 0x11, 0x0a, 0x94, 0x7e, 0x81, 0x09, 0x46, 0xa8, 0x6d, 0x80, 0x2e, 0xe6, 0x21, 0x4f,
	0x68, 0xb1, 0x37, 0x60, 0xd5, 0xf6, 0x5a, 0x5e, 0x07, 0x35, 0xfd, 0x97, 0xbe, 0xf6, 0x33, 0xfa,
	0x0f, 0xfd, 0xa3, 0x6a, 0x77, 0xed, 0xc4, 0x21, 0x51, 0xa5, 0xbe, 0xed, 0x9e, 0x39, 0x7b, 0x66,
	0xe7, 0xcc, 0x68, 0x60, 0x9b, 0xcf, 0xf2, 0x29, 0x09, 0xe8, 0x30, 0xcb, 0x59, 0xc1, 0x10, 0x54,
	0xd7, 0xa7, 0xb7, 0xfb, 0xff, 0x3f, 0x30, 0xf6, 0x10, 0xd3, 0x23, 0x19, 0xb9, 0x9f, 0x4d, 0x8f,
	0x48, 0x3a, 0x57, 0xb4, 0xde, 0xf7, 0x06, 0x98, 0xe7, 0x11, 0x8d, 0x43, 0x84, 0xc0, 0x48, 0x49,
	0x42, 0x1d, 0xad, 0xab, 0xf5, 0xdb, 0x58, 0x9e, 0x05, 0x56, 0xcc, 0x33, 0xea, 0x34, 0x14, 0x26,
	0xce, 0xe8, 0x00, 0x8c, 0xcf, 0x51, 0x1a, 0x3a, 0x7a, 0x57, 0xeb, 0xff, 0x7d, 0xf2, 0xef, 0x70,
	0x99, 0x67, 0x28, 0x85, 0x3e, 0x44, 0x69, 0x88, 0x25, 0x05, 0xed, 0x41, 0x73, 0xca, 0xf2, 0x84,
	0x14, 0x8e, 0x21, 0x05, 0xca, 0x1b, 0x3a, 0x06, 0x2b, 0x63, 0x3c, 0x2a, 0x22, 0x96, 0x3a, 0xa6,
	0x94, 0xd9, 0xad, 0xcb, 0x5c, 0x97, 0x31, 0xbc, 0x60, 0xa1, 0x97, 0x00, 0x29, 0x29, 0xa2, 0x27,
	0xea, 0x8b, 0xef, 0x34, 0xa5, 0x5a, 0x0d, 0x41, 0x2f, 0xa0, 0x3d, 0x15, 0xc9, 0x2f, 0x45, 0x05,
	0x2d, 0x19, 0x5e, 0x02, 0xe8, 0x15, 0x6c, 0x67, 0x24, 0x27, 0x09, 0x2d, 0x68, 0x2e, 0x19, 0x96,
	0x64, 0xac, 0x82, 0x42, 0x83, 0xd3, 0x3c, 0x22, 0x71, 0xf4, 0x95, 0x3a, 0xed, 0xae, 0xd6, 0xb7,
	0xf0, 0x12, 0xe8, 0xfd, 0x68, 0x80, 0x21, 0x53, 0x6d, 0xf2, 0xa9, 0x5f, 0x7a, 0xd2, 0x58, 0x2f,
	0x46, 0xbc, 0xa9, 0x59, 0xd2, 0x85, 0x4e, 0x48, 0x79, 0x90, 0x47, 0x99, 0xac, 0x5e, 0x97, 0x22,
	0x75, 0x48, 0x30, 0x02, 0x96, 0x16, 0x34, 0x2d, 0x64, 0xad, 0xca, 0xb9, 0x3a, 0x84, 0x0e, 0xa0,
	0x29, 0x6b, 0xe3, 0x8e, 0xd9, 0xd5, 0xfb, 0x9d, 0x93, 0x9d, 0xb5, 0x1e, 0xe0, 0x92, 0x80, 0xf6,
	0xc1, 0x12, 0x4d, 0x93, 0x45, 0x2b, 0xd7, 0x16, 0x77, 0x34, 0x86, 0xff, 0x78, 0x46, 0x83, 0x68,
	0x1a, 0x05, 0x44, 0x64, 0xbe, 0xa3, 0x5f, 0x0a, 0x9a, 0x72, 0xf1, 0xad, 0x96, 0xd4, 0x5d, 0xa9,
	0x43, 0x3c, 0x09, 0xdd, 0x74, 0x8e, 0xf7, 0x56, 0x1e, 0x79, 0xd5, 0x1b, 0xd1, 0x22, 0x21, 0x7d,
	0xc3, 0x66, 0x79, 0x40, 0x9d, 0x50, 0xb5, 0x68, 0x89, 0xf4, 0xbe, 0xe9, 0xd0, 0x1c, 0xd3, 0xe2,
	0x91, 0x85, 0xc2, 0x69, 0x96, 0xd1, 0x5c, 0x0a, 0x94, 0x3e, 0x2e, 0x01, 0x61, 0x70, 0x46, 0x8a,
	0xc7, 0x6a, 0xe8, 0xc4, 0x59, 0x4c, 0x52, 0x22, 0xdf, 0x96, 0x8e, 0x95, 0xb7, 0xe7, 0x76, 0x1a,
	0xeb, 0x76, 0x56, 0xed, 0x32, 0x6b, 0xed, 0xea, 0x42, 0xe7, 0x91, 0xa4, 0x61, 0x5c, 0x4e, 0x83,
	0x32, 0xa6, 0x0e, 0xc9, 0x89, 0xc9, 0x59, 0x40, 0x39, 0x67, 0x79, 0x6d, 0xa6, 0x56, 0x41, 0x51,
	0x72, 0x10, 0x47, 0x34, 0x2d, 0x6a, 0x43, 0x55, 0x43, 0xd0, 0x10, 0xd0, 0x62, 0xc4, 0xb8, 0x5f,
	0xf5, 0xa1, 0x2d, 0x79, 0x1b, 0x22, 0xe8, 0x10, 0x76, 0x72, 0xca, 0x33, 0x96, 0x72, 0xba, 0xa4,
	0x83, 0xa4, 0xaf, 0x07, 0x7e, 0xd7, 0xbf, 0xf0, 0xcf, 0xfb, 0xd7, 0xfb, 0xa9, 0x81, 0x39, 0x66,
	0x21, 0x8d, 0x37, 0x4e, 0xf8, 0x1b, 0x30, 0x45, 0x2f, 0xb9, 0xd3, 0x90, 0xd2, 0xf6, 0xf3, 0x11,
	0xc7, 0x2a, 0x8c, 0x0e, 0xa1, 0xa5, 0x5a, 0xc3, 0x1d, 0x5d, 0x32, 0x51, 0x9d, 0xa9, 0xfa, 0x8f,
	0x2b, 0x0a, 0x3a, 0x86, 0xdd, 0x80, 0x25, 0x09, 0x49, 0xc3, 0x3b, 0x5e, 0xe4, 0x94, 0x24, 0x77,
	0x2a, 0x89, 0x9a, 0x1e, 0x54, 0xc6, 0x6e, 0x64, 0xc8, 0x2f, 0xf5, 0xd1, 0x94, 0x04, 0x05, 0x5f,
	0xe5, 0x53, 0xc9, 0xb7, 0x65, 0xa4, 0xc6, 0xee, 0x79, 0x60, 0x55, 0x75, 0x6f, 0xac, 0xea, 0x35,
	0x98, 0x4f, 0x24, 0x9e, 0xa9, 0x05, 0xd7, 0x39, 0xf9, 0xa7, 0xfe, 0x57, 0xe1, 0x95, 0x8a, 0xf6,
	0x3c, 0xd0, 0x85, 0xc2, 0xa0, 0x62, 0x6b, 0x92, 0xbd, 0x3b, 0x54, 0x6b, 0x75, 0x58, 0xad, 0xd5,
	0xda, 0x13, 0x91, 0x6d, 0x4e, 0x92, 0xb8, 0x1a, 0x62, 0x71, 0x1e, 0x8c, 0xa0, 0xbd, 0xd8, 0x90,
	0x08, 0xa0, 0x79, 0x33, 0x72, 0x3f, 0xba, 0xd8, 0xde, 0x42, 0x2d, 0xd0, 0xc7, 0xee, 0xb5, 0xad,
	0xa1, 0x36, 0x98, 0x2e, 0xc6, 0xee, 0xc4, 0x6e, 0xa0, 0x6d, 0x68, 0x63, 0xef, 0xdc, 0xc3, 0xde,
	0xe5, 0xc8, 0xb3, 0x75, 0x41, 0x71, 0x2f, 0x27, 0xb6, 0x31, 0xe8, 0x81, 0x55, 0xad, 0x14, 0xa9,
	0xe1, 0xe3, 0xdb, 0x91, 0x6f, 0x6f, 0x89, 0xf3, 0xd5, 0xe9, 0x7b, 0x6f, 0xe4, 0xdb, 0xda, 0x60,
	0x04, 0x56, 0xb5, 0x43, 0x91, 0x05, 0xc6, 0xe9, 0xd5, 0xd9, 0x44, 0x31, 0x2e, 0x3c, 0xf7, 0xcc,
	0xc3, 0xb6, 0x86, 0xfe, 0x02, 0xeb, 0xfc, 0x0a, 0x8f, 0xcf, 0x5c, 0xdf, 0xb5, 0x1b, 0x22, 0xed,
	0xa7, 0x5b, 0x0f, 0x4f, 0x6c, 0x5d, 0xd0, 0xaf, 0x5d, 0xff, 0xc2, 0x36, 0xee, 0x9b, 0xb2, 0xac,
	0x77, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x61, 0x6f, 0x62, 0x60, 0x58, 0x06, 0x00, 0x00,
}
