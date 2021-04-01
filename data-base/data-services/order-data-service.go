package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/constants"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (dataBase *DataBase) CreateOrder(context context.Context, userId string, productId primitive.ObjectID, quantity uint32) (primitive.ObjectID, error) {

	if !dataBase.IsExistInUserData(context, "_id", userId) {
		return primitive.ObjectID{}, fmt.Errorf("user does not exist")
	}

	if !dataBase.IsExistProductExist(context, "_id", productId) {
		return primitive.ObjectID{}, fmt.Errorf("product does not exist")
	}

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	var order structs.OrderData

	productData := mongodb.OpenOrderDataCollection(dataBase.Data)
	order.TimeStamp = time.Now().UTC()
	order.UserId = userId
	order.Quantity = quantity
	order.ProductId = productId

	price, offer, err := dataBase.DecreaseStockToMakeOrderFromProductData(context, productId, quantity)

	order.Price = price - ((price / 100) * float64(offer))
	order.Status = constants.PENDING

	id, err := productData.InsertOne(context, order)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	err = dataBase.AddToOrdersUserData(context, userId, id.InsertedID.(primitive.ObjectID))
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id.InsertedID.(primitive.ObjectID), nil
}

func (dataBase *DataBase) UpdateOrderStatusInOrderData(context context.Context, orderId primitive.ObjectID, status constants.Status) error {

	orderData := mongodb.OpenOrderDataCollection(dataBase.Data)

	result, err := orderData.UpdateOne(context,
		bson.M{
			"_id": orderId,
		},
		bson.M{
			"$set": bson.M{
				"status": status,
			},
		},
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to update order")

}

func (dataBase *DataBase) GetOrderInfoFromOrderData(context context.Context, orderId primitive.ObjectID) (structs.OrderData, error) {

	orderData := mongodb.OpenOrderDataCollection(dataBase.Data)

	var data structs.OrderData

	err := orderData.FindOne(context,
		bson.M{
			"_id": orderId,
		},
	).Decode(&data)

	if err != nil {
		return structs.OrderData{}, err
	}

	return data, nil

}

func (dataBase *DataBase) IsExistOrderExist(context context.Context, key string, value interface{}) bool {
	productData := mongodb.OpenOrderDataCollection(dataBase.Data)

	filter := bson.D{{key, value}}
	singleCursor := productData.FindOne(context, filter)

	if singleCursor.Err() != nil {
		return false
	}

	return true

}

/*
	order := structs.OrderData{
		OrderId:   primitive.ObjectID{},
		UserId:    "",
		Status:    0,
		ProductId: primitive.ObjectID{},
		TimeStamp: time.Time{},
		Price:     0,
		Quantity:  0,
	}

*/
