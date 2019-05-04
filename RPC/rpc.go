package RPC

import (
	"github.com/btcsuite/btcd/rpcclient"
)

type Rpc struct {
	Client *rpcclient.Client
}

func (c *Rpc) Create(host, user, pass string) error {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host: host,
		User: user,
		Pass: pass,
		HTTPPostMode: true,
		DisableTLS: true,
	}, nil)

	if err != nil { return err }
	c.Client = client
	return nil
}