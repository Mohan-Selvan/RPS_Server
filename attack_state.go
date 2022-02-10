package main

import (
	"encoding/json"
	"fmt"
)

type AttackState struct {
	players      PlayersMap
	time_pending int

	matchEndMessage []MatchEndMessage
}

func (o *AttackState) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"players":   (o.players),
		"time_pen":  (o.time_pending),
		"match_end": (o.matchEndMessage),
	})
}

func (o *AttackState) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding AttackState")
	}
	return string(encoded)
}
