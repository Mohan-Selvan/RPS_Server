package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

func CreateAMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	//Parsing MatchParams
	params := make(map[string]interface{})
	if err := json.Unmarshal([]byte(payload), &params); err != nil {
		PrintError(logger, err, "Error decording payload")
		return "", err
	}

	if matchId, err := nk.MatchCreate(ctx, MATCH_MODULE_FATAL_FOUR, params); err != nil {

		PrintError(logger, err, "Error creating match")
		return Response{success: false, message: "Error creating match"}.GetEncodedObject(), errUnknownError

	} else {

		return Response{success: true, message: matchId}.GetEncodedObject(), nil
	}
}

type MatchState struct {
	presences map[string]runtime.Presence

	players   PlayersMap
	gameState GameState
}

type Match struct {
	moduleName string
}

func NewMatch_FatalFour(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error) {
	logger.Info("Created a new match!")
	return &Match{moduleName: MATCH_MODULE_FATAL_FOUR}, nil
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {

	state := &MatchState{
		presences: make(map[string]runtime.Presence),
		gameState: GAME_STATE_LOBBY,
		players:   make(PlayersMap),
	}

	tickRate := 1
	label := m.moduleName
	return state, tickRate, label
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	mState, _ := state.(*MatchState)
	//TODO:: Handle Player validations here.
	acceptUser := len(mState.presences) < 4

	return state, acceptUser, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	//Adding Player to presence.
	for _, p := range presences {
		mState.presences[p.GetUserId()] = p
	}

	//if len(presences) == 4 {
	InitializeMatch(ctx, logger, db, nk, dispatcher, tick, state)
	//}

	return mState
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	//Removing player from presences.
	for _, p := range presences {
		delete(mState.presences, p.GetUserId())
	}
	return mState
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	mState, _ := state.(*MatchState)

	//for _, presence := range mState.presences {
	//	logger.Info("Presence %v named %v", presence.GetUserId(), presence.GetUsername())
	//}

	switch mState.gameState {

	case GAME_STATE_LOBBY:
		HandleGameState_LOBBY(ctx, logger, db, nk, dispatcher, tick, state, messages)
		break

	case GAME_STATE_IN_PROGRESS:
		HandleGameState_INPROGRESS(ctx, logger, db, nk, dispatcher, tick, state, messages)
		break
	}
	return mState
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	message := "Server shutting down in " + strconv.Itoa(graceSeconds) + " seconds."
	reliable := true

	//Sending shutdown message.
	dispatcher.BroadcastMessage(2, []byte(message), nil, nil, reliable)
	return state
}

func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, "signal received: " + data
}
