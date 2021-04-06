package mapper

import (
	"aapanavyapar-service-viewprovider/data-base/structs"
	"time"
)

func MapLocationToSector(location *structs.Location) int64 {

	return 10
}

func CalculateDeliveryTime(distance int64) time.Time {

	return time.Now().UTC()
}

func CalculateDeliveryCost(distance int64, address *structs.Address) float64 {

	return float64(distance) * 5
}
