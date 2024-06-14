package main

func (blockchain *BlockChain) AddBlock(data string) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
