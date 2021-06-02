package services

import (
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"github.com/RediSearch/redisearch-go/redisearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (viewServer *ViewProviderService) CreateShopSchemaInCash() error {

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("shopName")).
		AddField(redisearch.NewTextFieldOptions("primaryImage", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
		AddField(redisearch.NewTextFieldOptions("categoryOfShop", redisearch.TextFieldOptions{Sortable: true, NoStem: true})).
		AddField(redisearch.NewNumericFieldOptions("ratingOfShop", redisearch.NumericFieldOptions{Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("shopkeeper", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
		AddField(redisearch.NewGeoField("location"))

	// Drop an existing index. If the index does not exist an error is returned
	_ = viewServer.Cash.ShopClient.Drop()

	// Create the index with the given schema
	if err := viewServer.Cash.ShopClient.CreateIndex(sc); err != nil {
		panic(err)
	}

	return nil
}

func (viewServer *ViewProviderService) CreateProductSchemaInCash() error {

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("shopId", redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("primaryImage", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
		AddField(redisearch.NewTextField("productName")).
		AddField(redisearch.NewTextFieldOptions("categoryOfProduct", redisearch.TextFieldOptions{Sortable: true, NoStem: true})).
		AddField(redisearch.NewNumericFieldOptions("likesOfProduct", redisearch.NumericFieldOptions{Sortable: true}))

	// Drop an existing index. If the index does not exist an error is returned
	_ = viewServer.Cash.ProductClient.Drop()

	// Create the index with the given schema
	if err := viewServer.Cash.ProductClient.CreateIndex(sc); err != nil {
		panic(err)
	}

	return nil
}

func (viewServer *ViewProviderService) LoadShopsInCash(ctx context.Context) error {

	err := viewServer.Data.GetAllShopsFromShopData(ctx, func(data structs.ShopData) error {

		latitude, err := strconv.ParseFloat(data.Location.Latitude, 64)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(data.Location.Longitude, 64)
		if err != nil {
			return err
		}

		var rating float32
		totalRatings := 0
		if data.Ratings != nil && len(*data.Ratings) > 0 {
			var sum int
			for _, rat := range *data.Ratings {
				sum += int(rat.Rating)
			}
			totalRatings = len(*data.Ratings)
			rating = float32(sum / totalRatings)
		}

		doc := viewServer.Cash.CreateShopDocument(data.ShopId, data.ShopName, data.PrimaryImage, data.Category, rating, totalRatings, data.ShopKeeperName, latitude, longitude)
		if err := viewServer.Cash.ShopClient.Index([]redisearch.Document{doc}...); err != nil {
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

		doc := viewServer.Cash.CreateProductDocument(data.ProductId.Hex(), data.ShopId, data.Title, data.Images[0], data.Category, data.Likes)
		if err := viewServer.Cash.ShopClient.Index([]redisearch.Document{doc}...); err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err

	}
	return nil

}

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

//
//func (viewServer *ViewProviderService) CreateShopSchemaInCash(ctx context.Context) error {
//	// Create a client. By default a client is schemaless
//	// unless a schema is provided when creating the index
//	c := redisearch.NewClient("localhost:6379", "shop")
//
//	// Create a schema
//	sc := redisearch.NewSchema(redisearch.DefaultOptions).
//		AddField(redisearch.NewTextFieldOptions("shopId", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
//		AddField(redisearch.NewTextField("shopName")).
//		AddField(redisearch.NewTextFieldOptions("primaryImage", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
//		AddField(redisearch.NewTextFieldOptions("categoryOfShop", redisearch.TextFieldOptions{Sortable: true, NoStem: true})).
//		AddField(redisearch.NewNumericFieldOptions("ratingOfShop", redisearch.NumericFieldOptions{Sortable: true})).
//		AddField(redisearch.NewTextFieldOptions("shopkeeper", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
//		AddField(redisearch.NewGeoField("location"))
//
//	// Drop an existing index. If the index does not exist an error is returned
//	_ = c.Drop()
//
//	// Create the index with the given schema
//	if err := c.CreateIndex(sc); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
//
//func (viewServer *ViewProviderService) CreateProductSchemaInCash(ctx context.Context) error {
//	// Create a client. By default a client is schemaless
//	// unless a schema is provided when creating the index
//	c := redisearch.NewClient("localhost:6379", "product")
//
//	// Create a schema
//	sc := redisearch.NewSchema(redisearch.DefaultOptions).
//		AddField(redisearch.NewTextFieldOptions("productId", redisearch.TextFieldOptions{NoStem: true, NoIndex: true})).
//		AddField(redisearch.NewTextField("productName")).
//		AddField(redisearch.NewTextFieldOptions("categoryOfProduct", redisearch.TextFieldOptions{Sortable: true, NoStem: true})).
//		AddField(redisearch.NewNumericFieldOptions("likesOfProduct", redisearch.NumericFieldOptions{Sortable: true}))
//
//	// Drop an existing index. If the index does not exist an error is returned
//	_ = c.Drop()
//
//	// Create the index with the given schema
//	if err := c.CreateIndex(sc); err != nil {
//		panic(err)
//	}
//
//	return nil
//}
//
//func (viewServer *ViewProviderService) LoadBasicCategoriesInCash(ctx context.Context) error {
//
//	err := viewServer.Data.GetAllBasicCategories(ctx, func(data structs.BasicCategoriesData) error {
//		err := viewServer.Cash.Cash.HSet(ctx, "categories", data.Category, data.Marshal()).Err()
//		if err != nil {
//			return status.Errorf(codes.Internal, "unable to add data to hash of Cash  : %w", err)
//		}
//		return nil
//
//	})
//	if err != nil {
//		return err
//
//	}
//	return nil
//
//}
//
//func (viewServer *ViewProviderService) LoadShopsInCash(ctx context.Context) error {
//
//	err := viewServer.Data.GetAllShopsFromShopData(ctx, func(data structs.ShopData) error {
//
//		err := viewServer.Cash.AddShopDataToCash(ctx, data.ShopId, data.Marshal())
//		if err != nil {
//			return err
//		}
//
//		return nil
//
//	})
//	if err != nil {
//		return err
//
//	}
//	return nil
//
//}
//
//func (viewServer *ViewProviderService) LoadProductsInCash(ctx context.Context) error {
//
//	err := viewServer.Data.GetAllProductsFromProductData(ctx, func(data structs.ProductData) error {
//
//		err := viewServer.Cash.AddProductDataToCash(ctx, data.ProductId.Hex(), data.Marshal())
//		if err != nil {
//			return err
//		}
//
//		err = viewServer.Cash.AddShopProductMapDataToCash(ctx, data.ShopId, data.ProductId.Hex())
//		if err != nil {
//			return err
//		}
//
//		return nil
//
//	})
//	if err != nil {
//		return err
//
//	}
//	return nil
//
//}
