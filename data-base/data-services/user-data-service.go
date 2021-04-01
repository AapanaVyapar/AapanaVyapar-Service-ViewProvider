package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dataBase *DataBase) CreateUser(context context.Context, userId string, userName string) (string, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	if dataBase.IsExistInUserData(context, "_id", userId) {
		return "", fmt.Errorf("already exist")
	}

	dataInsert := structs.UserData{
		UserId:    userId,
		UserName:  userName,
		Address:   nil,
		Cart:      nil,
		Favorites: nil,
		Orders:    nil,
	}

	id, err := userData.InsertOne(context, dataInsert)
	if err != nil {
		return "", err
	}

	return id.InsertedID.(string), nil
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

	if err := helpers.Validate(address); err != nil {
		return err
	}

	_, err := dataBase.CreateUser(context, userId, userName) // If already exist not going to create.
	if err != nil {
		return err
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	result, err := userData.UpdateOne(context,
		bson.M{"_id": userId},
		bson.D{
			{"$set", bson.D{{"address", address}}},
		},
	)
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to set the address")
}

func (dataBase *DataBase) DelAddressInUserData(context context.Context, userId string) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	result, err := userData.UpdateOne(context,
		bson.M{
			"_id": userId,
		},
		bson.D{
			{"$set",
				bson.D{
					{"address", nil},
				},
			},
		},
	)
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to delete address")
}

func (dataBase *DataBase) GetAddressUserData(context context.Context, userId string) (*structs.Address, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Address, nil
}

func (dataBase *DataBase) GetCartUserData(context context.Context, userId string) (*structs.ProductIdsForCart, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Cart, nil
}

func (dataBase *DataBase) GetFavoritesUserData(context context.Context, userId string) (*structs.ProductIdsForFavAndOrd, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Favorites, nil
}

func (dataBase *DataBase) GetOrdersUserData(context context.Context, userId string) (*structs.ProductIdsForFavAndOrd, error) {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	filter := bson.D{{"_id", userId}}

	data := structs.UserData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Orders, nil
}

func (dataBase *DataBase) AddToCartUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	if !dataBase.IsExistProductExist(context, "_id", productId) {
		return fmt.Errorf("product does not exist")
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	result := userData.FindOne(context, bson.M{"_id": userId, "cart.products.product_id": productId})

	// Error will be thrown if cart is null or product is not in cart in both cases we have to just add product
	if result.Err() != nil {
		result, err := userData.UpdateOne(context,
			bson.M{
				"_id": userId,
			},
			bson.M{
				"$push": bson.M{
					"cart.products": bson.D{
						{"product_id", productId},
						{"no_product", 1},
					},
				},
			},
		)
		if err != nil {
			return err
		}

		if result.ModifiedCount > 0 || result.MatchedCount > 0 {
			return nil
		}

		return fmt.Errorf("unable to add to cart")

	}

	res, err := userData.UpdateOne(context,
		bson.M{
			"_id":                      userId,
			"cart.products.product_id": productId,
		},
		bson.M{
			"$inc": bson.M{
				"cart.products.$.no_product": 1,
			},
		},
	)
	if err != nil {
		return err
	}

	if res.ModifiedCount > 0 || res.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to add to cart")
}

func (dataBase *DataBase) AddToFavoritesUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	if !dataBase.IsExistProductExist(context, "_id", productId) {
		return fmt.Errorf("product does not exist")
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	result := userData.FindOne(context, bson.M{"_id": userId, "favorites.products": productId})

	// Error will be thrown if favorites is null or product is not in favorites in both cases we have to just add product
	if result.Err() != nil {
		res, err := userData.UpdateOne(context,
			bson.M{
				"_id": userId,
			},
			bson.D{
				{"$push",
					bson.M{
						"favorites.products": productId,
					},
				},
			},
		)
		if err != nil {
			return err
		}

		if res.ModifiedCount > 0 || res.MatchedCount > 0 {
			return nil
		}

		return fmt.Errorf("unable to add to faviroute")
	}

	return fmt.Errorf("alredy exist in faviroute")
}

