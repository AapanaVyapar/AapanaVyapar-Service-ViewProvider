package services

import (
	data_base "aapanavyapar-service-viewprovider/data-base/data-services"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type ViewProviderService struct {
	data *data_base.DataBase
}

func (viewServer *ViewProviderService) GetTrendingCategories(context context.Context, request *pb.GetTrendingCategoriesRequest) (*pb.GetTrendingCategoriesResponse, error) {

	if !helpers.CheckForAPIKey(request.GetApiKey()) {
		return nil, status.Errorf(codes.Unauthenticated, "No API Key Is Specified")
	}

	receivedToken, err := helpers.ValidateToken(context, request.GetToken(), os.Getenv("AUTH_TOKEN_SECRETE"), helpers.External)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Request With Invalid Token")
	}

	return nil, nil
}

func (viewServer *ViewProviderService) GetTrendingDataByCategories(request *pb.GetTrendingDataByCategoriesRequest, stream pb.ViewProviderService_GetTrendingDataByCategoriesServer) error {

	return nil
}
