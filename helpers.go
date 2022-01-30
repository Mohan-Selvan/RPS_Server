package main

import (
	"math/rand"
	"time"
)

type HelperFunctions struct {
}

var Helpers HelperFunctions

func (h HelperFunctions) GetCurrentTime_Unix() int64 {
	return time.Now().UTC().Unix()
}

func (h HelperFunctions) GetRandomInt(minInclusive int, maxInclusive int) int {
	return rand.Intn(maxInclusive-minInclusive) + minInclusive
}

func (h HelperFunctions) GetHitFromPercentage(percentage float64) bool {
	rn := 0 + rand.Float64()*(100-0)
	//fmt.Println("Random Number generated : ", rn, " Hit : ", (rn <= percentage))
	if rn <= percentage {
		return true
	}
	return false
}

//#region HelperFunctions
func (h HelperFunctions) GetHitFromPercentage_1(percentage float64) bool {
	return h.GetHitFromPercentage(percentage * 100)
}

func (h HelperFunctions) GetRandomFloat(minValue float64, maxValue float64) float64 {
	rn := minValue + rand.Float64()*(maxValue-minValue)
	return rn
}

func (h HelperFunctions) ClampFloat(value float64, minValue float64, maxValue float64) float64 {

	if value > maxValue {
		value = maxValue
	} else if value < minValue {
		value = minValue
	}
	return value
}

func (h HelperFunctions) ClampInt(value int, minValue int, maxValue int) int {

	if value > maxValue {
		value = maxValue
	} else if value < minValue {
		value = minValue
	}
	return value
}

func (h HelperFunctions) ClampInt64(value int64, minValue int64, maxValue int64) int64 {

	if value > maxValue {
		value = maxValue
	} else if value < minValue {
		value = minValue
	}
	return value
}

// func (h HelperFunctions) CantonPair_Encode(id1 int, id2 int) int {
// 	key := int(pairing.Encode(uint64(id1), uint64(id2)))
// 	return key
// }
