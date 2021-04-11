package mapper

import (
	"aapanavyapar-service-viewprovider/data-base/structs"
	"time"
)

func MapLocationToSector(location *structs.Location) int32 {

	return 10
}

func CalculateDeliveryTime(distance int32) time.Time {

	return time.Now().UTC()
}

func CalculateDeliveryCost(distance int32, address *structs.Address) float32 {

	return float32(distance) * 5
}
