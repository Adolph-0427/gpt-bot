package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	APIKey = ""
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageArray []Message

func main() {
	messageSlice := make(MessageArray, 0)
	mes := Message{
		Role:    "user",
		Content: "吃饭了吗",
	}
	message := append(messageSlice, mes)
	// 构建请求数据
	data := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": message,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 设置请求头和URL
	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+APIKey)
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
