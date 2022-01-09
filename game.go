package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func HandleGameState_INPROGRESS(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	mState, _ := state.(*MatchState)

	for _, message := range messages {

		switch message.GetOpCode() {
		case c2s_PlayerMove:
			mState.HandlePlayerMove(message.GetUserId(), ctx, logger, db, nk, dispatcher, tick, state, messages)
			break

		default:
			logger.Info("Invalid Opcode received : %d", message.GetOpCode())
			break
		}
	}
}

func (mState *MatchState) HandlePlayerMove(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	//Update Player selection here.
	logger.Info("Player move made here by %v", userID)

}

func (mState *MatchState) ProcessAttack(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	logger.Info("Player move made here by %v", userID)

	//Claculate Attack State.
	mState.lastAttackState = *mState.currentAttackState
	mState.currentAttackState = mState.NewAttackStateObject(ctx, logger, db, nk, dispatcher, tick, state)
}
