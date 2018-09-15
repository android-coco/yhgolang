package core

//Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*Block
}

//AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain)AddBlock(data string)  {
	//最后一个区块作为新的区块的前一个区块
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.Blocks = append(bc.Blocks,newBlock)
}
//NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}