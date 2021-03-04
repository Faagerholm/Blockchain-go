package pkg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"strconv"
)

// Chain struct
type Chain struct {
	block []Block
}

// Singleton, there can only be one chain!
var instance *Chain

// Instance of a chian, made singleton
func Instance() *Chain {
	if instance == nil {
		// NOT THREAD SAFE !!
		key1, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			fmt.Println("Unable to generate RSA key")
			panic(err)
		}
		key2, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			fmt.Println("Unable to generate RSA key")
			panic(err)
		}
		instance = &Chain{block: []Block{NewBlock("", Transaction{100.0, &key1.PublicKey, &key2.PublicKey})}}
	}
	return instance
}

// LastBlock from the chain
func (chain *Chain) LastBlock() Block {
	// Return last block of chain
	return instance.block[len(instance.block)-1]
}

func (chain *Chain) mine(nonce int) {
	var solution = 1
	fmt.Println("Mining...")
	for {

		//TODO: fix
		// hash := sha256.Sum256([]byte(strconv.Itoa(nonce + solution)))
		attempt := sha256.Sum256([]byte(strconv.Itoa(nonce + solution)))

		if string(attempt[:4]) == "0000" {
			return
		}
		solution++
	}
}

func (chain *Chain) addBlock(transaction Transaction, senderKey *rsa.PublicKey, signature string) {
	// Add block to chain
	hashed := sha256.Sum256(transaction.toString())
	err := rsa.VerifyPKCS1v15(senderKey, crypto.SHA256, hashed[:], []byte(signature))
	if err != nil {
		fmt.Println(err)
		return
	}
	newBlock := NewBlock(chain.LastBlock().hash(), transaction)
	chain.mine(newBlock.nonce)
	chain.block = append(chain.block, newBlock)
}
