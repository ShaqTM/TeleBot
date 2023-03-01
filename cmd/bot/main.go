package main

import (
	"log"
	"os"

	"github.com/ShaqTM/telebot/cmd/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, `/help - help
	/list - list products`)
	bot.Send(msg)

}
func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMSG := "Here all the products: \n\n"
	productServiceList := productService.List()
	for _, p := range productServiceList {
		outputMSG += p.Title
		outputMSG += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMSG)
	bot.Send(msg)

}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)

}
