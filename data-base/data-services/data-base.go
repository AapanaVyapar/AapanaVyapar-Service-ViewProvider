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
	Data  *mongo.Client
	Cash  *redis.Client
}

func NewDataBase() *DataBase {

	client := mongodb.InitMongo()

	rdb := redisDb.InitRedis()

	return &DataBase{
		Data: client,
		Cash: rdb,
	}
}
