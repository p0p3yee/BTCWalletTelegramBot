package Commands

import (
	"BTCWalletTelegramBot/RPC"
	"fmt"
	"regexp"
	"strings"
)

var commands = [10]string{"start", "height", "help"}

type Handler struct {
	 RPC RPC.Rpc
}

func (h *Handler) IsCommand(cmd string) bool {
	cmd = strings.ToLower(cmd)
	for _, v := range commands{
		if v == cmd { return true }
	}
	return false
}

func (h *Handler) Handle(from int, cmd, args string) string{
	regex := regexp.MustCompile(" +")
	arguments := regex.Split(args, -1)
	switch cmd {
	case "start" :
		return fmt.Sprintf("Your User ID: <code>%d</code>", from)
	case "help":
		txt := "--Available Commands--\n"
		for _, v := range commands {
			if len(v) <= 0 { continue }
			txt += fmt.Sprintf("/%s, ", v)
		}
		return txt[: len(txt) - 2]
	case "height":
		v, err := h.RPC.Client.GetBlockCount()
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("Block Height: <b>%d</b>", v)
	default:
		return strings.Join(arguments, ", ")
	}
}

func CreateHandler(r RPC.Rpc) Handler{
	var h Handler
	h.RPC = r
	return h
}
