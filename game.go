package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func HandleGameState_INPROGRESS(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	for _, message := range messages {

		switch message.GetOpCode() {
		case c2s_PlayerMove:
			HandlePlayerMove(message.GetUserId(), ctx, logger, db, nk, dispatcher, tick, state, messages)
			break

		default:
			logger.Info("Invalid Opcode received : %d", message.GetOpCode())
			break
		}
	}
}

func HandlePlayerMove(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	//Calculate Attack here.
	logger.Info("Player move made here by %v", userID)
}
