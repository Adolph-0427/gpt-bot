package main

import (
	"context"
	"gpt-bot/bot"
	"log"
)

func main() {
	bot := bot.NewBot()
	answer, err := bot.Chat(context.Background(), "写一段golang采集cpu的代码")
	if err != nil {
		log.Println(err)
	}
	log.Println(answer)
}
