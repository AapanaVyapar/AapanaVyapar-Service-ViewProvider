package data_base

import (
	redisDb "aapanavyapar-service-viewprovider/configurations/redis"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"sync"
)

type CashDataBase struct {
	mutex         sync.RWMutex
	Cash          *redis.Client
	ShopClient    *redisearch.Client
	ProductClient *redisearch.Client
}

func NewDataBase() *CashDataBase {

	return &CashDataBase{
		Cash:          redisDb.InitRedis(),
		ShopClient:    redisearch.NewClient("localhost:6379", "shop"),    //shop is index
		ProductClient: redisearch.NewClient("localhost:6379", "product"), //product is index
	}
}
