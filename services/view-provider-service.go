package services

import (
	redisDataBase "aapanavyapar-service-viewprovider/data-base/cash-services"
	mongoDataBase "aapanavyapar-service-viewprovider/data-base/data-services"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ViewProviderService struct {
	Data *mongoDataBase.MongoDataBase
	Cash *redisDataBase.CashDataBase
}

func NewViewProviderService() *ViewProviderService {
	mongoData := mongoDataBase.NewDataBase()
	redisData := redisDataBase.NewDataBase()

	view := ViewProviderService{
		Data: mongoData,
		Cash: redisData,
	}

	err := view.CreateShopSchemaInCash()
	if err != nil {
		panic(err)
	}

	err = view.CreateProductSchemaInCash()
	if err != nil {
		panic(err)
	}

	doc := view.Cash.CreateShopDocument("f38d6a51-b961-474b-9be1-6de62ab5c57e", "Milap Store", "https://image.com", []pb.Category{pb.Category_MENS_CLOTHING}, 3.3, "ABC Person", 21.246435522726177, 75.29615236552934)
	doc1 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-9be1-jn362ab5c57e", "Rajkumar Coldrinks", "https://image.com", []pb.Category{pb.Category_FOOD}, 4.2, "ABC Person", 21.246703671726266, 75.29363042248966)
	doc2 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-d4g2-jn362ab5c57e", "Laxmi Offset", "https://image.com", []pb.Category{pb.Category_FOOD}, 5.0, "ABC Person", 21.246863857271922, 75.29460006862273)
	doc3 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-dv1e-jn362ab5c57e", "raza cutlery stores", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.24691956642532, 75.2936569970796)
	doc4 := view.Cash.CreateShopDocument("f38d6a51-b932-474b-dv1e-jn362ab5c57e", "joshi electrical works", "https://image.com", []pb.Category{pb.Category_ELECTRIC}, 3.0, "ABC Person", 21.24663669715862, 75.29432756403227)
	doc5 := view.Cash.CreateShopDocument("f38d6a51-4432-474b-dv1e-jn362ab5c57e", "chintamani kirana stores", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.24645703702037, 75.29347243866049)
	doc6 := view.Cash.CreateShopDocument("f38d6a32-4432-474b-dv1e-jn362ab5c57e", "Chandu Dairy", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.246983732987587, 75.29470401626448)
	doc7 := view.Cash.CreateShopDocument("f68d6a32-4432-474b-dv1e-jn362ab5c57e", "Vegetable Market", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.245561629488762, 75.29823368278593)

	// Index the document. The API accepts multiple documents at a time
	if err := view.Cash.ShopClient.Index([]redisearch.Document{doc, doc1, doc2, doc3, doc4, doc5, doc6, doc7}...); err != nil {
		log.Fatal(err)
	}

	docProd := view.Cash.CreateProductDocument("1", "f38d6a51-b961-474b-9be1-6de62ab5c57e", "T-Shirt", "https://image.com", []pb.Category{pb.Category_MENS_CLOTHING}, 20)
	doc1Prod := view.Cash.CreateProductDocument("2", "f38d6a51-b961-474b-9be1-jn362ab5c57e", "Samosa", "https://image.com", []pb.Category{pb.Category_FOOD}, 40)
	doc2Prod := view.Cash.CreateProductDocument("3", "f38d6a51-b961-474b-d4g2-jn362ab5c57e", "Jalabi", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	doc3Prod := view.Cash.CreateProductDocument("4", "f38d6a51-b961-474b-dv1e-jn362ab5c57e", "Kachori", "https://image.com", []pb.Category{pb.Category_FOOD}, 50)
	doc4Prod := view.Cash.CreateProductDocument("5", "f38d6a51-b932-474b-dv1e-jn362ab5c57e", "Motor", "https://image.com", []pb.Category{pb.Category_ELECTRIC}, 60)
	doc5Prod := view.Cash.CreateProductDocument("6", "f38d6a51-4432-474b-dv1e-jn362ab5c57e", "broom", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	doc6Prod := view.Cash.CreateProductDocument("7", "f38d6a32-4432-474b-dv1e-jn362ab5c57e", "milk", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	doc7Prod := view.Cash.CreateProductDocument("8", "f68d6a32-4432-474b-dv1e-jn362ab5c57e", "vegetable", "https://image.com", []pb.Category{pb.Category_FOOD}, 60)

	// Index the document. The API accepts multiple documents at a time
	if err := view.Cash.ProductClient.Index([]redisearch.Document{docProd, doc1Prod, doc2Prod, doc3Prod, doc4Prod, doc5Prod, doc6Prod, doc7Prod}...); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	err = view.LoadBasicCategoriesInCash(ctx)
	if err != nil {
		panic(err)
	}

	err = view.LoadShopsInCash(ctx)
	if err != nil {
		panic(err)
	}

	err = view.LoadProductsInCash(ctx)
	if err != nil {
		panic(err)
	}

	return &view
}

func (viewServer *ViewProviderService) GetTrendingShops(request *pb.GetTrendingShopsRequest, stream pb.ViewProviderService_GetTrendingShopsServer) error {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	latitude, err := strconv.ParseFloat(request.Location.Latitude, 64)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid Location")
	}

	longitude, err := strconv.ParseFloat(request.Location.Longitude, 64)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid Location")
	}

	meter, err := strconv.ParseFloat(request.GetDistanceInMeter(), 64)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid Location")
	}

	fmt.Println("HI")

	docs, err := viewServer.Cash.GetShopByLocation(latitude, longitude, meter, 20)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Unable To Get Data For Shop")
	}

	fmt.Println("HI")

	for _, doc := range docs {

		fmt.Println(doc)

		location := strings.Split(doc.Properties["location"].(string), ",")

		str := doc.Properties["categoryOfShop"].(string)[1:]
		str = str[:len(str)-1]
		data := strings.Split(str, ",")
		var category []pb.Category
		for _, cat := range data {
			category = append(category, pb.Category(pb.Category_value[cat]))
		}
		rating, err := strconv.ParseFloat(doc.Properties["ratingOfShop"].(string), 32)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.GetTrendingShopsResponse{
			Shops: &pb.ShopsNearBy{
				ShopId:       strings.ReplaceAll(doc.Id[5:], " ", "-"),
				ShopName:     doc.Properties["shopName"].(string),
				PrimaryImage: doc.Properties["primaryImage"].(string),
				Category:     category,
				Rating:       float32(rating),
				Shopkeeper:   doc.Properties["shopkeeper"].(string),
				Location: &pb.Location{
					Latitude:  location[0],
					Longitude: location[1],
				},
			},
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "Stream Error", err)
		}
	}

	return nil

}

