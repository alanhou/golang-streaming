package session

import (
	"golang-streaming/video_server/api/dbops"
	"golang-streaming/video_server/api/defs"
	"golang-streaming/video_server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func noInMilli() int64  {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := noInMilli()
	ttl := ct + 30 * 60 * 1000 // Server side session valid time: 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}