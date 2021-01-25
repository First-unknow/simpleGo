package src

import (
	"fmt"
	"strconv"
	"time"
)

func promotionDalta(priceRiseStartDate string, priceRiseEndDate string) bool {
	startDate, startDateErr := time.Parse(time.RFC3339, priceRiseStartDate)
	if startDateErr != nil {
		fmt.Println(startDateErr)
	}
	endDate, endDateErr := time.Parse(time.RFC3339, priceRiseEndDate)
	if endDateErr != nil {
		fmt.Println(startDateErr)
	}
	currentTime := time.Now().Local()
	return currentTime.After(startDate) && currentTime.Before(endDate)
}

func checkMaxCapabilityOil(volumn, maxCapabilityOil int) float64 {
	if volumn > maxCapabilityOil {
		return float64(maxCapabilityOil)
	}
	return float64(volumn)
}

func PointOil(volumn, maxCapabilityOil int, currentProduct map[string]string, memberCardType string) float64 {
	promotion := 1.0
	checkPromotion := promotionDalta(currentProduct["priceRiseStartDate"], currentProduct["priceRiseEndDate"])
	if checkPromotion {
		changeStringToFloat, err := strconv.ParseFloat(currentProduct["samePriseDelta"], 64)
		if err != nil {
			fmt.Println(err)
		}
		promotion = changeStringToFloat
	}
	newVolumn := checkMaxCapabilityOil(volumn, maxCapabilityOil)

	if currentProduct["productGroup"] == "Gasohol" {
		if memberCardType != "Normal" {
			return (newVolumn * 1.25) * promotion
		} else {
			return (newVolumn * 1) * promotion
		}
	} else if currentProduct["productGroup"] == "Diesel" {
		if memberCardType != "Normal" {
			return (newVolumn / 2) * 1 * promotion
		} else {
			return (newVolumn / 4) * 1 * promotion
		}
	}
	return -1
}

func PointNonOil(price, maxCapabilityNonOil int, currentProduct map[string]string) float64 {
	promotion := 1.0
	newPrice := checkMaxCapabilityOil(price, maxCapabilityNonOil)
	checkPromotion := promotionDalta(currentProduct["priceRiseStartDate"], currentProduct["priceRiseEndDate"])
	if checkPromotion {
		changeStringToFloat, err := strconv.ParseFloat(currentProduct["samePriseDelta"], 64)
		if err != nil {
			fmt.Println(err)
		}
		promotion = changeStringToFloat
	}
	promotionDalta(currentProduct["priceRiseStartDate"], currentProduct["priceRiseEndDate"])
	fmt.Println(currentProduct)
	return newPrice / 25 * promotion
}
