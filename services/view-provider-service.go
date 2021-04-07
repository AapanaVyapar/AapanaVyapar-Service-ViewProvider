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

	fmt.Println(receivedToken)

	return nil, nil
}

func (viewServer *ViewProviderService) RemoveFromLikeProduct(context context.Context, request *pb.RemoveFromLikeProductRequest) (*pb.RemoveFromLikeProductResponse, error) {
	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	fmt.Println(receivedToken)

	return nil, nil
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

	return nil, nil
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

	return nil, nil
}
