package pkg

import (
	"crypto/rsa"
	"encoding/json"
)

// Transaction inside our blockchain
type Transaction struct {
	Amount float32
	Payer  *rsa.PublicKey
	Payee  *rsa.PublicKey
}

func (t Transaction) toString() []byte {
	bytes, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return bytes
}
