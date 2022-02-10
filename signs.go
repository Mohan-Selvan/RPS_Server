package main

import (
	"encoding/json"
	"fmt"
)

type Sign struct {
	id      int
	damage  int
	defense int
}

func (o *Sign) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"id": o.id,
	})
}

func (o *Sign) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding Sign")
	}
	return string(encoded)
}

type SignsMap map[int]Sign

func (s *SignsMap) GetSign(signID int) Sign {
	if val, ok := (*s)[signID]; ok {
		return val
	}

	fmt.Println("ERROR : Invalid Sign requested.")
	return SIGNS.GetSign(SIGN_ID_EMPTY)
}

func (s *SignsMap) GetRandomSign() Sign {

	randomSignID := Helpers.GetRandomInt(1, len(SIGNS))
	return SIGNS.GetSign(randomSignID)
}

var SIGNS SignsMap = SignsMap{
	SIGN_ID_EMPTY:   SIGN_EMPTY,
	SIGN_ID_ROCK:    SIGN_ROCK,
	SIGN_ID_PAPER:   SIGN_PAPER,
	SIGN_ID_SCISSOR: SIGN_SCISSOR,
}

const SIGN_ID_EMPTY int = 0

var SIGN_EMPTY Sign = Sign{
	id:      SIGN_ID_EMPTY,
	damage:  0,
	defense: 0,
}

const SIGN_ID_ROCK int = 1

var SIGN_ROCK Sign = Sign{
	id:      SIGN_ID_ROCK,
	damage:  5,
	defense: 5,
}

const SIGN_ID_PAPER int = 2

var SIGN_PAPER Sign = Sign{
	id:      SIGN_ID_PAPER,
	damage:  5,
	defense: 5,
}

const SIGN_ID_SCISSOR int = 3

var SIGN_SCISSOR Sign = Sign{
	id:      SIGN_ID_SCISSOR,
	damage:  5,
	defense: 5,
}
