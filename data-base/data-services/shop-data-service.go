package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/constants"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/data-base/mapper"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"github.com/go-playground/locales/currency"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

func (dataBase *DataBase) CreateShop(context context.Context, dataInsert structs.ShopData) error {

	if err := helpers.Validate(dataInsert); err != nil {
		return err
	}

	for _, i := range dataInsert.Images {
		if _, err := url.ParseRequestURI(i); err != nil {
			return fmt.Errorf("invalid url")
		}
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataInsert.SectorNo = mapper.MapLocationToSector(dataInsert.Location)
	dataInsert.Timestamp = time.Now().UTC()

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.InsertOne(context, dataInsert)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) GetShopFromShopData(context context.Context, shopId primitive.ObjectID) (structs.ShopData, error) {

	userData := mongodb.OpenShopDataCollection(dataBase.Data)

	filter := bson.D{{"shop_id", shopId}}

	data := structs.ShopData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return structs.ShopData{}, err
	}

	return data, nil
}

func (dataBase *DataBase) AddRatingInShopData(context context.Context, shopId primitive.ObjectID, rating structs.Rating) error {

	if err := helpers.Validate(rating); err != nil {
		return err
	}

	userData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	result := userData.FindOne(context, bson.M{"shop_id": shopId, "ratings.user_id": rating.UserId})

	// Error will be thrown if rating is null or rating is already present in both cases we have to just add product
	if result.Err() != nil {
		_, err := userData.UpdateOne(context,
			bson.M{
				"shop_id": shopId,
			},
			bson.D{
				{"$push",
					bson.M{
						"ratings": rating,
					},
				},
			},
		)

		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("you already rated to the shop")

}

func (dataBase *DataBase) GetRatingsFromShopData(context context.Context, shopId primitive.ObjectID) (*[]structs.Rating, error) {

	userData := mongodb.OpenShopDataCollection(dataBase.Data)

	filter := bson.D{{"shop_id", shopId}}

	data := structs.ShopData{}
	err := userData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data.Ratings, nil
}

func (dataBase *DataBase) IsExistShopExist(context context.Context, key string, value interface{}) bool {
	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	filter := bson.D{{key, value}}
	singleCursor := shopData.FindOne(context, filter)

	if singleCursor.Err() != nil {
		return false
	}

	return true

}

func (dataBase *DataBase) DelShopFromShopData(context context.Context, shopId primitive.ObjectID) error {
	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	filter := bson.M{"shop_id": shopId}

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.DeleteOne(context, filter)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) DelShopImageFromShopData(context context.Context, shopId primitive.ObjectID, imageURL string) error {

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$pull": bson.M{
				"images": imageURL,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) AddShopImageInShopData(context context.Context, shopId primitive.ObjectID, imageURL string) error {

	if _, err := url.ParseRequestURI(imageURL); err != nil {
		return fmt.Errorf("invalid image url")
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$push": bson.M{
				"images": imageURL,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateShopPrimaryImageInShopData(context context.Context, shopId primitive.ObjectID, imageURL string) error {

	if _, err := url.ParseRequestURI(imageURL); err != nil {
		return fmt.Errorf("invalid image url")
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"primary_image": imageURL,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateShopKeeperNameInShopData(context context.Context, shopId primitive.ObjectID, name string) error {

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"shop_keeper_name": name,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateShopAddressAndLocationInShopData(context context.Context, shopId primitive.ObjectID, address structs.Address, location structs.Location) error {

	if err := helpers.Validate(address); err != nil {
		return err
	}
	if err := helpers.Validate(location); err != nil {
		return err
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)
	sectorNo := mapper.MapLocationToSector(&location)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"address":   address,
				"location":  location,
				"sector_no": sectorNo,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateShopLocationInShopData(context context.Context, shopId primitive.ObjectID, location structs.Location) error {

	if err := helpers.Validate(location); err != nil {
		return err
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"location": location,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateCategoryInShopData(context context.Context, shopId primitive.ObjectID, category []constants.Categories) error {

	if len(category) == 0 {
		return fmt.Errorf("category can not be empty")
	}

	for _, c := range category {
		if c > 0 && c <= constants.MAX_CATEGORY {
			continue
		} else {
			return fmt.Errorf("invalid category")
		}
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"category": category,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) GetCategoryFromShopData(context context.Context, shopId primitive.ObjectID) ([]constants.Categories, error) {

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	shop := structs.ShopData{}
	err := shopData.FindOne(context,
		bson.M{
			"shop_id": shopId,
		},
	).Decode(&shop)

	if err != nil {
		return []constants.Categories{}, err
	}

	return shop.Category, nil
}

func (dataBase *DataBase) UpdateBusinessInfoInShopData(context context.Context, shopId primitive.ObjectID, info string) error {

	if info == "" {
		return fmt.Errorf("can not accept empty information")
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"business_information": info,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateCurrencyInShopData(context context.Context, shopId primitive.ObjectID, currency currency.Type) error {

	if currency > constants.MAX_CURRENCY {
		return fmt.Errorf("currency not supported")
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"currency": currency,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateOperationalHoursInShopData(context context.Context, shopId primitive.ObjectID, operationalHours structs.OperationalHours) error {

	if err := helpers.Validate(operationalHours); err != nil {
		return err
	}

	shopData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := shopData.UpdateOne(context,
		bson.M{
			"shop_id": shopId,
		},
		bson.M{
			"$set": bson.M{
				"operational_hours": operationalHours,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

/*
	dataInsert := structs.ShopData{
		ShopId:         primitive.NewObjectID(),
		ShopName:       "Milap Stores",
		ShopKeeperName: "ABC Person",
		Images:         []string{"https://www.google.com", "https://www.google.com"},
		PrimaryImage:  "https://www.google.com",
		Address: &structs.Address{
			FullName:      "ABC Person",
			HouseDetails:  "Milap Store",
			StreetDetails: "Mustufa Chishti Colony Main Rd, Panchshil Nagar",
			LandMark:      "",
			PinCode:       "425107",
			City:          "chopda",
			State:         "maharastra",
			Country:       "india",
			PhoneNo:       "+919890713171",
		},
		Location: &structs.Location{
			Longitude: "21.246435522726177",
			Latitude:  "75.29615236552934",
		},
		SectorNo:            10,
		Category:            []constants.Categories{constants.MENS_ACCSSORIES, constants.WONENS_CLOTHING},
		BusinessInformation: "Famous Seller Of Cloths In Chopda",
		Currency:            currency.INR,
		OperationalHours: &structs.OperationalHours{
			Sunday:    [2]string{"", ""},
			Monday:    [2]string{"9AM", "9PM"},
			Tuesday:   [2]string{"9AM", "9PM"},
			Wednesday: [2]string{"9AM", "9PM"},
			Thursday:  [2]string{"9AM", "9PM"},
			Friday:    [2]string{"9AM", "9PM"},
			Saturday:  [2]string{"9AM", "9PM"},
		},
		Ratings: &structs.Rating{
			UserId:    "1",
			UserName:  "ABC",
			Comment:   "Excellent",
			Rating:    constants.GOOD,
			Timestamp: time.Now().UTC(),
		},
		Timestamp: time.Now().UTC(),
	}

*/
/*
	Declarations : Once the shop is created you can not change its name but you can delete your shop.

*/
