package pkg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// Wallet for transactions
type Wallet struct {
	PublicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewWallet for blockchain transactons
func NewWallet() Wallet {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Could not create wallet, unable to generate RSA key")
		panic(err)
	}

	wallet := Wallet{}
	wallet.PrivateKey = key
	wallet.PublicKey = &key.PublicKey

	return wallet
}

// SendMoney send money from one walled to another
func (w Wallet) SendMoney(amount float32, payeeKey *rsa.PublicKey) {
	transaction := Transaction{amount, w.PublicKey, payeeKey}

	hashed := sha256.Sum256(transaction.toString())
	signature, err := rsa.SignPKCS1v15(rand.Reader, w.privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	Instance().addBlock(transaction, w.PublicKey, string(signature))
}
