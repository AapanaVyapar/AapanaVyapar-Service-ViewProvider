package data_base

import (
	"aapanavyapar-service-viewprovider/pb"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (dataBase *CashDataBase) CreateShopDocument(shopId, shopName, primaryImage string, categoryOfShop []pb.Category, ratingOfShop float32, shopkeeper string, latitude, longitude float64) redisearch.Document {

	doc := redisearch.NewDocument("shop:"+strings.ReplaceAll(shopId, SHOP_ID_CHAR_TO_REPLACE, SHOP_ID_CHAR_TO_REPLACE_WITH), 1.0)
	doc.Set("shopName", shopName).
		Set("primaryImage", primaryImage).
		Set("categoryOfShop", categoryOfShop).
		Set("ratingOfShop", ratingOfShop).
		Set("shopkeeper", shopkeeper).
		Set("location", fmt.Sprintf("%v,%v", latitude, longitude))

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
		redisearch.NewQuery("@shopName:(" + searchString + ")").AddFilter(
			redisearch.Filter{
				Field: "location",
				Options: redisearch.GeoFilterOptions{
					Lon:    latitude,
					Lat:    longitude,
					Radius: distanceInMeter,
					Unit:   redisearch.METERS,
				},
			},
		),
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
					Lon:    latitude,
					Lat:    longitude,
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
					Lon:    latitude,
					Lat:    longitude,
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
