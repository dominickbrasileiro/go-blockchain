package main

import "time"

type Block struct {
	Data         map[string]interface{}
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Pow          int
}
