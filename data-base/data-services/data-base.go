package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	redisDb "aapanavyapar-service-viewprovider/configurations/redis"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type DataBase struct {
	mutex sync.RWMutex
	Data  *mongo.Database
	Cash  *redis.Client
}

func NewDataBase() *DataBase {

	database := mongodb.InitMongo()

	rdb := redisDb.InitRedis()

	return &DataBase{
		Data: database,
		Cash: rdb,
	}
}
