package blockchain

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"polaris/blockchain/block"
)

type Blockchain struct {
	Blocks []*block.Block
}

func (b *Blockchain) AppendBlock(nblock *block.Block) {
	if len(b.Blocks) == 0 {
		b.Blocks = append(b.Blocks, nblock)
		return
	}
	if isValid(*nblock, *b.Blocks[len(b.Blocks)-1 ]) {
		b.Blocks = append(b.Blocks, nblock)
		return
	}
	log.Fatal("Append block fail, invalid block")
}

func (b *Blockchain) SendData(data string) {
	prev := b.Blocks[len(b.Blocks)-1]
	nblock := prev.GenNewBlock(data)
	b.AppendBlock(nblock)
}

func NewBlockChain() *Blockchain{
	genesisBlock := block.GenGenesisBlock()
	blockchain := Blockchain{}
	blockchain.AppendBlock(genesisBlock)
	return &blockchain
}

func (b *Blockchain) Print() {
	for _, v := range b.Blocks {
		fmt.Printf("Index: %d \n", v.Index)
		fmt.Printf("PrevHash: %s \n", v.PrevBlockHash)
		fmt.Printf("Hash: %s \n", v.Hash)
		fmt.Printf("Data: %s \n", v.Data)
		fmt.Printf("Timestamp: %v \n", v.TimeStamp)
	}
}


func isValid(nblock block.Block, oblock block.Block) bool {
	if nblock.Index-1 != oblock.Index {
		return false
	}
	if nblock.PrevBlockHash != oblock.Hash {
		return false
	}
	if nblock.CalcuHash() != nblock.Hash {
		return false
	}
	return true
}
