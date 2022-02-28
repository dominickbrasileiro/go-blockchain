package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
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
