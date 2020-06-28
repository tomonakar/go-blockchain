package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Blockの構造体
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// Blockを新規作成
func NewBlock(nonce int, previousHash string) *Block {
	// こう書いても良いし
	b := new(Block) // newを使うとポインタが返る
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

	// こう書いても良い
	// return &Block{
	// 	timestamp: time.Now().UnixNano(),
	// }
}

// Block Structの中身を単に見やすくするためのもの
func (b *Block) Print() {
	fmt.Printf("timestamp          %d\n", b.timestamp)
	fmt.Printf("nonce              %d\n", b.nonce)
	fmt.Printf("previousHash       %s\n", b.previousHash)
	fmt.Printf("transactions       %s\n", b.transactions)
}

// Blockchain構造体
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash2")
	blockChain.Print()
}
