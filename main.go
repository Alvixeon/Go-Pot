package main

import (
	"GO-POT/http"
	"GO-POT/ssh"
)

func main() {
	go http.StartServer()
	go ssh.Start()

	block := make(chan struct{})
	<-block
}
