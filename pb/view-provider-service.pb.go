// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: view-provider-service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_view_provider_service_proto protoreflect.FileDescriptor

var file_view_provider_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74,
	0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xf2, 0x06, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x18, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0f, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x42, 0x79,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70,
	0x73, 0x42, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x42, 0x79, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x47, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x18, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1d, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x69, 0x6b, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x18, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1d, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x30, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x70,
	0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61,
	0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50,
	0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_view_provider_service_proto_goTypes = []interface{}{
	(*GetTrendingShopsRequest)(nil),           // 0: GetTrendingShopsRequest
	(*GetTrendingProductsByShopRequest)(nil),  // 1: GetTrendingProductsByShopRequest
	(*GetProductRequest)(nil),                 // 2: GetProductRequest
	(*GetShopRequest)(nil),                    // 3: GetShopRequest
	(*GetProductsBySearchRequest)(nil),        // 4: GetProductsBySearchRequest
	(*GetShopsBySearchRequest)(nil),           // 5: GetShopsBySearchRequest
	(*AddToLikeProductRequest)(nil),           // 6: AddToLikeProductRequest
	(*RemoveFromLikeProductRequest)(nil),      // 7: RemoveFromLikeProductRequest
	(*AddToCartProductRequest)(nil),           // 8: AddToCartProductRequest
	(*RemoveFromCartProductRequest)(nil),      // 9: RemoveFromCartProductRequest
	(*GetOrdersRequest)(nil),                  // 10: GetOrdersRequest
	(*GetCartRequest)(nil),                    // 11: GetCartRequest
	(*GetTrendingShopsResponse)(nil),          // 12: GetTrendingShopsResponse
	(*GetTrendingProductsByShopResponse)(nil), // 13: GetTrendingProductsByShopResponse
	(*GetProductResponse)(nil),                // 14: GetProductResponse
	(*GetShopResponse)(nil),                   // 15: GetShopResponse
	(*GetProductsBySearchResponse)(nil),       // 16: GetProductsBySearchResponse
	(*GetShopsBySearchResponse)(nil),          // 17: GetShopsBySearchResponse
	(*AddToLikeProductResponse)(nil),          // 18: AddToLikeProductResponse
	(*RemoveFromLikeProductResponse)(nil),     // 19: RemoveFromLikeProductResponse
	(*AddToCartProductResponse)(nil),          // 20: AddToCartProductResponse
	(*RemoveFromCartProductResponse)(nil),     // 21: RemoveFromCartProductResponse
	(*GetOrdersResponse)(nil),                 // 22: GetOrdersResponse
	(*GetCartResponse)(nil),                   // 23: GetCartResponse
}
var file_view_provider_service_proto_depIdxs = []int32{
	0,  // 0: ViewProviderService.GetTrendingShops:input_type -> GetTrendingShopsRequest
	1,  // 1: ViewProviderService.GetTrendingProductsByShop:input_type -> GetTrendingProductsByShopRequest
	2,  // 2: ViewProviderService.GetProduct:input_type -> GetProductRequest
	3,  // 3: ViewProviderService.GetShop:input_type -> GetShopRequest
	4,  // 4: ViewProviderService.GetProductsBySearch:input_type -> GetProductsBySearchRequest
	5,  // 5: ViewProviderService.GetShopsBySearch:input_type -> GetShopsBySearchRequest
	6,  // 6: ViewProviderService.AddToLikeProduct:input_type -> AddToLikeProductRequest
	7,  // 7: ViewProviderService.RemoveFromLikeProduct:input_type -> RemoveFromLikeProductRequest
	8,  // 8: ViewProviderService.AddToCartProduct:input_type -> AddToCartProductRequest
	9,  // 9: ViewProviderService.RemoveFromCartProduct:input_type -> RemoveFromCartProductRequest
	10, // 10: ViewProviderService.GetOrders:input_type -> GetOrdersRequest
	11, // 11: ViewProviderService.GetCart:input_type -> GetCartRequest
	12, // 12: ViewProviderService.GetTrendingShops:output_type -> GetTrendingShopsResponse
	13, // 13: ViewProviderService.GetTrendingProductsByShop:output_type -> GetTrendingProductsByShopResponse
	14, // 14: ViewProviderService.GetProduct:output_type -> GetProductResponse
	15, // 15: ViewProviderService.GetShop:output_type -> GetShopResponse
	16, // 16: ViewProviderService.GetProductsBySearch:output_type -> GetProductsBySearchResponse
	17, // 17: ViewProviderService.GetShopsBySearch:output_type -> GetShopsBySearchResponse
	18, // 18: ViewProviderService.AddToLikeProduct:output_type -> AddToLikeProductResponse
	19, // 19: ViewProviderService.RemoveFromLikeProduct:output_type -> RemoveFromLikeProductResponse
	20, // 20: ViewProviderService.AddToCartProduct:output_type -> AddToCartProductResponse
	21, // 21: ViewProviderService.RemoveFromCartProduct:output_type -> RemoveFromCartProductResponse
	22, // 22: ViewProviderService.GetOrders:output_type -> GetOrdersResponse
	23, // 23: ViewProviderService.GetCart:output_type -> GetCartResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_view_provider_service_proto_init() }
func file_view_provider_service_proto_init() {
	if File_view_provider_service_proto != nil {
		return
	}
	file_trending_message_proto_init()
	file_detailed_message_proto_init()
	file_user_commands_message_proto_init()
	file_search_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_view_provider_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_view_provider_service_proto_goTypes,
		DependencyIndexes: file_view_provider_service_proto_depIdxs,
	}.Build()
	File_view_provider_service_proto = out.File
	file_view_provider_service_proto_rawDesc = nil
	file_view_provider_service_proto_goTypes = nil
	file_view_provider_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ViewProviderServiceClient is the client API for ViewProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ViewProviderServiceClient interface {
	GetTrendingShops(ctx context.Context, in *GetTrendingShopsRequest, opts ...grpc.CallOption) (ViewProviderService_GetTrendingShopsClient, error)
	GetTrendingProductsByShop(ctx context.Context, in *GetTrendingProductsByShopRequest, opts ...grpc.CallOption) (ViewProviderService_GetTrendingProductsByShopClient, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetShop(ctx context.Context, in *GetShopRequest, opts ...grpc.CallOption) (*GetShopResponse, error)
	GetProductsBySearch(ctx context.Context, in *GetProductsBySearchRequest, opts ...grpc.CallOption) (ViewProviderService_GetProductsBySearchClient, error)
	GetShopsBySearch(ctx context.Context, in *GetShopsBySearchRequest, opts ...grpc.CallOption) (ViewProviderService_GetShopsBySearchClient, error)
	AddToLikeProduct(ctx context.Context, in *AddToLikeProductRequest, opts ...grpc.CallOption) (*AddToLikeProductResponse, error)
	RemoveFromLikeProduct(ctx context.Context, in *RemoveFromLikeProductRequest, opts ...grpc.CallOption) (*RemoveFromLikeProductResponse, error)
	AddToCartProduct(ctx context.Context, in *AddToCartProductRequest, opts ...grpc.CallOption) (*AddToCartProductResponse, error)
	RemoveFromCartProduct(ctx context.Context, in *RemoveFromCartProductRequest, opts ...grpc.CallOption) (*RemoveFromCartProductResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (ViewProviderService_GetOrdersClient, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (ViewProviderService_GetCartClient, error)
}

type viewProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewProviderServiceClient(cc grpc.ClientConnInterface) ViewProviderServiceClient {
	return &viewProviderServiceClient{cc}
}

func (c *viewProviderServiceClient) GetTrendingShops(ctx context.Context, in *GetTrendingShopsRequest, opts ...grpc.CallOption) (ViewProviderService_GetTrendingShopsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[0], "/ViewProviderService/GetTrendingShops", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetTrendingShopsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetTrendingShopsClient interface {
	Recv() (*GetTrendingShopsResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetTrendingShopsClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetTrendingShopsClient) Recv() (*GetTrendingShopsResponse, error) {
	m := new(GetTrendingShopsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewProviderServiceClient) GetTrendingProductsByShop(ctx context.Context, in *GetTrendingProductsByShopRequest, opts ...grpc.CallOption) (ViewProviderService_GetTrendingProductsByShopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[1], "/ViewProviderService/GetTrendingProductsByShop", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetTrendingProductsByShopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetTrendingProductsByShopClient interface {
	Recv() (*GetTrendingProductsByShopResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetTrendingProductsByShopClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetTrendingProductsByShopClient) Recv() (*GetTrendingProductsByShopResponse, error) {
	m := new(GetTrendingProductsByShopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewProviderServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) GetShop(ctx context.Context, in *GetShopRequest, opts ...grpc.CallOption) (*GetShopResponse, error) {
	out := new(GetShopResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) GetProductsBySearch(ctx context.Context, in *GetProductsBySearchRequest, opts ...grpc.CallOption) (ViewProviderService_GetProductsBySearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[2], "/ViewProviderService/GetProductsBySearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetProductsBySearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetProductsBySearchClient interface {
	Recv() (*GetProductsBySearchResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetProductsBySearchClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetProductsBySearchClient) Recv() (*GetProductsBySearchResponse, error) {
	m := new(GetProductsBySearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewProviderServiceClient) GetShopsBySearch(ctx context.Context, in *GetShopsBySearchRequest, opts ...grpc.CallOption) (ViewProviderService_GetShopsBySearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[3], "/ViewProviderService/GetShopsBySearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetShopsBySearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetShopsBySearchClient interface {
	Recv() (*GetShopsBySearchResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetShopsBySearchClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetShopsBySearchClient) Recv() (*GetShopsBySearchResponse, error) {
	m := new(GetShopsBySearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewProviderServiceClient) AddToLikeProduct(ctx context.Context, in *AddToLikeProductRequest, opts ...grpc.CallOption) (*AddToLikeProductResponse, error) {
	out := new(AddToLikeProductResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/AddToLikeProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) RemoveFromLikeProduct(ctx context.Context, in *RemoveFromLikeProductRequest, opts ...grpc.CallOption) (*RemoveFromLikeProductResponse, error) {
	out := new(RemoveFromLikeProductResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/RemoveFromLikeProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) AddToCartProduct(ctx context.Context, in *AddToCartProductRequest, opts ...grpc.CallOption) (*AddToCartProductResponse, error) {
	out := new(AddToCartProductResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/AddToCartProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) RemoveFromCartProduct(ctx context.Context, in *RemoveFromCartProductRequest, opts ...grpc.CallOption) (*RemoveFromCartProductResponse, error) {
	out := new(RemoveFromCartProductResponse)
	err := c.cc.Invoke(ctx, "/ViewProviderService/RemoveFromCartProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewProviderServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (ViewProviderService_GetOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[4], "/ViewProviderService/GetOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetOrdersClient interface {
	Recv() (*GetOrdersResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetOrdersClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetOrdersClient) Recv() (*GetOrdersResponse, error) {
	m := new(GetOrdersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewProviderServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (ViewProviderService_GetCartClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewProviderService_serviceDesc.Streams[5], "/ViewProviderService/GetCart", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewProviderServiceGetCartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewProviderService_GetCartClient interface {
	Recv() (*GetCartResponse, error)
	grpc.ClientStream
}

type viewProviderServiceGetCartClient struct {
	grpc.ClientStream
}

func (x *viewProviderServiceGetCartClient) Recv() (*GetCartResponse, error) {
	m := new(GetCartResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ViewProviderServiceServer is the server API for ViewProviderService service.
type ViewProviderServiceServer interface {
	GetTrendingShops(*GetTrendingShopsRequest, ViewProviderService_GetTrendingShopsServer) error
	GetTrendingProductsByShop(*GetTrendingProductsByShopRequest, ViewProviderService_GetTrendingProductsByShopServer) error
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetShop(context.Context, *GetShopRequest) (*GetShopResponse, error)
	GetProductsBySearch(*GetProductsBySearchRequest, ViewProviderService_GetProductsBySearchServer) error
	GetShopsBySearch(*GetShopsBySearchRequest, ViewProviderService_GetShopsBySearchServer) error
	AddToLikeProduct(context.Context, *AddToLikeProductRequest) (*AddToLikeProductResponse, error)
	RemoveFromLikeProduct(context.Context, *RemoveFromLikeProductRequest) (*RemoveFromLikeProductResponse, error)
	AddToCartProduct(context.Context, *AddToCartProductRequest) (*AddToCartProductResponse, error)
	RemoveFromCartProduct(context.Context, *RemoveFromCartProductRequest) (*RemoveFromCartProductResponse, error)
	GetOrders(*GetOrdersRequest, ViewProviderService_GetOrdersServer) error
	GetCart(*GetCartRequest, ViewProviderService_GetCartServer) error
}

// UnimplementedViewProviderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedViewProviderServiceServer struct {
}

func (*UnimplementedViewProviderServiceServer) GetTrendingShops(*GetTrendingShopsRequest, ViewProviderService_GetTrendingShopsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTrendingShops not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetTrendingProductsByShop(*GetTrendingProductsByShopRequest, ViewProviderService_GetTrendingProductsByShopServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTrendingProductsByShop not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetShop(context.Context, *GetShopRequest) (*GetShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShop not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetProductsBySearch(*GetProductsBySearchRequest, ViewProviderService_GetProductsBySearchServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductsBySearch not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetShopsBySearch(*GetShopsBySearchRequest, ViewProviderService_GetShopsBySearchServer) error {
	return status.Errorf(codes.Unimplemented, "method GetShopsBySearch not implemented")
}
func (*UnimplementedViewProviderServiceServer) AddToLikeProduct(context.Context, *AddToLikeProductRequest) (*AddToLikeProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToLikeProduct not implemented")
}
func (*UnimplementedViewProviderServiceServer) RemoveFromLikeProduct(context.Context, *RemoveFromLikeProductRequest) (*RemoveFromLikeProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromLikeProduct not implemented")
}
func (*UnimplementedViewProviderServiceServer) AddToCartProduct(context.Context, *AddToCartProductRequest) (*AddToCartProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCartProduct not implemented")
}
func (*UnimplementedViewProviderServiceServer) RemoveFromCartProduct(context.Context, *RemoveFromCartProductRequest) (*RemoveFromCartProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCartProduct not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetOrders(*GetOrdersRequest, ViewProviderService_GetOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (*UnimplementedViewProviderServiceServer) GetCart(*GetCartRequest, ViewProviderService_GetCartServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}

func RegisterViewProviderServiceServer(s *grpc.Server, srv ViewProviderServiceServer) {
	s.RegisterService(&_ViewProviderService_serviceDesc, srv)
}

func _ViewProviderService_GetTrendingShops_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTrendingShopsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetTrendingShops(m, &viewProviderServiceGetTrendingShopsServer{stream})
}

type ViewProviderService_GetTrendingShopsServer interface {
	Send(*GetTrendingShopsResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetTrendingShopsServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetTrendingShopsServer) Send(m *GetTrendingShopsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewProviderService_GetTrendingProductsByShop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTrendingProductsByShopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetTrendingProductsByShop(m, &viewProviderServiceGetTrendingProductsByShopServer{stream})
}

type ViewProviderService_GetTrendingProductsByShopServer interface {
	Send(*GetTrendingProductsByShopResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetTrendingProductsByShopServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetTrendingProductsByShopServer) Send(m *GetTrendingProductsByShopResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewProviderService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).GetShop(ctx, req.(*GetShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_GetProductsBySearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProductsBySearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetProductsBySearch(m, &viewProviderServiceGetProductsBySearchServer{stream})
}

type ViewProviderService_GetProductsBySearchServer interface {
	Send(*GetProductsBySearchResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetProductsBySearchServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetProductsBySearchServer) Send(m *GetProductsBySearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewProviderService_GetShopsBySearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetShopsBySearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetShopsBySearch(m, &viewProviderServiceGetShopsBySearchServer{stream})
}

type ViewProviderService_GetShopsBySearchServer interface {
	Send(*GetShopsBySearchResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetShopsBySearchServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetShopsBySearchServer) Send(m *GetShopsBySearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewProviderService_AddToLikeProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToLikeProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).AddToLikeProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/AddToLikeProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).AddToLikeProduct(ctx, req.(*AddToLikeProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_RemoveFromLikeProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromLikeProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).RemoveFromLikeProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/RemoveFromLikeProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).RemoveFromLikeProduct(ctx, req.(*RemoveFromLikeProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_AddToCartProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).AddToCartProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/AddToCartProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).AddToCartProduct(ctx, req.(*AddToCartProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_RemoveFromCartProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewProviderServiceServer).RemoveFromCartProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ViewProviderService/RemoveFromCartProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewProviderServiceServer).RemoveFromCartProduct(ctx, req.(*RemoveFromCartProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewProviderService_GetOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetOrders(m, &viewProviderServiceGetOrdersServer{stream})
}

type ViewProviderService_GetOrdersServer interface {
	Send(*GetOrdersResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetOrdersServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetOrdersServer) Send(m *GetOrdersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewProviderService_GetCart_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewProviderServiceServer).GetCart(m, &viewProviderServiceGetCartServer{stream})
}

type ViewProviderService_GetCartServer interface {
	Send(*GetCartResponse) error
	grpc.ServerStream
}

type viewProviderServiceGetCartServer struct {
	grpc.ServerStream
}

func (x *viewProviderServiceGetCartServer) Send(m *GetCartResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ViewProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ViewProviderService",
	HandlerType: (*ViewProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ViewProviderService_GetProduct_Handler,
		},
		{
			MethodName: "GetShop",
			Handler:    _ViewProviderService_GetShop_Handler,
		},
		{
			MethodName: "AddToLikeProduct",
			Handler:    _ViewProviderService_AddToLikeProduct_Handler,
		},
		{
			MethodName: "RemoveFromLikeProduct",
			Handler:    _ViewProviderService_RemoveFromLikeProduct_Handler,
		},
		{
			MethodName: "AddToCartProduct",
			Handler:    _ViewProviderService_AddToCartProduct_Handler,
		},
		{
			MethodName: "RemoveFromCartProduct",
			Handler:    _ViewProviderService_RemoveFromCartProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTrendingShops",
			Handler:       _ViewProviderService_GetTrendingShops_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTrendingProductsByShop",
			Handler:       _ViewProviderService_GetTrendingProductsByShop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetProductsBySearch",
			Handler:       _ViewProviderService_GetProductsBySearch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetShopsBySearch",
			Handler:       _ViewProviderService_GetShopsBySearch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrders",
			Handler:       _ViewProviderService_GetOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCart",
			Handler:       _ViewProviderService_GetCart_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "view-provider-service.proto",
}
