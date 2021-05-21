package main

import (
	redisDataBase "aapanavyapar-service-viewprovider/data-base/cash-services"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	redisData := redisDataBase.NewDataBase()

	productId := "60a76bb7f5ab012effedcc5e"

	data, err := redisData.GetProductById(productId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	val, err := redisData.Cash.HGetAll(context.Background(), "product:"+productId).Result()
	switch {
	case err == redis.Nil:
		fmt.Println(codes.NotFound, "Value Not Exist %v", err)
	case err != nil:
		fmt.Println(codes.Internal, "Unable To Fetch Value %v", err)
	}
	fmt.Println(val)
	fmt.Println(err)

	//
	//err = UpdateProductInCache(redisData, productId, "My Product", "https://www.gizchina.com/wp-content/uploads/images/2021/01/android-12.jpg", []pb.Category{pb.Category_AGRICULTURAL, pb.Category_MENS_CLOTHING})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//val, err = redisData.Cash.HGetAll(context.Background(), "product:" + productId).Result()
	//switch {
	//case err == redis.Nil:
	//	fmt.Println(codes.NotFound, "Value Not Exist %v", err)
	//case err != nil:
	//	fmt.Println(codes.Internal, "Unable To Fetch Value %v", err)
	//}
	//fmt.Println(val)
	//fmt.Println(err)

}

//doc.Set("shopId", strings.ReplaceAll(shopId, "-", " ")).
//Set("productName", productName).
//Set("primaryImage", primaryImage).
//Set("categoryOfProduct", categoryOfProduct).
//Set("likesOfProduct", likesOfProduct)

func UpdateProductInCache(redisData *redisDataBase.CashDataBase, productId, productName string, primaryImage string, categoryOfProduct []pb.Category) error {

	category := "["
	for ind, cat := range categoryOfProduct {
		if ind > 0 {
			category += ","
		}
		category += cat.String()
	}
	category += "]"

	err := redisData.Cash.HSet(context.Background(), "product:"+productId,
		"productName", productName,
		"primaryImage", primaryImage,
		"categoryOfProduct", category,
	).Err()

	if err != nil {
		return status.Errorf(codes.Internal, "unable to add data to hash of Cash  : %w", err)
	}
	return nil
}
