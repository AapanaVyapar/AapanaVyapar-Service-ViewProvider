package main

import (
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("Stating server on port  :  %d", os.Getenv("Port"))

	fmt.Println("Environmental Variables Loaded .. !!")

	serverAddress := "0.0.0.0:4356"
	log.Printf("dialing to server  : %s", serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server ", err)
	}

	client := pb.NewViewProviderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	token, err := helpers.GenerateAuthToken("f38d6a51-b961-474b-9be1-6de62ab5c57e", "Shitij", "319dc46b-e193-4212-9fb7-0b05fcf5d65c", true, []int{helpers.External})
	if err != nil {
		panic(err)
	}

	//resp, err := client.InitUser(context.TODO(), &pb.InitUserRequest{
	//	Token:  token,
	//	ApiKey: os.Getenv("API_KEY_FOR_WEB"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.GetStatus())

	addCartResp, err := client.AddToCartProduct(ctx, &pb.AddToCartProductRequest{
		Token:     token,
		ApiKey:    os.Getenv("API_KEY_FOR_WEB"),
		ProductId: "6097b2e1a9eae2a42076834f",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(addCartResp.GetStatus())

	getOrderResp, err := client.GetOrders(ctx, &pb.GetOrdersRequest{
		Token:  token,
		ApiKey: os.Getenv("API_KEY_FOR_WEB"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(getOrderResp.Recv())

	stream, err := client.GetTrendingShops(ctx, &pb.GetTrendingShopsRequest{
		ApiKey: os.Getenv("API_KEY_FOR_WEB"),
		Token:  token,
		Location: &pb.Location{
			Latitude:  "21.243784586218112",
			Longitude: "75.29856196795996",
		},
		DistanceInMeter: "20",
	})
	if err != nil {
		panic(err)
	}

	var shopIds []string

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		shopIds = append(shopIds, res.Shops.ShopId)
		fmt.Println(res.String())
		fmt.Println(res.Shops.Category)

	}
	//
	//streamProd, err := client.GetTrendingProductsByShop(ctx, &pb.GetTrendingProductsByShopRequest{
	//	ApiKey: os.Getenv("API_KEY_FOR_WEB"),
	//	Token:  token,
	//	ShopId: shopIds,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//var productIds []string
	//
	//for {
	//	res, err := streamProd.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		panic(err)
	//	}
	//	productIds = append(productIds, res.CategoryData.ProductId)
	//	fmt.Println(res.String())
	//
	//}

	//status, err := client.AddToLikeProduct(ctx, &pb.AddToLikeProductRequest{
	//	Token:     token,
	//	ApiKey:    os.Getenv("API_KEY_FOR_WEB"),
	//	ProductId: productIds[0],
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(status)

}
