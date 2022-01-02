package main

import "github.com/heroiclabs/nakama-common/runtime"

func (mState *MatchState) ChangeGameState(targetGameState GameState, logger runtime.Logger) {

	previousGameState := mState.gameState
	mState.gameState = targetGameState

	logger.Info("Changing game state %v ==> %v", previousGameState, mState.gameState)

	//Handle GameState change here..
}
