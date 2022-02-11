package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
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

func GetSplitMinMaxFloats(value string, logger runtime.Logger) (float64, float64) {

	splits := strings.Split(value, "_")

	minValue := float64(0)
	maxValue := float64(0)

	minValue, e1 := strconv.ParseFloat(splits[0], 64)

	if e1 != nil {
		logger.Error("Error occured while splitting MinMaxFloats : %v", e1)
	}

	maxValue, e2 := strconv.ParseFloat(splits[1], 64)
	if e2 != nil {
		logger.Error("Error occured while splitting MinMaxFloats : %v", e1)
	}

	return minValue, maxValue
}

func GetFirstValueReplacedInSplitFloats(value string, replacementValue float64, logger runtime.Logger) string {

	_, secondValue := GetSplitMinMaxFloats(value, logger)

	return SplitFloatsToString(replacementValue, secondValue)
}

func GetSecondValueReplacedInSplitFloats(value string, replacementValue float64, logger runtime.Logger) string {

	firstValue, _ := GetSplitMinMaxFloats(value, logger)

	return SplitFloatsToString(firstValue, replacementValue)
}

func SplitFloatsToString(value1 float64, value2 float64) string {
	return fmt.Sprintf("%.2f_%.2f", value1, value2)
}

func Float64ToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func TimeToString(unixTime int64) string {
	return fmt.Sprint(unixTime)
}

func StringToTime(timeString string) int64 {
	return StringToInt64(timeString)
}

func GetAddedSplitFloat(value string, logger runtime.Logger) float64 {

	firstValue, secondValue := GetSplitMinMaxFloats(value, logger)

	return (firstValue + secondValue)
}

func StringToFloat64(value string) float64 {

	result := float64(0)
	result, e := strconv.ParseFloat(value, 64)

	if e != nil {
		fmt.Printf("Error StringToFloat64, String : %s, Float64 : %v", value, result)
	}

	return result
}

// func (h HelperFunctions) CantonPair_Encode(id1 int, id2 int) int {
// 	key := int(pairing.Encode(uint64(id1), uint64(id2)))
// 	return key
// }
