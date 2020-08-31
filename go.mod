module github.com/go_web
//使用 go mod 就无需将项目建立在gopath下

//项目迁移使用go mod步骤：以 go_web 项目为例

//1、项目外 执行 go mod init go_web  会向当前目录下的go_web内添加go mod文件

//2、cd go_web ;  go mod tidy  拉取必须模块 移出不用模块
go 1.15

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.2.0 // indirect
)