func (viewServer *ViewProviderService) GetTrendingProductsByShop(request *pb.GetTrendingProductsByShopRequest, stream pb.ViewProviderService_GetTrendingProductsByShopServer) error {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	docs, err := viewServer.Cash.GetProductsByShopsSortByRating(request.GetShopId(), 100)
	if err != nil {
		return status.Errorf(codes.Unknown, "Unable To Provide Data For Given Shops")
	}

	for _, doc := range docs {

		fmt.Println(doc)

		str := doc.Properties["categoryOfProduct"].(string)[1:]
		str = str[:len(str)-1]
		data := strings.Split(str, ",")
		var category []pb.Category
		for _, cat := range data {
			category = append(category, pb.Category(pb.Category_value[cat]))
		}

		likes, err := strconv.ParseUint(doc.Properties["likesOfProduct"].(string), 10, 64)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.GetTrendingProductsByShopResponse{
			CategoryData: &pb.ProductsOfShopsNearBy{
				ProductId:    doc.Id[7:],
				ShopId:       strings.ReplaceAll(doc.Properties["shopId"].(string), " ", "-"),
				ProductName:  doc.Properties["productName"].(string),
				PrimaryImage: doc.Properties["primaryImage"].(string),
				Category:     category,
				Likes:        likes,
			},
		},
		)
		if err != nil {
			return status.Errorf(codes.Unknown, "Stream Error", err)
		}
	}

	return nil
}

func (viewServer *ViewProviderService) AddToLikeProduct(context context.Context, request *pb.AddToLikeProductRequest) (*pb.AddToLikeProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println("Add Like : ", receivedToken)

	productByte, err := viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
	}

	err = viewServer.Cash.AddToFavStream(context, receivedToken.Audience, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Add Like To Product")
	}

	var product structs.ProductData
	structs.UnmarshalProductData([]byte(productByte), &product)
	product.Likes += 1

	err = viewServer.Cash.AddProductDataToCash(context, request.GetProductId(), product.Marshal())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Add Like")
	}
	return &pb.AddToLikeProductResponse{Status: true}, nil

}

func (viewServer *ViewProviderService) RemoveFromLikeProduct(context context.Context, request *pb.RemoveFromLikeProductRequest) (*pb.RemoveFromLikeProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println("Add Like : ", receivedToken)

	productByte, err := viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
	}

	err = viewServer.Cash.DelToFavStream(context, receivedToken.Audience, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Remove Like To Product")
	}

	var product structs.ProductData
	structs.UnmarshalProductData([]byte(productByte), &product)
	product.Likes -= 1

	err = viewServer.Cash.AddProductDataToCash(context, request.GetProductId(), product.Marshal())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To UnLike")
	}
	return &pb.RemoveFromLikeProductResponse{Status: true}, nil

}

func (viewServer *ViewProviderService) AddToCartProduct(context context.Context, request *pb.AddToCartProductRequest) (*pb.AddToCartProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println(receivedToken)

	_, err = viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
	}

	err = viewServer.Cash.AddToCartStream(context, receivedToken.Audience, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Add Product To Cart")
	}

	return &pb.AddToCartProductResponse{Status: true}, nil
}

