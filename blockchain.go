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
	prevHash     [32]byte
	timestamp    int64
	transactions []*Transaction
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.prevHash = prevHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("nonce               %d\n", b.nonce)
	fmt.Printf("previous_hash       %x\n", b.prevHash)
	fmt.Printf("timestamp           %d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	//fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"prev_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address   %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipinet_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value   %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	prevHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, prevHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	prevHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, prevHash)
	blockChain.Print()

	/*
		block := &Block{nonce: 1}
		fmt.Printf("%x\n", block.Hash())
	*/
}
