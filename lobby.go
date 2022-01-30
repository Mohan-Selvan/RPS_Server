package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

//Game State Handling.
func HandleGameState_LOBBY(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	for _, message := range messages {

		switch message.GetOpCode() {
		case c2s_PlayerReady:
			HandlePlayerReady(ctx, logger, db, nk, dispatcher, tick, state, messages)
			break

		default:
			logger.Info("Invalid Opcode received : %d", message.GetOpCode())
			break
		}
	}
}

func HandlePlayerReady(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
	logger.Error("HandlePlayerReady : Not implemented")
}

func InitializeMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) {
	mState, _ := state.(*MatchState)

	logger.Info("Initialize Match called, Starting match!")

	//Initializing Player for all presences in match.
	for _, value := range mState.presences {
		player, err := InitPlayer(value.GetUserId(), ctx, logger, db, nk, dispatcher, tick, state)
		if err != nil {
			PrintError(logger, err, "Error initializing player")
		}

		mState.players.AddPlayer(player)
	}

	//initializing AI for remaining required players
	for _, value := range AI_ACCOUNTS {

		player, err := InitPlayer(value, ctx, logger, db, nk, dispatcher, tick, state)
		if err != nil {
			PrintError(logger, err, "Error initializing player")
		}

		mState.players.AddPlayer(player)
	}

	for _, player := range mState.players {
		logger.Info("Player in match : %s", player.userID)
	}

	//TODO :: Target implementation pending.
	//Setting target circularly.
	lastKey := -1
	for key, val := range mState.players {

		if lastKey >= 0 {
			val.target = mState.players[lastKey]
		}

		lastKey = key
	}

	//Initial turn goes to 1st Player.
	mState.turnCounter = 1

	mState.currentAttackState = mState.NewAttackStateObject(mState.GetPlayerOfCurrentTurn(), ctx, logger, db, nk, dispatcher, tick, state)
	mState.lastAttackState = *mState.currentAttackState

	encodedMessage := InitMatchMessage{
		players:              mState.players,
		attackState:          mState.currentAttackState,
		timePerTurnInSeconds: TIME_PER_TURN,
	}.GetEncodedObject()

	mState.DispatchMessage(s2c_InitMatch, encodedMessage, ctx, logger, db, nk, dispatcher, tick, state)

	mState.ChangeGameState(GAME_STATE_IN_PROGRESS, logger)
}

func InitPlayer(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) (*Player, error) {

	player := Player{
		userID:        userID,
		currentHealth: 100,
		maxHealth:     100,
	}

	return &player, nil
}
