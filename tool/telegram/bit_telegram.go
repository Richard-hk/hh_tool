package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func TelegramSendText(sendText string) {
	token := viper.GetString("telegram.token")
	channelChatId := viper.GetInt64("telegram.channel.id")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	msg := tgbotapi.NewMessage(channelChatId, sendText)
	bot.Send(msg)
}
