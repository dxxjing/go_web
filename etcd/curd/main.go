package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
//运行etcd 报错  undefined: balancer.PickOptions 在go.mod require上面添加如下
//replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

var cli *clientv3.Client

func NewClientv3() (err error){
	cli,err = clientv3.New(clientv3.Config{
		Endpoints:[]string{"192.168.56.101:2379","192.168.56.101:22379","192.168.56.101:32379"},
		DialTimeout:time.Second*5,
	})
	return
}

func main(){
	if err := NewClientv3();err != nil{
		fmt.Println(err)
	}
	fmt.Println("etcd connect success")
	defer cli.Close()
	/*
	put("/server/node1","192.168.56.101:2379")
	put("/server/node2","192.168.56.101:22379")
	put("/server/node3","192.168.56.101:32379")
	get("/server/")
	 */

	/*
	操作客户端
	./etcdctl --endpoints=192.168.56.101:2379 put user xiaoming
	watch("user")
	 */

	//lease("user","jdx")
	keepalive("user","jdx")
}


func put(key,val string) {
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	rsp,err := cli.Put(ctx,key,val)
	cancel()
	if err != nil {
		fmt.Println("etcd put err:",err)
		return
	}
	fmt.Printf("put:key=%s,val=%s,version=%d\n",rsp.PrevKv.Key,rsp.PrevKv.Value,rsp.PrevKv.Version)
}

func get(key string){
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	//clientv3.WithPrefix() 取前缀为key的所有key
	rsp,err := cli.Get(ctx,key,clientv3.WithPrefix())
	cancel()
	if err != nil{
		fmt.Println("etcd get err:",err)
		return
	}
	fmt.Println(rsp.Kvs)
}


//阻塞等待 监听key变化
//./etcdctl --endpoints=192.168.56.101:2379 put user xiaoming
func watch(key string){

	watchChan := cli.Watch(context.Background(),key,clientv3.WithPrefix())

	for val := range watchChan {
		for _,row := range val.Events{
			fmt.Printf("type:%s,key:%s,val:%s\n",row.Type,row.Kv.Key,row.Kv.Value)
		}
	}
}

func lease(key,val string){
	//创建租约
	leaseRsp,err := cli.Grant(context.TODO(),2*60)
	if err != nil{
		fmt.Println("ectd create lease err:",err)
		return
	}
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	_,err = cli.Put(ctx,key,val,clientv3.WithLease(leaseRsp.ID))
	cancel()
	if err != nil{
		fmt.Println("etcd put err:",err)
		return
	}
	fmt.Printf("%s 2min后将被删除",key)
}

func keepalive(key,val string){
	leaseRsp,err := cli.Grant(context.TODO(),5)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*5)
	if _,err := cli.Put(ctx,key,val,clientv3.WithLease(leaseRsp.ID));err != nil{
		fmt.Println(err)
		return
	}
	cancel()
	fmt.Println("etcd put success")

	ch,err := cli.KeepAlive(context.TODO(),leaseRsp.ID)
	if err != nil{
		fmt.Println(err)
		return
	}
	for{
		rsp := <-ch
		fmt.Println("续租成功:",rsp.TTL)
	}

}