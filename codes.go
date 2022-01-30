package main

import "github.com/heroiclabs/nakama-common/runtime"

const MAX_PLAYER_PVA int = 4

var (
	errUnknownError error = runtime.NewError("Unknown error", 500)
)

//Lobby State OpCodes
var (
	c2s_PlayerReady   int64 = 0
	c2s_PlayerUnReady int64 = 1

	//InGame
	c2s_PlayerMove int64 = 2
)

var (
	s2c_InitMatch  int64 = 1
	s2c_PlayerMove int64 = 2
)

type GameState int

var (
	GAME_STATE_LOBBY       GameState = 1
	GAME_STATE_IN_PROGRESS GameState = 2
	GAME_STATE_END         GameState = 3
)
