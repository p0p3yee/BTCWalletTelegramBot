package Bot

import (
	bot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	Bot *bot.BotAPI
	Config bot.UpdateConfig
}

func (b *Bot) Create(token string) error {
	newBot, err := bot.NewBotAPI(token)
	if err != nil { return err }
	b.Config = bot.NewUpdate(0)
	b.Config.Timeout = 10

	b.Bot = newBot
	return nil
}

func (b *Bot) NewMsg(txt string, update bot.Update) bot.MessageConfig{
	msg := bot.NewMessage(update.Message.Chat.ID, txt)
	msg.ParseMode = bot.ModeHTML
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}