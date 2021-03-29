package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	redisDb "aapanavyapar-service-viewprovider/configurations/redis"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type DataBase struct {
	mutex          sync.RWMutex
	UserData       *mongo.Collection
	OrderData      *mongo.Collection
	ShopData       *mongo.Collection
	ProductData    *mongo.Collection
	AnalyticalData *mongo.Collection
	Cash           *redis.Client
}

func NewDataBase() *DataBase {

	rdb := redisDb.InitRedis()

	database := mongodb.InitMongo()

	return &DataBase{
		UserData:       mongodb.OpenUserDataCollection(database),
		OrderData:      mongodb.OpenOrderDataCollection(database),
		ShopData:       mongodb.OpenShopDataCollection(database),
		ProductData:    mongodb.OpenProductDataCollection(database),
		AnalyticalData: mongodb.OpenAnalyticalDataCollection(database),
		Cash:           rdb,
	}
}
