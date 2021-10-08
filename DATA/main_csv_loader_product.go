package main

import (
	"aapanavyapar-service-viewprovider/data-base/structs"
	"aapanavyapar-service-viewprovider/pb"
	"aapanavyapar-service-viewprovider/services"
	"context"
	"encoding/csv"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"strings"
)

func main() {

	service := services.NewViewProviderService()

	lines, err := ReadCsv("/home/shitij/go/src/aapanavyapar-service-viewprovider/DATA/productData.csv")
	if err != nil {
		panic(err)
	}

	counter := 0
	for _, line := range lines {
		if counter > 0 {

			// 0 :_id, 1 : shop_id, 2 : shop_name, 3 : title, 4 : description, 5 : shipping_info
			// 6 : stock, 7 : likes, 8 : price, 9 : offer, 10 : images, 11 : category, 12 : timestamp

			str := line[11][1:]
			str = str[:len(str)-1]
			data := strings.Split(str, ",")
			var category []pb.Category
			for _, cat := range data {
				cat, _ := strconv.ParseInt(cat, 10, 64)
				category = append(category, pb.Category(cat))
			}

			fmt.Println(category)

			str = line[10][1:]
			str = str[:len(str)-1]
			data = strings.Split(str, ",")
			var images []string
			for _, cat := range data {
				cat = cat[1:]
				cat = cat[:len(cat)-1]
				images = append(images, cat)
			}

			stock, err := strconv.ParseUint(line[6], 10, 64)
			if err != nil {
				panic(status.Errorf(codes.Internal, "Unable To Parse Data"))
			}

			likes, err := strconv.ParseUint(line[7], 10, 32)
			if err != nil {
				panic(status.Errorf(codes.Internal, "Unable To Parse Data"))
			}

			price, err := strconv.ParseFloat(line[8], 32)
			if err != nil {
				panic(status.Errorf(codes.Internal, "Unable To Parse Data"))
			}

			offer, err := strconv.ParseUint(line[9], 10, 32)
			if err != nil {
				panic(status.Errorf(codes.Internal, "Unable To Parse Data"))
			}

			productData := structs.ProductData{
				ShopId:       line[1],
				ShopName:     line[2],
				Title:        line[3],
				Description:  line[4],
				ShippingInfo: line[5],
				Stock:        uint32(stock),
				Likes:        likes,
				Price:        float32(price),
				Offer:        uint32(offer),
				Images:       images,
				Category:     category,
			}

			id, err := service.Data.CreateProduct(context.Background(), productData)
			if err != nil {
				panic(err)
			}
			fmt.Println(id.Hex())

		}
		counter += 1
	}

}
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
