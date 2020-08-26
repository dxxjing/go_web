package main

import (
	"fmt"
	"go_web/day01/proto"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
/*
func main(){
	conn,err := net.Dial("tcp","0.0.0.0:20000")
	if err != nil {
		fmt.Println("dial err")
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for{
		input,_ := inputReader.ReadString('\n')
		input = strings.Trim(input,"\r\n")
		if strings.ToUpper(input) == "Q" {
			return
		}
		fmt.Println("send ready:",input,len(input))
		n,err := conn.Write([]byte(input))
		fmt.Println(n)
		if err != nil {
			fmt.Println("write data err")
			return
		}
		buf := [1024]byte{}
		n,err = conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read data err:%v\n",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}*/
