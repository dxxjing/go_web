package main

import "github.com/go_web/gin/session"

func main(){

	s := session.NewMemorySession("1")
	s.Set("jdx","xxx")
}
