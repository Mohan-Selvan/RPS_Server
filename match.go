package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

	futureActions *FutureActions

	tick                  int
	turnCounter           int
	timerPerTurnInSeconds int

	currentTurn  int
	previousTurn int

	matchStartTime int64

	currentAttackState *AttackState
	lastAttackState    AttackState
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

		futureActions: &FutureActions{},

		timerPerTurnInSeconds: TIME_PER_TURN,
		turnCounter:           1,
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

	if len(presences) == MAX_NUMBER_OF_PLAYERS {
		InitializeMatch(ctx, logger, db, nk, dispatcher, tick, state)
	}

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

	//Checking if Match is set to TERMINATE
	if mState.gameState == GAME_STATE_END {
		logger.Info("Ending game : ")
		//Handle Game end here.
		//For example, Storing match data to database for future reference.

		return nil
	}

	mState.ProcessMatchTimers(ctx, logger, db, nk, dispatcher, tick, state, messages)

	mState.ProcessFutureActions(ctx, logger, db, nk, dispatcher, tick, state, messages)

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

func (m *MatchState) TerminateMatch(logger runtime.Logger) {
	m.ChangeGameState(GAME_STATE_END, logger)
	fmt.Println("End match set..")
}

func (m *MatchState) AddFutureAction(f *FutureAction) {
	m.futureActions.Enqueue(f)
}

func (mState *MatchState) NewAttackStateObject(attacker *Player, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) *AttackState {

	logger.Info("Initializing New AttackState")

	attackState := AttackState{
		time_pending: mState.timerPerTurnInSeconds,

		attacker:     attacker,
		attackerSign: SIGN_ROCK,
	}

	return &attackState
}

func (mState *MatchState) DispatchMessage(opcode int64, encodedMessage string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}) {

	logger.Info("Dispatching message : %s", encodedMessage)

	err := dispatcher.BroadcastMessage(opcode, []byte(encodedMessage), nil, nil, true)
	if err != nil {
		PrintError(logger, err, "Error dispatching message")
	}
}

func (m *MatchState) ProcessMatchTimers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	m.tick += 1

	// if m.gameState == GAME_STATE_IN_PROGRESS {
	// 	m.tickCounter -= 1
	// }
}

func (m *MatchState) ProcessFutureActions(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

	for i := 0; i < len(*(m.futureActions)); i++ {
		action := (*m.futureActions)[i]

		if action.followGameState && m.gameState != GAME_STATE_IN_PROGRESS {
			continue
		}

		action.timerInTicks -= 1
		//logger.Info("Future Action update : (%v) countdown : (%v)", action.name, action.timerInTicks)

		if action.timerInTicks <= 0 {
			//Processing action if timer < 0;
			logger.Info("Future Action update Executing : (%v) ", action.name)
			action.action(ctx, logger, db, nk, dispatcher, tick, state, messages)
			m.futureActions.Dequeue()

			//Modifying iterator since collection is modified during iteration.
			i--
		}
	}
}

func (mState *MatchState) GetPlayerOfCurrentTurn() *Player {

	if val, ok := mState.players[mState.turnCounter]; ok {
		return val
	}

	fmt.Println("ERROR : Invalid turn player requested")
	return nil
}

func (m *MatchState) ChangeTurn() {

	//TODO :: No of players to be decided dynamically.
	var targetTurn = ((m.currentTurn) % (2)) + 1

	m.previousTurn = m.currentTurn
	m.currentTurn = targetTurn

	OnChangeTurn(m.currentTurn)
}

func OnChangeTurn(currentTurn int) {
	fmt.Println("Turn changed : ", currentTurn)
}
