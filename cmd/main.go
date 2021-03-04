package cmd

import (
	"fmt"

	"github.com/faagerholm/blockchain-go/pkg/wallet"
)

func main() {
	bob := wallet.NewWallet()
	fmt.Println(bob.Publickey)
}
