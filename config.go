package BTCWalletTelegramBot

import (
	"encoding/json"
	"os"
)

type Config struct {
	RPCUser string
	RPCHost string
	RPCPass string
}

var MyConfig Config

func LoadConfig() error{
	file, err := os.Open("config.json")
	if err != nil { return err }
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&MyConfig)
	if err != nil { return err }
	err = file.Close()
	if err != nil { return err }
	return nil
}