package session

import (
	"fmt"
	"github.com/satori/go.uuid"
	"sync"
)

//实现SessionMgr接口

type MemorySessionMgr struct {
	sync.RWMutex

	sessionMap map[string]Session
}

func NewMemorySessionMgr() *MemorySessionMgr{
	return &MemorySessionMgr{
		sessionMap:make(map[string]Session),
	}
}

func (m *MemorySessionMgr) Init(addr string,options ...interface{}) (err error){
	return
}


func (m *MemorySessionMgr) Create() (s Session,err error){
	m.Lock()
	defer m.Unlock()

	//用uuid作为sessionID
	sessionID := uuid.NewV4().String()

	s = NewMemorySession(sessionID)
	m.sessionMap[sessionID] = s
	return
}

func (m *MemorySessionMgr) Get(sessionID string) (s Session,err error){
	m.Lock()
	defer m.Unlock()

	s,ok := m.sessionMap[sessionID]
	if !ok{
		fmt.Println("session id not existed")
		return
	}
	return
}