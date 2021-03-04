package main

import (
	"fmt"

	"github.com/faagerholm/blockchain-go/pkg"
)

func main() {
	bob := pkg.NewWallet()
	alice := pkg.NewWallet()
	steve := pkg.NewWallet()

	// Bob sends money to alice
	fmt.Println("Bob sending 10 'coins' to alice..")
	alice.SendMoney(10, bob.PublicKey)

	// Bob sends money to steve
	fmt.Println("Bob sending 50 'coins' to steve..")
	steve.SendMoney(50, bob.PublicKey)

	// Fix: pretty print last block
	fmt.Printf("%v", pkg.Instance().LastBlock())
}
