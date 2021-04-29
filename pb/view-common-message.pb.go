// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: view-common-message.proto

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

type Category int32

const (
	Category_SPORTS_AND_FITNESS       Category = 0
	Category_ELECTRIC                 Category = 1
	Category_DEVOTIONAL               Category = 2
	Category_AGRICULTURAL             Category = 3
	Category_WOMENS_CLOTHING          Category = 4
	Category_WOMENS_ACCESSORIES       Category = 5
	Category_MENS_CLOTHING            Category = 6
	Category_MENS_ACCESSORIES         Category = 7
	Category_HOME_GADGETS             Category = 8
	Category_TOYS                     Category = 9
	Category_ELECTRONIC               Category = 10
	Category_DECORATION               Category = 11
	Category_FOOD                     Category = 12
	Category_STATIONERY               Category = 13
	Category_BAGS                     Category = 14
	Category_HARDWARE                 Category = 15
	Category_FURNITURE                Category = 16
	Category_PACKAGING_AND_PRINTING   Category = 17
	Category_BEAUTY_AND_PERSONAL_CARE Category = 18
	Category_CHEMICALS                Category = 19
	Category_GARDEN                   Category = 20
	Category_KITCHEN                  Category = 21
	Category_MACHINERY                Category = 22
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0:  "SPORTS_AND_FITNESS",
		1:  "ELECTRIC",
		2:  "DEVOTIONAL",
		3:  "AGRICULTURAL",
		4:  "WOMENS_CLOTHING",
		5:  "WOMENS_ACCESSORIES",
		6:  "MENS_CLOTHING",
		7:  "MENS_ACCESSORIES",
		8:  "HOME_GADGETS",
		9:  "TOYS",
		10: "ELECTRONIC",
		11: "DECORATION",
		12: "FOOD",
		13: "STATIONERY",
		14: "BAGS",
		15: "HARDWARE",
		16: "FURNITURE",
		17: "PACKAGING_AND_PRINTING",
		18: "BEAUTY_AND_PERSONAL_CARE",
		19: "CHEMICALS",
		20: "GARDEN",
		21: "KITCHEN",
		22: "MACHINERY",
	}
	Category_value = map[string]int32{
		"SPORTS_AND_FITNESS":       0,
		"ELECTRIC":                 1,
		"DEVOTIONAL":               2,
		"AGRICULTURAL":             3,
		"WOMENS_CLOTHING":          4,
		"WOMENS_ACCESSORIES":       5,
		"MENS_CLOTHING":            6,
		"MENS_ACCESSORIES":         7,
		"HOME_GADGETS":             8,
		"TOYS":                     9,
		"ELECTRONIC":               10,
		"DECORATION":               11,
		"FOOD":                     12,
		"STATIONERY":               13,
		"BAGS":                     14,
		"HARDWARE":                 15,
		"FURNITURE":                16,
		"PACKAGING_AND_PRINTING":   17,
		"BEAUTY_AND_PERSONAL_CARE": 18,
		"CHEMICALS":                19,
		"GARDEN":                   20,
		"KITCHEN":                  21,
		"MACHINERY":                22,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_view_common_message_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_view_common_message_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{0}
}

type Ratings int32

const (
	Ratings_VERY_BAD  Ratings = 0
	Ratings_BAD       Ratings = 1
	Ratings_OKAY      Ratings = 2
	Ratings_GOOD      Ratings = 3
	Ratings_VERY_GOOD Ratings = 4
)

// Enum value maps for Ratings.
var (
	Ratings_name = map[int32]string{
		0: "VERY_BAD",
		1: "BAD",
		2: "OKAY",
		3: "GOOD",
		4: "VERY_GOOD",
	}
	Ratings_value = map[string]int32{
		"VERY_BAD":  0,
		"BAD":       1,
		"OKAY":      2,
		"GOOD":      3,
		"VERY_GOOD": 4,
	}
)

func (x Ratings) Enum() *Ratings {
	p := new(Ratings)
	*p = x
	return p
}

func (x Ratings) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ratings) Descriptor() protoreflect.EnumDescriptor {
	return file_view_common_message_proto_enumTypes[1].Descriptor()
}

func (Ratings) Type() protoreflect.EnumType {
	return &file_view_common_message_proto_enumTypes[1]
}

func (x Ratings) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ratings.Descriptor instead.
func (Ratings) EnumDescriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{1}
}

type Status int32

