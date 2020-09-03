package main

import (
	"context"
	"fmt"
	"go_web/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//仅仅为了绑定方法 实现pb getUserInfo接口
type server struct {

}

func (s *server) GetUserInfo(ctx context.Context,req *pb.UserInfoReq)(rsp *pb.UserInfoRsp,err error){
	name := req.Name
	//从存储中查询该参数的数据
	rsp = &pb.UserInfoRsp{
		Id:1,
		Name:name,
		Hobby:[]string{"song","run"},
	}
	return
}


func main(){

	lis,err := net.Listen("tcp",":8999")
	if err != nil {
		fmt.Println("listen tcp err",err)
		return
	}
	//创建rpc服务
	rpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterUserServer(rpcServer,&server{})
	//在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(rpcServer)
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	if err := rpcServer.Serve(lis);err != nil{
		fmt.Println("failed serve",err)
		return
	}
}
