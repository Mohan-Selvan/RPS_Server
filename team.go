package main

type PlayersMap map[int]*Player

func (o PlayersMap) AddPlayer(p *Player) {

	if !o.ContainsPlayer(p) && o.HasEmptySlot() {
		var id = o.GetEmptySlot()
		(o)[id] = p
	}
}

func (o PlayersMap) RemovePlayer(p *Player) {

	for key, value := range o {
		if value.userID == p.userID {
			delete(o, key)
		}
	}
}

func (o PlayersMap) ContainsPlayer(p *Player) bool {

	for _, value := range o {
		if value.userID == p.userID {
			return true
		}
	}

	return false
}

func (o PlayersMap) HasEmptySlot() bool {

	maxNumberOfPlayer := 4

	for i := 1; i <= maxNumberOfPlayer; i++ {
		_, ok := (o)[i]

		if !ok {
			return true
		}
	}

	return false
}

func (o PlayersMap) GetEmptySlot() int {

	maxNumberOfPlayer := 4

	for i := 1; i <= maxNumberOfPlayer; i++ {
		_, isEmpty := (o)[i]

		if isEmpty {
			return i
		}
	}

	return -1
}
