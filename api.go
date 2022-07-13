package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

// api body 형식

type InGameNoti struct {
	Message string `json:"Message"`
}

func SendInGameNoti(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// var reqMap map[string]string
	// jsonErr := json.Unmarshal([]byte(payload), &reqMap)
	// if jsonErr != nil {
	// 	logger.Error("json UnMarshal payload (Chat Entered) Error")
	// }
	logger.Info("match ID: ", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID))

	subject := payload

	logger.Info(subject)

	content := map[string]interface{}{}
	code := 101
	persistent := false

	err := nk.NotificationSendAll(ctx, subject, content, code, persistent)

	return subject, err
}

type InGameNotiToOneMatch struct {
	MatchId string `json:"MatchId"`
	Message string `json:"Message"`
}

func SendInGameNotiToOneMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	var reqMap map[string]interface{}
	jsonErr := json.Unmarshal([]byte(payload), &reqMap)

	if jsonErr != nil {
		logger.Error("SendInGameNotiToOneMatch Unmarshal Error")
	}
	matchId := reqMap["MatchId"].(string)
	subject := reqMap["Message"].(string)
	code := 101
	content := map[string]interface{}{}
	persistent := false

	//매치 접속자 추출 비효율적인것같음
	matchState := GetMatchState(matchId)
	presences := GetPresences(matchState)

	inGameNoti := InGameNoti{subject}
	jsonData, _ := json.Marshal(inGameNoti)

	for _, p := range presences {
		logger.Warn("here here here here")
		logger.Debug(p.GetUserId())
		nk.NotificationSend(ctx, p.GetUserId(), string(jsonData), content, code, "", persistent)
	}

	return "", nil
}
