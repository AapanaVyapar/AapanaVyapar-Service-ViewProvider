package main

import (
	data_base2 "aapanavyapar-service-viewprovider/data-base/data-services"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func main() {
	database := data_base2.NewDataBase()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := database.UserData.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}

	id := res.InsertedID
	fmt.Println(id)
}
