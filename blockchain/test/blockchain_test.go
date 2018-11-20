package test

import (
	"polaris/blockchain"
	"testing"
)

func TestBlockchain(t *testing.T) {
	b := blockchain.NewBlockChain()
	b.SendData("Send 1 BTC to 1302岁的龙猫")
	b.SendData("Send 1 EOS to 1302岁的龙猫")
	b.Print()
}
