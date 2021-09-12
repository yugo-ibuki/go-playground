package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	//bool型のchannelのdoneを生成する。
	done := make(chan bool)

	//生成したdoneを関数helloに渡す
	go hello(done)

	<-done
	//main
	fmt.Println("main function")

}
