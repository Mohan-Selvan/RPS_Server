package main

import (
	"encoding/json"
	"fmt"
)

type InitMatchMessage struct {
	players              PlayersMap
	attackState          *AttackState
	timePerTurnInSeconds int
}

func (o InitMatchMessage) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"players":              o.players,
		"att_sta":              o.attackState,
		"time_per_turn_in_sec": (o.timePerTurnInSeconds),
	})
}

func (o InitMatchMessage) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding InitMatchMessage")
	}
	return string(encoded)
}
