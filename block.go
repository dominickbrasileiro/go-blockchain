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
	Data         map[string]interface{}
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	POW          int
}

func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)

	dataToHash := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.POW)

	hash := sha256.Sum256([]byte(dataToHash))

	return fmt.Sprintf("%x", hash)
}

func (b *Block) Mine(difficulty int) {
	desiredPrefix := strings.Repeat("0", difficulty)

	for !strings.HasPrefix(b.Hash, desiredPrefix) {
		b.POW++
		b.Hash = b.CalculateHash()
	}
}
