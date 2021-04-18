package data_base

import (
	"aapanavyapar-service-viewprovider/pb"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"strings"
)

func (dataBase *CashDataBase) CreateProductDocument(productId, shopId, productName, primaryImage string, categoryOfProduct []pb.Category, likesOfProduct uint64) redisearch.Document {

	doc := redisearch.NewDocument("product:"+productId, 1.0)
	doc.Set("shopId", strings.ReplaceAll(shopId, "-", " ")).
		Set("productName", productName).
		Set("primaryImage", primaryImage).
		Set("categoryOfProduct", categoryOfProduct).
		Set("likesOfProduct", likesOfProduct)

	return doc
}

func (dataBase *CashDataBase) GetProductById(productId string) (*redisearch.Document, error) {
	data, err := dataBase.ProductClient.Get("product:" + productId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dataBase *CashDataBase) GetProductByName(shopId []string, productName string) ([]redisearch.Document, error) {

	shodIdSearchString, err := prepareShopIds(shopId)
	if err != nil {
		return nil, err
	}
	fmt.Println(shodIdSearchString)

	words := strings.Fields(productName)

	if len(words) == 0 {
		return nil, fmt.Errorf("empty")
	}

	searchString := ""
	for _, i := range words {
		searchString += "%" + i + "%"
	}

	docs, total, err := dataBase.ProductClient.Search(
		redisearch.NewQuery("@shopId:(" + shodIdSearchString + ") @productName:(" + searchString + ")"),
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(total)

	return docs, nil
}

func (dataBase *CashDataBase) GetProductsByShopsSortByRating(shopId []string, limit int) ([]redisearch.Document, error) {

	shodIdSearchString, err := prepareShopIds(shopId)
	if err != nil {
		return nil, err
	}
	fmt.Println(shodIdSearchString)

	query := redisearch.Query{
		Raw:    "@shopId:(" + shodIdSearchString + ")",
		Paging: redisearch.Paging{Offset: redisearch.DefaultOffset, Num: redisearch.DefaultNum},
		SortBy: &redisearch.SortingKey{
			Field:     "likesOfProduct",
			Ascending: false,
		},
	}

	docs, total, err := dataBase.ProductClient.Search(
		query.Limit(0, limit),
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(total)

	return docs, nil
}

func (dataBase *CashDataBase) GetProductByCategory(shopId []string, categoryOfProduct []pb.Category, limit int) ([]redisearch.Document, error) {

	shodIdSearchString, err := prepareShopIds(shopId)
	if err != nil {
		return nil, err
	}

	searchString := ""

	switch len(categoryOfProduct) {
	case 0:
		return nil, fmt.Errorf("empty")
	case 1:
		searchString = categoryOfProduct[0].String()
		break
	default:
		i := 0
		for ; i < len(categoryOfProduct)-1; i++ {
			searchString += categoryOfProduct[i].String() + "|"
		}
		searchString += categoryOfProduct[i].String()
	}

	query := redisearch.Query{
		Raw:    "@shopId:(" + shodIdSearchString + ") @categoryOfProduct:(" + searchString + ")",
		Paging: redisearch.Paging{Offset: redisearch.DefaultOffset, Num: redisearch.DefaultNum},
		SortBy: &redisearch.SortingKey{
			Field:     "likesOfProduct",
			Ascending: false,
		},
	}

	docs, total, err := dataBase.ProductClient.Search(
		query.Limit(0, limit),
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(total)

	return docs, nil
}

func prepareShopIds(shopId []string) (string, error) {

	shodIdSearchString := ""

	switch len(shopId) {
	case 0:
		return "", fmt.Errorf("empty")
	case 1:
		shodIdSearchString = shopId[0]
		break
	default:
		i := 0
		for ; i < len(shopId)-1; i++ {
			shodIdSearchString += shopId[i] + "|"
		}
		shodIdSearchString += shopId[i]
	}

	shodIdSearchString = strings.ReplaceAll(shodIdSearchString, "-", " ")
	return shodIdSearchString, nil
}
