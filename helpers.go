package main

import (
	"fmt"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

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

//#region Error

func PrintError(logger runtime.Logger, err error, message string) {
	logger.WithField("err", err).Error(message)
}

//#endregion

//#region Encoding Decoding

//#endregion
