package main

import (
	"encoding/json"
	"fmt"
)

type AttackState struct {
	time_pending int
}

func (o *AttackState) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"time_pen": IntToString(o.time_pending),
	})
}

func (o *AttackState) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding AttackState")
	}
	return string(encoded)
}
