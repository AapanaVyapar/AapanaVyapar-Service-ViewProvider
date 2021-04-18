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

	stream, err := client.GetTrendingShops(ctx, &pb.GetTrendingShopsRequest{
		ApiKey: os.Getenv("API_KEY_FOR_WEB"),
		Token:  token,
		Location: &pb.Location{
			Latitude:  "21.246534558436874",
			Longitude: "75.29436710722217",
		},
		DistanceInMeter: "10",
	})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(res.String())

	}
}
