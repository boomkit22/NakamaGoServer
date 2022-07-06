package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UserChat struct {
	UserId      string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Message     string    `json:"Message"`
	UserName    string    `json:"UserName"`
	CreatedTime time.Time `json:"CreatedTime"`
}

//testversion
// func ChatEntered(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
// 	UserId := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

// 	Message := "testMessage"
// 	UserName := "TestUserName"
// 	CreateTime := time.Now()

// 	UserChatData := UserChat{}

// 	UserChatData.UserId = UserId
// 	UserChatData.UserName = UserName
// 	UserChatData.CreatedTime = CreateTime
// 	UserChatData.Message = Message

// 	jsonData, err := json.Marshal(UserChatData)

// 	if err != nil {
// 		logger.Error("json Marshal (UserChatData) Error")
// 	}

// 	//todo db에 삽입하기

// 	return string(jsonData), nil
// }

func ChatEntered(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	UserId := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	var reqMap map[string]interface{}
	jsonErr := json.Unmarshal([]byte(payload), &reqMap)
	if jsonErr != nil {
		logger.Error("json UnMarshal payload (Chat Entered) Error")
	}

	Message := reqMap["Message"].(string)
	UserName := reqMap["UserName"].(string)
	CreateTime := time.Now()

	UserChatData := UserChat{}
	UserChatData.UserId = UserId
	UserChatData.UserName = UserName
	UserChatData.CreatedTime = CreateTime
	UserChatData.Message = Message

	insert_query := "INSERT INTO chat_test(userid,created_time,username,message) VALUES($1, $2, $3, $4)"
	_, queryErr := db.ExecContext(ctx, insert_query, UserId, CreateTime, UserName, Message)

	if queryErr != nil {
		logger.Error("Insert Query (UserChatData) Error")
	}
	jsonData, err := json.Marshal(UserChatData)

	if err != nil {
		logger.Error("json Marshal (UserChatData) Error")
	}

	//todo db에 삽입하기

	return string(jsonData), nil
}
