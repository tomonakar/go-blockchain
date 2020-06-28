package main

import (
	"fmt"
	"goblockchain/block"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
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

	fmt.Printf("my %.1f\n", blockChain.CaluculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockChain.CaluculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CaluculateTotalAmount("D"))
}
