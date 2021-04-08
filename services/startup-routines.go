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

func (viewServer *ViewProviderService) LoadShopsInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllShopsFromShopData(ctx, func(data structs.ShopData) error {

		err := viewServer.Cash.AddShopDataToCash(ctx, data.ShopId.Hex(), data.Marshal())
		if err != nil {
			return err
		}

		array := structs.CashStructureProductArray{
			Products: []string{},
		}

		err = viewServer.Cash.AddShopProductMapDataToCash(ctx, data.ShopId.Hex(), array.Marshal())
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err

	}
	return nil

}

func (viewServer *ViewProviderService) LoadProductsInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllProductsFromProductData(ctx, func(data structs.ProductData) error {

		err := viewServer.Cash.AddProductDataToCash(ctx, data.ProductId.Hex(), data.Marshal())
		if err != nil {
			return err
		}

		val, err := viewServer.Cash.GetShopProductMapDataFromCash(ctx, data.ShopId.Hex())
		if err != nil {
			return err
		}

		array := structs.CashStructureProductArray{}
		structs.UnmarshalCashStructureProductArray([]byte(val), &array)

		array.Products = append(array.Products, data.ProductId.Hex())

		err = viewServer.Cash.AddShopProductMapDataToCash(ctx, data.ShopId.Hex(), array.Marshal())
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err

	}
	return nil

}
