package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-Difficulty)
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			fmt.Appendf(nil, "%d", pow.Block.Timestamp),
			fmt.Appendf(nil, "%d", nonce),
		},
		[]byte{},
	)
}
func (pow *ProofOfWork) Run() (int, []byte) {
	for nonce := 0; ; nonce++ {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)

		if new(big.Int).SetBytes(hash[:]).Cmp(pow.Target) < 0 {
			return nonce, hash[:]
		}
	}
}