package main

import (
	"encoding/json"
	"fmt"
)

type AttackState struct {
	time_pending int

	attacker     *Player
	attackerSign Sign
}

func (o *AttackState) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"time_pen": (o.time_pending),
		"att":      o.attacker.userID,
		"att_sign": o.attackerSign,
	})
}

func (o *AttackState) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding AttackState")
	}
	return string(encoded)
}
