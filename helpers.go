package main

import (
	"math/rand"
)

type HelperFunctions struct {
}

var Helpers HelperFunctions

func (h HelperFunctions) GetRandomInt(minInclusive int, maxInclusive int) int {
	return rand.Intn(maxInclusive-minInclusive) + minInclusive
}
