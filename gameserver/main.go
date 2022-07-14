package main

import (
	Api "Boomkit/nakama/src/Api"
	Chat "Boomkit/nakama/src/Chat"
	Match "Boomkit/nakama/src/Match"
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

	err := initializer.RegisterRpc("Chat_Entered", Chat.ChatEntered)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Chat_Deleted", Chat.ChatDelete)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Load_Recent_Chat", Chat.LoadRecentChat)
	if err != nil {
		return err
	}

	err = initializer.RegisterMatch("Init_Test_Match", Match.InitTestMatch)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Match_Create", Match.MatchCreate)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Send_InGame_Noti", Api.SendInGameNoti)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Send_InGame_Noti_To_One_Match", Api.SendInGameNotiToOneMatch)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("Get_Match_List", Match.GetMatchList)
	if err != nil {
		return err
	}

	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())

	return nil
}
