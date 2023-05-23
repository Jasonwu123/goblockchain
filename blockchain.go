package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	prevHash     string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.prevHash = prevHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("nonce               %d\n", b.nonce)
	fmt.Printf("previous_hash       %s\n", b.prevHash)
	fmt.Printf("timestamp           %d\n", b.timestamp)
	fmt.Printf("transactions        %s\n", b.transactions)
}

func (b Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(m)
	return sha256.Sum256([]byte(m))
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash string) *Block {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	/*
		blockChain := NewBlockchain()
		blockChain.Print()
		blockChain.CreateBlock(5, "hash 1")
		blockChain.Print()
	*/
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
}
