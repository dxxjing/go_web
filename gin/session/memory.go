package session

import (
	"errors"
	"sync"
)
//实现Session接口
type MemorySession struct {
	sessionID string

	data map[string]interface{}

	sync.RWMutex
}

//MemorySession实现Session接口 所以可以返回接口类型
func NewMemorySession(id string) Session {
	return &MemorySession{
		sessionID: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (s *MemorySession) Set(key string,val interface{}) (err error){
	s.Lock()
	defer s.Unlock()

	s.data[key] = val

	return
}

func (s *MemorySession) Get(key string) (val interface{},err error){
	s.Lock()
	defer s.Unlock()
	val,ok := s.data[key]
	if !ok {
		err := errors.New("key not existed")
		return val,err
	}
	return
}

func (s *MemorySession) Del(key string) (err error){
	s.Lock()
	defer s.Unlock()

	delete(s.data,key)
	return
}

func (s *MemorySession) Save() (err error){
	return
}
