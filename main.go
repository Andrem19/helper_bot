package main

import (
	"log"
	"strconv"

	"github.com/Andrem19/Timer_Job/helpers"
	"github.com/Andrem19/Timer_Job/switcher"
	"github.com/Andrem19/Timer_Job/variables"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	helpers.LoadConfig(".")
	variables.StartWithDb()
	var err error
	variables.Bot, err = tgbotapi.NewBotAPI(variables.Config.TELEGTRAM_BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	variables.Bot.Debug = true

	log.Printf("Authorized on account %s", variables.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := variables.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			chat_id := strconv.Itoa(int(update.Message.Chat.ID))
			answer, err := switcher.Switcher(update.Message.Text, chat_id)
			if err != nil {
				helpers.AddToLog(err.Error())
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			msg.ReplyToMessageID = update.Message.MessageID
			variables.Bot.Send(msg)
		}
	}
}