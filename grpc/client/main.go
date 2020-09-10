package main

import (
	"context"
	"fmt"
	"go_web/grpc/pb"
	"google.golang.org/grpc"
)

func main(){
	//grpc.WithInsecure() 禁用传输安全性
	conn,err := grpc.Dial(":8999",grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc dial err",err)
		return
	}
	defer conn.Close()

	cli := pb.NewUserClient(conn)

	name := "jdx"
	rsp,err := cli.GetUserInfo(context.Background(),&pb.UserInfoReq{Name:name})
	if err != nil {
		fmt.Println("getuserinfo err")
		return
	}

	fmt.Println(rsp.GetId(),rsp.GetName(),rsp.GetHobby())
}
