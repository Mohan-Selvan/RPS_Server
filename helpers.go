package main

import (
	"fmt"
	"strconv"
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

//#region Encoding Decoding

//#endregion
