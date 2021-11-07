// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: usermgmt.proto

package grpcpb

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

type UnaryCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appeal string `protobuf:"bytes,1,opt,name=appeal,proto3" json:"appeal,omitempty"`
}

func (x *UnaryCall) Reset() {
	*x = UnaryCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryCall) ProtoMessage() {}

func (x *UnaryCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryCall.ProtoReflect.Descriptor instead.
func (*UnaryCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryCall) GetAppeal() string {
	if x != nil {
		return x.Appeal
	}
	return ""
}

type UnaryReturnCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *UnaryReturnCall) Reset() {
	*x = UnaryReturnCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryReturnCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryReturnCall) ProtoMessage() {}

func (x *UnaryReturnCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryReturnCall.ProtoReflect.Descriptor instead.
func (*UnaryReturnCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryReturnCall) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type ServerStreamingCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appeal string `protobuf:"bytes,1,opt,name=appeal,proto3" json:"appeal,omitempty"`
}

func (x *ServerStreamingCall) Reset() {
	*x = ServerStreamingCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamingCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamingCall) ProtoMessage() {}

func (x *ServerStreamingCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamingCall.ProtoReflect.Descriptor instead.
func (*ServerStreamingCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStreamingCall) GetAppeal() string {
	if x != nil {
		return x.Appeal
	}
	return ""
}

type ServerStreamingReturnCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *ServerStreamingReturnCall) Reset() {
	*x = ServerStreamingReturnCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamingReturnCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamingReturnCall) ProtoMessage() {}

func (x *ServerStreamingReturnCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamingReturnCall.ProtoReflect.Descriptor instead.
func (*ServerStreamingReturnCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{3}
}

func (x *ServerStreamingReturnCall) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type ClientStreamingCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appeal string `protobuf:"bytes,1,opt,name=appeal,proto3" json:"appeal,omitempty"`
}

func (x *ClientStreamingCall) Reset() {
	*x = ClientStreamingCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamingCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamingCall) ProtoMessage() {}

func (x *ClientStreamingCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamingCall.ProtoReflect.Descriptor instead.
func (*ClientStreamingCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{4}
}

func (x *ClientStreamingCall) GetAppeal() string {
	if x != nil {
		return x.Appeal
	}
	return ""
}

type ClientStreamingReturnCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *ClientStreamingReturnCall) Reset() {
	*x = ClientStreamingReturnCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamingReturnCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamingReturnCall) ProtoMessage() {}

func (x *ClientStreamingReturnCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamingReturnCall.ProtoReflect.Descriptor instead.
func (*ClientStreamingReturnCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{5}
}

func (x *ClientStreamingReturnCall) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type DuoStreamingCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appeal string `protobuf:"bytes,1,opt,name=appeal,proto3" json:"appeal,omitempty"`
}

func (x *DuoStreamingCall) Reset() {
	*x = DuoStreamingCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DuoStreamingCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuoStreamingCall) ProtoMessage() {}

func (x *DuoStreamingCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuoStreamingCall.ProtoReflect.Descriptor instead.
func (*DuoStreamingCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{6}
}

func (x *DuoStreamingCall) GetAppeal() string {
	if x != nil {
		return x.Appeal
	}
	return ""
}

type DuoStreamingReturnCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *DuoStreamingReturnCall) Reset() {
	*x = DuoStreamingReturnCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DuoStreamingReturnCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DuoStreamingReturnCall) ProtoMessage() {}

func (x *DuoStreamingReturnCall) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DuoStreamingReturnCall.ProtoReflect.Descriptor instead.
func (*DuoStreamingReturnCall) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{7}
}

func (x *DuoStreamingReturnCall) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_usermgmt_proto protoreflect.FileDescriptor

var file_usermgmt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x23, 0x0a, 0x09, 0x55, 0x6e,
	0x61, 0x72, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x22,
	0x29, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61,
	0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x22, 0x33, 0x0a, 0x19, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2d,
	0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x22, 0x33, 0x0a,
	0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x22, 0x2a, 0x0a, 0x10, 0x44, 0x75, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x22, 0x30,
	0x0a, 0x16, 0x44, 0x75, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x32, 0x6e, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4d, 0x73, 0x67, 0x4d,
	0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x6c, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usermgmt_proto_rawDescOnce sync.Once
	file_usermgmt_proto_rawDescData = file_usermgmt_proto_rawDesc
)

func file_usermgmt_proto_rawDescGZIP() []byte {
	file_usermgmt_proto_rawDescOnce.Do(func() {
		file_usermgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_usermgmt_proto_rawDescData)
	})
	return file_usermgmt_proto_rawDescData
}

var file_usermgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_usermgmt_proto_goTypes = []interface{}{
	(*UnaryCall)(nil),                 // 0: usermgmt.UnaryCall
	(*UnaryReturnCall)(nil),           // 1: usermgmt.UnaryReturnCall
	(*ServerStreamingCall)(nil),       // 2: usermgmt.ServerStreamingCall
	(*ServerStreamingReturnCall)(nil), // 3: usermgmt.ServerStreamingReturnCall
	(*ClientStreamingCall)(nil),       // 4: usermgmt.ClientStreamingCall
	(*ClientStreamingReturnCall)(nil), // 5: usermgmt.ClientStreamingReturnCall
	(*DuoStreamingCall)(nil),          // 6: usermgmt.DuoStreamingCall
	(*DuoStreamingReturnCall)(nil),    // 7: usermgmt.DuoStreamingReturnCall
}
var file_usermgmt_proto_depIdxs = []int32{
	2, // 0: usermgmt.UserManagement.ReturnMsgManyTimes:input_type -> usermgmt.ServerStreamingCall
	3, // 1: usermgmt.UserManagement.ReturnMsgManyTimes:output_type -> usermgmt.ServerStreamingReturnCall
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usermgmt_proto_init() }
func file_usermgmt_proto_init() {
	if File_usermgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usermgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryCall); i {
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
		file_usermgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryReturnCall); i {
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
		file_usermgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamingCall); i {
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
		file_usermgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamingReturnCall); i {
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
		file_usermgmt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamingCall); i {
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
		file_usermgmt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamingReturnCall); i {
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
		file_usermgmt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DuoStreamingCall); i {
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
		file_usermgmt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DuoStreamingReturnCall); i {
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
			RawDescriptor: file_usermgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usermgmt_proto_goTypes,
		DependencyIndexes: file_usermgmt_proto_depIdxs,
		MessageInfos:      file_usermgmt_proto_msgTypes,
	}.Build()
	File_usermgmt_proto = out.File
	file_usermgmt_proto_rawDesc = nil
	file_usermgmt_proto_goTypes = nil
	file_usermgmt_proto_depIdxs = nil
}
