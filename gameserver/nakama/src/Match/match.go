package match

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type Test_Match struct {
}

type MatchState struct {
	presences map[string]runtime.Presence
}

type MatchInfoList struct {
	MatchInfoList []MatchInfo `json:"MatchInfoList"` // 슬라이스 변수 vector 같은거 가변길이
}

type MatchInfo struct {
	MatchId   string `json:"MatchId"`
	MatchSize int    `json:"MatchSize"`
}

// type Match struct {
// 	matchId   string
// 	testMatch *Test_Match
// }

var _matchMap = map[string]*MatchState{}

// json:"UserName"`
// 	Message     string    `json:"Message"`
func GetMatchList(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	limit := 10
	authoritative := true
	label := ""
	minSize := 0
	maxSize := 10
	query := "" //Additional query parameters to shortlist matches

	matches, err := nk.MatchList(ctx, limit, authoritative, label, &minSize, &maxSize, query)

	matchInfoList := MatchInfoList{}

	if err != nil {
		logger.WithField("err", err).Error("Match list error.")
	} else {
		for _, match := range matches {
			logger.Info("Found match with id: %s", match.GetMatchId())
			logger.Info("Found match size : %d", match.GetSize())
			matchInfo := MatchInfo{match.GetMatchId(), int(match.GetSize())}
			matchInfoList.MatchInfoList = append(matchInfoList.MatchInfoList, matchInfo)
		}
	}

	jsonData, marshalError := json.Marshal(matchInfoList)

	if marshalError != nil {
		logger.Error(("json.Marshal(userChatDto) Error"))
	}

	logger.Debug("GetMatchList Finish")
	return string(jsonData), err
}

func MatchCreate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	modulename := "Init_Test_Match"

	matchId, err := nk.MatchCreate(ctx, modulename, nil)

	if err != nil {
		logger.Debug("Create Match Error")
	}

	logger.Debug("Create match Finish")
	return matchId, nil
}

func InitTestMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	// matchId := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)

	// _matchMap[matchId] = &matchState
	return &Test_Match{}, nil

}

//idea Test_Match 전용 함수들
// match_Id에 해당하는 Test_Match를 찾아서 주면됨
func GetMatchState(matchId string) *MatchState {
	return _matchMap[matchId]
}

func GetPresencesByState(matchState *MatchState) map[string]runtime.Presence {
	return matchState.presences
}

func GetPresencesByMathId(matchId string) map[string]runtime.Presence {
	matchState := _matchMap[matchId]
	return matchState.presences
}

func (m *Test_Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		presences: make(map[string]runtime.Presence),
	} // Define custom MatchState in the code as per your game's requirements
	matchId := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
	_matchMap[matchId] = state

	// state.matchId = ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)

	tickRate := 60 // Call MatchLoop() every 1s.
	label := ""    // Custom label that will be used to filter match listings.

	return state, tickRate, label
}

func (m *Test_Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// Custom code to process match join and send updated state to a joining or re-joining user.
	// dispathcer: Ex poses useful functions to the match, and may be used by the server to send messages to the participants of the match.
	// state: Custom match state interface, use this to manage the state of your game. You may choose any structure for this interface depending on your game's needs.
	// presences: A list of presences that have successfully completed the match join process.

	mState, _ := state.(*MatchState)

	for _, p := range presences {
		presences = append(presences, p)
		mState.presences[p.GetUserId()] = p
		logger.Debug(mState.presences[p.GetUserId()].GetUserId())
	}

	return mState
}

func (m *Test_Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	acceptUser := true

	return state, acceptUser, ""
}

func (m *Test_Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	for _, p := range presences {
		delete(mState.presences, p.GetUserId())
	}

	return mState
}

func (m *Test_Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	mState, _ := state.(*MatchState)

	// for _, presence := range mState.presences {
	// 	logger.Info("Presence %v named %v", presence.GetUserId(), presence.GetUsername())
	// }

	for _, message := range messages {
		// logger.Info("Received %v from %v", string(message.GetData()), message.GetUserId())
		dispatcher.BroadcastMessage(message.GetOpCode(), message.GetData(), nil, nil, true)
		// dispatcher.BroadcastMessage(message.GetOpCode(), message.GetData(), mState.presences, message., true)
	}

	return mState
}

func (m *Test_Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, "signal received: " + data
}

func (m *Test_Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	// Custom code to process the termination of match.
	return state
}
