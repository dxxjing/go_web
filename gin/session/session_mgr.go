package session

//session管理
type SessionMgr interface {
	
	Init(addr string,options ...interface{}) error

	Create() (Session,error)

	Get(sessionID string) (Session,error)
}
