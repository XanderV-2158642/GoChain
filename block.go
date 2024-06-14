package main

import (
	"fmt"

	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func CreateBlock(Header, Body string) {
	fmt.Println("Header: ", Header)
	fmt.Println("Body: ", Body)
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PrevBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
