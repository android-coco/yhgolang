package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)


var (
	maxNonce = math.MaxInt64
)
//目标位20
const targetBits = 20

//工作量证明结构
//ProofofWork  represents a proof-of-work
type ProofofWork struct {
	block  *Block   //区块
	target *big.Int //目标
}

// NewProofofWork builds and returns a ProofofWork
func NewProofofWork(b *Block) *ProofofWork {
	target := big.NewInt(1)
	//移位操作
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofofWork{b, target}

	return pow
}

func (pow *ProofofWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//Run performs a proof-of-work
func (pow *ProofofWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing %s\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

//校验
//Validate  validates blok's PoW
func (pow *ProofofWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
