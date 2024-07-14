package data_base

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

func (dataBase *CashDataBase) AddToFavStream(ctx context.Context, uId string, likes uint64, productId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_FAV_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"uId", uId, "prodId", productId, "likes", strconv.FormatUint(likes, 10), "operation", "+"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) DelToFavStream(ctx context.Context, uId string, likes uint64, productId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_FAV_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"uId", uId, "prodId", productId, "likes", strconv.FormatUint(likes, 10), "operation", "-"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) AddToCartStream(ctx context.Context, uId, productId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_CART_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"uId", uId, "prodId", productId, "operation", "+"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *CashDataBase) DelToCartStream(ctx context.Context, uId, productId string) error {

	err := dataBase.Cash.XAdd(ctx, &redis.XAddArgs{
		Stream:       os.Getenv("REDIS_STREAM_CART_NAME"),
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       []string{"uId", uId, "prodId", productId, "operation", "-"},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}
