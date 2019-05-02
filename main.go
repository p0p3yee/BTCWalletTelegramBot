package BTCWalletTelegramBot

import (
	"log"
	"github.com/btcsuite/btcd/rpcclient"
)

func main(){

	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}else{
		log.Println("Config Loaded.")
	}


	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host: MyConfig.RPCHost,
		User: MyConfig.RPCUser,
		Pass: MyConfig.RPCPass,
		HTTPPostMode: true,
		DisableTLS: true,
	}, nil)

	if err != nil { log.Fatal(err) }

	c, err := client.GetBlockCount()
	if err != nil { log.Fatal(err) }

	log.Println(c)
	client.Shutdown()
}