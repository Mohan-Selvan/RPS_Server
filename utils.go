package main

import (
	"fmt"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

//#region Convertors

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StringToInt(value string) int {
	response, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Error Parsing int from string, string : %s, Error : %v", value, err)
	}

	return response
}

func StringToInt64(value string) int64 {
	response, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Printf("Error Parsing int64 from string, string : %s, Error : %v", value, err)
	}

	return response
}

//Bool
func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

func StringToBool(value string) bool {
	response, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Error Parsing Bool, Input : %v", value)
	}

	return response
}

//#endregion

//#region Error

func PrintError(logger runtime.Logger, err error, message string) {
	logger.WithField("err", err).Error(message)
}

//#endregion

//#region Encoding Decoding

//#endregion
