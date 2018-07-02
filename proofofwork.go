package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const (
	targetBits = 24
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}

		nonce++
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			intToHex(pow.block.Timestamp),
			intToHex(int64(targetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}
