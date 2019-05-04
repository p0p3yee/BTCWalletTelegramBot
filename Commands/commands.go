package Commands

import (
	"BTCWalletTelegramBot/RPC"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"regexp"
	"strings"
)

var commands = [11]string{
	"ping",
	"start",
	"height",
	"help",
	"listacc",
	"getaddrbyac",
	"getbalancebyacc",
	"getreceivedbyacc",
	"getnewaddr",
	"gettrans",
	"listtrans",
}

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
	case "ping":
		err := h.RPC.Client.Ping()
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return "Pong!"

	case "start" :
		return fmt.Sprintf("Your User ID: <code>%d</code>", from)

	case "help":
		txt := "--<b>Available Commands</b>--\n"
		for _, v := range commands {
			if len(v) <= 0 { continue }
			txt += fmt.Sprintf("/%s, ", v)
		}
		return txt[: len(txt) - 2]

	case "height":
		v, err := h.RPC.Client.GetBlockCount()
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("Block Height: <b>%d</b>", v)

	case "listacc":
		m, err := h.RPC.Client.ListAccounts()
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		txt := "--<b>Accounts</b>--\n"
		for v, k := range m {
			txt += fmt.Sprintf("<code>%s</code>: <b>%.8f BTC</b>\n", v, k.ToBTC())
		}
		return txt

	case "getaddrbyac":
		param := ""
		if len(arguments) > 0 {
			param = arguments[0]
		}
		acc, err := h.RPC.Client.GetAddressesByAccount(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		txt := ""
		if len(acc) == 0 {
			txt = fmt.Sprintf("There is <b>NO</b> Addresses of Account: <b>%s</b>", param)
		}else {
			txt = fmt.Sprintf("Addresses of Account: <b>%s</b>\n", param)
			if param == ""{
				txt = fmt.Sprintf("All Addresses:\n", param)
			}
			for i, v := range acc {
				txt += fmt.Sprintf("%d. <code>%s</code>\n", i, v)
			}
		}
		return txt

	case "getbalancebyacc":
		param := ""
		if len(arguments) == 1 {
			param = arguments[0]
		}

		bal, err := h.RPC.Client.GetBalance(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("Balance of Account: <b>%s</b>: <b>%s</b>", param, bal.String())

	case "getreceivedbyacc":
		param := ""
		if len(arguments) == 1 {
			param = arguments[0]
		}
		bal, err := h.RPC.Client.GetReceivedByAccount(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("Account: <b>%s</b> Received Total: <b>%s</b>", param, bal.String())

	case "getnewaddr":
		param := ""
		if len(arguments) == 1 {
			param = arguments[0]
		}
		addr, err := h.RPC.Client.GetNewAddress(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("Address with account: <b>%s</b>, Generated: <code>%s</code>", param, addr.String())

	case "gettrans":
		param := ""
		if len(arguments) <= 0 {
			return "<b>Error</b>: No TxHash Found.\nUsage: /gettrans <code>TxHash</code>"
		}
		hash, err := chainhash.NewHashFromStr(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		trans, err := h.RPC.Client.GetTransaction(hash)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		return fmt.Sprintf("TxID: <code>%s</code>\nBlock Hash: <code>%s</code>\nBlock Index: <code>%d</code>\nBlock Timestamp: <code>%d</code>\nAmount: <b>%.8f BTC</b>\n\nFee: <b>%.8f BTC</b>\nConfirmation: <b>%d</b>", trans.TxID, trans.BlockHash, trans.BlockIndex, trans.BlockTime, trans.Amount, trans.Fee, trans.Confirmations)

	case "listtrans":
		param := ""
		if len(arguments) == 1 {
			param = arguments[0]
		}
		trans, err := h.RPC.Client.ListTransactions(param)
		if err != nil { return fmt.Sprintf("<b>Error</b>: %s", err) }
		txt := "--<b>Transactions</b>--\n"
		for i, t := range trans {
			txt += fmt.Sprintf("<b>%d</b>. <code>%s</code>", i, t.TxID)
		}
		return txt
	default:
		return strings.Join(arguments, ", ")

	}
}

func CreateHandler(r RPC.Rpc) Handler{
	var h Handler
	h.RPC = r
	return h
}
