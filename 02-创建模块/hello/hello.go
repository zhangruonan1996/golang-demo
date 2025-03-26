package main

import (
	"fmt"

	"zhangruonan.org/greetings"
)

func main() {
	// 获取问候信息并打印出来.
	message := greetings.Hello("章若楠")
	fmt.Println(message)
}