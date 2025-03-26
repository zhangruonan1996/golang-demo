package main

import (
	"fmt"
	"log"

	"zhangruonan.org/greetings"
)

func main() {
	// 设置预定义Logger的属性，包括
    // 日志条目前缀和禁用打印的标志
    //  时间、源文件和行号.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 一个名字切片.
	names := []string{"章若楠", "白敬亭", "杨洋"}

	// 请求问候消息.
	// message, err := greetings.Hello("章若楠")
	messages, err := greetings.Hellos(names)
	// 如果返回错误，则将其打印到控制台并
    // 退出程序.
    if err != nil {
        log.Fatal(err)
    }

	// 如果没有返回错误，则打印返回的消息
    // 到控制台.
	fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}