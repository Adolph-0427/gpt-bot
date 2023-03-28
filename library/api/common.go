package api

import (
	"bytes"
	"io"
	"net/http"
)

type Api struct {
}

const (
	methodPost string = "POST"
	methodGet  string = "GET"
)

func request(method string, url string, header map[string][]string, data []byte) ([]byte, error) {
	// 设置请求头和URL
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header = header
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 处理响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}
