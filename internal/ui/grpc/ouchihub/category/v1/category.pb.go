// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ouchihub/category/v1/category.proto

package categoryv1

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

type ListCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedsRestricted bool `protobuf:"varint,1,opt,name=needs_restricted,json=needsRestricted,proto3" json:"needs_restricted,omitempty"`
}

func (x *ListCategoriesRequest) Reset() {
	*x = ListCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_category_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesRequest) ProtoMessage() {}

func (x *ListCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_category_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesRequest.ProtoReflect.Descriptor instead.
func (*ListCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_ouchihub_category_v1_category_proto_rawDescGZIP(), []int{0}
}

func (x *ListCategoriesRequest) GetNeedsRestricted() bool {
	if x != nil {
		return x.NeedsRestricted
	}
	return false
}

type ListCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ListCategoriesResponse) Reset() {
	*x = ListCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_category_v1_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesResponse) ProtoMessage() {}

func (x *ListCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_category_v1_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ListCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_ouchihub_category_v1_category_proto_rawDescGZIP(), []int{1}
}

func (x *ListCategoriesResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsRestricted bool   `protobuf:"varint,3,opt,name=is_restricted,json=isRestricted,proto3" json:"is_restricted,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ouchihub_category_v1_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_ouchihub_category_v1_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_ouchihub_category_v1_category_proto_rawDescGZIP(), []int{2}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetIsRestricted() bool {
	if x != nil {
		return x.IsRestricted
	}
	return false
}

var File_ouchihub_category_v1_category_proto protoreflect.FileDescriptor

var file_ouchihub_category_v1_category_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x42, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x5f, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x6e, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x22,
	0x58, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x32, 0x80,
	0x01, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xf8, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x69, 0x6d,
	0x6f, 0x74, 0x6f, 0x2d, 0x6a, 0x75, 0x6e, 0x79, 0x61, 0x2f, 0x6f, 0x75, 0x63, 0x68, 0x69, 0x2d,
	0x68, 0x75, 0x62, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x75, 0x63,
	0x68, 0x69, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f,
	0x43, 0x58, 0xaa, 0x02, 0x14, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x4f, 0x75, 0x63, 0x68,
	0x69, 0x68, 0x75, 0x62, 0x5c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x20, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x5c, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x4f, 0x75, 0x63, 0x68, 0x69, 0x68, 0x75, 0x62, 0x3a, 0x3a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ouchihub_category_v1_category_proto_rawDescOnce sync.Once
	file_ouchihub_category_v1_category_proto_rawDescData = file_ouchihub_category_v1_category_proto_rawDesc
)

func file_ouchihub_category_v1_category_proto_rawDescGZIP() []byte {
	file_ouchihub_category_v1_category_proto_rawDescOnce.Do(func() {
		file_ouchihub_category_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_ouchihub_category_v1_category_proto_rawDescData)
	})
	return file_ouchihub_category_v1_category_proto_rawDescData
}

var file_ouchihub_category_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ouchihub_category_v1_category_proto_goTypes = []interface{}{
	(*ListCategoriesRequest)(nil),  // 0: ouchihub.category.v1.ListCategoriesRequest
	(*ListCategoriesResponse)(nil), // 1: ouchihub.category.v1.ListCategoriesResponse
	(*Category)(nil),               // 2: ouchihub.category.v1.Category
}
var file_ouchihub_category_v1_category_proto_depIdxs = []int32{
	2, // 0: ouchihub.category.v1.ListCategoriesResponse.categories:type_name -> ouchihub.category.v1.Category
	0, // 1: ouchihub.category.v1.CategoryService.ListCategories:input_type -> ouchihub.category.v1.ListCategoriesRequest
	1, // 2: ouchihub.category.v1.CategoryService.ListCategories:output_type -> ouchihub.category.v1.ListCategoriesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ouchihub_category_v1_category_proto_init() }
func file_ouchihub_category_v1_category_proto_init() {
	if File_ouchihub_category_v1_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ouchihub_category_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesRequest); i {
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
		file_ouchihub_category_v1_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesResponse); i {
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
		file_ouchihub_category_v1_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_ouchihub_category_v1_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ouchihub_category_v1_category_proto_goTypes,
		DependencyIndexes: file_ouchihub_category_v1_category_proto_depIdxs,
		MessageInfos:      file_ouchihub_category_v1_category_proto_msgTypes,
	}.Build()
	File_ouchihub_category_v1_category_proto = out.File
	file_ouchihub_category_v1_category_proto_rawDesc = nil
	file_ouchihub_category_v1_category_proto_goTypes = nil
	file_ouchihub_category_v1_category_proto_depIdxs = nil
}
