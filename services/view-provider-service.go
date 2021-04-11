package services

import (
	redisDataBase "aapanavyapar-service-viewprovider/data-base/cash-services"
	mongoDataBase "aapanavyapar-service-viewprovider/data-base/data-services"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type ViewProviderService struct {
	Data *mongoDataBase.MongoDataBase
	Cash *redisDataBase.CashDataBase
}

func NewViewProviderService(ctx context.Context) *ViewProviderService {
	mongoData := mongoDataBase.NewDataBase()
	redisData := redisDataBase.NewDataBase()

	view := ViewProviderService{
		Data: mongoData,
		Cash: redisData,
	}

	err := view.LoadBasicCategoriesInCash(ctx)
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

func (viewServer *ViewProviderService) GetTrendingCategories(context context.Context, request *pb.GetTrendingCategoriesRequest) (*pb.GetTrendingCategoriesResponse, error) {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	//receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	//if err != nil {
	//	return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	//}
	//
	//fmt.Println(receivedToken)

	keys := viewServer.Cash.Cash.HKeys(context, "categories")

	return &pb.GetTrendingCategoriesResponse{
		CategoryData: keys.Val(),
	}, nil
}

func (viewServer *ViewProviderService) GetTrendingDataByCategories(request *pb.GetTrendingDataByCategoriesRequest, stream pb.ViewProviderService_GetTrendingDataByCategoriesServer) error {

	location := structs.Location{
		Longitude: request.Location.Longitude,
		Latitude:  request.Location.Latitude,
	}

	if err := helpers.Validate(location); err != nil {
		return err
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

	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
	//}
	//
	//err = viewServer.Data.AddToFavoritesUserData(context, receivedToken.Audience, productId)
	//if err != nil {
	//	return nil, status.Errorf(codes.AlreadyExists, "You Already Like To Product Or Unable To Process Request At A Movement")
	//}

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

	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
	//}
	//
	//err = viewServer.Data.DelFromFavoritesUserData(context, receivedToken.Audience, productId)
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Unlike The Product")
	//}

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

	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
	//}
	//
	//err = viewServer.Data.AddToCartUserData(context, receivedToken.Audience, productId)
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Add To Cart")
	//}

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

	//productId, err := primitive.ObjectIDFromHex(request.GetProductId())
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Invalid Product Id")
	//}
	//
	//err = viewServer.Data.DelFromCartUserData(context, receivedToken.Audience, productId)
	//if err != nil {
	//	return nil, status.Errorf(codes.InvalidArgument, "Unable To Process Request Of Remove From Cart")
	//}

	err = viewServer.Cash.DelToCartStream(context, receivedToken.Audience, request.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable To Remove Product From Cart")
	}

	return &pb.RemoveFromCartProductResponse{Status: true}, nil
}
