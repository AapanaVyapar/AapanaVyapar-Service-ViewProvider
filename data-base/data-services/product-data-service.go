package data_base

import (
	"aapanavyapar-service-viewprovider/configurations/mongodb"
	"aapanavyapar-service-viewprovider/data-base/helpers"
	"aapanavyapar-service-viewprovider/data-base/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"time"
)

func (dataBase *DataBase) CreateProduct(context context.Context, shopId primitive.ObjectID, dataInsert structs.ProductData) error {

	if err := helpers.Validate(dataInsert); err != nil {
		return err
	}

	if !dataBase.IsExistShopExist(context, "shop_id", shopId) {
		return fmt.Errorf("shop with specified id does not exist")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataInsert.Timestamp = time.Now().UTC()

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.InsertOne(context, dataInsert)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) GetAllProductsOfShopByFunctionFromProductData(context context.Context, shopId primitive.ObjectID, sendData func(data structs.ProductData) error) error {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	filter := bson.D{{"shop_id", shopId}}
	cursor, err := productData.Find(context, filter)

	if err != nil {
		return err
	}
	defer cursor.Close(context)

	for cursor.Next(context) {
		result := structs.ProductData{}
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

func (dataBase *DataBase) GetAllProductsOfShopByArrayFromProductData(context context.Context, shopId primitive.ObjectID) ([]structs.ProductData, error) {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	filter := bson.D{{"shop_id", shopId}}
	cursor, err := productData.Find(context, filter)

	if err != nil {
		return []structs.ProductData{}, err
	}
	defer cursor.Close(context)

	if err := cursor.Err(); err != nil {
		return []structs.ProductData{}, err
	}

	var results []structs.ProductData
	err = cursor.All(context, &results)
	if err != nil {
		return []structs.ProductData{}, err
	}

	return results, nil

}

func (dataBase *DataBase) GetSpecificProductsOfShopFromProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID) (structs.ProductData, error) {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	filter := bson.D{{"shop_id", shopId}, {"product_id", productId}}

	var data structs.ProductData
	err := productData.FindOne(context, filter).Decode(&data)

	if err != nil {
		return structs.ProductData{}, err
	}

	return data, nil

}

func (dataBase *DataBase) DelProductFromProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID) error {
	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	filter := bson.M{"shop_id": shopId, "product_id": productId}

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.DeleteOne(context, filter)
	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) AddProductImageInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, imageURL string) error {

	if _, err := url.ParseRequestURI(imageURL); err != nil {
		return fmt.Errorf("invalid image url")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
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

func (dataBase *DataBase) DelProductImageFromProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, imageURL string) error {

	if _, err := url.ParseRequestURI(imageURL); err != nil {
		return fmt.Errorf("invalid image url")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
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

func (dataBase *DataBase) UpdateProductTitleInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, title string) error {

	if title == "" {
		return fmt.Errorf("title can not be empty")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"title": title,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateProductDescriptionInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, description string) error {

	if description == "" {
		return fmt.Errorf("description can not be empty")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"description": description,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateProductShippingInfoInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, shippingInfo string) error {

	if shippingInfo == "" {
		return fmt.Errorf("shipping info can not be empty")
	}

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"shipping_info": shippingInfo,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateProductStockInfoInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, stock uint32) error {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"stock": stock,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateProductPriceInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, price float64) error {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"price": price,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (dataBase *DataBase) UpdateProductOfferInProductData(context context.Context, shopId primitive.ObjectID, productId primitive.ObjectID, offer uint8) error {

	productData := mongodb.OpenProductDataCollection(dataBase.Data)

	dataBase.mutex.Lock()
	defer dataBase.mutex.Unlock()

	_, err := productData.UpdateOne(context,
		bson.M{
			"shop_id":    shopId,
			"product_id": productId,
		},
		bson.M{
			"$set": bson.M{
				"offer": offer,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

/*
	dataProduct := structs.ProductData{
		ShopId:       dataInsert.ShopId,
		ProductId:    primitive.NewObjectID(),
		Title:        "Yellow Shirt",
		Description:  "Best in Class Size XL",
		ShippingInfo: "200x70x10",
		Stock:        10,
		Price:        1000,
		Offer:        10,
		Images:       []string{"https://image.com"},
		Timestamp:    time.Now().UTC(),
	}

*/
