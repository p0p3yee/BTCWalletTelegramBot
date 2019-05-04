# BTC Wallet Telegram Bot

A Telegram bot for you to interact with btc wallet.

## Libraries Needed:

RPC Client

["github.com/btcsuite/btcd/rpcclient"]("github.com/btcsuite/btcd/rpcclient")

`go get -u github.com/btcsuite/btcd/rpcclient`

---

Telegram Bot API

["github.com/go-telegram-bot-api/telegram-bot-api"]("github.com/go-telegram-bot-api/telegram-bot-api")

`go get github.com/go-telegram-bot-api/telegram-bot-api`

---

## Configuration

Create `config.json`

```json
{
  "Host": "host:port",
  "User": "rpcuser",
  "Pass": "rpcuserpassword",
  "token": "telegram bot token",
  "ownerID": 0
}
```

Pass `-deprecatedrpc=accounts` when you start BTC Daemon.

or put `deprecatedrpc=accounts` in `bitcoin.conf`

## Build & Run

1. `go install`

2. `go run main.go ` OR `go build main.go`

3. `./BTCWalletTelegramBot`