const (
	Status_PENDING    Status = 0
	Status_CANCELED   Status = 1
	Status_CONFORM    Status = 2
	Status_DISPATCHED Status = 3
	Status_DELIVERED  Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "CANCELED",
		2: "CONFORM",
		3: "DISPATCHED",
		4: "DELIVERED",
	}
	Status_value = map[string]int32{
		"PENDING":    0,
		"CANCELED":   1,
		"CONFORM":    2,
		"DISPATCHED": 3,
		"DELIVERED":  4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_view_common_message_proto_enumTypes[2].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_view_common_message_proto_enumTypes[2]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{2}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude string `protobuf:"bytes,1,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude  string `protobuf:"bytes,2,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_common_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_view_common_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *Location) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

type ShopsNearBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId       string     `protobuf:"bytes,3,opt,name=shopId,proto3" json:"shopId,omitempty"`
	ShopName     string     `protobuf:"bytes,4,opt,name=shopName,proto3" json:"shopName,omitempty"`
	PrimaryImage string     `protobuf:"bytes,5,opt,name=primaryImage,proto3" json:"primaryImage,omitempty"`
	Category     []Category `protobuf:"varint,6,rep,packed,name=category,proto3,enum=Category" json:"category,omitempty"`
	Rating       float32    `protobuf:"fixed32,7,opt,name=rating,proto3" json:"rating,omitempty"`
	Shopkeeper   string     `protobuf:"bytes,8,opt,name=shopkeeper,proto3" json:"shopkeeper,omitempty"`
	Location     *Location  `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *ShopsNearBy) Reset() {
	*x = ShopsNearBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_common_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopsNearBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopsNearBy) ProtoMessage() {}

func (x *ShopsNearBy) ProtoReflect() protoreflect.Message {
	mi := &file_view_common_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopsNearBy.ProtoReflect.Descriptor instead.
func (*ShopsNearBy) Descriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{1}
}

func (x *ShopsNearBy) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *ShopsNearBy) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *ShopsNearBy) GetPrimaryImage() string {
	if x != nil {
		return x.PrimaryImage
	}
	return ""
}

func (x *ShopsNearBy) GetCategory() []Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ShopsNearBy) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ShopsNearBy) GetShopkeeper() string {
	if x != nil {
		return x.Shopkeeper
	}
	return ""
}

func (x *ShopsNearBy) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type ProductsOfShopsNearBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    string     `protobuf:"bytes,15,opt,name=productId,proto3" json:"productId,omitempty"`
	ShopId       string     `protobuf:"bytes,16,opt,name=shopId,proto3" json:"shopId,omitempty"`
	ProductName  string     `protobuf:"bytes,17,opt,name=productName,proto3" json:"productName,omitempty"`
	PrimaryImage string     `protobuf:"bytes,18,opt,name=primaryImage,proto3" json:"primaryImage,omitempty"`
	Category     []Category `protobuf:"varint,19,rep,packed,name=category,proto3,enum=Category" json:"category,omitempty"`
	Likes        uint64     `protobuf:"varint,20,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *ProductsOfShopsNearBy) Reset() {
	*x = ProductsOfShopsNearBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_common_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsOfShopsNearBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsOfShopsNearBy) ProtoMessage() {}

func (x *ProductsOfShopsNearBy) ProtoReflect() protoreflect.Message {
	mi := &file_view_common_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsOfShopsNearBy.ProtoReflect.Descriptor instead.
func (*ProductsOfShopsNearBy) Descriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{2}
}

func (x *ProductsOfShopsNearBy) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductsOfShopsNearBy) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *ProductsOfShopsNearBy) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductsOfShopsNearBy) GetPrimaryImage() string {
	if x != nil {
		return x.PrimaryImage
	}
	return ""
}

