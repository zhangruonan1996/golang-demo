package greetings

import (
	"errors"
	"fmt"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	if name == "" {
		return "", errors.New("empty name")
	}

	// 如果接收到名称，则返回一个嵌入名称的值
	// 在问候消息中.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}