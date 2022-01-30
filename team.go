package main

import (
	"encoding/json"
	"fmt"
)

type PlayersMap map[int]*Player

func (o *PlayersMap) MarshalJSON() ([]byte, error) {

	return json.Marshal(o)
}

func (o *PlayersMap) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding PlayersMap")
	}
	return string(encoded)
}

func (o *PlayersMap) AddPlayer(p *Player) {

	if !o.ContainsPlayer(p) && o.HasEmptySlot() {
		var id = o.GetEmptySlotID()
		(*o)[id] = p
	}
}

func (o *PlayersMap) RemovePlayer(p *Player) {

	for key, value := range *o {
		if value.userID == p.userID {
			delete(*o, key)
		}
	}
}

func (o *PlayersMap) ContainsPlayer(p *Player) bool {

	for _, value := range *o {
		if value.userID == p.userID {
			return true
		}
	}

	return false
}

func (o *PlayersMap) GetPlayer(userID string) *Player {

	for _, value := range *o {
		if value.userID == userID {
			return value
		}
	}

	return nil
}

func (o *PlayersMap) HasEmptySlot() bool {

	for i := 1; i <= MAX_NUMBER_OF_PLAYERS; i++ {
		_, ok := (*o)[i]

		if !ok {
			return true
		}
	}

	return false
}

func (o *PlayersMap) GetEmptySlotID() int {

	maxNumberOfPlayer := 4

	for i := 1; i <= maxNumberOfPlayer; i++ {
		_, isEmpty := (*o)[i]

		if !isEmpty {
			return i
		}
	}

	return -1
}
