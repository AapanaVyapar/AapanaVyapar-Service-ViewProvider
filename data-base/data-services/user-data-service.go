package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dataBase *DataBase) CreateUser(context context.Context, userId string, userName string) (primitive.ObjectID, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	if dataBase.IsExistInUserData(context, "user_id", userId) {
		return primitive.ObjectID{}, fmt.Errorf("already exist")
	}

	dataInsert := structs.UserData{
		Id:        primitive.NewObjectID(),
		UserId:    userId,
		UserName:  userName,
		Address:   nil,
		Cart:      nil,
		Favorites: nil,
		Orders:    nil,
	}

	id, err := userData.InsertOne(context, dataInsert)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id.InsertedID.(primitive.ObjectID), nil
}

func (dataBase *DataBase) IsExistInUserData(context context.Context, key string, value interface{}) bool {
	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{key, value}}
	singleCursor := userData.FindOne(context, filter)

	if singleCursor.Err() != nil {
		return false
	}

	return true
}

func (dataBase *DataBase) SetAddressInUserData(context context.Context, userId string, userName string, address structs.Address) error {

	_, err := dataBase.CreateUser(context, userId, userName) // If already exist not going to create.
	if err != nil {
		return err
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err = userData.UpdateOne(context,
		bson.M{"user_id": userId},
		bson.D{
			{"$set", bson.D{{"address", address}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) DelAddressInUserData(context context.Context, userId string) error {

	if !dataBase.IsExistInUserData(context, "user_id", userId) {
		return fmt.Errorf("user does not exist")
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := userData.UpdateOne(context,
		bson.M{"user_id": userId},
		bson.D{
			{"$set", bson.D{{"address", structs.Address{}}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) GetAddressUserData(context context.Context, userId string) *structs.Address {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"user_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil
	}

	return data.Address
}

/*
	dataInsert := structs.UserData{
		Id:        primitive.ObjectID{},
		UserId:    "1",
		UserName:  "Shitij",
		Address:   structs.Address{
			FullName:      "Shitij Shailendra Agrawal",
			HouseDetails:  "B.K Road Chopda",
			StreetDetails: "B.K Road Chopda",
			LandMark:      "HDFC Bank",
			PinCode:       "425107",
			City:          "Chopda",
			State:         "Maharastra",
			Country:       "India",
			PhoneNo:       "9172879779",
		},
		Cart:      structs.ShopAndProductIds{
			Product: []string{"1", "2"},
		},
		Favorites: structs.ShopAndProductIds{
			Product: []string{"1", "2"},
		},
		Orders:    structs.ShopAndProductIds{
			Product: []string{"1", "2"},
		},
	}

*/
