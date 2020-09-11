package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	//client()
	//client2()
	clientPost()
}

func clientPost(){
	var u = User{
		Name:"jdx",
		Age:18,
	}
	param,_ := json.Marshal(u)
	req,err := http.NewRequest("post","http://localhost:8080/user",strings.NewReader(string(param)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("content-type","application/json")

	client := &http.Client{}
	rsp,err := client.Do(req)
	if err !=nil{
		fmt.Println(err)
	}
	defer rsp.Body.Close()
	res,_ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(res))
}

func client2(){
	req,err := http.NewRequest("get","http://localhost:8080/user",nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("kfuin-cli","1131807740")
	req.Header.Add("content-type","application/json")

	cli := &http.Client{}
	rsp,err := cli.Do(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()

	res,err := ioutil.ReadAll(rsp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(res))
	fmt.Println(rsp.Header)

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