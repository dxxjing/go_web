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
//直接返回接口类型
func NewDog(name string) Animal {
	return Dog{Name:name}
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
//返回接口类型
func NewCat(name string) Animal {
	return Cat{Name:name}
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
	a := NewDog("阿黄")
	a.Say()
	a = NewCat("小咪")
	a.Say()
}
