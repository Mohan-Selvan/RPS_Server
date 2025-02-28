package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Initializing Init..")

	InitializeSignTable(ctx, logger)

	fmt.Println("Registering match_handler")

	//User Init
	// Register as after hook.
	if err := initializer.RegisterAfterAuthenticateEmail(InitializeUser_Email); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	// Register as after hook.
	if err := initializer.RegisterAfterAuthenticateDevice(InitializeUser_Device); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	if err := initializer.RegisterMatch(MATCH_MODULE_FATAL_FOUR, NewMatch_FatalFour); err != nil {
		logger.Error("Unable to register: %v", err)
		fmt.Println("Error registering match_handler!!")
		return err
	}

	if err := initializer.RegisterRpc("create_match_rpc", CreateMatchRPC); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	if err := initializer.RegisterRpc("setup_account_rpc", SetupANewUserAccount_RPC); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	return nil
}
