package main

import (
	"context"
	"github.com/go-acme/lego/v3/log"
	"github.com/micro/go-micro"
	"go_web/go-micro/pb"
)

type Server struct {

}

func (g *Server) GetUserInfo(ctx context.Context, req *pb.UserInfoReq, rsp *pb.UserInfoRsp) error{
	name := req.Name
	//从存储中查询该参数的数据
	rsp = &pb.UserInfoRsp{
		Id:1,
		Name:name,
		Hobby:[]string{"song","run"},
	}
	return nil
}

func main(){
	//初始化服务 并设置服务名称
	s := micro.NewService(micro.Name("userserver"))

	s.Init()

	pb.RegisterUserHandler(s.Server(),new(Server))

	log.Fatal(s.Run())
}
