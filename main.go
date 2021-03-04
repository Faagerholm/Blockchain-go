package main

import (
	"fmt"

	"github.com/faagerholm/blockchain-go/pkg/wallet"
)

func main() {
	fmt.Println("Hello world")
	bob := wallet.NewWallet()
	fmt.Println(bob.Publickey)
}
