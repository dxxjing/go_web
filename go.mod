module go_web

go 1.15

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-acme/lego/v3 v3.9.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/nsqio/go-nsq v1.0.8
	github.com/satori/go.uuid v1.2.0
	go.etcd.io/etcd v3.3.25+incompatible
	google.golang.org/grpc v1.32.0
	google.golang.org/grpc/examples v0.0.0-20200910201057-6591123024b3 // indirect
)
//为解决etcd 报错问题
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0