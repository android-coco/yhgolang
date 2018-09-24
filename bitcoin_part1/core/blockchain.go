package core

//BlockChan keeps a sequence of Blocks
type BlockChan struct {
	Blocks []*Block
}

//AddBlock saves provided data as a block in the blockchain
func (bc *BlockChan)AddBlock(data string)  {
	//最后一个区块作为新的区块的前一个区块
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.Blocks = append(bc.Blocks,newBlock)
}
//NewBlockChain creates a new BlockChan with genesis Block
func NewBlockChain() *BlockChan {
	return &BlockChan{[]*Block{NewGenesisBlock()}}
}