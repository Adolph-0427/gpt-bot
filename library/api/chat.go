package api

import (
	"context"
	"encoding/json"
)

const (
	chatUrl    string = "https://api.openai.com/v1/chat/completions"
	role       string = "user"
	apiKey     string = ""
	modelGPT35 string = "gpt-3.5-turbo"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type messageArray []message

type ChatResult struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Usage   map[string]interface{} `json:"usage"`
	Choices []Choice               `json:"choices"`
}

type Choice struct {
	Message      message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}

func (a *Api) Chat(ctx context.Context, content string) (*ChatResult, error) {
	header := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + apiKey},
	}
	messageSlice := make(messageArray, 0)
	mes := message{
		Role:    role,
		Content: content,
	}
	message := append(messageSlice, mes)
	// 构建请求数据
	data := map[string]interface{}{
		"model":    modelGPT35,
		"messages": message,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := request(methodPost, chatUrl, header, payload)
	if err != nil {
		return nil, err
	}
	chatRes := &ChatResult{}
	err = json.Unmarshal(resp, chatRes)
	if err != nil {
		return nil, err
	}
	return chatRes, nil
}
