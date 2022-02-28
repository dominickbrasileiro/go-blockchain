package main

import "time"

type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Difficulty   int
}

func (b *Blockchain) AddBlock(data BlockData) {
	lastBlock := b.Chain[len(b.Chain)-1]

	newBlock := Block{
		Data:         data,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}

	newBlock.Mine(b.Difficulty)

	b.Chain = append(b.Chain, newBlock)
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}

	return Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []Block{genesisBlock},
		Difficulty:   difficulty,
	}
}
