package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UserChat struct {
	UserId      string    `sql:"type:uuid;default:uuid_generate_v4()"`
	Message     string    `json:"Message"`
	UserName    string    `json:"UserName"`
	CreatedTime time.Time `json:"CreatedTime"`
}

type UserChatDto struct {
	UserName    string    `json:"UserName"`
	Message     string    `json:"Message"`
	CreatedTime time.Time `json:"CreatedTime"`
}

func LoadRecentChat(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	logger.Debug("RPC Called : LoadRecentChat")
	select_query := "SELECT username,message,created_time FROM chat_test ORDER BY created_time DESC LIMIT 100"
	userChat, queryErr := db.QueryContext(ctx, select_query)

	if queryErr != nil {
		logger.Error("(LoadRecentChat) SELECT query error")
	}

	// 	If using db.QueryContext() or db.Query(), you must call row.Close() after you are finished with the database rows data.

	// If using db.QueryRow() or db.QueryRowContext(), you must call either row.Scan or row.Close() after you are finished with the database rows data.
	// row.Close()

	var message string
	var username string
	var created_time time.Time
	var userChatDto []UserChatDto

	count := 0
	for userChat.Next() {
		err := userChat.Scan(&username, &message, &created_time)
		if err != nil {
			logger.Error("(LoadRecentChat) user.Next() Error")
		}
		userChatDto = append(userChatDto, UserChatDto{username, message, created_time})
		count++
	}

	jsonData, marshalError := json.Marshal(userChatDto)

	if marshalError != nil {
		logger.Error(("json.Marshal(userChatDto) Error"))
	}

	return string(jsonData), nil

}

func ChatDelete(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	delete_query := "DELETE FROM chat_test"
	_, queryErr := db.QueryContext(ctx, delete_query)

	if queryErr != nil {
		logger.Error("(ChatDelete) DELETE query error")
	}

	return string("test"), nil
}

func ChatEntered(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	UserId := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	var reqMap map[string]interface{}
	jsonErr := json.Unmarshal([]byte(payload), &reqMap)
	if jsonErr != nil {
		logger.Error("json UnMarshal payload (Chat Entered) Error")
	}

	Message := reqMap["Message"].(string)
	UserName := reqMap["UserName"].(string)

	loc, _ := time.LoadLocation("Asia/Seoul")

	CreateTime := time.Now()
	CreateTime = CreateTime.In(loc)

	logger.Debug(CreateTime.String())

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
