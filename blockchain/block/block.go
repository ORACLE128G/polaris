package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64
	TimeStamp     int64
	PrevBlockHash string
	Hash          string

	Data string
}

func (b *Block) CalcuHash() string {
	blockData := string(b.Index) + string(b.TimeStamp) + b.PrevBlockHash + b.Data
	r := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(r[:])
}

func (b *Block) GenNewBlock(data string) *Block {
	newBlock := Block{
		Index:         b.Index + 1,
		PrevBlockHash: b.Hash,
		TimeStamp:     time.Now().Unix(),
	}
	newBlock.Data = data
	newBlock.Hash = newBlock.CalcuHash()
	return &newBlock
}

func GenGenesisBlock() *Block{
	newBlock := Block{
		Index: -1,
		Hash:  "",
	}
	return newBlock.GenNewBlock("Genesis Block")
}
