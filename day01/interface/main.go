package main

import "fmt"

// Animal interface
type Animal interface {
	Say()
	Run()
}

//Dog 类
type Dog struct {
	Name string
}

//Say say
func (d Dog) Say() {
	fmt.Println(d.Name, " can say")
}

//Run 跑
func (d Dog) Run() {
	fmt.Println(d.Name, " can Run")
}

//Cat 类
type Cat struct {
	Name string
}

//Say 说
func (c Cat) Say() {
	fmt.Println(c.Name, " can say")
}

//Run 跑
func (c Cat) Run() {
	fmt.Println(c.Name, " can Run")
}

//Say ...
func Say(s Animal) {
	s.Say()
}

func main() {
	var (
		d Dog
		c Cat
	)
	d = Dog{
		Name: "阿黄",
	}
	c = Cat{
		Name: "小咪",
	}
	Say(d)
	Say(c)
}
