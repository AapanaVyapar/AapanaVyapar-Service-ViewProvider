package structs

import (
	"aapanavyapar-service-viewprovider/data-base/constants"
	"github.com/go-playground/locales/currency"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserData struct {
	UserId    string                         `bson:"_id" json:"_id" validate:"required"`
	UserName  string                         `bson:"user_name" json:"user_name" validate:"required"`
	Address   *Address                       `bson:"address,omitempty" json:"address"`
	Cart      *ShopAndProductIdsForCart      `bson:"cart,omitempty" json:"cart"`
	Favorites *ShopAndProductIdsForFavAndOrd `bson:"favorites,omitempty" json:"favorites"`
	Orders    *ShopAndProductIdsForFavAndOrd `bson:"orders,omitempty" json:"orders"`
}

type ProductCartData struct {
	ProductId   string `bson:"product_id" json:"product_id"`
	NoOfProduct uint8  `bson:"no_product" json:"no_product" validate:"required,min=1,max=200"`
}

type ShopAndProductIdsForCart struct {
	Product []ProductCartData `bson:"products,omitempty" json:"products"`
}

type ShopAndProductIdsForFavAndOrd struct {
	Product []string `bson:"products,omitempty" json:"products"`
}

type OrderData struct {
	OrderId          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserId           string             `bson:"user_id" json:"user_id" validate:"required"`
	Status           constants.Status   `bson:"status" json:"status" validate:"required"`
	ShopAndProductId string             `bson:"shop_and_product_id" json:"shop_and_product_id" validate:"required"`
	TimeStamp        time.Time          `bson:"timestamp" json:"timestamp" validate:"required"`
	Price            float64            `bson:"price" json:"price" validate:"required"`
	Quantity         int32              `bson:"quantity" json:"quantity" validate:"required"`
}

type Address struct {
	FullName      string `bson:"full_name" json:"name" validate:"required,min=2,max=100"`
	HouseDetails  string `bson:"house_details" json:"house_details" validate:"required,max=100"`
	StreetDetails string `bson:"street_details" json:"street_details" validate:"required,max=100"`
	LandMark      string `bson:"land_mark" json:"land_mark" validate:"required,max=100"`
	PinCode       string `bson:"code" json:"code" validate:"required,max=6"`
	City          string `bson:"city" json:"city" validate:"required,max=90"`
	State         string `bson:"state" json:"state" validate:"required,max=90"`
	Country       string `bson:"country" json:"country" validate:"required,max=90"`
	PhoneNo       string `bson:"phone_no" json:"phone_no" validate:"required,max=10"`
}

type Location struct {
	Longitude string `bson:"longitude" json:"longitude" validate:"required,longitude"`
	Latitude  string `bson:"latitude" json:"latitude" validate:"required,latitude"`
}

type OperationalHours struct {
	Sunday    [2]string `json:"sunday" validate:"required"`
	Monday    [2]string `json:"monday" validate:"required"`
	Tuesday   [2]string `json:"tuesday" validate:"required"`
	Wednesday [2]string `json:"wednesday" validate:"required"`
	Thursday  [2]string `json:"thursday" validate:"required"`
	Friday    [2]string `json:"friday" validate:"required"`
	Saturday  [2]string `json:"saturday" validate:"required"`
}

type Rating struct {
	UserId    string            `bson:"user_id" json:"user_id" validate:"required"`
	UserName  string            `bson:"user_name" json:"user_name" validate:"required"`
	Comment   string            `bson:"comment" json:"comment" validate:"required,max=100"`
	Rating    constants.Ratings `bson:"rating" json:"rating" validate:"required"`
	Timestamp time.Time         `bson:"timestamp" json:"timestamp" validate:"required"`
}

type ShopData struct {
	ShopId              primitive.ObjectID     `bson:"_id,omitempty" json:"_id"`
	ShopName            string                 `bson:"shop_name" json:"shop_name" validate:"required,max=50"`
	ShopKeeperName      string                 `bson:"shop_keeper_name" json:"shop_keeper_name" validate:"required,min=2,max=100"`
	Images              []string               `bson:"images" json:"images" validate:"required"`
	PrimaryImage        string                 `bson:"primary_image" json:"primary_images" validate:"required,url"`
	Address             *Address               `bson:"address" json:"address" validate:"required"`
	Location            *Location              `bson:"location" json:"location" validate:"required"`
	SectorNo            int64                  `bson:"sector_no" json:"sector_no"`
	Category            []constants.Categories `bson:"category" json:"category" validate:"required"`
	BusinessInformation string                 `bson:"business_information" json:"business_information" validate:"required,max=500"`
	Currency            currency.Type          `bson:"currency" json:"currency" validate:"required"`
	OperationalHours    *OperationalHours      `bson:"operational_hours" json:"operational_hours" validate:"required"`
	Ratings             *[]Rating              `bson:"ratings,omitempty" json:"ratings"`
	Timestamp           time.Time              `bson:"timestamp" json:"timestamp" validate:"required"`
}

type ProductData struct {
	ProductId    primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ShopId       primitive.ObjectID `bson:"shop_id" json:"shop_id" validate:"required"`
	Title        string             `bson:"title" json:"title" validate:"required"`
	Description  string             `bson:"description" json:"description" validate:"required"`
	ShippingInfo string             `bson:"shipping_info" json:"shipping_info" validate:"required"`
	Stock        uint32             `bson:"stock" json:"stock" validate:"required"`
	Price        float64            `bson:"price" json:"price" validate:"required"`
	Offer        uint8              `bson:"offer" json:"offer" validate:"required,max=100"`
	Images       []string           `bson:"images" json:"images" validate:"required"`
	Timestamp    time.Time          `bson:"timestamp" json:"timestamp" validate:"required"`
}

type AnalyticalClickData struct {
	ShopAndProductId string                 `bson:"shop_and_product_id" json:"shop_and_product_id" validate:"required"`
	Timestamp        time.Time              `bson:"timestamp" json:"timestamp" validate:"required"`
	Category         []constants.Categories `bson:"category" json:"category" validate:"required"`
}

type MostVisited struct {
	Product []AnalyticalClickData `bson:"analytical_click_data,omitempty" json:"analytical_click_data" validate:"required"`
}

type AnalyticalData struct {
	UserId      string       `bson:"user_id" json:"user_id" validate:"required"`
	MostVisited *MostVisited `bson:"most_visited,omitempty" json:"most_visited" validate:"required"`
}
