package session

//sessiong 接口定义
type Session interface {

	Set(key string,val interface{}) error

	Get(key string) (interface{},error)

	Del(key string) error

	Save() error
}
