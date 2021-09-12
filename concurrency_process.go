package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	go process(2,"A") // 並列処理にする go を付与することでprocessが完了する前にmainが完了する。
	go process(2,"B")
	fmt.Println("Finish!")
}

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}
