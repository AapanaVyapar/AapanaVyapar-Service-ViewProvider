package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func InitMongo() *mongo.Database {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	credential := options.Credential{
		Username: os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	database := client.Database("db_aapanavypar")
	return database

}

func OpenUserDataCollection(database *mongo.Database) *mongo.Collection {
	userData := database.Collection("userData")
	return userData
}

func OpenOrderDataCollection(database *mongo.Database) *mongo.Collection {
	userData := database.Collection("orderData")
	return userData
}

func OpenShopDataCollection(database *mongo.Database) *mongo.Collection {
	userData := database.Collection("shopData")
	return userData
}

func OpenProductDataCollection(database *mongo.Database) *mongo.Collection {
	userData := database.Collection("productData")
	return userData
}

func OpenAnalyticalDataCollection(database *mongo.Database) *mongo.Collection {
	userData := database.Collection("analyticalData")
	return userData
}
