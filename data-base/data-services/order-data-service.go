package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/mapper"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"aapanavyapar-service-viewprovider/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (dataBase *MongoDataBase) CreateOrder(context context.Context, userId string, shopId string, productId primitive.ObjectID, quantity uint32, distance int32, address *structs.Address) (primitive.ObjectID, error) {

	if !dataBase.IsExistInUserData(context, "_id", userId) {
		return primitive.ObjectID{}, fmt.Errorf("user does not exist")
	}

	shopId, productName, productImage, err := dataBase.GetShopDetailsFromProductData(context, productId)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("product does not exist")
	}

	var order structs.OrderData

	productData := mongodb.OpenOrderDataCollection(dataBase.Data)
	order.ShopId = shopId
	order.ProductName = productName
	order.ProductImage = productImage
	order.OrderTimeStamp = time.Now().UTC()
	order.UserId = userId
	order.Quantity = quantity
	order.ShopId = shopId
	order.ProductId = productId
	order.Address = address
	order.DeliveryTimeStamp = mapper.CalculateDeliveryTime(distance)
	order.DeliveryCost = mapper.CalculateDeliveryCost(distance, address)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := dataBase.Data.StartSession()
	if err != nil {
		return primitive.ObjectID{}, err
	}
	defer session.EndSession(context)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		price, offer, err := dataBase.DecreaseStockToMakeOrderFromProductData(sessCtx, productId, quantity)
		if err != nil {
			return primitive.ObjectID{}, err
		}

		order.Price = price - ((price / 100) * float32(offer))
		order.Price += order.DeliveryCost

		order.Status = pb.Status_PENDING

		id, err := productData.InsertOne(sessCtx, order)
		if err != nil {
			return primitive.ObjectID{}, err
		}

		err = dataBase.AddToOrdersUserData(sessCtx, userId, id.InsertedID.(primitive.ObjectID))
		if err != nil {
			return primitive.ObjectID{}, err
		}

		return id.InsertedID.(primitive.ObjectID), nil
	}

	result, err := session.WithTransaction(context, callback, txnOpts)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.(primitive.ObjectID), nil
}

func (dataBase *MongoDataBase) UpdateOrderStatusInOrderData(context context.Context, orderId primitive.ObjectID, status pb.Status) error {

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

func (dataBase *MongoDataBase) GetOrderInfoFromOrderData(context context.Context, orderId primitive.ObjectID) (structs.OrderData, error) {

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

func (dataBase *MongoDataBase) GetOrderInfoByShopIdFromOrderData(context context.Context, shopId string, sendData func(data structs.OrderData) error) error {

	orderData := mongodb.OpenOrderDataCollection(dataBase.Data)

	cursor, err := orderData.Find(context,
		bson.M{
			"shop_id": shopId,
		},
	)

	if err != nil {
		return err
	}
	defer cursor.Close(context)

	for cursor.Next(context) {
		result := structs.OrderData{}
		err = cursor.Decode(&result)

		if err != nil {
			return err
		}

		if err = sendData(result); err != nil {
			return err
		}

	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil

}

func (dataBase *MongoDataBase) GetMultipleOrdersInfoByUserIdFromOrderData(context context.Context, userId string, sendData func(data structs.OrderData) error) error {

	orderData := mongodb.OpenOrderDataCollection(dataBase.Data)

	cursor, err := orderData.Find(context,
		bson.M{
			"user_id": userId,
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Unable To Get Data")
	}
	defer cursor.Close(context)

	for cursor.Next(context) {
		result := structs.OrderData{}

		err = cursor.Decode(&result)
		if err != nil {
			return status.Errorf(codes.Internal, "Unable To Get Data")
		}

		if err = sendData(result); err != nil {
			return err
		}

	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Problem Occur While Retrieving Data")
	}

	return nil

}

func (dataBase *MongoDataBase) IsExistOrderExist(context context.Context, key string, value interface{}) bool {
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
