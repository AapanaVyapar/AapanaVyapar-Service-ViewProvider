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
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"regexp"
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

	//doc := view.Cash.CreateShopDocument("f38d6a51-b961-474b-9be1-6de62ab5c57e", "Milap Store", "https://image.com", []pb.Category{pb.Category_MENS_CLOTHING}, 3.3, "ABC Person", 21.246435522726177, 75.29615236552934)
	//doc1 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-9be1-jn362ab5c57e", "Rajkumar Coldrinks", "https://image.com", []pb.Category{pb.Category_FOOD}, 4.2, "ABC Person", 21.246703671726266, 75.29363042248966)
	//doc2 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-d4g2-jn362ab5c57e", "Laxmi Offset", "https://image.com", []pb.Category{pb.Category_FOOD}, 5.0, "ABC Person", 21.246863857271922, 75.29460006862273)
	//doc3 := view.Cash.CreateShopDocument("f38d6a51-b961-474b-dv1e-jn362ab5c57e", "raza cutlery stores", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.24691956642532, 75.2936569970796)
	//doc4 := view.Cash.CreateShopDocument("f38d6a51-b932-474b-dv1e-jn362ab5c57e", "joshi electrical works", "https://image.com", []pb.Category{pb.Category_ELECTRIC}, 3.0, "ABC Person", 21.24663669715862, 75.29432756403227)
	//doc5 := view.Cash.CreateShopDocument("f38d6a51-4432-474b-dv1e-jn362ab5c57e", "chintamani kirana stores", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.24645703702037, 75.29347243866049)
	//doc6 := view.Cash.CreateShopDocument("f38d6a32-4432-474b-dv1e-jn362ab5c57e", "Chandu Dairy", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.246983732987587, 75.29470401626448)
	//doc7 := view.Cash.CreateShopDocument("f68d6a32-4432-474b-dv1e-jn362ab5c57e", "Vegetable Market", "https://image.com", []pb.Category{pb.Category_FOOD}, 3.0, "ABC Person", 21.245561629488762, 75.29823368278593)
	//
	//// Index the document. The API accepts multiple documents at a time
	//if err := view.Cash.ShopClient.Index([]redisearch.Document{doc, doc1, doc2, doc3, doc4, doc5, doc6, doc7}...); err != nil {
	//	log.Fatal(err)
	//}

	//docProd := view.Cash.CreateProductDocument("1", "f38d6a51-b961-474b-9be1-6de62ab5c57e", "T-Shirt", "https://image.com", []pb.Category{pb.Category_MENS_CLOTHING}, 20)
	//doc1Prod := view.Cash.CreateProductDocument("2", "f38d6a51-b961-474b-9be1-jn362ab5c57e", "Samosa", "https://image.com", []pb.Category{pb.Category_FOOD}, 40)
	//doc2Prod := view.Cash.CreateProductDocument("3", "f38d6a51-b961-474b-d4g2-jn362ab5c57e", "Jalabi", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	//doc3Prod := view.Cash.CreateProductDocument("4", "f38d6a51-b961-474b-dv1e-jn362ab5c57e", "Kachori", "https://image.com", []pb.Category{pb.Category_FOOD}, 50)
	//doc4Prod := view.Cash.CreateProductDocument("5", "f38d6a51-b932-474b-dv1e-jn362ab5c57e", "Motor", "https://image.com", []pb.Category{pb.Category_ELECTRIC}, 60)
	//doc5Prod := view.Cash.CreateProductDocument("6", "f38d6a51-4432-474b-dv1e-jn362ab5c57e", "broom", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	//doc6Prod := view.Cash.CreateProductDocument("7", "f38d6a32-4432-474b-dv1e-jn362ab5c57e", "milk", "https://image.com", []pb.Category{pb.Category_FOOD}, 30)
	//doc7Prod := view.Cash.CreateProductDocument("8", "f68d6a32-4432-474b-dv1e-jn362ab5c57e", "vegetable", "https://image.com", []pb.Category{pb.Category_FOOD}, 60)
	//
	//// Index the document. The API accepts multiple documents at a time
	//if err := view.Cash.ProductClient.Index([]redisearch.Document{docProd, doc1Prod, doc2Prod, doc3Prod, doc4Prod, doc5Prod, doc6Prod, doc7Prod}...); err != nil {
	//	log.Fatal(err)
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	err = view.LoadBasicCategoriesInCash(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded Basic Category")

	err = view.LoadShopsInCash(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded Shop")

	err = view.LoadProductsInCash(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded Products")

	return &view
}

func (viewServer *ViewProviderService) RateShop(ctx context.Context, request *pb.RateShopRequest) (*pb.RateShopResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println("\n\n\nShopID : " + request.ShopId)

	shopIdForCache := strings.ReplaceAll(request.GetShopId(), redisDataBase.SHOP_ID_CHAR_TO_REPLACE, redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH)

	fmt.Println("\n\n\nShopID Cache : " + shopIdForCache)

	totalRating, err := viewServer.Cash.GetTotalRating(ctx, shopIdForCache)
	if err != nil {
		fmt.Println("Total Rating")
		return nil, err
	}

	rating, err := viewServer.Cash.GetRating(ctx, shopIdForCache)
	if err != nil {
		fmt.Println("Get Rating")
		return nil, err
	}

	err = viewServer.Data.AddRatingInShopData(ctx, request.GetShopId(), structs.Rating{
		UserId:    receivedToken.Audience,
		UserName:  receivedToken.Subject,
		Comment:   request.GetComment(),
		Rating:    request.GetRatings(),
		Timestamp: time.Now().UTC(),
	})
	if err != nil {
		fmt.Println("Add Rating")
		return nil, err
	}

	i := float32(request.GetRatings().Enum().Number())
	newRating := rating + ((i - rating) / float32(totalRating+1))

	err = viewServer.Cash.SetTotalRating(ctx, shopIdForCache, totalRating+1)
	if err != nil {
		fmt.Println("Total Rating Set")
		return nil, err
	}

	err = viewServer.Cash.SetRating(ctx, shopIdForCache, newRating)
	if err != nil {
		fmt.Println("Set Rating")
		return nil, err
	}

	return &pb.RateShopResponse{Status: true}, err
}

func (viewServer *ViewProviderService) InitUser(ctx context.Context, request *pb.InitUserRequest) (*pb.InitUserResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	_, err = viewServer.Data.CreateUser(ctx, receivedToken.Audience, receivedToken.Subject)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Unable To Init User")
	}

	return &pb.InitUserResponse{Status: true}, nil
}

func (viewServer *ViewProviderService) GetProfile(ctx context.Context, request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	user, err := viewServer.Data.GetUserData(ctx, receivedToken.Audience)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Data Not Found")
	}

	var fav []string
	if user.Favorites != nil && len(user.Favorites.Products) > 0 {
		for _, product := range user.Favorites.Products {
			fav = append(fav, product.Hex())
		}
	}

	var cart []string
	if user.Cart != nil && len(user.Cart.Products) > 0 {
		for _, cartItem := range user.Cart.Products {
			cart = append(cart, cartItem.Hex())
		}
	}

	if user.Address == nil {
		return &pb.GetProfileResponse{
			UserName:  user.UserName,
			Address:   nil,
			Favourite: fav,
			Cart:      cart,
		}, nil
	}
	return &pb.GetProfileResponse{
		UserName: user.UserName,
		Address: &pb.Address{
			FullName:      user.Address.FullName,
			HouseDetails:  user.Address.HouseDetails,
			StreetDetails: user.Address.StreetDetails,
			LandMark:      user.Address.LandMark,
			PinCode:       user.Address.PinCode,
			City:          user.Address.City,
			State:         user.Address.State,
			Country:       user.Address.Country,
			PhoneNo:       user.Address.PhoneNo,
		},
		Favourite: fav,
		Cart:      cart,
	}, nil
}

