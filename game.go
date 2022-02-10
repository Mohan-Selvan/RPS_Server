package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

func HandleGameState_INPROGRESS(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	mState, _ := state.(*MatchState)

	for _, message := range messages {

		data := string(message.GetData())
		//senderID := message.GetUserId()

		switch message.GetOpCode() {
		case c2s_PlayerMove:

			//Validation.
			//if mState.GetPlayerOfCurrentTurn().userID != senderID {
			//	logger.Info("Player move message received at Invalid turn.")
			//	continue
			//}

			sign := SIGNS.GetSign(StringToInt(data))

			mState.HandlePlayerMove(message.GetUserId(), sign, ctx, logger, db, nk, dispatcher, tick, state, messages)
			break

		default:
			logger.Info("Invalid Opcode received : %d", message.GetOpCode())
			break
		}
	}
}

func (mState *MatchState) HandlePlayerMove(userID string, sign Sign, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	//Update Player selection here.
	player := mState.players.GetPlayer(userID)
	if player == nil {
		logger.Error("Error :: Player not found!")
		return
	}

	if player.IsDead() {
		logger.Info("Invalid move : Player dead")
		return
	}

	if player.hasSelectedMoveForThisTurn {
		logger.Info("Move has already been selected")
		return
	}

	player.selectedSign = sign
	player.hasSelectedMoveForThisTurn = true

	if mState.players.DidEveryoneSelectMove() {
		mState.ExecuteTurn(ctx, logger, db, nk, dispatcher, tick, state, messages)
	}
}

func (mState *MatchState) ExecuteTurn(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	//Update Player selection here.
	logger.Info("Executing Moves (%v)", mState.moveCounter)

	//target := mState.players.GetPlayer(userID).target

	attackState := mState.currentAttackState

	for _, p := range mState.players.GetNonDeadPlayers() {
		pSign := p.selectedSign

		for _, o := range mState.players.GetNonDeadPlayersExcept(p.userID) {
			oSign := o.selectedSign

			resultantDamage := pSign.damage - oSign.defense

			if resultantDamage > 0 {
				//When damage is still valid.
				o.ModifyHealth(resultantDamage)

			} else if resultantDamage < 0 {
				//When deflected.
				//Do nothing.
			}
		}
	}

	mState.CheckMatchEnd(ctx, logger, db, nk, dispatcher, tick, state, messages)

	mState.DispatchMessage(s2c_PlayerMove, mState.currentAttackState.GetEncodedObject(), ctx, logger, db, nk, dispatcher, tick, state)

	mState.moveCounter++

	if len(attackState.matchEndMessage) > 0 {
		mState.TerminateMatch(logger)
	}

	mState.lastAttackState = *mState.currentAttackState
	mState.currentAttackState = mState.NewAttackStateObject(ctx, logger, db, nk, dispatcher, tick, state)

	//Process AI Move.
	//mState.ProcessAIMove(ctx, logger, db, nk, dispatcher, tick, state)
}

func (mState *MatchState) ProcessAIMove(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) {

	mState.AddFutureAction(mState.CreateFutureAction(
		func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
			logger.Info("Creating future action (AIMove)")

			for _, val := range mState.players.GetNonDead_AI_Players() {
				randomSign := SIGNS.GetRandomSign()
				logger.Info("AI Move (%v) update : Attack (%v).", val.userID, randomSign)
				mState.HandlePlayerMove(val.userID, randomSign, ctx, logger, db, nk, dispatcher, tick, state, messages)
			}
		},
		2,    //Delay
		true, //Follow game timer.
		fmt.Sprintf("AI Move (Move Counter : %v)", mState.moveCounter), //Action name.
	))
}
