package chat

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageArray []Message
