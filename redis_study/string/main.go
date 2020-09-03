package main

import (
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func main(){

}

func op1(){
	rdb.Set("key1","jdx",time.Second*10).Result()

}
