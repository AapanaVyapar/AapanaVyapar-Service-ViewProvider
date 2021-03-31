package main

import (
	"aapanavyapar-service-viewprovider/data-base/constants"
	data_base "aapanavyapar-service-viewprovider/data-base/data-services"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"github.com/go-playground/locales/currency"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {
	database := data_base.NewDataBase()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	//_, err := database.CreateProduct(context.Background(), "1")
	//_, err = database.CreateShop(context.Background(), "1")
	//_, err = database.CreateUser(ctx, "1")
	//if err != nil {
	//	panic(err)
	//}

	//address := structs.Address{
	//	FullName:      "Shitij Shailendra Agrawal",
	//	HouseDetails:  "B.K Road Chopda",
	//	StreetDetails: "B.K Road Chopda",
	//	LandMark:      "HDFC Bank",
	//	PinCode:       "425107",
	//	City:          "Chopda",
	//	State:         "Maharastra",
	//	Country:       "India",
	//	PhoneNo:       "9172879779",
	//}
	//
	//err := database.SetAddressInUserData(ctx, "101", "", address)
	//if err != nil {
	//	panic(err)
	//}
	//
	//address1 := database.GetAddressUserData(ctx, "1")
	//if address1 == nil {
	//	fmt.Println("Party")
	//}
	//
	//err = database.DelAddressInUserData(ctx, "100")
	//if err != nil {
	//	panic(err)
	//}

	//_, err := database.CreateUser(ctx, "11", "test")
	//if err != nil {
	//			panic(err)
	//}

	//err = database.AddToCartUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddToCartUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddToCartUserData(ctx, "11", "2")
	//if err != nil {
	//	panic(err)
	//}

	//err = database.RemoveFromCartUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}

	//err = database.DelFromCartUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}

	//products, err := database.GetCartUserData(ctx, "11")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(products)

	//Favorite
	//err = database.AddToFavoritesUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelFromFavoritesUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}

	//Order
	//err = database.AddToOrdersUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelFromOrdersUserData(ctx, "11", "1")
	//if err != nil {
	//	panic(err)
	//}

	dataInsert := structs.ShopData{
		ShopId:         primitive.NewObjectID(),
		ShopName:       "Milap Stores",
		ShopKeeperName: "ABC Person",
		Images:         []string{"https://google.com"},
		PrimaryImage:   "https://www.google.com",
		Address: &structs.Address{
			FullName:      "ABC Person",
			HouseDetails:  "Milap Store",
			StreetDetails: "Mustufa Chishti Colony Main Rd, Panchshil Nagar",
			LandMark:      "Milap Store",
			PinCode:       "425107",
			City:          "chopda",
			State:         "maharastra",
			Country:       "india",
			PhoneNo:       "9890713171",
		},
		Location: &structs.Location{
			Longitude: "21.246435522726177",
			Latitude:  "75.29615236552934",
		},
		Category:            []constants.Categories{constants.MENS_ACCSSORIES, constants.WONENS_CLOTHING},
		BusinessInformation: "Famous Seller Of Cloths In Chopda",
		Currency:            currency.INR,
		OperationalHours: &structs.OperationalHours{
			Sunday:    [2]string{"0AM", "0PM"},
			Monday:    [2]string{"9AM", "9PM"},
			Tuesday:   [2]string{"9AM", "9PM"},
			Wednesday: [2]string{"9AM", "9PM"},
			Thursday:  [2]string{"9AM", "9PM"},
			Friday:    [2]string{"9AM", "9PM"},
			Saturday:  [2]string{"9AM", "9PM"},
		},
		Ratings:   nil,
		Timestamp: time.Now().UTC(),
	}

	err := database.CreateShop(ctx, dataInsert)
	if err != nil {
		panic(err)
	}

	err = database.AddRatingInShopData(ctx, dataInsert.ShopId, structs.Rating{
		UserId:    "2",
		UserName:  "test",
		Comment:   "Hi",
		Rating:    constants.VERY_GOOD,
		Timestamp: time.Now().UTC(),
	})
	if err != nil {
		panic(err)
	}

	err = database.AddRatingInShopData(ctx, dataInsert.ShopId, structs.Rating{
		UserId:    "3",
		UserName:  "test",
		Comment:   "Hi",
		Rating:    constants.VERY_GOOD,
		Timestamp: time.Now().UTC(),
	})
	if err != nil {
		panic(err)
	}

	rating, err := database.GetRatingsFromShopData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}
	fmt.Println(rating)

	shop, err := database.GetShopFromShopData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}
	fmt.Println(shop)

	err = database.DelShopImageFromShopData(ctx, dataInsert.ShopId, "https://www.gooogle.com")
	if err != nil {
		panic(err)
	}

	err = database.UpdateShopKeeperNameInShopData(ctx, dataInsert.ShopId, "temp")
	if err != nil {
		panic(err)
	}

	address1 := structs.Address{
		FullName:      "Shitij Shailendra Agrawal",
		HouseDetails:  "B.K Road Chopda",
		StreetDetails: "B.K Road Chopda",
		LandMark:      "HDFC Bank",
		PinCode:       "425107",
		City:          "Chopda",
		State:         "Maharastra",
		Country:       "India",
		PhoneNo:       "9172879779",
	}

	location := structs.Location{
		Longitude: "22.246435522726177",
		Latitude:  "75.29615236552934",
	}

	err = database.UpdateShopAddressAndLocationInShopData(ctx, dataInsert.ShopId, address1, location)
	if err != nil {
		panic(err)
	}

	location1 := structs.Location{
		Longitude: "23.246435522726177",
		Latitude:  "75.29615236552934",
	}

	err = database.UpdateShopLocationInShopData(ctx, dataInsert.ShopId, location1)
	if err != nil {
		panic(err)
	}

	err = database.UpdateCategoryInShopData(ctx, dataInsert.ShopId, []constants.Categories{constants.ELECTRIC, constants.ELECTRONIC})
	if err != nil {
		panic(err)
	}

	categories, err := database.GetCategoryFromShopData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}
	fmt.Println(categories)

	err = database.UpdateBusinessInfoInShopData(ctx, dataInsert.ShopId, "HELLO BUSINESS")
	if err != nil {
		panic(err)
	}

	err = database.UpdateCurrencyInShopData(ctx, dataInsert.ShopId, currency.USD)
	if err != nil {
		panic(err)
	}

	hours := structs.OperationalHours{
		Sunday:    [2]string{"0AM", "0PM"},
		Monday:    [2]string{"12AM", "9PM"},
		Tuesday:   [2]string{"12AM", "9PM"},
		Wednesday: [2]string{"12AM", "9PM"},
		Thursday:  [2]string{"12AM", "9PM"},
		Friday:    [2]string{"12AM", "9PM"},
		Saturday:  [2]string{"12AM", "9PM"},
	}

	err = database.UpdateOperationalHoursInShopData(ctx, dataInsert.ShopId, hours)
	if err != nil {
		panic(err)
	}

	shop, err = database.GetShopFromShopData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}
	fmt.Println(shop)

	dataProduct1 := structs.ProductData{
		ShopId:       dataInsert.ShopId,
		ProductId:    primitive.NewObjectID(),
		Title:        "Yellow Shirt",
		Description:  "Best in Class Size XL",
		ShippingInfo: "200x70x10",
		Stock:        10,
		Price:        100,
		Offer:        10,
		Images:       []string{"https://image.com"},
		Timestamp:    time.Now().UTC(),
	}
	err = database.CreateProduct(ctx, dataInsert.ShopId, dataProduct1)
	if err != nil {
		panic(err)
	}

	dataProduct2 := structs.ProductData{
		ShopId:       dataInsert.ShopId,
		ProductId:    primitive.NewObjectID(),
		Title:        "BLACK Shirt",
		Description:  "Best in Class Size XL",
		ShippingInfo: "200x70x10",
		Stock:        10,
		Price:        100,
		Offer:        10,
		Images:       []string{"https://image.com"},
		Timestamp:    time.Now().UTC(),
	}

	err = database.CreateProduct(ctx, dataInsert.ShopId, dataProduct2)
	if err != nil {
		panic(err)
	}

	err = database.GetAllProductsOfShopByFunctionFromProductData(ctx, dataInsert.ShopId, func(data structs.ProductData) error {
		// Here send the data to client in stream one by one if error occurred while sending then return form here.
		fmt.Println(data)
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = database.AddProductImageInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, "https://imageurl.in")
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductTitleInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, "Orange Shirt")
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductDescriptionInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, "Best")
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductOfferInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, 25)
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductPriceInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, 100)
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductShippingInfoInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, "500x500")
	if err != nil {
		panic(err)
	}

	err = database.UpdateProductStockInfoInProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, 50)
	if err != nil {
		panic(err)
	}

	data, err := database.GetAllProductsOfShopByArrayFromProductData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	err = database.DelProductImageFromProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId, "https://imageurl.in")
	if err != nil {
		panic(err)
	}

	dataP, err := database.GetSpecificProductsOfShopFromProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId)
	if err != nil {
		panic(err)
	}
	fmt.Println(dataP)

	err = database.DelProductFromProductData(ctx, dataInsert.ShopId, dataProduct1.ProductId)
	if err != nil {
		panic(err)
	}

	err = database.DelShopFromShopData(ctx, dataInsert.ShopId)
	if err != nil {
		panic(err)
	}

}
