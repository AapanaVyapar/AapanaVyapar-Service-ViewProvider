package redis

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"os"
	"strconv"
)

func InitRedisShop() *redisearch.Client {
	dbName, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	cert, err := tls.LoadX509KeyPair("./redis-sharding/redis-tls-container/certs/client.crt", "./redis-sharding/redis-tls-container/certs/client.key")
	if err != nil {
		panic(err)
	}

	caCert, err := ioutil.ReadFile("./redis-sharding/redis-tls-container/certs/ca.crt")
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_ADDRESS"),
				redis.DialPassword(os.Getenv("REDIS_PASSWORD")),
				redis.DialDatabase(dbName),
				redis.DialTLSConfig(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					RootCAs:            caCertPool,
					InsecureSkipVerify: true,
				}),
				redis.DialUseTLS(true),
				redis.DialTLSSkipVerify(true),
			)
		},
	}

	client := redisearch.NewClientFromPool(pool, "shop") //shop is index

	return client

}

func InitRedisProduct() *redisearch.Client {
	dbName, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	cert, err := tls.LoadX509KeyPair("./redis-sharding/redis-tls-container/certs/client.crt", "./redis-sharding/redis-tls-container/certs/client.key")
	if err != nil {
		panic(err)
	}

	caCert, err := ioutil.ReadFile("./redis-sharding/redis-tls-container/certs/ca.crt")
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_ADDRESS"),
				redis.DialPassword(os.Getenv("REDIS_PASSWORD")),
				redis.DialDatabase(dbName),
				redis.DialTLSConfig(&tls.Config{
					Certificates:       []tls.Certificate{cert},
					RootCAs:            caCertPool,
					InsecureSkipVerify: true,
				}),
				redis.DialUseTLS(true),
				redis.DialTLSSkipVerify(true),
			)
		},
	}

	client := redisearch.NewClientFromPool(pool, "product") //product is index

	return client

}
