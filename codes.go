package main

import "github.com/heroiclabs/nakama-common/runtime"

const MAX_PLAYER_PVA int = 4

var (
	errUnknownError  error = runtime.NewError("Unknown error", 500)
	errAccountFetch  error = runtime.NewError("Account not found", 501)
	errAccountUpdate error = runtime.NewError("Account update error", 502)
	errInvalidData   error = runtime.NewError("Invalid request", 503)
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

type MatchEndState int

var (
	MATCH_END_STATE_WIN  MatchEndState = 1
	MATCH_END_STATE_DRAW MatchEndState = 2
)

type AccountState int

var (
	ACCOUNT_STATE_COMPLETE   AccountState = 0
	ACCOUNT_STATE_INCOMPLETE AccountState = 1
)
