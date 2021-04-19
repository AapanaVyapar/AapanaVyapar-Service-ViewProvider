// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: trending-message.proto

package pb

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

type GetTrendingShopsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey          string    `protobuf:"bytes,10,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Token           string    `protobuf:"bytes,11,opt,name=token,proto3" json:"token,omitempty"`
	Location        *Location `protobuf:"bytes,12,opt,name=location,proto3" json:"location,omitempty"`
	DistanceInMeter string    `protobuf:"bytes,13,opt,name=distanceInMeter,proto3" json:"distanceInMeter,omitempty"`
}

func (x *GetTrendingShopsRequest) Reset() {
	*x = GetTrendingShopsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trending_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendingShopsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendingShopsRequest) ProtoMessage() {}

func (x *GetTrendingShopsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trending_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendingShopsRequest.ProtoReflect.Descriptor instead.
func (*GetTrendingShopsRequest) Descriptor() ([]byte, []int) {
	return file_trending_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetTrendingShopsRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetTrendingShopsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTrendingShopsRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *GetTrendingShopsRequest) GetDistanceInMeter() string {
	if x != nil {
		return x.DistanceInMeter
	}
	return ""
}

type GetTrendingShopsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shops *ShopsNearBy `protobuf:"bytes,14,opt,name=shops,proto3" json:"shops,omitempty"`
}

func (x *GetTrendingShopsResponse) Reset() {
	*x = GetTrendingShopsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trending_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendingShopsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendingShopsResponse) ProtoMessage() {}

func (x *GetTrendingShopsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trending_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendingShopsResponse.ProtoReflect.Descriptor instead.
func (*GetTrendingShopsResponse) Descriptor() ([]byte, []int) {
	return file_trending_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetTrendingShopsResponse) GetShops() *ShopsNearBy {
	if x != nil {
		return x.Shops
	}
	return nil
}

type GetTrendingProductsByShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string   `protobuf:"bytes,21,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Token  string   `protobuf:"bytes,22,opt,name=token,proto3" json:"token,omitempty"`
	ShopId []string `protobuf:"bytes,23,rep,name=shopId,proto3" json:"shopId,omitempty"`
}

func (x *GetTrendingProductsByShopRequest) Reset() {
	*x = GetTrendingProductsByShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trending_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendingProductsByShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendingProductsByShopRequest) ProtoMessage() {}

func (x *GetTrendingProductsByShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trending_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendingProductsByShopRequest.ProtoReflect.Descriptor instead.
func (*GetTrendingProductsByShopRequest) Descriptor() ([]byte, []int) {
	return file_trending_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrendingProductsByShopRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *GetTrendingProductsByShopRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTrendingProductsByShopRequest) GetShopId() []string {
	if x != nil {
		return x.ShopId
	}
	return nil
}

type GetTrendingProductsByShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryData *ProductsOfShopsNearBy `protobuf:"bytes,24,opt,name=categoryData,proto3" json:"categoryData,omitempty"`
}

func (x *GetTrendingProductsByShopResponse) Reset() {
	*x = GetTrendingProductsByShopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trending_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrendingProductsByShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrendingProductsByShopResponse) ProtoMessage() {}

func (x *GetTrendingProductsByShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trending_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrendingProductsByShopResponse.ProtoReflect.Descriptor instead.
func (*GetTrendingProductsByShopResponse) Descriptor() ([]byte, []int) {
	return file_trending_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetTrendingProductsByShopResponse) GetCategoryData() *ProductsOfShopsNearBy {
	if x != nil {
		return x.CategoryData
	}
	return nil
}

var File_trending_message_proto protoreflect.FileDescriptor

var file_trending_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98,
	0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x68,
	0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x6e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72,
	0x42, 0x79, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x22, 0x68, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x17, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4f, 0x66, 0x53, 0x68, 0x6f, 0x70, 0x73,
	0x4e, 0x65, 0x61, 0x72, 0x42, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x30, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x70, 0x61,
	0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76,
	0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trending_message_proto_rawDescOnce sync.Once
	file_trending_message_proto_rawDescData = file_trending_message_proto_rawDesc
)

func file_trending_message_proto_rawDescGZIP() []byte {
	file_trending_message_proto_rawDescOnce.Do(func() {
		file_trending_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_trending_message_proto_rawDescData)
	})
	return file_trending_message_proto_rawDescData
}

var file_trending_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trending_message_proto_goTypes = []interface{}{
	(*GetTrendingShopsRequest)(nil),           // 0: GetTrendingShopsRequest
	(*GetTrendingShopsResponse)(nil),          // 1: GetTrendingShopsResponse
	(*GetTrendingProductsByShopRequest)(nil),  // 2: GetTrendingProductsByShopRequest
	(*GetTrendingProductsByShopResponse)(nil), // 3: GetTrendingProductsByShopResponse
	(*Location)(nil),                          // 4: Location
	(*ShopsNearBy)(nil),                       // 5: ShopsNearBy
	(*ProductsOfShopsNearBy)(nil),             // 6: ProductsOfShopsNearBy
}
var file_trending_message_proto_depIdxs = []int32{
	4, // 0: GetTrendingShopsRequest.location:type_name -> Location
	5, // 1: GetTrendingShopsResponse.shops:type_name -> ShopsNearBy
	6, // 2: GetTrendingProductsByShopResponse.categoryData:type_name -> ProductsOfShopsNearBy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_trending_message_proto_init() }
func file_trending_message_proto_init() {
	if File_trending_message_proto != nil {
		return
	}
	file_common_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_trending_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendingShopsRequest); i {
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
		file_trending_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendingShopsResponse); i {
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
		file_trending_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendingProductsByShopRequest); i {
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
		file_trending_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrendingProductsByShopResponse); i {
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
			RawDescriptor: file_trending_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trending_message_proto_goTypes,
		DependencyIndexes: file_trending_message_proto_depIdxs,
		MessageInfos:      file_trending_message_proto_msgTypes,
	}.Build()
	File_trending_message_proto = out.File
	file_trending_message_proto_rawDesc = nil
	file_trending_message_proto_goTypes = nil
	file_trending_message_proto_depIdxs = nil
}
