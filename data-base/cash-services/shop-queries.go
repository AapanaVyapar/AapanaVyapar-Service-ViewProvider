package data_base

import (
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func (dataBase *CashDataBase) CreateShopDocument(shopId, shopName, primaryImage string, categoryOfShop []pb.Category, ratingOfShop float32, totalRatingOfShop int, shopkeeper string, latitude, longitude float64) redisearch.Document {

	doc := redisearch.NewDocument("shop:"+strings.ReplaceAll(shopId, SHOP_ID_CHAR_TO_REPLACE, SHOP_ID_CHAR_TO_REPLACE_WITH), 1.0)
	doc.Set("shopName", shopName).
		Set("primaryImage", primaryImage).
		Set("categoryOfShop", categoryOfShop).
		Set("ratingOfShop", ratingOfShop).
		Set("totalRatingOfShop", totalRatingOfShop).
		Set("shopkeeper", shopkeeper).
		Set("location", fmt.Sprintf("%v,%v", longitude, latitude))

	return doc
}

func (dataBase *CashDataBase) GetShopById(shopId string) (*redisearch.Document, error) {
	data, err := dataBase.ShopClient.Get("shop:" + strings.ReplaceAll(shopId, SHOP_ID_CHAR_TO_REPLACE, SHOP_ID_CHAR_TO_REPLACE_WITH))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dataBase *CashDataBase) GetShopByName(shopName string, latitude, longitude, distanceInMeter float64, send func(document redisearch.Document) error) error {
	words := strings.Fields(shopName)

	if len(words) == 0 {
		return fmt.Errorf("empty")
	}

	searchString := ""
	for _, i := range words {
		searchString += "%" + i + "%"
	}

	//queryString := "@shopName:(" + searchString + ") GEOFILTER location " + strconv.FormatFloat(latitude, 'f', -1, 64) + " " + strconv.FormatFloat(longitude, 'f', -1, 64) + " " + strconv.FormatFloat(distanceInMeter, 'f', -1, 64) + " " + "m"
	//fmt.Println(queryString)

	//docs, total, err := dataBase.ShopClient.Search(redisearch.NewQuery(queryString))

	docs, total, err := dataBase.ShopClient.Search(
		redisearch.NewQuery("@shopName:(" + searchString + ")"),
	)
	if err != nil {
		return err
	}
	fmt.Println(total)
	for _, doc := range docs {
		err = send(doc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dataBase *CashDataBase) GetShopByLocation(latitude, longitude, distanceInMeter float64, limit int) ([]redisearch.Document, error) {

	query := redisearch.Query{
		Raw: "*",
		Filters: []redisearch.Filter{
			{
				Field: "location",
				Options: redisearch.GeoFilterOptions{
					Lon:    longitude,
					Lat:    latitude,
					Radius: distanceInMeter,
					Unit:   redisearch.METERS,
				},
			},
		},
		Paging: redisearch.Paging{Offset: redisearch.DefaultOffset, Num: redisearch.DefaultNum},
		SortBy: &redisearch.SortingKey{
			Field:     "ratingOfShop",
			Ascending: false,
		},
	}

	fmt.Println(query)

	docs, total, err := dataBase.ShopClient.Search(query.Limit(0, limit))

	if err != nil {
		return nil, err
	}

	fmt.Println(total)

	return docs, nil
}

func (dataBase *CashDataBase) GetShopByCategory(categoryOfShop []pb.Category, latitude, longitude, distanceInMeter float64, limit int) ([]redisearch.Document, error) {

	searchString := ""

	switch len(categoryOfShop) {
	case 0:
		return nil, fmt.Errorf("empty")
	case 1:
		searchString = categoryOfShop[0].String()
		break
	default:
		i := 0
		for ; i < len(categoryOfShop)-1; i++ {
			searchString += categoryOfShop[i].String() + "|"
		}
		searchString += categoryOfShop[i].String()
	}

	fmt.Println(searchString)

	query := redisearch.Query{
		Raw: "@categoryOfShop:(" + searchString + ")",
		Filters: []redisearch.Filter{
			{
				Field: "location",
				Options: redisearch.GeoFilterOptions{
					Lon:    longitude,
					Lat:    latitude,
					Radius: distanceInMeter,
					Unit:   redisearch.METERS,
				},
			},
		},
		Paging: redisearch.Paging{Offset: redisearch.DefaultOffset, Num: redisearch.DefaultNum},
		SortBy: &redisearch.SortingKey{
			Field:     "ratingOfShop",
			Ascending: false,
		},
	}

	docs, total, err := dataBase.ShopClient.Search(
		query.Limit(0, limit),
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(total)

	return docs, nil
}

func (dataBase *CashDataBase) DelShop(shopId string) error {
	err := dataBase.ShopClient.DeleteDocument("shop:" + strings.ReplaceAll(shopId, SHOP_ID_CHAR_TO_REPLACE, SHOP_ID_CHAR_TO_REPLACE_WITH))
	if err != nil {
		return status.Errorf(codes.Unknown, "Unable To Delete Data From Cash", err)
	}
	return nil

}

func (dataBase *CashDataBase) GetTotalRating(ctx context.Context, shopId string) (int, error) {

	totalRatingString, err := dataBase.Cash.HGet(ctx, "shop:"+shopId, "totalRatingOfShop").Result()
	switch {
	case err == redis.Nil:
		return 0, status.Errorf(codes.NotFound, "Value Not Exist %v", err)
	case err != nil:
		return 0, status.Errorf(codes.Internal, "Unable To Fetch Value %v", err)
	}

	totalRating, err := strconv.Atoi(totalRatingString)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "Unable To Parse Data")
	}

	return totalRating, nil

}

func (dataBase *CashDataBase) SetTotalRating(ctx context.Context, shopId string, rating int) error {

	err := dataBase.Cash.HSet(ctx, "shop:"+shopId, "totalRatingOfShop", rating).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "unable to set rating ", err)
	}
	return nil
}

func (dataBase *CashDataBase) GetRating(ctx context.Context, shopId string) (float32, error) {

	ratingString, err := dataBase.Cash.HGet(ctx, "shop:"+shopId, "ratingOfShop").Result()
	switch {
	case err == redis.Nil:
		return 0, status.Errorf(codes.NotFound, "Value Not Exist %v", err)
	case err != nil:
		return 0, status.Errorf(codes.Internal, "Unable To Fetch Value %v", err)
	}

	rating, err := strconv.ParseFloat(ratingString, 32)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "Unable To Parse Data")
	}

	return float32(rating), nil

}

func (dataBase *CashDataBase) SetRating(ctx context.Context, shopId string, rating float32) error {

	err := dataBase.Cash.HSet(ctx, "shop:"+shopId, "ratingOfShop", rating).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "unable to set rating ", err)
	}
	return nil
}
