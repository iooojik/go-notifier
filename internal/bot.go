package internal

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)

func SendMessage(goVersion, u, token string, chatIds []string) error {
	template, err := os.ReadFile("templates/msg_template.md")
	if err != nil {
		return err
	}
	newMessage := strings.Replace(strings.Replace(string(template), "${VERSION}", goVersion, -1), "${DOWNLOAD_URL}", u, -1)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	for _, chatId := range chatIds {
		num, err := strconv.ParseInt(chatId, 10, 64)
		if err != nil {
			log.Println(err)
		}
		msgConf := tgbotapi.NewMessage(num, newMessage)
		send, err := bot.Send(msgConf)
		if err != nil {
			log.Println(err)
		}
		log.Println("message sent to", send.Chat.ID, "message id", send.MessageID)
	}
	return nil
}
