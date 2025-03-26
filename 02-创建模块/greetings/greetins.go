package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	if name == "" {
		return "", errors.New("empty name")
	}

	// 如果接收到名称，则返回一个嵌入名称的值
	// 在问候消息中.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init 为函数中使用的变量设置初始值.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 返回一组问候消息中的一个。返回的
// 消息是随机选择的.
func randomFormat() string {
	// 消息格式的切片.
	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

	// 通过为格式切片指定随机索引
    // 来返回随机选择的消息格式.
	return formats[rand.Intn(len(formats))]
}