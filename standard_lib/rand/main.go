package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//random()
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
	fmt.Println(n)
}

func random(){
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Println(n)
}