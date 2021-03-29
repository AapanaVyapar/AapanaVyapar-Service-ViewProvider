package structs

import (
	"aapanavyapar-service-viewprovider/data-base/constants"
	"github.com/go-playground/locales/currency"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

type UserData struct {
	Id        primitive.ObjectID `bson:"_id" validate:"required"`
	UserId    string             `bson:"user_id" validate:"required"`
	UserName  string             `bson:"user_name" validate:"required"`
	Address   Address            `bson:"address"`
	Cart      ShopAndProductIds  `bson:"cart,omitempty"`
	Favorites ShopAndProductIds  `bson:"favorites,omitempty"`
	Orders    ShopAndProductIds  `bson:"orders,omitempty"`
}

type ShopAndProductIds struct {
	Product []string `bson:"products"`
}

type OrderData struct {
	OrderId          primitive.ObjectID  `bson:"_id" validate:"required"`
	UserId           string              `bson:"user_id" validate:"required"`
	Status           constants.Status    `bson:"status" validate:"required"`
	ShopAndProductId string              `bson:"shop_and_product_id" validate:"required"`
	TimeStamp        primitive.Timestamp `bson:"timestamp" validate:"required"`
	Price            string              `bson:"price" validate:"required"`
	Quantity         int64               `bson:"price" validate:"required"`
}

type Address struct {
	FullName      string `bson:"name" validate:"required,min=2,max=100"`
	HouseDetails  string `bson:"house_details" validate:"required,max=100"`
	StreetDetails string `bson:"street_details" validate:"required,max=100"`
	LandMark      string `bson:"land_mark" validate:"required,max=100"`
	PinCode       string `bson:"code" validate:"required,max=6"`
	City          string `bson:"city" validate:"required,max=90"`
	State         string `bson:"state" validate:"required,max=90"`
	Country       string `bson:"country" validate:"required,max=90"`
	PhoneNo       string `bson:"phone_no" validate:"required,max=10"`
}

type Location struct {
	Longitude string `bson:"longitude" validate:"required,longitude"`
	Latitude  string `bson:"latitude" validate:"required,latitude"`
}

type OperationalHours struct {
	Sunday    [2]time.Time `validate:"required"`
	Monday    [2]time.Time `validate:"required"`
	Tuesday   [2]time.Time `validate:"required"`
	Wednesday [2]time.Time `validate:"required"`
	Thursday  [2]time.Time `validate:"required"`
	Friday    [2]time.Time `validate:"required"`
	Saturday  [2]time.Time `validate:"required"`
}

type Rating struct {
	UserId    string              `bson:"user_id" validate:"required"`
	UserName  string              `bson:"user_name" validate:"required"`
	Comment   string              `bson:"comment" validate:"required,max=100"`
	Rating    constants.Ratings   `bson:"rating" validate:"required"`
	Timestamp primitive.Timestamp `bson:"timestamp" validate:"required"`
}

type ShopData struct {
	ShopId              primitive.ObjectID     `bson:"shop_id" validate:"required"`
	ShopName            string                 `bson:"shop_name" validate:"required,max=50"`
	ShopKeeperName      string                 `bson:"shop_keeper_name" validate:"required,min=2,max=100"`
	Images              []url.URL              `bson:"images" validate:"required"`
	PrimaryImages       url.URL                `bson:"primary_images" validate:"required"`
	Address             Address                `bson:"address" validate:"required"`
	Location            Location               `bson:"location" validate:"required"`
	SectorNo            int64                  `bson:"sector_no" validate:"required"`
	Category            []constants.Categories `bson:"category" validate:"required"`
	BusinessInformation string                 `bson:"business_information" validate:"required,max=500"`
	Currency            currency.Type          `bson:"currency" validate:"required"`
	OperationalHours    OperationalHours       `bson:"operational_hours" validate:"required"`
	Ratings             Rating                 `bson:"ratings"`
	Timestamp           primitive.Timestamp    `bson:"timestamp" validate:"required"`
}

type ProductData struct {
	ShopId       primitive.ObjectID  `bson:"shop_id" validate:"required"`
	ProductId    primitive.ObjectID  `bson:"product_id" validate:"required"`
	Title        string              `bson:"title" validate:"required"`
	Description  string              `bson:"description" validate:"required"`
	ShippingInfo string              `bson:"shipping_info" validate:"required"`
	Stock        string              `bson:"stock" validate:"required"`
	Price        string              `bson:"price" validate:"required"`
	Offer        uint8               `bson:"offer" validate:"required"`
	Images       []url.URL           `bson:"images" validate:"required"`
	Timestamp    primitive.Timestamp `bson:"timestamp" validate:"required"`
}

type AnalyticalClickData struct {
	ShopAndProductId string                 `bson:"shop_and_product_id" validate:"required"`
	Timestamp        primitive.Timestamp    `bson:"timestamp" validate:"required"`
	Category         []constants.Categories `bson:"category" validate:"required"`
}

type MostVisited struct {
	Product []AnalyticalClickData `bson:"analytical_click_data" validate:"required"`
}

type AnalyticalData struct {
	UserId      string      `bson:"user_id" validate:"required"`
	MostVisited MostVisited `bson:"most_visited" validate:"required"`
}
