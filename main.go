package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"

	blockChain := NewBlockchain(myBlockchainAddress)
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

}
