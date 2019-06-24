package process

import (
	//"coinford_process/models"
	"coinford_process/configs"
	"github.com/btcsuite/btcrpcclient"
	//"github.com/btcsuite/btcd/btcjson"
	//"github.com/btcsuite/btcd/chaincfg"
	//"github.com/btcsuite/btcutil"
	//"fmt"
	//"time"
	//"log"
)

type BTCProcess struct {
	client 		*btcrpcclient.Client
	currencyid 	int64
}

func (btcp *BTCProcess) GetNewAddress() (string, error) {
	var err error
	btcp.client, err = btcp.Connect()
	if err == nil {
		address, err := btcp.client.GetNewAddress(configs.BTC_IN_ACCOUNT)
		if err == nil {
			return address.String(), nil
		}
		return *configs.NullString, err
	}
	return *configs.NullString, err
}

func (btcp *BTCProcess) Connect() (*btcrpcclient.Client, error) {
	wpassphrase := "suman123"
	var wtimeout int64
	wtimeout = 60 //seconds
	client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         configs.BTCAuthUrl,
		User:         "bitcoinrpc",
		Pass:         "7003dd926de77d98d64d9c8152a3ec68",
	}, nil)
	if err == nil {
		client.WalletPassphrase(wpassphrase, wtimeout)
	}
	return client, err
}