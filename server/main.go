package main

import (
	data_base "aapanavyapar-service-viewprovider/data-base/data-services"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"time"
)

func main() {
	database := data_base.NewDataBase()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	//_, err := database.CreateProduct(context.Background(), "1")
	//_, err = database.CreateShop(context.Background(), "1")
	//_, err = database.CreateUser(ctx, "1")
	//if err != nil {
	//	panic(err)
	//}

	address := structs.Address{
		FullName:      "Shitij Shailendra Agrawal",
		HouseDetails:  "B.K Road Chopda",
		StreetDetails: "B.K Road Chopda",
		LandMark:      "HDFC Bank",
		PinCode:       "425107",
		City:          "Chopda",
		State:         "Maharastra",
		Country:       "India",
		PhoneNo:       "9172879779",
	}

	err := database.SetAddressInUserData(ctx, "101", "", address)
	if err != nil {
		panic(err)
	}

	address1 := database.GetAddressUserData(ctx, "1")
	if address1 == nil {
		fmt.Println("Party")
	}

	err = database.DelAddressInUserData(ctx, "100")
	if err != nil {
		panic(err)
	}

}