func (x *ProductsOfShopsNearBy) GetCategory() []Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ProductsOfShopsNearBy) GetLikes() uint64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName      string `protobuf:"bytes,118,opt,name=FullName,proto3" json:"FullName,omitempty"`
	HouseDetails  string `protobuf:"bytes,119,opt,name=HouseDetails,proto3" json:"HouseDetails,omitempty"`
	StreetDetails string `protobuf:"bytes,120,opt,name=StreetDetails,proto3" json:"StreetDetails,omitempty"`
	LandMark      string `protobuf:"bytes,121,opt,name=LandMark,proto3" json:"LandMark,omitempty"`
	PinCode       string `protobuf:"bytes,122,opt,name=PinCode,proto3" json:"PinCode,omitempty"`
	City          string `protobuf:"bytes,123,opt,name=City,proto3" json:"City,omitempty"`
	State         string `protobuf:"bytes,124,opt,name=State,proto3" json:"State,omitempty"`
	Country       string `protobuf:"bytes,125,opt,name=Country,proto3" json:"Country,omitempty"`
	PhoneNo       string `protobuf:"bytes,126,opt,name=PhoneNo,proto3" json:"PhoneNo,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_common_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_view_common_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_view_common_message_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Address) GetHouseDetails() string {
	if x != nil {
		return x.HouseDetails
	}
	return ""
}

func (x *Address) GetStreetDetails() string {
	if x != nil {
		return x.StreetDetails
	}
	return ""
}

func (x *Address) GetLandMark() string {
	if x != nil {
		return x.LandMark
	}
	return ""
}

func (x *Address) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

var File_view_common_message_proto protoreflect.FileDescriptor

var file_view_common_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0xeb, 0x01, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x42,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x70,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68,
	0x6f, 0x70, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xd0, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x4f, 0x66, 0x53, 0x68,
	0x6f, 0x70, 0x73, 0x4e, 0x65, 0x61, 0x72, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x76, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x77, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x72, 0x6b,
	0x18, 0x79, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x72, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x7a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x50, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69,
	0x74, 0x79, 0x18, 0x7b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x7c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x7d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x18, 0x7e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x2a, 0x96, 0x03, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x53, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x46, 0x49, 0x54, 0x4e, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44,
	0x45, 0x56, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x47, 0x52, 0x49, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x57, 0x4f, 0x4d, 0x45, 0x4e, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x4f, 0x4d, 0x45, 0x4e, 0x53, 0x5f, 0x41, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45,
	0x4e, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x14, 0x0a,
	0x10, 0x4d, 0x45, 0x4e, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x4f, 0x52, 0x49, 0x45,
	0x53, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x47, 0x41, 0x44, 0x47,
	0x45, 0x54, 0x53, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x59, 0x53, 0x10, 0x09, 0x12,
	0x0e, 0x0a, 0x0a, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x4f, 0x4e, 0x49, 0x43, 0x10, 0x0a, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x43, 0x4f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4f, 0x44, 0x10, 0x0c, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x45, 0x52, 0x59, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x47,
	0x53, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x41, 0x52, 0x44, 0x57, 0x41, 0x52, 0x45, 0x10,
	0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x52, 0x4e, 0x49, 0x54, 0x55, 0x52, 0x45, 0x10, 0x10,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x50, 0x52, 0x49, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x1c, 0x0a, 0x18,
	0x42, 0x45, 0x41, 0x55, 0x54, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x10, 0x12, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48,
	0x45, 0x4d, 0x49, 0x43, 0x41, 0x4c, 0x53, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x41, 0x52,
	0x44, 0x45, 0x4e, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x49, 0x54, 0x43, 0x48, 0x45, 0x4e,
	0x10, 0x15, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x52, 0x59, 0x10,
	0x16, 0x2a, 0x43, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x56, 0x45, 0x52, 0x59, 0x5f, 0x42, 0x41, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41,
	0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x4b, 0x41, 0x59, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x45, 0x52, 0x59, 0x5f,
	0x47, 0x4f, 0x4f, 0x44, 0x10, 0x04, 0x2a, 0x4f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x4e, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x04, 0x42, 0x30, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x61, 0x61, 0x70, 0x61,
	0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_view_common_message_proto_rawDescOnce sync.Once
	file_view_common_message_proto_rawDescData = file_view_common_message_proto_rawDesc
)

func file_view_common_message_proto_rawDescGZIP() []byte {
	file_view_common_message_proto_rawDescOnce.Do(func() {
		file_view_common_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_view_common_message_proto_rawDescData)
	})
	return file_view_common_message_proto_rawDescData
}

var file_view_common_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_view_common_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_view_common_message_proto_goTypes = []interface{}{
	(Category)(0),                 // 0: Category
	(Ratings)(0),                  // 1: Ratings
	(Status)(0),                   // 2: Status
	(*Location)(nil),              // 3: Location
	(*ShopsNearBy)(nil),           // 4: ShopsNearBy
	(*ProductsOfShopsNearBy)(nil), // 5: ProductsOfShopsNearBy
	(*Address)(nil),               // 6: Address
}
var file_view_common_message_proto_depIdxs = []int32{
	0, // 0: ShopsNearBy.category:type_name -> Category
	3, // 1: ShopsNearBy.location:type_name -> Location
	0, // 2: ProductsOfShopsNearBy.category:type_name -> Category
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_view_common_message_proto_init() }
func file_view_common_message_proto_init() {
	if File_view_common_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_view_common_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_view_common_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopsNearBy); i {
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
		file_view_common_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsOfShopsNearBy); i {
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
		file_view_common_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_view_common_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_view_common_message_proto_goTypes,
		DependencyIndexes: file_view_common_message_proto_depIdxs,
		EnumInfos:         file_view_common_message_proto_enumTypes,
		MessageInfos:      file_view_common_message_proto_msgTypes,
	}.Build()
	File_view_common_message_proto = out.File
	file_view_common_message_proto_rawDesc = nil
	file_view_common_message_proto_goTypes = nil
	file_view_common_message_proto_depIdxs = nil
}