func (viewServer *ViewProviderService) UpdateAddress(ctx context.Context, request *pb.UpdateAddressRequest) (*pb.UpdateAddressResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	err = viewServer.Data.SetAddressInUserData(ctx, receivedToken.Audience, receivedToken.Subject, structs.Address{
		FullName:      request.GetAddress().FullName,
		HouseDetails:  request.GetAddress().HouseDetails,
		StreetDetails: request.GetAddress().StreetDetails,
		LandMark:      request.GetAddress().LandMark,
		PinCode:       request.GetAddress().PinCode,
		City:          request.GetAddress().City,
		State:         request.GetAddress().State,
		Country:       request.GetAddress().Country,
		PhoneNo:       request.GetAddress().PhoneNo,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Unable To Update Address")
	}

	return &pb.UpdateAddressResponse{Status: true}, nil
}

func (viewServer *ViewProviderService) GetCart(request *pb.GetCartRequest, stream pb.ViewProviderService_GetCartServer) error {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	cart, err := viewServer.Data.GetCartUserData(stream.Context(), receivedToken.Audience)
	if err != nil {
		return status.Errorf(codes.NotFound, "Unable To Get Cart")
	}

	if cart != nil {
		for _, product := range cart.Products {
			data, err := viewServer.Cash.GetProductById(product.Hex())
			if err != nil {
				return status.Errorf(codes.Internal, "Unable To Get Product Info")
			}

			str := data.Properties["categoryOfProduct"].(string)[1:]
			str = str[:len(str)-1]
			catData := strings.Split(str, ",")
			var category []pb.Category
			for _, cat := range catData {
				category = append(category, pb.Category(pb.Category_value[cat]))
			}

			likes, err := strconv.ParseUint(data.Properties["likesOfProduct"].(string), 10, 64)
			if err != nil {
				return status.Errorf(codes.Internal, "Unable To Parse Data")
			}

			err = stream.Send(&pb.GetCartResponse{Products: &pb.ProductsOfShopsNearBy{
				ProductId:    product.Hex(),
				ShopId:       strings.ReplaceAll(data.Properties["shopId"].(string), redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH, redisDataBase.SHOP_ID_CHAR_TO_REPLACE),
				ProductName:  data.Properties["productName"].(string),
				PrimaryImage: data.Properties["primaryImage"].(string),
				Category:     category,
				Likes:        likes,
			}})
			if err != nil {
				return status.Errorf(codes.Unknown, "Stream Error")
			}

		}
	}

	return nil
}

func (viewServer *ViewProviderService) GetOrders(request *pb.GetOrdersRequest, stream pb.ViewProviderService_GetOrdersServer) error {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	err = viewServer.Data.GetMultipleOrdersInfoByUserIdFromOrderData(stream.Context(), receivedToken.Audience, func(data structs.OrderData) error {

		err = stream.Send(&pb.GetOrdersResponse{
			OrderId:           data.OrderId.Hex(),
			Status:            data.Status,
			ProductId:         data.ProductId.Hex(),
			DeliveryTimeStamp: data.DeliveryTimeStamp.String(),
			OrderTimeStamp:    data.OrderTimeStamp.String(),
			Price:             data.Price,
			Quantity:          data.Quantity,
			ProductName:       data.ProductName,
			ProductImage:      data.ProductImage,
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "Error While Sending Data")
		}
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "Stream Error")
	}

	return nil
}

func (viewServer *ViewProviderService) GetProductsBySearch(request *pb.GetProductsBySearchRequest, stream pb.ViewProviderService_GetProductsBySearchServer) error {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	var processedShopIds []string
	count := 0
	for _, id := range request.GetShopIds() {
		if count > 100 {
			break
		}
		count += 1
		_, err = uuid.Parse(id)
		if err != nil {
			continue
		}
		processedShopIds = append(processedShopIds, id)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return status.Errorf(codes.Internal, "Problem Occurred")
	}
	processedSearch := reg.ReplaceAllString(request.GetSearch(), "")

	err = viewServer.Cash.GetProductByName(processedShopIds, processedSearch, func(document redisearch.Document) error {

		str := document.Properties["categoryOfProduct"].(string)[1:]
		str = str[:len(str)-1]
		data := strings.Split(str, ",")
		var category []pb.Category
		for _, cat := range data {
			category = append(category, pb.Category(pb.Category_value[cat]))
		}

		likes, err := strconv.ParseUint(document.Properties["likesOfProduct"].(string), 10, 64)
		if err != nil {
			return status.Errorf(codes.Internal, "Unable To Parse Data")
		}

		err = stream.Send(&pb.GetProductsBySearchResponse{Products: &pb.ProductsOfShopsNearBy{
			ProductId:    document.Id[8:],
			ShopId:       strings.ReplaceAll(document.Properties["shopId"].(string), redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH, redisDataBase.SHOP_ID_CHAR_TO_REPLACE),
			ProductName:  document.Properties["productName"].(string),
			PrimaryImage: document.Properties["primaryImage"].(string),
			Category:     category,
			Likes:        likes,
		}})
		return err

	})
	if err != nil {
		return status.Errorf(codes.NotFound, "Unable To Get Data")
	}

	return nil

}

func (viewServer *ViewProviderService) GetShopsBySearch(request *pb.GetShopsBySearchRequest, stream pb.ViewProviderService_GetShopsBySearchServer) error {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(stream.Context(), request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return status.Errorf(codes.Internal, "Problem Occurred")
	}
	processedSearch := reg.ReplaceAllString(request.GetSearch(), "")

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
		return status.Errorf(codes.InvalidArgument, "Invalid Distance")
	}

	err = viewServer.Cash.GetShopByName(processedSearch, latitude, longitude, meter, func(document redisearch.Document) error {

		location := strings.Split(document.Properties["location"].(string), ",")

		str := document.Properties["categoryOfShop"].(string)[1:]
		str = str[:len(str)-1]
		data := strings.Split(str, ",")
		var category []pb.Category
		for _, cat := range data {
			category = append(category, pb.Category(pb.Category_value[cat]))
		}

		rating, err := strconv.ParseFloat(document.Properties["ratingOfShop"].(string), 32)
		if err != nil {
			return status.Errorf(codes.Internal, "Unable To Parse Data")
		}

		err = stream.Send(&pb.GetShopsBySearchResponse{Shops: &pb.ShopsNearBy{
			ShopId:       strings.ReplaceAll(document.Id[5:], redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH, redisDataBase.SHOP_ID_CHAR_TO_REPLACE),
			ShopName:     document.Properties["shopName"].(string),
			PrimaryImage: document.Properties["primaryImage"].(string),
			Category:     category,
			Rating:       float32(rating),
			Shopkeeper:   document.Properties["shopkeeper"].(string),
			Location: &pb.Location{
				Latitude:  location[1],
				Longitude: location[0],
			},
		}})
		return err

	})
	if err != nil {
		return status.Errorf(codes.Internal, "Error While Searching")
	}

	return nil
}

