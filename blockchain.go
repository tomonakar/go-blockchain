package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlack(nonce int, previousHash string) *Block {
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

// Printメソッドは、BlockStructの中身を単に見やすくするためのもの
func (b *Block) Print() {
	fmt.Printf("timestamp          %d\n", b.timestamp)
	fmt.Printf("nonce              %d\n", b.nonce)
	fmt.Printf("previousHash       %s\n", b.previousHash)
	fmt.Printf("transactions       %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	log.Println("test")
	fmt.Println("test2")
	b := NewBlack(0, "init hash")
	b.Print()
}
