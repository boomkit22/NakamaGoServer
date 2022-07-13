package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	initStart := time.Now()

	// err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)
	// if err != nil {
	// 	return err
	// }

	err := initializer.RegisterRpc("Chat_Entered", ChatEntered)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Chat_Deleted", ChatDelete)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Load_Recent_Chat", LoadRecentChat)
	if err != nil {
		return err
	}

	err = initializer.RegisterMatch("Init_Test_Match", InitTestMatch)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Match_Create", MatchCreate)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Send_InGame_Noti", SendInGameNoti)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Send_InGame_Noti_To_One_Match", SendInGameNotiToOneMatch)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Get_Match_List", GetMatchList)
	if err != nil {
		return err
	}

	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())
	return nil
}
