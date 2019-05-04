package main

import (
	"log"
	"github.com/btcsuite/btcd/rpcclient"
	"BTCWalletTelegramBot/Config"
)

func main(){

	if err := Config.LoadConfig(); err != nil {
		log.Fatal(err)
	}else{
		log.Println("Config Loaded.")
	}

	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host: Config.MyConfig.RPCHost,
		User: Config.MyConfig.RPCUser,
		Pass: Config.MyConfig.RPCPass,
		HTTPPostMode: true,
		DisableTLS: true,
	}, nil)

	if err != nil { log.Fatal(err) }

	c, err := client.GetBlockCount()
	if err != nil { log.Fatal(err) }

	log.Println(c)
	client.Shutdown()
}