func (viewServer *ViewProviderService) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	productId, err := primitive.ObjectIDFromHex(request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
	}

	data, err := viewServer.Data.GetProductFromProductData(ctx, productId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable To Get Product")
	}

	return &pb.GetProductResponse{
		ProductId:          data.ProductId.Hex(),
		ShopId:             data.ShopId,
		ShopName:           data.ShopName,
		ProductName:        data.Title,
		ProductDescription: data.Description,
		ShippingInfo:       data.ShippingInfo,
		Stock:              data.Stock,
		Likes:              data.Likes,
		Price:              data.Price,
		Offer:              data.Offer,
		Images:             data.Images,
		Category:           data.Category,
		Timestamp:          data.Timestamp.String(),
	}, nil
}

func (viewServer *ViewProviderService) GetShop(ctx context.Context, request *pb.GetShopRequest) (*pb.GetShopResponse, error) {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	_, err := helpers.ValidateToken(ctx, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	data, err := viewServer.Data.GetShopFromShopData(ctx, request.GetShopId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable To Get Shop")
	}

	totalRating, err := viewServer.Cash.GetRating(ctx, strings.ReplaceAll(request.GetShopId(), redisDataBase.SHOP_ID_CHAR_TO_REPLACE, redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH))
	if err != nil {
		return nil, err
	}

	var rating []*pb.RatingOfShop
	if data.Ratings != nil && len(*data.Ratings) > 0 {
		for _, rat := range *data.Ratings {
			ra := pb.RatingOfShop{
				UserName:  rat.UserName,
				Comment:   rat.Comment,
				Rating:    rat.Rating,
				Timestamp: rat.Timestamp.String(),
			}
			rating = append(rating, &ra)
		}
	}

	return &pb.GetShopResponse{
		ShopId:         data.ShopId,
		ShopName:       data.ShopName,
		ShopKeeperName: data.ShopKeeperName,
		Images:         data.Images,
		PrimaryImage:   data.PrimaryImage,
		Location: &pb.Location{
			Longitude: data.Location.Longitude,
			Latitude:  data.Location.Latitude,
		},
		Category:            data.Category,
		BusinessInformation: data.BusinessInformation,
		OperationalHours: &pb.OperationalHours{
			Sunday:    []string{data.OperationalHours.Sunday[0], data.OperationalHours.Sunday[1]},
			Monday:    []string{data.OperationalHours.Monday[0], data.OperationalHours.Monday[1]},
			Tuesday:   []string{data.OperationalHours.Tuesday[0], data.OperationalHours.Tuesday[1]},
			Wednesday: []string{data.OperationalHours.Wednesday[0], data.OperationalHours.Wednesday[1]},
			Thursday:  []string{data.OperationalHours.Thursday[0], data.OperationalHours.Thursday[1]},
			Friday:    []string{data.OperationalHours.Friday[0], data.OperationalHours.Friday[1]},
			Saturday:  []string{data.OperationalHours.Saturday[0], data.OperationalHours.Saturday[1]},
		},
		Ratings:     rating,
		TotalRating: totalRating,
		Timestamp:   data.Timestamp.String(),
	}, nil
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
		return status.Errorf(codes.InvalidArgument, "Invalid Distance")
	}

	docs, err := viewServer.Cash.GetShopByLocation(latitude, longitude, meter, 100)
	if err != nil {
		return status.Errorf(codes.NotFound, "Unable To Get Data For Shop")
	}

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
			return status.Errorf(codes.Internal, "Unable To Parse Data")
		}

		err = stream.Send(&pb.GetTrendingShopsResponse{
			Shops: &pb.ShopsNearBy{
				ShopId:       strings.ReplaceAll(doc.Id[5:], redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH, redisDataBase.SHOP_ID_CHAR_TO_REPLACE),
				ShopName:     doc.Properties["shopName"].(string),
				PrimaryImage: doc.Properties["primaryImage"].(string),
				Category:     category,
				Rating:       float32(rating),
				Shopkeeper:   doc.Properties["shopkeeper"].(string),
				Location: &pb.Location{
					Latitude:  location[1],
					Longitude: location[0],
				},
			},
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "Stream Error")
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
			return status.Errorf(codes.Internal, "Unable To Parse Data")
		}

		err = stream.Send(&pb.GetTrendingProductsByShopResponse{
			CategoryData: &pb.ProductsOfShopsNearBy{
				ProductId:    doc.Id[8:],
				ShopId:       strings.ReplaceAll(doc.Properties["shopId"].(string), redisDataBase.SHOP_ID_CHAR_TO_REPLACE_WITH, redisDataBase.SHOP_ID_CHAR_TO_REPLACE),
				ProductName:  doc.Properties["productName"].(string),
				PrimaryImage: doc.Properties["primaryImage"].(string),
				Category:     category,
				Likes:        likes,
			},
		},
		)
		if err != nil {
			return status.Errorf(codes.Unknown, "Stream Error")
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

	productData, err := viewServer.Cash.GetProductById(request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
	}

	likes, err := strconv.ParseUint(productData.Properties["likesOfProduct"].(string), 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Add Like To Product")
	}

	err = viewServer.Cash.AddToFavStream(context, receivedToken.Audience, likes, request.GetProductId())
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

	fmt.Println("Remove Like : ", receivedToken)

	productData, err := viewServer.Cash.GetProductById(request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Product Does Not Exist")
	}

	likes, err := strconv.ParseUint(productData.Properties["likesOfProduct"].(string), 10, 64)
	if err != nil {
		return &pb.RemoveFromLikeProductResponse{Status: true}, nil
		//return nil, status.Errorf(codes.Internal, "Unable To Parse Data")
	}

	//if likes == 0 {
	//	return &pb.RemoveFromLikeProductResponse{Status: true}, nil
	//}

	err = viewServer.Cash.DelToFavStream(context, receivedToken.Audience, likes, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Remove Like To Product")
	}

	//likes -= 1

	//err = viewServer.Cash.UpdateLikeOfProduct(context, request.GetProductId(), likes)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "Unable To UnLike")
	//}
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

	_, err = viewServer.Cash.GetProductById(request.GetProductId())
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
