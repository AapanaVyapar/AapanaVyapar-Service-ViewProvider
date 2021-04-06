package services

import (
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (viewServer *ViewProviderService) LoadBasicCategoriesInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllBasicCategories(ctx, func(data structs.BasicCategoriesData) error {
		err := viewServer.Cash.Cash.HSet(ctx, "categories", data.Category, data.Marshal()).Err()
		if err != nil {
			return status.Errorf(codes.Internal, "unable to add data to hash of Cash  : %w", err)
		}
		return nil

	})
	if err != nil {
		return err

	}
	return nil

}
