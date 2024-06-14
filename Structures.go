package main

type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	MyBlockHash   []byte
	AllData       []byte
}

type BlockChain struct {
	Blocks []*Block
}
