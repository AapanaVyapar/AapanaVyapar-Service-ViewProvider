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
		bson.M{
			"user_id": userId,
		},
		bson.D{
			{"$set",
				bson.D{
					{"address", structs.Address{}},
				},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) GetAddressUserData(context context.Context, userId string) (*structs.Address, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"user_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Address, nil
}

func (dataBase *DataBase) GetCartUserData(context context.Context, userId string) (*structs.ShopAndProductIds, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"user_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Cart, nil
}

func (dataBase *DataBase) AddToCartUserData(context context.Context, userId string, productId string) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	data := structs.UserData{}
	err := userData.FindOne(context, bson.M{"user_id": userId}).Decode(&data)
	if err != nil {
		return err
	}

	if data.Cart == nil {
		// User Want To Add Product First Time
		cart := &structs.ShopAndProductIds{
			Product: []structs.ProductCartData{
				{
					ProductId:   productId,
					NoOfProduct: 0,
				},
			},
		}
		_, err = userData.UpdateOne(context, bson.M{"user_id": userId}, bson.M{"$set": bson.M{"cart": cart}})
		if err != nil {
			return err
		}
		return nil
	}

	for _, prod := range data.Cart.Product {
		if prod.ProductId == productId {
			// User Want To Increase Count Of Product
			_, err = userData.UpdateOne(context,
				bson.M{
					"user_id":                  userId,
					"cart.products.product_id": productId,
				},
				bson.D{
					{"$set",
						bson.M{
							"cart.products.$": bson.D{
								{"product_id", productId},
								{"no_product", prod.NoOfProduct + 1},
							},
						},
					},
				},
			)

			if err != nil {
				return err
			}

			return nil
		}
	}

	// Means User Want To Add New Product In Cart

	_, err = userData.UpdateOne(context,
		bson.M{
			"user_id": userId,
		},
		bson.D{
			{"$push",
				bson.M{
					"cart.products": bson.D{
						{"product_id", productId},
						{"no_product", 0},
					},
				},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) DelFromCartUserData(context context.Context, userId string, productId string) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	_, err := userData.UpdateOne(context,
		bson.M{
			"user_id": userId,
		},
		bson.M{"$pull": bson.M{
			"cart.products": bson.D{
				{"product_id", productId},
			},
		},
		},
	)

	if err != nil {
		return err
	}

	return nil
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