func (viewServer *ViewProviderService) RemoveFromCartProduct(context context.Context, request *pb.RemoveFromCartProductRequest) (*pb.RemoveFromCartProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println(receivedToken)

	err = viewServer.Cash.DelToCartStream(context, receivedToken.Audience, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Remove Product From Cart")
	}

	return &pb.RemoveFromCartProductResponse{Status: true}, nil
}

//
//func (viewServer *ViewProviderService) GetTrendingCategories(context context.Context, request *pb.GetTrendingCategoriesRequest) (*pb.GetTrendingCategoriesResponse, error) {
//
//	if !helpers.CheckForAPIKey(request.GetApiKey()) {
//		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
//	}
//
//	//receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
//	//}
//	//
//	//fmt.Println(receivedToken)
//
//	keys := viewServer.Cash.Cash.HKeys(context, "categories")
//
//	return &pb.GetTrendingCategoriesResponse{
//		CategoryData: keys.Val(),
//	}, nil
//}
//
//func (viewServer *ViewProviderService) GetTrendingDataByCategories(request *pb.GetTrendingDataByCategoriesRequest, stream pb.ViewProviderService_GetTrendingDataByCategoriesServer) error {
//
//	location := structs.Location{
//		Longitude: request.Location.Longitude,
//		Latitude:  request.Location.Latitude,
//	}
//
//	if err := helpers.Validate(location); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (viewServer *ViewProviderService) AddToLikeProduct(context context.Context, request *pb.AddToLikeProductRequest) (*pb.AddToLikeProductResponse, error) {
//	if !helpers.CheckForAPIKey(request.GetApiKey()) {
//		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
//	}
//
//	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
//	}
//
//	fmt.Println("Add Like : ", receivedToken)
//
//	productByte, err := viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
//	}
//
//	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
//	//}
//	//
//	//err = viewServer.Data.AddToFavoritesUserData(context, receivedToken.Audience, productId)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.AlreadyExists, "You Already Like To Product Or Unable To Process Request At A Movement")
//	//}
//
//	err = viewServer.Cash.AddToFavStream(context, receivedToken.Audience, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To Add Like To Product")
//	}
//
//	var product structs.ProductData
//	structs.UnmarshalProductData([]byte(productByte), &product)
//	product.Likes += 1
//
//	err = viewServer.Cash.AddProductDataToCash(context, request.GetProductId(), product.Marshal())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To Add Like")
//	}
//	return &pb.AddToLikeProductResponse{Status: true}, nil
//
//}
//
//func (viewServer *ViewProviderService) RemoveFromLikeProduct(context context.Context, request *pb.RemoveFromLikeProductRequest) (*pb.RemoveFromLikeProductResponse, error) {
//	if !helpers.CheckForAPIKey(request.GetApiKey()) {
//		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
//	}
//
//	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
//	}
//
//	fmt.Println("Add Like : ", receivedToken)
//
//	productByte, err := viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
//	}
//
//	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
//	//}
//	//
//	//err = viewServer.Data.DelFromFavoritesUserData(context, receivedToken.Audience, productId)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Unlike The Product")
//	//}
//
//	err = viewServer.Cash.DelToFavStream(context, receivedToken.Audience, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To Remove Like To Product")
//	}
//
//	var product structs.ProductData
//	structs.UnmarshalProductData([]byte(productByte), &product)
//	product.Likes -= 1
//
//	err = viewServer.Cash.AddProductDataToCash(context, request.GetProductId(), product.Marshal())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To UnLike")
//	}
//	return &pb.RemoveFromLikeProductResponse{Status: true}, nil
//
//}
//
//func (viewServer *ViewProviderService) AddToCartProduct(context context.Context, request *pb.AddToCartProductRequest) (*pb.AddToCartProductResponse, error) {
//	if !helpers.CheckForAPIKey(request.GetApiKey()) {
//		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
//	}
//
//	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
//	}
//
//	fmt.Println(receivedToken)
//
//	_, err = viewServer.Cash.GetProductDataFromCash(context, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
//	}
//
//	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
//	//}
//	//
//	//err = viewServer.Data.AddToCartUserData(context, receivedToken.Audience, productId)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Add To Cart")
//	//}
//
//	err = viewServer.Cash.AddToCartStream(context, receivedToken.Audience, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To Add Product To Cart")
//	}
//
//	return &pb.AddToCartProductResponse{Status: true}, nil
//}
//
//func (viewServer *ViewProviderService) RemoveFromCartProduct(context context.Context, request *pb.RemoveFromCartProductRequest) (*pb.RemoveFromCartProductResponse, error) {
//	if !helpers.CheckForAPIKey(request.GetApiKey()) {
//		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
//	}
//
//	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
//	}
//
//	fmt.Println(receivedToken)
//
//	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
//	//}
//	//
//	//err = viewServer.Data.DelFromCartUserData(context, receivedToken.Audience, productId)
//	//if err != nil {
//	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Remove From Cart")
//	//}
//
//	err = viewServer.Cash.DelToCartStream(context, receivedToken.Audience, request.GetProductId())
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "Unable To Remove Product From Cart")
//	}
//
//	return &pb.RemoveFromCartProductResponse{Status: true}, nil
//}
