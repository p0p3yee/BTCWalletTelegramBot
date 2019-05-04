package main

import (
	"BTCWalletTelegramBot/Commands"
	"log"
	"BTCWalletTelegramBot/Config"
	"BTCWalletTelegramBot/RPC"
	bot "BTCWalletTelegramBot/Bot"
)

func main(){

	if err := Config.LoadConfig(); err != nil {
		log.Fatal("Error in Loading Config: ", err)
	}else{
		log.Println("Config Loaded.")
	}

	var Bot bot.Bot
	if err := Bot.Create(Config.MyConfig.BotToken); err != nil {
		log.Fatal("Error in creating Telegram Bot: ", err)
	}else{
		log.Println("Telegram Bot Authorized.")
	}

	var rpc RPC.Rpc
	if err := rpc.Create(Config.MyConfig.RPCHost, Config.MyConfig.RPCUser, Config.MyConfig.RPCPass); err != nil {
		log.Fatal("Error in creating RPC Client: ", err)
	}else{
		log.Println("RPC Client Created.")
	}

	commandHandler := Commands.CreateHandler(rpc)

	if updates, err := Bot.Bot.GetUpdatesChan(Bot.Config); err != nil{
		log.Fatal("Error in getting Telegram Bot Updates: ", err)
	}else {
		for each := range updates {
			if !each.Message.Chat.IsPrivate() { continue }
			if !each.Message.IsCommand() { continue }

			cmd := each.Message.Command()

			if !Config.MyConfig.IsOwner(int64(each.Message.From.ID)) && cmd != "start"{
				log.Printf("User: @%s, ID: %d Trying to Execute Command.", each.Message.From.UserName, each.Message.From.ID)
				continue
			}

			var response string

			if commandHandler.IsCommand(cmd) {
				response = commandHandler.Handle(each.Message.From.ID, cmd, each.Message.CommandArguments())
			}else {
				response = "<b>Error</b>: Command Not Exists.\nType /help to see all available commands."
			}

			if len(response) <= 0 { continue }

			if _, err := Bot.Bot.Send(Bot.NewMsg(response, each)); err != nil {
				log.Println("Error sending Message: ", err)
			}
		}
	}
}