package main

import "fmt"

func main() {
	//channelの作成
	messages := make(chan string)

	//作成されたchannelに値(str)を送信
	go func() { messages <- "str" }()

	//channelから値を受信
	msg := <-messages
	fmt.Println(msg) //=> str
}
