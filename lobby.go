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

	logger.Info("Initializing match..")

	//Initializing Player for all presences in match.

	for _, value := range mState.presences {

		isAi := false

		player, err := InitPlayer(value.GetUserId(), isAi, ctx, logger, db, nk, dispatcher, tick, state)
		if err != nil {
			PrintError(logger, err, "Error initializing player")
		}

		mState.players.AddPlayer(player)
	}

	//initializing AI for remaining required players
	for _, value := range AI_ACCOUNTS {
		isAi := true
		player, err := InitPlayer(value, isAi, ctx, logger, db, nk, dispatcher, tick, state)
		if err != nil {
			PrintError(logger, err, "Error initializing player")
		}

		mState.players.AddPlayer(player)
	}

	//Assigning Player ID here.
	for playerID, player := range mState.players {
		player.playerID = playerID

		logger.Info("Player in match (Id : %v) : %s", player.playerID, player.userID)
	}

	//TODO :: Target implementation pending.
	//Setting target circularly.
	previousKey := -1
	for key, val := range mState.players {

		if previousKey >= 0 {
			val.target = mState.players[previousKey]
		}

		previousKey = key
	}
	mState.players[1].target = mState.players[len(mState.players)]

	//Initial turn goes to 1st Player.

	mState.currentAttackState = mState.NewAttackStateObject(ctx, logger, db, nk, dispatcher, tick, state)
	mState.lastAttackState = *mState.currentAttackState

	encodedMessage := InitMatchMessage{
		attackState:          mState.currentAttackState,
		timePerTurnInSeconds: TIME_PER_TURN,
	}.GetEncodedObject()

	mState.DispatchMessage(s2c_InitMatch, encodedMessage, ctx, logger, db, nk, dispatcher, tick, state)

	mState.ChangeGameState(GAME_STATE_IN_PROGRESS, logger)
}

//Initializes Player with default values.
//NOTE :: PlayerID assigned here is a placeholder.
func InitPlayer(userID string, isAI bool, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) (*Player, error) {

	playerAccount, err := nk.AccountGetId(ctx, userID)
	if err != nil {
		PrintError(logger, err, "Error fetching account")
	}

	player := Player{
		playerID: 0,
		userID:   userID,

		account: playerAccount,

		isAI: isAI,

		currentHealth: 100,
		maxHealth:     100,
	}

	return &player, nil
}
