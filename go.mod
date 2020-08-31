module github.com/go_web
//使用 go mod 就无需将项目建立在gopath下

//项目迁移使用go mod步骤：以 go_web 项目为例

//1、项目外 执行 go mod init go_web  会向当前目录下的go_web内添加go mod文件

//2、cd go_web ;  go mod tidy  拉取必须模块 移出不用模块
go 1.15

require (
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nsqio/go-nsq v1.0.8 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20200828194041-157a740278f4 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
