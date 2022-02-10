package main

import (
	"encoding/json"
	"fmt"
)

type PlayersMap map[int]*Player

// func (o *PlayersMap) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(o)
// }

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
	} else {
		fmt.Sprintf("Couldn't add player (%v) to playersMap \n", p.userID)
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

func (o *PlayersMap) ContainsPlayerWithUserID(userID string) bool {

	for _, value := range *o {
		if value.userID == userID {
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

func (o *PlayersMap) GetNoOfPlayersAlive() int {

	var noOfPlayersAlive int = 0

	for _, p := range *o {
		if !p.IsDead() {
			noOfPlayersAlive += 1
		}
	}

	return noOfPlayersAlive
}

func (o *PlayersMap) GetNonDeadPlayers() []*Player {

	alivePlayers := make([]*Player, 0)

	for _, p := range *o {
		if !p.IsDead() {
			alivePlayers = append(alivePlayers, p)
		}
	}

	return alivePlayers
}

func (o *PlayersMap) GetNonDead_AI_Players() []*Player {

	alivePlayers := make([]*Player, 0)

	for _, p := range *o {
		if !p.IsDead() && p.isAI {
			alivePlayers = append(alivePlayers, p)
		}
	}

	return alivePlayers
}

func (o *PlayersMap) GetNonDeadPlayersExcept(userID string) []*Player {

	alivePlayers := make([]*Player, 0)

	for _, p := range *o {
		if !p.IsDead() {
			if userID == p.userID {
				continue
			}
			alivePlayers = append(alivePlayers, p)
		}
	}

	return alivePlayers
}

func (o *PlayersMap) DidEveryoneSelectMove() bool {

	for _, p := range *o {

		if p.IsDead() {
			continue
		}

		if !p.hasSelectedMoveForThisTurn {
			return false
		}
	}

	return true
}
