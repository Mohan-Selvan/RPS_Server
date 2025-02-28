package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

//#region DecodersType

type DecoderFunctions struct {
}

var Decoders DecoderFunctions

//#endregion

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
	playerID int
	userID   string

	account *api.Account

	isAI bool

	currentHealth int
	maxHealth     int

	target                     *Player
	hasSelectedMoveForThisTurn bool
	selectedSign               Sign
}

func (o *Player) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"player_id":    (o.playerID),
		"user_id":      (o.userID),
		"is_ai":        (o.isAI),
		"curr_hp":      (o.currentHealth),
		"max_hp":       (o.maxHealth),
		"has_selected": (o.hasSelectedMoveForThisTurn),
		"sign_id":      (o.selectedSign.id),
	})
}

func (o *Player) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding Player")
	}
	return string(encoded)
}

func (p *Player) ModifyHealth(amount int) {
	p.currentHealth = Helpers.ClampInt(p.currentHealth+amount, 0, p.maxHealth)
}

func (p *Player) IsDead() bool {
	return p.currentHealth <= 0
}

//#endregion

//#region MatchEndMessage

type MatchEndMessage struct {
	matchEndState  MatchEndState
	winning_player *Player
	endTime        int64
}

func (o *MatchEndMessage) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"match_end_state": o.matchEndState,
		"winner":          o.winning_player,
	})
}

func (o *MatchEndMessage) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding MatchEndMessage")
	}
	return string(encoded)
}

//#endregion

//#region Metadata

type Metadata struct {
	accountState AccountState
	skinID       int
}

func (o *Metadata) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"state":   o.accountState,
		"skin_id": o.skinID,
	}
}

func (o *Metadata) MarshalJSON() ([]byte, error) {

	return json.Marshal(o.ToMap())
}

func (m *Metadata) UnmarshalJSON(data []byte) error {

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	m.accountState = AccountState(int(v["state"].(float64)))
	m.skinID = int(v["skin_id"].(float64))

	return nil
}

func (d *DecoderFunctions) ParseMetadata(s string) Metadata {

	var metadata Metadata

	e := json.Unmarshal([]byte(s), &metadata)

	if e != nil {
		fmt.Printf("Error Unmarshalling Metadata : %v, Error : %v", s, e)
	}

	return metadata
}

func (o *Metadata) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding Metadata")
	}
	return string(encoded)
}

type Skin struct {
	skinID int
}

func (o *Skin) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"id": o.skinID,
	})
}

func (o *Skin) GetEncodedObject() string {
	encoded, err := json.Marshal(o)

	if err != nil {
		fmt.Println("Error encoding Skin")
	}
	return string(encoded)
}

//#endregion

//#region Payload_EncodersDecoders

type PayloadMap map[string]interface{}

func GetDecodedObject(payload string, logger runtime.Logger) PayloadMap {

	response := make(PayloadMap)

	e := json.Unmarshal([]byte(payload), &response)

	if e != nil {
		logger.Error("Error UnMarshalling Payload : %v, Error : %v", payload, e)
	}

	return response
}

func (p *PayloadMap) Contains(key string) bool {
	if _, ok := (*p)[key]; ok {
		return true
	}
	return false
}

func (p *PayloadMap) ContainsAll(keys ...string) bool {

	for _, key := range keys {
		if _, ok := (*p)[key]; !ok {
			return false
		}
	}

	return true
}

//#endregion
