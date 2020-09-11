package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Res struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/user",func(rsp http.ResponseWriter,req *http.Request){
		//读取client header
		fmt.Println(req.Header)

		//读取post json数据
		content,_ := ioutil.ReadAll(req.Body)
		fmt.Println(string(content))
		var u User
		json.Unmarshal(content,&u)
		fmt.Println(u)

		//设置reponse header
		rsp.Header().Set("kfuin-server","1131807740")
		rsp.Header().Set("Content-Type","application/json; charset=utf-8")

		msg,_ := json.Marshal(Res{Code:0,Msg:"ok"})
		//设置header 状态码
		rsp.WriteHeader(200)
		rsp.Write(msg)
	})
	log.Fatal(http.ListenAndServe(":8080",mux))
}
