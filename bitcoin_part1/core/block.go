package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block keeps block headers
type Block struct {
	Timestamp     int64  //区块创建时间
	Data          []byte //区块包含的数据
	PrevBlockHash []byte //前一个区块的哈希值
	Hash          []byte //区块自身的哈希值,用于校验区块数据有效
	Nonce         int    //工作量证明
}

//NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash}

	pow := NewProofofWork(block) //工作量证明
	nonce, hash := pow.Run()     //挖矿

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	//上个区块的Hash 自己的数据  和时间戳
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	//sha256
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//NewGenesisBlock create and returns genesis Block
//创世纪块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
