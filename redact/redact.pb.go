// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.1
// source: redact/redact.proto

// Package redact provides interfaces and methods to help implement redaction.

package redact

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// FieldRules encapsulates options to change the redacted values of any type of field.
// Depending on the field, the correct type value should be used.
type FieldRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// values for redacted field
	//
	// Types that are assignable to Values:
	//	*FieldRules_Float
	//	*FieldRules_Double
	//	*FieldRules_Int32
	//	*FieldRules_Int64
	//	*FieldRules_Uint32
	//	*FieldRules_Uint64
	//	*FieldRules_Sint32
	//	*FieldRules_Sint64
	//	*FieldRules_Fixed32
	//	*FieldRules_Fixed64
	//	*FieldRules_Sfixed32
	//	*FieldRules_Sfixed64
	//	*FieldRules_Bool
	//	*FieldRules_String_
	//	*FieldRules_Bytes
	//	*FieldRules_Enum
	//	*FieldRules_Message
	//	*FieldRules_Element
	Values isFieldRules_Values `protobuf_oneof:"values"`
}

func (x *FieldRules) Reset() {
	*x = FieldRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redact_redact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldRules) ProtoMessage() {}

func (x *FieldRules) ProtoReflect() protoreflect.Message {
	mi := &file_redact_redact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldRules.ProtoReflect.Descriptor instead.
func (*FieldRules) Descriptor() ([]byte, []int) {
	return file_redact_redact_proto_rawDescGZIP(), []int{0}
}

func (m *FieldRules) GetValues() isFieldRules_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

func (x *FieldRules) GetFloat() float32 {
	if x, ok := x.GetValues().(*FieldRules_Float); ok {
		return x.Float
	}
	return 0
}

func (x *FieldRules) GetDouble() float64 {
	if x, ok := x.GetValues().(*FieldRules_Double); ok {
		return x.Double
	}
	return 0
}

func (x *FieldRules) GetInt32() int32 {
	if x, ok := x.GetValues().(*FieldRules_Int32); ok {
		return x.Int32
	}
	return 0
}

func (x *FieldRules) GetInt64() int64 {
	if x, ok := x.GetValues().(*FieldRules_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *FieldRules) GetUint32() uint32 {
	if x, ok := x.GetValues().(*FieldRules_Uint32); ok {
		return x.Uint32
	}
	return 0
}

func (x *FieldRules) GetUint64() uint64 {
	if x, ok := x.GetValues().(*FieldRules_Uint64); ok {
		return x.Uint64
	}
	return 0
}

func (x *FieldRules) GetSint32() int32 {
	if x, ok := x.GetValues().(*FieldRules_Sint32); ok {
		return x.Sint32
	}
	return 0
}

func (x *FieldRules) GetSint64() int64 {
	if x, ok := x.GetValues().(*FieldRules_Sint64); ok {
		return x.Sint64
	}
	return 0
}

func (x *FieldRules) GetFixed32() uint32 {
	if x, ok := x.GetValues().(*FieldRules_Fixed32); ok {
		return x.Fixed32
	}
	return 0
}

func (x *FieldRules) GetFixed64() uint64 {
	if x, ok := x.GetValues().(*FieldRules_Fixed64); ok {
		return x.Fixed64
	}
	return 0
}

func (x *FieldRules) GetSfixed32() int32 {
	if x, ok := x.GetValues().(*FieldRules_Sfixed32); ok {
		return x.Sfixed32
	}
	return 0
}

func (x *FieldRules) GetSfixed64() int64 {
	if x, ok := x.GetValues().(*FieldRules_Sfixed64); ok {
		return x.Sfixed64
	}
	return 0
}

func (x *FieldRules) GetBool() bool {
	if x, ok := x.GetValues().(*FieldRules_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *FieldRules) GetString_() string {
	if x, ok := x.GetValues().(*FieldRules_String_); ok {
		return x.String_
	}
	return ""
}

func (x *FieldRules) GetBytes() []byte {
	if x, ok := x.GetValues().(*FieldRules_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *FieldRules) GetEnum() int32 {
	if x, ok := x.GetValues().(*FieldRules_Enum); ok {
		return x.Enum
	}
	return 0
}

func (x *FieldRules) GetMessage() *MessageRules {
	if x, ok := x.GetValues().(*FieldRules_Message); ok {
		return x.Message
	}
	return nil
}

func (x *FieldRules) GetElement() *ElementRules {
	if x, ok := x.GetValues().(*FieldRules_Element); ok {
		return x.Element
	}
	return nil
}

type isFieldRules_Values interface {
	isFieldRules_Values()
}

type FieldRules_Float struct {
	// Scalar Field Types values
	Float float32 `protobuf:"fixed32,2,opt,name=float,proto3,oneof"`
}

type FieldRules_Double struct {
	Double float64 `protobuf:"fixed64,3,opt,name=double,proto3,oneof"`
}

type FieldRules_Int32 struct {
	Int32 int32 `protobuf:"varint,4,opt,name=int32,proto3,oneof"`
}

type FieldRules_Int64 struct {
	Int64 int64 `protobuf:"varint,5,opt,name=int64,proto3,oneof"`
}

type FieldRules_Uint32 struct {
	Uint32 uint32 `protobuf:"varint,6,opt,name=uint32,proto3,oneof"`
}

type FieldRules_Uint64 struct {
	Uint64 uint64 `protobuf:"varint,7,opt,name=uint64,proto3,oneof"`
}

type FieldRules_Sint32 struct {
	Sint32 int32 `protobuf:"zigzag32,8,opt,name=sint32,proto3,oneof"`
}

type FieldRules_Sint64 struct {
	Sint64 int64 `protobuf:"zigzag64,9,opt,name=sint64,proto3,oneof"`
}

type FieldRules_Fixed32 struct {
	Fixed32 uint32 `protobuf:"fixed32,10,opt,name=fixed32,proto3,oneof"`
}

type FieldRules_Fixed64 struct {
	Fixed64 uint64 `protobuf:"fixed64,11,opt,name=fixed64,proto3,oneof"`
}

type FieldRules_Sfixed32 struct {
	Sfixed32 int32 `protobuf:"fixed32,12,opt,name=sfixed32,proto3,oneof"`
}

type FieldRules_Sfixed64 struct {
	Sfixed64 int64 `protobuf:"fixed64,13,opt,name=sfixed64,proto3,oneof"`
}

type FieldRules_Bool struct {
	Bool bool `protobuf:"varint,14,opt,name=bool,proto3,oneof"`
}

type FieldRules_String_ struct {
	String_ string `protobuf:"bytes,15,opt,name=string,proto3,oneof"`
}

type FieldRules_Bytes struct {
	Bytes []byte `protobuf:"bytes,16,opt,name=bytes,proto3,oneof"`
}

type FieldRules_Enum struct {
	Enum int32 `protobuf:"varint,17,opt,name=enum,proto3,oneof"`
}

type FieldRules_Message struct {
	// Message defines rules for singular message type fields only
	Message *MessageRules `protobuf:"bytes,19,opt,name=message,proto3,oneof"`
}

type FieldRules_Element struct {
	// Element defines rules for repeated or map type fields
	Element *ElementRules `protobuf:"bytes,20,opt,name=element,proto3,oneof"`
}

func (*FieldRules_Float) isFieldRules_Values() {}

func (*FieldRules_Double) isFieldRules_Values() {}

func (*FieldRules_Int32) isFieldRules_Values() {}

func (*FieldRules_Int64) isFieldRules_Values() {}

func (*FieldRules_Uint32) isFieldRules_Values() {}

func (*FieldRules_Uint64) isFieldRules_Values() {}

func (*FieldRules_Sint32) isFieldRules_Values() {}

func (*FieldRules_Sint64) isFieldRules_Values() {}

func (*FieldRules_Fixed32) isFieldRules_Values() {}

func (*FieldRules_Fixed64) isFieldRules_Values() {}

func (*FieldRules_Sfixed32) isFieldRules_Values() {}

func (*FieldRules_Sfixed64) isFieldRules_Values() {}

func (*FieldRules_Bool) isFieldRules_Values() {}

func (*FieldRules_String_) isFieldRules_Values() {}

func (*FieldRules_Bytes) isFieldRules_Values() {}

func (*FieldRules_Enum) isFieldRules_Values() {}

func (*FieldRules_Message) isFieldRules_Values() {}

func (*FieldRules_Element) isFieldRules_Values() {}

// MessageRules describe the constraints applied to embedded message for redaction.
// For message-type fields, rules are performed recursively.
type MessageRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Skip specifies that the redaction rules of this field should not be evaluated
	Skip bool `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	// Empty specifies that after redaction this field should be set to empty object
	// instead of nil
	Empty bool `protobuf:"varint,2,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *MessageRules) Reset() {
	*x = MessageRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redact_redact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRules) ProtoMessage() {}

func (x *MessageRules) ProtoReflect() protoreflect.Message {
	mi := &file_redact_redact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRules.ProtoReflect.Descriptor instead.
func (*MessageRules) Descriptor() ([]byte, []int) {
	return file_redact_redact_proto_rawDescGZIP(), []int{1}
}

func (x *MessageRules) GetSkip() bool {
	if x != nil {
		return x.Skip
	}
	return false
}

func (x *MessageRules) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

// ElementRules describe the constraints applied to `repeated` or `map` values
type ElementRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Empty specifies that after redaction the list/map will be empty
	Empty bool `protobuf:"varint,1,opt,name=empty,proto3" json:"empty,omitempty"`
	// Nested specifies that default rules is to be applied on each item of map/list
	Nested bool `protobuf:"varint,2,opt,name=nested,proto3" json:"nested,omitempty"`
	// Item specifies the redaction rules to be applied to each item in map/list.
	Item *FieldRules `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *ElementRules) Reset() {
	*x = ElementRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redact_redact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElementRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElementRules) ProtoMessage() {}

func (x *ElementRules) ProtoReflect() protoreflect.Message {
	mi := &file_redact_redact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElementRules.ProtoReflect.Descriptor instead.
func (*ElementRules) Descriptor() ([]byte, []int) {
	return file_redact_redact_proto_rawDescGZIP(), []int{2}
}

func (x *ElementRules) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

func (x *ElementRules) GetNested() bool {
	if x != nil {
		return x.Nested
	}
	return false
}

func (x *ElementRules) GetItem() *FieldRules {
	if x != nil {
		return x.Item
	}
	return nil
}

var file_redact_redact_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         90102,
		Name:          "redact.file_skip",
		Tag:           "varint,90102,opt,name=file_skip",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54123,
		Name:          "redact.service_skip",
		Tag:           "varint,54123,opt,name=service_skip",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54124,
		Name:          "redact.internal_service",
		Tag:           "varint,54124,opt,name=internal_service",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         54125,
		Name:          "redact.internal_service_code",
		Tag:           "varint,54125,opt,name=internal_service_code",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54126,
		Name:          "redact.internal_service_err_message",
		Tag:           "bytes,54126,opt,name=internal_service_err_message",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54123,
		Name:          "redact.method_skip",
		Tag:           "varint,54123,opt,name=method_skip",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54124,
		Name:          "redact.internal_method",
		Tag:           "varint,54124,opt,name=internal_method",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         54125,
		Name:          "redact.internal_method_code",
		Tag:           "varint,54125,opt,name=internal_method_code",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54126,
		Name:          "redact.internal_method_err_message",
		Tag:           "bytes,54126,opt,name=internal_method_err_message",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54123,
		Name:          "redact.nil",
		Tag:           "varint,54123,opt,name=nil",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54124,
		Name:          "redact.empty",
		Tag:           "varint,54124,opt,name=empty",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54125,
		Name:          "redact.ignored",
		Tag:           "varint,54125,opt,name=ignored",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54123,
		Name:          "redact.redact",
		Tag:           "varint,54123,opt,name=redact",
		Filename:      "redact/redact.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldRules)(nil),
		Field:         54124,
		Name:          "redact.custom",
		Tag:           "bytes,54124,opt,name=custom",
		Filename:      "redact/redact.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// FileSkip is used to skip generation of any redaction for proto file
	//
	// optional bool file_skip = 90102;
	E_FileSkip = &file_redact_redact_proto_extTypes[0]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// ServiceSkip is used to skip the redaction in grpc service in the server
	//
	// optional bool service_skip = 54123;
	E_ServiceSkip = &file_redact_redact_proto_extTypes[1]
	// InternalService will make this service private and client will not be
	// able to receive any response for any of it's method, (unless skipped
	// explicitly) and will get PermissionDenied(7) error by default, to set
	// any other code set it in InternalServiceCode, it should be one of the
	// defined GRPC status code, and InternalServiceErrMessage for error
	// message, in which, one can use `%service%` or `%method%` tags to include
	// corresponding service name or method name, respectively.
	//
	// optional bool internal_service = 54124;
	E_InternalService = &file_redact_redact_proto_extTypes[2]
	// optional uint32 internal_service_code = 54125;
	E_InternalServiceCode = &file_redact_redact_proto_extTypes[3]
	// optional string internal_service_err_message = 54126;
	E_InternalServiceErrMessage = &file_redact_redact_proto_extTypes[4]
)

// Extension fields to descriptor.MethodOptions.
var (
	// MethodSkip is used to skip the redactions for this method in the grpc server
	//
	// optional bool method_skip = 54123;
	E_MethodSkip = &file_redact_redact_proto_extTypes[5]
	// InternalMethod, InternalMethodCode and InternalMethodErrMessage works same
	// as that of service level options: InternalService, InternalServiceCode and
	// InternalServiceErrMessage, but at Method level. All the validations and
	// functionalities are same for both. But the method level options will be used
	// whenever both are specified.
	//
	// optional bool internal_method = 54124;
	E_InternalMethod = &file_redact_redact_proto_extTypes[6]
	// optional uint32 internal_method_code = 54125;
	E_InternalMethodCode = &file_redact_redact_proto_extTypes[7]
	// optional string internal_method_err_message = 54126;
	E_InternalMethodErrMessage = &file_redact_redact_proto_extTypes[8]
)

// Extension fields to descriptor.MessageOptions.
var (
	// Nil will redact message to nil (can be override by field level, `empty` option)
	//
	// optional bool nil = 54123;
	E_Nil = &file_redact_redact_proto_extTypes[9]
	// Empty will redact message to it's empty object
	//
	// optional bool empty = 54124;
	E_Empty = &file_redact_redact_proto_extTypes[10]
	// Ignored skips generation of any redaction for this message.
	//
	// optional bool ignored = 54125;
	E_Ignored = &file_redact_redact_proto_extTypes[11]
)

// Extension fields to descriptor.FieldOptions.
var (
	// Redact: it redact the field with predefined defaults:
	//  * `0` for any number type
	//  * `"REDACTED"` for string type
	//  * `nil` for byte type
	//  * `0th value` for enum type
	//  * `nil` for message type
	//  * `nil` map for map type
	//  * for repeated field type, default rules are applied `recursively` to each item
	//
	// optional bool redact = 54123;
	E_Redact = &file_redact_redact_proto_extTypes[12]
	// Custom: specify the different values to be used for redaction on this field. By
	// default, if Custom value is not defined Redact should be true to apply redaction.
	// And if Custom value is to be assigned, one can skip the Redact field.
	//
	// optional redact.FieldRules custom = 54124;
	E_Custom = &file_redact_redact_proto_extTypes[13]
)

var File_redact_redact_proto protoreflect.FileDescriptor

var file_redact_redact_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x96, 0x04, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x18, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x06, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x06, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x12, 0x1c, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0f, 0x48, 0x00, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1c,
	0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x10,
	0x48, 0x00, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x64, 0x61, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x07,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x64, 0x0a, 0x0c, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x12, 0x26, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x3a, 0x3b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf6, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x6b, 0x69, 0x70, 0x3a, 0x44, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x3a, 0x4c, 0x0a, 0x10, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xec, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x55, 0x0a, 0x15, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xed, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x3a, 0x62, 0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xee, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x41, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x73,
	0x6b, 0x69, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x53, 0x6b, 0x69, 0x70, 0x3a, 0x49, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0xa6, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x3a, 0x52, 0x0a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xed, 0xa6, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x5f, 0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xee, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x72, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x33, 0x0a, 0x03, 0x6e, 0x69, 0x6c, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xeb, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6e, 0x69, 0x6c, 0x3a, 0x37, 0x0a, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x3a, 0x3b, 0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xed, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x64, 0x3a, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xa6, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x3a, 0x4b, 0x0a, 0x06, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72,
	0x65, 0x64, 0x61, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x69, 0x76, 0x61, 0x6d, 0x30, 0x31, 0x30,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x72, 0x65, 0x64, 0x61,
	0x63, 0x74, 0x2f, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74, 0x3b, 0x72, 0x65, 0x64, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redact_redact_proto_rawDescOnce sync.Once
	file_redact_redact_proto_rawDescData = file_redact_redact_proto_rawDesc
)

func file_redact_redact_proto_rawDescGZIP() []byte {
	file_redact_redact_proto_rawDescOnce.Do(func() {
		file_redact_redact_proto_rawDescData = protoimpl.X.CompressGZIP(file_redact_redact_proto_rawDescData)
	})
	return file_redact_redact_proto_rawDescData
}

var file_redact_redact_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_redact_redact_proto_goTypes = []interface{}{
	(*FieldRules)(nil),                // 0: redact.FieldRules
	(*MessageRules)(nil),              // 1: redact.MessageRules
	(*ElementRules)(nil),              // 2: redact.ElementRules
	(*descriptor.FileOptions)(nil),    // 3: google.protobuf.FileOptions
	(*descriptor.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 5: google.protobuf.MethodOptions
	(*descriptor.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
}
var file_redact_redact_proto_depIdxs = []int32{
	1,  // 0: redact.FieldRules.message:type_name -> redact.MessageRules
	2,  // 1: redact.FieldRules.element:type_name -> redact.ElementRules
	0,  // 2: redact.ElementRules.item:type_name -> redact.FieldRules
	3,  // 3: redact.file_skip:extendee -> google.protobuf.FileOptions
	4,  // 4: redact.service_skip:extendee -> google.protobuf.ServiceOptions
	4,  // 5: redact.internal_service:extendee -> google.protobuf.ServiceOptions
	4,  // 6: redact.internal_service_code:extendee -> google.protobuf.ServiceOptions
	4,  // 7: redact.internal_service_err_message:extendee -> google.protobuf.ServiceOptions
	5,  // 8: redact.method_skip:extendee -> google.protobuf.MethodOptions
	5,  // 9: redact.internal_method:extendee -> google.protobuf.MethodOptions
	5,  // 10: redact.internal_method_code:extendee -> google.protobuf.MethodOptions
	5,  // 11: redact.internal_method_err_message:extendee -> google.protobuf.MethodOptions
	6,  // 12: redact.nil:extendee -> google.protobuf.MessageOptions
	6,  // 13: redact.empty:extendee -> google.protobuf.MessageOptions
	6,  // 14: redact.ignored:extendee -> google.protobuf.MessageOptions
	7,  // 15: redact.redact:extendee -> google.protobuf.FieldOptions
	7,  // 16: redact.custom:extendee -> google.protobuf.FieldOptions
	0,  // 17: redact.custom:type_name -> redact.FieldRules
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	17, // [17:18] is the sub-list for extension type_name
	3,  // [3:17] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_redact_redact_proto_init() }
func file_redact_redact_proto_init() {
	if File_redact_redact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redact_redact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldRules); i {
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
		file_redact_redact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRules); i {
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
		file_redact_redact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElementRules); i {
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
	file_redact_redact_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FieldRules_Float)(nil),
		(*FieldRules_Double)(nil),
		(*FieldRules_Int32)(nil),
		(*FieldRules_Int64)(nil),
		(*FieldRules_Uint32)(nil),
		(*FieldRules_Uint64)(nil),
		(*FieldRules_Sint32)(nil),
		(*FieldRules_Sint64)(nil),
		(*FieldRules_Fixed32)(nil),
		(*FieldRules_Fixed64)(nil),
		(*FieldRules_Sfixed32)(nil),
		(*FieldRules_Sfixed64)(nil),
		(*FieldRules_Bool)(nil),
		(*FieldRules_String_)(nil),
		(*FieldRules_Bytes)(nil),
		(*FieldRules_Enum)(nil),
		(*FieldRules_Message)(nil),
		(*FieldRules_Element)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redact_redact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 14,
			NumServices:   0,
		},
		GoTypes:           file_redact_redact_proto_goTypes,
		DependencyIndexes: file_redact_redact_proto_depIdxs,
		MessageInfos:      file_redact_redact_proto_msgTypes,
		ExtensionInfos:    file_redact_redact_proto_extTypes,
	}.Build()
	File_redact_redact_proto = out.File
	file_redact_redact_proto_rawDesc = nil
	file_redact_redact_proto_goTypes = nil
	file_redact_redact_proto_depIdxs = nil
}
