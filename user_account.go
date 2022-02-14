package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func GetUserIDFromContext(ctx context.Context, logger runtime.Logger) string {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	if !ok {
		logger.Error("UserID NOT FOUND IN CONTEXT.")
	}

	return userId
}

func InitializeUser_Email(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error {

	if out.Created {
		// Only run this logic if the account that has authenticated is new.
		userID := GetUserIDFromContext(ctx, logger)

		err := HandleNewUserAccountCreated(userID, ctx, logger, db, nk)
		return err
	}

	return nil
}

func InitializeUser_Device(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateDeviceRequest) error {

	if out.Created {
		// Only run this logic if the account that has authenticated is new.
		userID := GetUserIDFromContext(ctx, logger)

		err := HandleNewUserAccountCreated(userID, ctx, logger, db, nk)
		return err
	}

	return nil
}

func HandleNewUserAccountCreated(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {

	metadata := Metadata{
		accountState: ACCOUNT_STATE_INCOMPLETE,
		skinID:       1,
	}

	account, err := nk.AccountGetId(ctx, userID)
	if err != nil {
		PrintError(logger, err, "Error fetching account")
		return errAccountFetch
	}

	userName := ""

	if len(account.GetEmail()) > 0 {
		userName = account.GetEmail()
	}

	username := userName
	displayName := username
	metadataMap := metadata.ToMap()
	timezone := ""
	location := ""
	langTag := ""
	avatarUrl := ""

	if err := nk.AccountUpdateId(ctx, userID, username, metadataMap, displayName, timezone, location, langTag, avatarUrl); err != nil {
		// Handle error.
		PrintError(logger, err, "Error updating account")
		return err
	}

	logger.Info("Account updated successfully!")
	return nil
}

func SetupANewUserAccount_RPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	userID := GetUserIDFromContext(ctx, logger)

	payloadMap := GetDecodedObject(payload, logger)

	if !payloadMap.ContainsAll("username", "skin_id") {
		logger.Error("Invalid data provided")
		return Response{success: false, message: errInvalidData.Error()}.GetEncodedObject(), errInvalidData
	}

	newUsername := payloadMap["username"].(string)
	skinID := int(payloadMap["skin_id"].(float64))

	account, err := nk.AccountGetId(ctx, userID)
	if err != nil {
		PrintError(logger, err, "Error fetching account")
		return Response{success: false, message: errAccountFetch.Error()}.GetEncodedObject(), errAccountFetch
	}

	metadata := Decoders.ParseMetadata(account.User.Metadata)

	metadata.skinID = skinID
	metadata.accountState = ACCOUNT_STATE_COMPLETE

	username := newUsername
	displayName := username
	metadataMap := metadata.ToMap()
	timezone := ""
	location := ""
	langTag := ""
	avatarUrl := ""

	if err := nk.AccountUpdateId(ctx, userID, username, metadataMap, displayName, timezone, location, langTag, avatarUrl); err != nil {
		// Handle error.
		PrintError(logger, err, "Error updating account")
		return "", err
	}

	logger.Info("Account updated successfully!")

	return Response{success: true, message: "Account updated successfully!"}.GetEncodedObject(), nil
}

// func SetupCollectionForNewUser(userID string, ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) error {

// 	inventoryCollectionID := "INVENTORY"

// 	storage := map[string]interface{}{
// 		"inv": GetNewUserInventory(),
// 	}

// 	//encodedStorageObject, marshalError := json.Marshal(storage)
// 	//if marshalError != nil {
// 	//	logger.Error("Error Marshalling StorageObject : %v, Error : %v", marshalError)
// 	//}

// 	objectIDs := []*runtime.StorageWrite{

// 		&runtime.StorageWrite{
// 			Collection:      inventoryCollectionID,
// 			Key:             userID,
// 			UserID:          userID,
// 			Value:           string(encodedStorageObject), // Value must be a valid encoded JSON object.
// 			PermissionRead:  2,
// 			PermissionWrite: 1,
// 		},
// 	}

// 	storageObjectAck, err := nk.StorageWrite(ctx, objectIDs)
// 	if err != nil {
// 		logger.WithField("err", err).Error("Storage write error.")
// 	}

// 	logger.Info("StorageObjectAck : %v", storageObjectAck)

// 	return nil
// }
