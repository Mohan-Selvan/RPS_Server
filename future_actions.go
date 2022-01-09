package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

type FutureAction struct {
	action          func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData)
	timerInTicks    int
	followGameState bool
}

type FutureActions []FutureAction

func (m *MatchState) CreateFutureAction(action func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData),
	timeInSeconds int, followGameState bool) FutureAction {
	return FutureAction{
		action:          action,
		timerInTicks:    SecondsToTicks(timeInSeconds),
		followGameState: followGameState,
	}
}

func (o *FutureActions) Enqueue(action FutureAction) {
	*o = append(([]FutureAction(*o)), action)
}

func (slice *FutureActions) Dequeue() FutureAction {
	if len([]FutureAction(*slice)) <= 0 {
		fmt.Println("No Element Present")
	}

	var s int = 0
	d := FutureAction((*slice)[0])
	*slice = append((*slice)[:s], (*slice)[s+1:]...)

	return d
}

func SecondsToTicks(seconds int) int {
	return seconds * TICK_RATE
}
