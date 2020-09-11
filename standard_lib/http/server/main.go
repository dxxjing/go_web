package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)
//自定义handler
type myhandler struct {

}

func (m myhandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(rsp, "Hello, %q", html.EscapeString(req.URL.Path))
}
var m myhandler


func main(){
	//srv1()
	//srv2()
	srv3()
}

//简单使用
func srv1(){
	http.HandleFunc("/test",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/test1",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello1, %q", html.EscapeString(r.URL.Path))
	})


	http.Handle("/test2",m)

	log.Fatal(http.ListenAndServe(":8080",nil))
}

//自定义服务
func srv2(){
	mux := http.NewServeMux() //已实现handler接口  也可以直接使用 http.DefaultServeMux
	mux.HandleFunc("/test4",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10*time.Second,
		ReadHeaderTimeout: 10*time.Second,
		WriteTimeout:      10*time.Second,
		IdleTimeout:       10*time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

//自定义handler
func srv3(){
	/*
	mux := http.NewServeMux() // 等价  http.DefaultServeMux
	mux.Handle("/test5",m)
	mux.HandleFunc("/test6",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080",mux))
	*/
	http.DefaultServeMux.Handle("/test5",m)
	http.DefaultServeMux.HandleFunc("/test6",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	//启动监听 必须指定handler
	log.Fatal(http.ListenAndServe(":8080",http.DefaultServeMux))
}


