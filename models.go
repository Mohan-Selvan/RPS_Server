package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//#region Response

type Response struct {
	success bool
	message string
}

func (r Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"success": BoolToString(r.success),
		"message": r.message,
	})
}

func (r Response) GetEncodedObject() string {
	encoded, err := json.Marshal(r)

	if err != nil {
		fmt.Println("Error encodingf RPC Response")
	}

	return string(encoded)
}

//#endregion

//#region Float4
type Floats4 struct {
	a float64
	b float64
	c float64
	d float64
}

var StringToFloats4 = func(value string) Floats4 {

	splits := strings.Split(value, "_")

	splitFloats := make([]float64, 4)

	for i := 0; i < 4; i++ {
		floatValue, e1 := strconv.ParseFloat(splits[i], 64)

		if e1 != nil {
			fmt.Println("Error during Float Conversion")
			floatValue = float64(0)
		}

		splitFloats[i] = floatValue
	}

	return Floats4{splitFloats[0], splitFloats[1], splitFloats[2], splitFloats[3]}
}

func (f Floats4) ToString() string {
	return fmt.Sprintf("%.2f_%.2f_%.2f_%.2f", f.a, f.b, f.c, f.d)
}

//#endregion

//#region Float3
type Floats3 struct {
	a float64
	b float64
	c float64
}

var StringToFloats3 = func(value string) Floats3 {

	splits := strings.Split(value, "_")

	splitFloats := make([]float64, 3)

	for i := 0; i < 3; i++ {
		floatValue, e1 := strconv.ParseFloat(splits[i], 64)

		if e1 != nil {
			fmt.Println("Error duringh Float Conversion")
			floatValue = float64(0)
		}

		splitFloats[i] = floatValue
	}

	return Floats3{splitFloats[0], splitFloats[1], splitFloats[2]}
}

func (f Floats3) ToString() string {
	return fmt.Sprintf("%.2f_%.2f_%.2f", f.a, f.b, f.c)
}

//#endregion

//#region Floats2
type Floats2 struct {
	a float64
	b float64
}

var StringToFloats2 = func(value string) Floats2 {

	splits := strings.Split(value, "_")

	splitFloats := make([]float64, 2)

	for i := 0; i < 2; i++ {
		floatValue, e1 := strconv.ParseFloat(splits[i], 64)

		if e1 != nil {
			fmt.Println("Error during Float Conversion")
			floatValue = float64(0)
		}

		splitFloats[i] = floatValue
	}

	return Floats2{splitFloats[0], splitFloats[1]}
}

func (f Floats2) ToString() string {
	return fmt.Sprintf("%.2f_%.2f", f.a, f.b)
}

func (f Floats2) GetAddedValue() float64 {
	return (f.a + f.b)
}

//#endregion

//#region Player

type Player struct {
	userID string

	currentHealth int
	maxHealth     int
}

//#endregion
