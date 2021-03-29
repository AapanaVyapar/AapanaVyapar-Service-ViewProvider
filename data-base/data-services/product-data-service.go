package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (dataBase *DataBase) CreateProduct(context context.Context, userId string) (*structs.UserData, error) {

	userData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataInsert := structs.ProductData{
		ShopId:       primitive.ObjectID{1},
		ProductId:    primitive.ObjectID{},
		Title:        "Yellow Shirt",
		Description:  "Best in Class Size XL",
		ShippingInfo: "200x70x10",
		Stock:        "10",
		Price:        "1000",
		Offer:        10,
		Images:       []string{"https://image.com"},
		Timestamp:    time.Now().UTC(),
	}
	id, err := userData.InsertOne(context, dataInsert)
	if err != nil {
		return nil, err
	}

	fmt.Println(id)

	filter := bson.D{{"user_id", userId}}

	data := structs.UserData{}
	err = userData.FindOne(context, filter).Decode(&data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return &data, nil
}
