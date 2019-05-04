package Config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	RPCUser string `json:"User"`
	RPCHost string `json:"Host"`
	RPCPass string `json:"Pass"`
}

var MyConfig Config

func LoadConfig() error{
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil { return err }
	byteVal, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal(byteVal, &MyConfig); err != nil {
		return err
	}
	return nil
}