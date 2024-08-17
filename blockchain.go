package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
   )
	
	type Block struct {
		data	map[string]interface{}
		hash 	string
		previousHash string
		timestamp time.Time
		pow 	  int
	}

	type Blockchain struct {
		genesisBlock Block
		chain 		[]Block
		difficulty  int
	}

	func (b Block) calculeHash() string {
		data, _ := json.Marshal(b.data)
		blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
		blockHash := sha256.Sum256([]byte(blockData))
		return fmt.Sprintf("%x", blockHash)
	}

	func (b *Block) mine(difficulty int) {
		for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
			b.pow++
			b.hash = b.calculeHash()
		}
	}

	func CreateBlockchain(difficulty int) Blockchain {
		genesisBlock := Block {
			hash: "0",
			timestamp: time.Now(),
		}
		return Blockchain{
			genesisBlock,
			[]Block{genesisBlock},
			difficulty,
		}
	}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:     blockData,
		previousHash: lastBlock.hash,
		timestamp:   time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculeHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	// cria nova instância de blockchain com a dificuldade da mineração 2
	blockchain := CreateBlockchain(2)

	blockchain.addBlock("Caique", "Aline", 5)
	blockchain.addBlock("João", "Gaby", 2)

	fmt.Println(blockchain.isValid())
}