func (dataBase *DataBase) AddToOrdersUserData(context context.Context, userId string, orderId primitive.ObjectID) error {

	if !dataBase.IsExistOrderExist(context, "_id", orderId) {
		return fmt.Errorf("order is not created")
	}

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	//dataBase.mutex.Lock()
	//defer dataBase.mutex.Unlock()

	result := userData.FindOne(context, bson.M{"_id": userId, "orders.products": orderId})

	// Error will be thrown if favorites in null or product is not in favorites in both cases we have to just add product
	if result.Err() != nil {
		res, err := userData.UpdateOne(context,
			bson.M{
				"_id": userId,
			},
			bson.D{
				{"$push",
					bson.M{
						"orders.products": orderId,
					},
				},
			},
		)
		if err != nil {
			return err
		}

		if res.ModifiedCount > 0 || res.MatchedCount > 0 {
			return nil
		}

		return fmt.Errorf("unable to add order")
	}

	return fmt.Errorf("order alredy exist")
}

func (dataBase *DataBase) DelFromCartUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	result, err := userData.UpdateOne(context,
		bson.M{
			"_id": userId,
		},
		bson.M{
			"$pull": bson.M{
				"cart.products": bson.D{
					{"product_id", productId},
				},
			},
		},
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to delete from cart")
}

func (dataBase *DataBase) DelFromFavoritesUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	result, err := userData.UpdateOne(context,
		bson.M{
			"_id": userId,
		},
		bson.M{
			"$pull": bson.M{
				"favorites.products": productId,
			},
		},
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to delete from faviroute")
}

func (dataBase *DataBase) DelFromOrdersUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	result, err := userData.UpdateOne(context,
		bson.M{
			"_id": userId,
		},
		bson.M{
			"$pull": bson.M{
				"orders.products": productId,
			},
		},
	)

	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("unable to delete from order")
}

func (dataBase *DataBase) RemoveFromCartUserData(context context.Context, userId string, productId primitive.ObjectID) error {

	userData := mongodb.OpenUserDataCollection(dataBase.Data)

	result, err := userData.UpdateOne(context,
		bson.M{
			"_id":                      userId,
			"cart.products.product_id": productId,
			"cart.products.no_product": bson.M{"$gt": 1},
		},
		bson.M{
			"$inc": bson.M{
				"cart.products.$.no_product": -1,
			},
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(result.ModifiedCount)

	if result.ModifiedCount > 0 || result.MatchedCount > 0 {
		return nil
	}

	return fmt.Errorf("minimum quantity for cart is one")

}

/*
	dataInsert := structs.UserData{
		_Id:        "1", //userid
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

//func (dataBase *DataBase) AddToCartUserData(context context.Context, userId string, productId primitive.ObjectID) error {
//
//	userData := mongodb.OpenUserDataCollection(dataBase.Data)
//
//	data := structs.UserData{}
//	err := userData.FindOne(context, bson.M{"user_id": userId}).Decode(&data)
//	if err != nil {
//		return err
//	}
//
//	dataBase.mutex.Lock()
//	defer dataBase.mutex.Unlock()
//
//	if data.Cart == nil {
//		// User Want To Add Product First Time
//		cart := &structs.ShopAndProductIdsForCart{
//			Product: []structs.ProductCartData{
//				{
//					ProductId:   productId,
//					NoOfProduct: 1,
//				},
//			},
//		}
//		_, err = userData.UpdateOne(context, bson.M{"user_id": userId}, bson.M{"$set": bson.M{"cart": cart}})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	// User Want To Increase Count Of Product
//	// Only going to increment if product is there.
//	result, err := userData.UpdateOne(context,
//		bson.M{
//			"user_id":                  userId,
//			"cart.products.product_id": productId,
//		},
//		bson.M{"$inc": bson.M{
//			"cart.products.$.no_product": 1,
//		},
//		},
//	)
//
//	if err != nil {
//		return err
//	}
//	if result.ModifiedCount > 0 {
//		return nil
//	}
//
//	// Means User Want To Add New Product In Cart
//
//	_, err = userData.UpdateOne(context,
//		bson.M{
//			"user_id": userId,
//		},
//		bson.D{
//			{"$push",
//				bson.M{
//					"cart.products": bson.D{
//						{"product_id", productId},
//						{"no_product", 1},
//					},
//				},
//			},
//		},
//	)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
