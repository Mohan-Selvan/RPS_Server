package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

func (mState *MatchState) ChangeGameState(targetGameState GameState, logger runtime.Logger) {

	previousGameState := mState.gameState
	mState.gameState = targetGameState

	logger.Info("Changing game state %v ==> %v", previousGameState, mState.gameState)

	//Handle GameState change here..
	switch mState.gameState {
	case GAME_STATE_LOBBY:
		logger.Info("************-----GAME_STATE_LOBBY-----************")

		mState.AddFutureAction(mState.CreateFutureAction(

			func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

				logger.Info("Checking if the match is Idle..")

				if len(mState.presences) == 0 {
					logger.Info("Ending Idle Match..")
					mState.TerminateMatch(logger)
				}

			}, 60, false, fmt.Sprintf("Idle match check"),
		))

		mState.matchStartTime = Helpers.GetCurrentTime_Unix()

		break
	case GAME_STATE_IN_PROGRESS:
		logger.Info("************-----GAME_STATE_INPROGRESS-----************")
		break
	//case GAME_STATE_PAUSED:
	//	logger.Info("************-----GAME_STATE_PAUSED-----************")
	//	break
	case GAME_STATE_END:
		logger.Info("************-----GAME_STATE_ENDED-----************")
		break
	}
}
