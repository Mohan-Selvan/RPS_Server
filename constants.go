package main

var (
	TICK_RATE int = 2
)

var (
	MAX_NUMBER_OF_PLAYERS int = 2
	TIME_PER_TURN         int = 20
)

var (
	MATCH_MODULE_FATAL_FOUR string = "module_fatal_four"
)

type Sign int

var (
	SIGN_ROCK    Sign = 1
	SIGN_PAPER   Sign = 2
	SIGN_SCISSOR Sign = 3
)
