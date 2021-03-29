package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/constants"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"github.com/go-playground/locales/currency"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (dataBase *DataBase) CreateShop(context context.Context, userId string) (*structs.UserData, error) {

	userData := mongodb.OpenShopDataCollection(dataBase.Data)

	dataInsert := structs.ShopData{
		ShopId:         primitive.ObjectID{},
		ShopName:       "Milap Stores",
		ShopKeeperName: "ABC Person",
		Images:         []string{"https://www.google.com", "https://www.google.com"},
		PrimaryImages:  "https://www.google.com",
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
