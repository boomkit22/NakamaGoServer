package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// api body 형식

// type InGameNoti struct {
// 	Message string `json:"Message"`
// }

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
	persistent := true

	err := nk.NotificationSendAll(ctx, subject, content, code, persistent)

	return subject, err
}
