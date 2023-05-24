package main

import (
	"fmt"
	"goblockchain/block"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	/*
		myBlockchainAddress := "my_blockchain_address"

		blockChain := block.NewBlockchain(myBlockchainAddress)
		blockChain.Print()

		blockChain.AddTransaction("A", "B", 1.0)
		blockChain.Mining()
		blockChain.Print()

		blockChain.AddTransaction("C", "D", 2.0)
		blockChain.AddTransaction("X", "Y", 3.0)
		blockChain.Mining()
		blockChain.Print()

		fmt.Printf("my %.1f\n", blockChain.CalculdateTotalAmount("my_blockchain_address"))
		fmt.Printf("A %.1f\n", blockChain.CalculdateTotalAmount("A"))
		fmt.Printf("B %.1f\n", blockChain.CalculdateTotalAmount("B"))
		fmt.Printf("C %.1f\n", blockChain.CalculdateTotalAmount("C"))
		fmt.Printf("D %.1f\n", blockChain.CalculdateTotalAmount("D"))
	*/

	wm := wallet.NewWallet()
	wa := wallet.NewWallet()
	wb := wallet.NewWallet()
	t := wallet.NewTransaction(wa.PrivateKey(), wa.PublicKey(), wa.BlockchainAddress(), wb.BlockchainAddress(), 1.0)

	bc := block.NewBlockchain(wm.BlockchainAddress())

	isAdded := bc.AddTransaction(wa.BlockchainAddress(), wb.BlockchainAddress(), 1.0, wa.PublicKey(), t.GenerateSignature())

	fmt.Println("Added? ", isAdded)

	bc.Mining()
	bc.Print()

	fmt.Printf("A %.1f\n", bc.CalculdateTotalAmount(wa.BlockchainAddress()))
	fmt.Printf("B %.1f\n", bc.CalculdateTotalAmount(wb.BlockchainAddress()))
	fmt.Printf("M %.1f\n", bc.CalculdateTotalAmount(wm.BlockchainAddress()))

}
