package util

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/nicktaobo/go_tool/btcx"
)

func ParseAddress(addr string, net btcx.Net) (btcutil.Address, error) {
	return btcutil.DecodeAddress(addr, net.Params())
}

func ParseTxHash(txid string) (*chainhash.Hash, error) {
	return chainhash.NewHashFromStr(txid)
}
