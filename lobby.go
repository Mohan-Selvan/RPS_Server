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

	for _, value := range mState.presences {
		player, err := InitPlayer(value.GetUserId(), ctx, logger, db, nk, dispatcher, tick, state)
		if err != nil {
			PrintError(logger, err, "Error initializing player")
		}

		mState.players.AddPlayer(player)
	}

	for _, player := range mState.players {
		logger.Info("Player in match : %s", player.userID)
	}

	mState.currentAttackState = mState.NewAttackStateObject(ctx, logger, db, nk, dispatcher, tick, state)
	mState.lastAttackState = *mState.currentAttackState

	encodedMessage := InitMatchMessage{
		players:     mState.players,
		attackState: mState.currentAttackState,
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
