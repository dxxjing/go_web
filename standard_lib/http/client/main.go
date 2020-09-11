package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	client()
}

func client2(){
	req,err := http.NewRequest("get","http://localhost:8080/user",nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("kfuin-cli","1131807740")
	req.Header.Add("content-type","application/json")
	req.Header.Del("Accept-Encoding")

	cli := &http.Client{}
	rsp,err := cli.Do(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()
/*
	res,err := ioutil.ReadAll(rsp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
	fmt.Println(rsp.Header)
*/
}


func client(){
	rsp,err := http.Get("http://localhost:8080/user")
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()
	p,err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p))
	header := rsp.Header
	fmt.Println(header)
}