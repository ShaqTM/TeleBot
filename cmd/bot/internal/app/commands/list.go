package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMSG := "Here all the products: \n\n"
	productServiceList := c.productService.List()
	for _, p := range productServiceList {
		outputMSG += p.Title
		outputMSG += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMSG)
	c.bot.Send(msg)

}

func init() {
	registredCommands["list"] = (*Commander).List
}
