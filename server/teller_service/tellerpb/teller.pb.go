// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: teller.proto

package tellerpb

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

type Teller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Teller) Reset() {
	*x = Teller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teller) ProtoMessage() {}

func (x *Teller) ProtoReflect() protoreflect.Message {
	mi := &file_teller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teller.ProtoReflect.Descriptor instead.
func (*Teller) Descriptor() ([]byte, []int) {
	return file_teller_proto_rawDescGZIP(), []int{0}
}

func (x *Teller) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Teller) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Teller) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Teller) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Teller) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Teller) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetTellersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTellersRequest) Reset() {
	*x = GetTellersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTellersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTellersRequest) ProtoMessage() {}

func (x *GetTellersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTellersRequest.ProtoReflect.Descriptor instead.
func (*GetTellersRequest) Descriptor() ([]byte, []int) {
	return file_teller_proto_rawDescGZIP(), []int{1}
}

type GetTellersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tellers []*Teller `protobuf:"bytes,1,rep,name=Tellers,proto3" json:"Tellers,omitempty"`
}

func (x *GetTellersResponse) Reset() {
	*x = GetTellersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTellersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTellersResponse) ProtoMessage() {}

func (x *GetTellersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTellersResponse.ProtoReflect.Descriptor instead.
func (*GetTellersResponse) Descriptor() ([]byte, []int) {
	return file_teller_proto_rawDescGZIP(), []int{2}
}

func (x *GetTellersResponse) GetTellers() []*Teller {
	if x != nil {
		return x.Tellers
	}
	return nil
}

type GetTellerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTellerRequest) Reset() {
	*x = GetTellerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTellerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTellerRequest) ProtoMessage() {}

func (x *GetTellerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTellerRequest.ProtoReflect.Descriptor instead.
func (*GetTellerRequest) Descriptor() ([]byte, []int) {
	return file_teller_proto_rawDescGZIP(), []int{3}
}

func (x *GetTellerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTellerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teller *Teller `protobuf:"bytes,1,opt,name=Teller,proto3" json:"Teller,omitempty"`
}

func (x *GetTellerResponse) Reset() {
	*x = GetTellerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTellerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTellerResponse) ProtoMessage() {}

func (x *GetTellerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTellerResponse.ProtoReflect.Descriptor instead.
func (*GetTellerResponse) Descriptor() ([]byte, []int) {
	return file_teller_proto_rawDescGZIP(), []int{4}
}

func (x *GetTellerResponse) GetTeller() *Teller {
	if x != nil {
		return x.Teller
	}
	return nil
}

var File_teller_proto protoreflect.FileDescriptor

var file_teller_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x07, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x06, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x32, 0x92, 0x01, 0x0a, 0x0d, 0x54, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teller_proto_rawDescOnce sync.Once
	file_teller_proto_rawDescData = file_teller_proto_rawDesc
)

func file_teller_proto_rawDescGZIP() []byte {
	file_teller_proto_rawDescOnce.Do(func() {
		file_teller_proto_rawDescData = protoimpl.X.CompressGZIP(file_teller_proto_rawDescData)
	})
	return file_teller_proto_rawDescData
}

var file_teller_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_teller_proto_goTypes = []interface{}{
	(*Teller)(nil),             // 0: main.Teller
	(*GetTellersRequest)(nil),  // 1: main.GetTellersRequest
	(*GetTellersResponse)(nil), // 2: main.GetTellersResponse
	(*GetTellerRequest)(nil),   // 3: main.GetTellerRequest
	(*GetTellerResponse)(nil),  // 4: main.GetTellerResponse
}
var file_teller_proto_depIdxs = []int32{
	0, // 0: main.GetTellersResponse.Tellers:type_name -> main.Teller
	0, // 1: main.GetTellerResponse.Teller:type_name -> main.Teller
	3, // 2: main.TellerService.GetTeller:input_type -> main.GetTellerRequest
	1, // 3: main.TellerService.GetTellers:input_type -> main.GetTellersRequest
	4, // 4: main.TellerService.GetTeller:output_type -> main.GetTellerResponse
	2, // 5: main.TellerService.GetTellers:output_type -> main.GetTellersResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teller_proto_init() }
func file_teller_proto_init() {
	if File_teller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teller); i {
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
		file_teller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTellersRequest); i {
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
		file_teller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTellersResponse); i {
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
		file_teller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTellerRequest); i {
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
		file_teller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTellerResponse); i {
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
			RawDescriptor: file_teller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teller_proto_goTypes,
		DependencyIndexes: file_teller_proto_depIdxs,
		MessageInfos:      file_teller_proto_msgTypes,
	}.Build()
	File_teller_proto = out.File
	file_teller_proto_rawDesc = nil
	file_teller_proto_goTypes = nil
	file_teller_proto_depIdxs = nil
}
