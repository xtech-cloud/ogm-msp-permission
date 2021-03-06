// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/permission/shared.proto

package permission

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

// 状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 状态信息
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 空白请求
type BlankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankRequest) Reset() {
	*x = BlankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankRequest) ProtoMessage() {}

func (x *BlankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankRequest.ProtoReflect.Descriptor instead.
func (*BlankRequest) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{1}
}

// 空白回复
type BlankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
}

func (x *BlankResponse) Reset() {
	*x = BlankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankResponse) ProtoMessage() {}

func (x *BlankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankResponse.ProtoReflect.Descriptor instead.
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{2}
}

func (x *BlankResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 带UUID的回复
type UuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Uuid   string  `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`     // uuid
}

func (x *UuidResponse) Reset() {
	*x = UuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidResponse) ProtoMessage() {}

func (x *UuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidResponse.ProtoReflect.Descriptor instead.
func (*UuidResponse) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{3}
}

func (x *UuidResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *UuidResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 作用域实体
type ScopeEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`   // 键
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // 名称
}

func (x *ScopeEntity) Reset() {
	*x = ScopeEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopeEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeEntity) ProtoMessage() {}

func (x *ScopeEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeEntity.ProtoReflect.Descriptor instead.
func (*ScopeEntity) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{4}
}

func (x *ScopeEntity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ScopeEntity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ScopeEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 规则实体
type RuleEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`    // uuid
	Scope string `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`  // 作用域的uuid
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`      // 键
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`    // 名称
	State int32  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"` // 状态
}

func (x *RuleEntity) Reset() {
	*x = RuleEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_permission_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleEntity) ProtoMessage() {}

func (x *RuleEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_permission_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleEntity.ProtoReflect.Descriptor instead.
func (*RuleEntity) Descriptor() ([]byte, []int) {
	return file_proto_permission_shared_proto_rawDescGZIP(), []int{5}
}

func (x *RuleEntity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RuleEntity) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *RuleEntity) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RuleEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RuleEntity) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_proto_permission_shared_proto protoreflect.FileDescriptor

var file_proto_permission_shared_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x4e, 0x0a, 0x0c, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x47, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x0a, 0x52, 0x75, 0x6c,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x1d, 0x5a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x3b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_permission_shared_proto_rawDescOnce sync.Once
	file_proto_permission_shared_proto_rawDescData = file_proto_permission_shared_proto_rawDesc
)

func file_proto_permission_shared_proto_rawDescGZIP() []byte {
	file_proto_permission_shared_proto_rawDescOnce.Do(func() {
		file_proto_permission_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_permission_shared_proto_rawDescData)
	})
	return file_proto_permission_shared_proto_rawDescData
}

var file_proto_permission_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_permission_shared_proto_goTypes = []interface{}{
	(*Status)(nil),        // 0: permission.Status
	(*BlankRequest)(nil),  // 1: permission.BlankRequest
	(*BlankResponse)(nil), // 2: permission.BlankResponse
	(*UuidResponse)(nil),  // 3: permission.UuidResponse
	(*ScopeEntity)(nil),   // 4: permission.ScopeEntity
	(*RuleEntity)(nil),    // 5: permission.RuleEntity
}
var file_proto_permission_shared_proto_depIdxs = []int32{
	0, // 0: permission.BlankResponse.status:type_name -> permission.Status
	0, // 1: permission.UuidResponse.status:type_name -> permission.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_permission_shared_proto_init() }
func file_proto_permission_shared_proto_init() {
	if File_proto_permission_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_permission_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_permission_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankRequest); i {
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
		file_proto_permission_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankResponse); i {
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
		file_proto_permission_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidResponse); i {
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
		file_proto_permission_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopeEntity); i {
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
		file_proto_permission_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleEntity); i {
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
			RawDescriptor: file_proto_permission_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_permission_shared_proto_goTypes,
		DependencyIndexes: file_proto_permission_shared_proto_depIdxs,
		MessageInfos:      file_proto_permission_shared_proto_msgTypes,
	}.Build()
	File_proto_permission_shared_proto = out.File
	file_proto_permission_shared_proto_rawDesc = nil
	file_proto_permission_shared_proto_goTypes = nil
	file_proto_permission_shared_proto_depIdxs = nil
}
