package api

import (
	Match "Boomkit/nakama/src/Match"
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
	// logger.Info("match ID: ", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID))

	subject := payload
	// logger.Info(subject)

	//content는 무엇이지
	content := map[string]interface{}{}
	// map을 notificationSend로 보냈을때 유니티에서 json 형식으로 받게된다
	// ex: {"exp":500,"item":"집행검","reward_coins":1000}

	// content := map[string]interface{}{
	// 	"reward_coins": 1000,
	// 	"item":         "집행검",
	// 	"exp":          500,
	// }

	code := 101
	//Whether to record this in the database for later listing. 데이터베이스에 저장할것인지 안할것인지 notification table
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

	//RequestBody의 data를 map으로 변환
	// map["Message"] = "testMessage"
	// map["MatchId"] = "testMatchId" 형식
	jsonErr := json.Unmarshal([]byte(payload), &reqMap)

	if jsonErr != nil {
		logger.Error("SendInGameNotiToOneMatch Unmarshal Error")
	}

	//reqMap에서 추출 후
	matchId := reqMap["MatchId"].(string)
	subject := reqMap["Message"].(string)

	//matchId를 사용하여 현재 매치내에 유저들 추출
	// matchState := GetMatchState(matchId)
	// presences := GetPresencesByState(matchState)

	presences := Match.GetPresencesByMathId(matchId)

	code := 101
	content := map[string]interface{}{}

	//Whether to record this in the database for later listing.
	persistent := false
	//매치 접속자 추출 비효율적인것같음///

	/////////////////////////////////////

	inGameNoti := InGameNoti{subject}
	jsonData, _ := json.Marshal(inGameNoti)

	for _, p := range presences {
		logger.Warn("here here here here")
		logger.Debug(p.GetUserId())
		nk.NotificationSend(ctx, p.GetUserId(), string(jsonData), content, code, "", persistent)
	}

	return "", nil
}
