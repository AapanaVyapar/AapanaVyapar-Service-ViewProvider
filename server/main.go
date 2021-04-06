package main

import (
	"aapanavyapar-service-viewprovider/pb"
	"aapanavyapar-service-viewprovider/services"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	service := services.NewViewProviderService(ctx)

	data, err := service.GetTrendingCategories(ctx, &pb.GetTrendingCategoriesRequest{
		ApiKey: os.Getenv("API_KEY_FOR_WEB"),
		Token:  "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	time.Sleep(time.Second * 5)

	data, err = service.GetTrendingCategories(ctx, &pb.GetTrendingCategoriesRequest{
		ApiKey: os.Getenv("API_KEY_FOR_WEB"),
		Token:  "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	//// Shop Testing Start
	//
	//dataInsert := structs.ShopData{
	//	ShopName:       "Milap Stores",
	//	ShopKeeperName: "ABC Person",
	//	Images:         []string{"https://google.com"},
	//	PrimaryImage:   "https://www.google.com",
	//	Address: &structs.Address{
	//		FullName:      "ABC Person",
	//		HouseDetails:  "Milap Store",
	//		StreetDetails: "Mustufa Chishti Colony Main Rd, Panchshil Nagar",
	//		LandMark:      "Milap Store",
	//		PinCode:       "425107",
	//		City:          "chopda",
	//		State:         "maharastra",
	//		Country:       "india",
	//		PhoneNo:       "9890713171",
	//	},
	//	Location: &structs.Location{
	//		Longitude: "21.246435522726177",
	//		Latitude:  "75.29615236552934",
	//	},
	//	Category:            []constants.Categories{constants.MENS_ACCSSORIES, constants.WONENS_CLOTHING},
	//	BusinessInformation: "Famous Seller Of Cloths In Chopda",
	//	OperationalHours: &structs.OperationalHours{
	//		Sunday:    [2]string{"0AM", "0PM"},
	//		Monday:    [2]string{"9AM", "9PM"},
	//		Tuesday:   [2]string{"9AM", "9PM"},
	//		Wednesday: [2]string{"9AM", "9PM"},
	//		Thursday:  [2]string{"9AM", "9PM"},
	//		Friday:    [2]string{"9AM", "9PM"},
	//		Saturday:  [2]string{"9AM", "9PM"},
	//	},
	//	Ratings:   nil,
	//	Timestamp: time.Now().UTC(),
	//}
	//
	//shopId, err := database.CreateShop(ctx, dataInsert)
	//if err != nil {
	//	panic(err)
	//}
	//dataInsert.ShopId = shopId
	//
	//err = database.AddRatingInShopData(ctx, dataInsert.ShopId, structs.Rating{
	//	UserId:    "2",
	//	UserName:  "test",
	//	Comment:   "Hi",
	//	Rating:    constants.VERY_GOOD,
	//	Timestamp: time.Now().UTC(),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddRatingInShopData(ctx, dataInsert.ShopId, structs.Rating{
	//	UserId:    "3",
	//	UserName:  "test",
	//	Comment:   "Hi",
	//	Rating:    constants.VERY_GOOD,
	//	Timestamp: time.Now().UTC(),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//rating, err := database.GetRatingsFromShopData(ctx, dataInsert.ShopId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rating)
	//
	//shop, err := database.GetShopFromShopData(ctx, dataInsert.ShopId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(shop)
	//
	//err = database.DelShopImageFromShopData(ctx, dataInsert.ShopId, "https://www.gooogle.com")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateShopKeeperNameInShopData(ctx, dataInsert.ShopId, "temp")
	//if err != nil {
	//	panic(err)
	//}
	//
	//address2 := structs.Address{
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
	//location := structs.Location{
	//	Longitude: "22.246435522726177",
	//	Latitude:  "75.29615236552934",
	//}
	//
	//err = database.UpdateShopAddressAndLocationInShopData(ctx, dataInsert.ShopId, address2, location)
	//if err != nil {
	//	panic(err)
	//}
	//
	//location1 := structs.Location{
	//	Longitude: "23.246435522726177",
	//	Latitude:  "75.29615236552934",
	//}
	//
	//err = database.UpdateShopLocationInShopData(ctx, dataInsert.ShopId, location1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateCategoryInShopData(ctx, dataInsert.ShopId, []constants.Categories{constants.ELECTRIC, constants.ELECTRONIC})
	//if err != nil {
	//	panic(err)
	//}
	//
	//categories, err := database.GetCategoryFromShopData(ctx, dataInsert.ShopId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(categories)
	//
	//err = database.UpdateBusinessInfoInShopData(ctx, dataInsert.ShopId, "HELLO BUSINESS")
	//if err != nil {
	//	panic(err)
	//}
	//
	//hours := structs.OperationalHours{
	//	Sunday:    [2]string{"0AM", "0PM"},
	//	Monday:    [2]string{"12AM", "9PM"},
	//	Tuesday:   [2]string{"12AM", "9PM"},
	//	Wednesday: [2]string{"12AM", "9PM"},
	//	Thursday:  [2]string{"12AM", "9PM"},
	//	Friday:    [2]string{"12AM", "9PM"},
	//	Saturday:  [2]string{"12AM", "9PM"},
	//}
	//
	//err = database.UpdateOperationalHoursInShopData(ctx, dataInsert.ShopId, hours)
	//if err != nil {
	//	panic(err)
	//}
	//
	//shop, err = database.GetShopFromShopData(ctx, dataInsert.ShopId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(shop)
	//
	////err = database.DelShopFromShopData(ctx, dataInsert.ShopId)
	////if err != nil {
	////	panic(err)
	////}
	//
	//// Shop Testing Ends
	//
	//// Products Testing Starts
	//
	//dataProduct1 := structs.ProductData{
	//	ShopId:           dataInsert.ShopId,
	//	Title:            "Yellow Shirt",
	//	ShortDescription: "Best In Class Only",
	//	Description:      "Cotton Fabric Size XL",
	//	ShippingInfo:     "200x70x10",
	//	Stock:            10,
	//	Price:            100,
	//	Offer:            10,
	//	Category:         []constants.Categories{constants.MENS_CLOTHING},
	//	Images:           []string{"https://image.com"},
	//	Timestamp:        time.Now().UTC(),
	//}
	//productId1, err := database.CreateProduct(ctx, dataProduct1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//dataProduct2 := structs.ProductData{
	//	ShopId:           dataInsert.ShopId,
	//	Title:            "BLACK Shirt",
	//	ShortDescription: "Best In Class Only",
	//	Description:      "Silk Fabric Size XL",
	//	ShippingInfo:     "200x70x10",
	//	Stock:            10,
	//	Price:            100,
	//	Offer:            10,
	//	Images:           []string{"https://image.com"},
	//	Category:         []constants.Categories{constants.MENS_CLOTHING},
	//	Timestamp:        time.Now().UTC(),
	//}
	//
	//productId2, err := database.CreateProduct(ctx, dataProduct2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.IncreaseStockFromProductData(ctx, productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.GetAllProductsOfShopByFunctionFromProductData(ctx, dataInsert.ShopId, func(data structs.ProductData) error {
	//	// Here send the data to client in stream one by one if error occurred while sending then return form here.
	//	fmt.Println(data)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddProductImageInProductData(ctx, dataInsert.ShopId, productId1, "https://imageurl.in")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductTitleInProductData(ctx, dataInsert.ShopId, productId1, "Orange Shirt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductDescriptionInProductData(ctx, dataInsert.ShopId, productId1, "Best")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductShortDescriptionInProductData(ctx, dataInsert.ShopId, productId1, "short Description")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductOfferInProductData(ctx, dataInsert.ShopId, productId1, 25)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductPriceInProductData(ctx, dataInsert.ShopId, productId1, 100)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductShippingInfoInProductData(ctx, dataInsert.ShopId, productId1, "500x500")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductStockInfoInProductData(ctx, dataInsert.ShopId, productId1, 10)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelProductImageFromProductData(ctx, dataInsert.ShopId, productId1, "https://imageurl.in")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.GetAllProductsByCategoryOfShopsByFunctionFromProductData(ctx, []primitive.ObjectID{dataInsert.ShopId}, []constants.Categories{constants.WONENS_CLOTHING, constants.MENS_CLOTHING}, func(data structs.ProductData) error {
	//	// Here send the data to client in stream one by one if error occurred while sending then return form here.
	//	fmt.Println("Category wise :", data)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.UpdateProductCategoryInProductData(ctx, dataInsert.ShopId, productId1, []constants.Categories{constants.MENS_CLOTHING, constants.WONENS_CLOTHING})
	//if err != nil {
	//	panic(err)
	//}
	//
	//dataP, err := database.GetProductFromProductData(ctx, productId1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(dataP)
	//
	//dataP, err = database.GetSpecificProductsOfShopFromProductData(ctx, dataInsert.ShopId, productId1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(dataP)
	//
	////err = database.DelProductFromProductData(ctx, dataInsert.ShopId, productId1)
	////if err != nil {
	////	panic(err)
	////}
	////
	//
	//// Products Testing Ends
	//
	//// User Testing Starts
	//
	//_, err = database.CreateUser(ctx, "1", "test")
	//if err != nil {
	//	panic(err)
	//}
	//
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
	//err = database.SetAddressInUserData(ctx, "101", "abc", address)
	//if err != nil {
	//	panic(err)
	//}
	//
	//address1, err := database.GetAddressUserData(ctx, "1")
	//if err != nil {
	//	panic(err)
	//}
	//if address1 == nil {
	//	fmt.Println("Party")
	//}
	//
	//err = database.DelAddressInUserData(ctx, "101")
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = database.CreateUser(ctx, "11", "test")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddToCartUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddToCartUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.AddToCartUserData(ctx, "11", productId2)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.RemoveFromCartUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelFromCartUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//products, err := database.GetCartUserData(ctx, "11")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(products)
	//
	////	Favorite
	//err = database.AddToFavoritesUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelFromFavoritesUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = database.DelFromOrdersUserData(ctx, "11", productId1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// user testing ends
	//
	//// Order Testing Starts
	//
	//orderId, err := database.CreateOrder(ctx, "1", productId1, 3, 10, &address)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(orderId)
	//
	//err = database.UpdateOrderStatusInOrderData(ctx, orderId, constants.CONFORM)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(orderId)
	//
	//orderData, err := database.GetOrderInfoFromOrderData(ctx, orderId)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Order Data : ", orderData)
	//
	////	Order
	//err = database.AddToOrdersUserData(ctx, "11", orderId)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Order Testing Ends
	//
	//err = database.AddAnalyticalDataToAnalyticalData(ctx, "1", structs.AnalyticalClickData{
	//	ProductId: primitive.NewObjectID(),
	//	Timestamp: time.Now().UTC(),
	//	Category:  []constants.Categories{constants.MENS_CLOTHING},
	//},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
}
