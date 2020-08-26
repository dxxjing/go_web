package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)
/*
//解决粘包
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}

 */
func work(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf = make([]byte,128)
	//var buf [128]byte //如果buf定义为数组
	for{
		n,err := reader.Read(buf)//读取数据
		fmt.Printf("read %d bytes\n",n)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read err",err)
			return
		}
		fmt.Println("recv data:",string(buf[:n]))
		conn.Write(buf[:n])//发送数据
	}

	//fmt.Println("recv all data:",string(allBuf))
	//conn.Write([]byte(allBuf))//发送数据
}

func main(){
	listen,err := net.Listen("tcp","0.0.0.0:20000")//监听
	if err != nil {
		fmt.Println("listen tcp err")
		return
	}

	for{
		conn,err := listen.Accept()//接收请求
		if err != nil {
			fmt.Println("accept err")
			continue
		}

		go work(conn)
	}
}
