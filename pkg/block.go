package pkg

import (
	"bytes"
	"encoding/gob"
	"math"
	"math/rand"
	"time"
)

// Block inside our blockchain
type Block struct {
	PrevHash    string
	Transaction Transaction
	timestamp   time.Time
	nonce       int
}

// NewBlock of type Block
func NewBlock(prevHash string, transaction Transaction) Block {
	block := Block{}
	block.PrevHash = prevHash
	block.Transaction = transaction
	block.timestamp = time.Now()
	block.nonce = int(math.Round(rand.Float64() * 99999999))
	return block
}

func (b Block) hash() string {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(b)
	return buf.String()
}
