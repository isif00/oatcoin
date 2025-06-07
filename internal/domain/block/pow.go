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
func (pow *ProofOfWork) Run(workers int) (int, []byte) {
	type result struct {
		nonce int
		hash  []byte
	}

	resultChan := make(chan result)
	stopChan := make(chan struct{})

	for w := range workers {
		go func(workerID int) {
			var hash [32]byte
			nonce := workerID

			for {
				select {
				case <-stopChan:
					return
				default:
					data := pow.prepareData(nonce)
					hash = sha256.Sum256(data)

					var hashInt big.Int
					hashInt.SetBytes(hash[:])

					if hashInt.Cmp(pow.Target) == -1 {
						select {
						case resultChan <- result{nonce, hash[:]}:
						case <-stopChan:
						}
						return
					}
					nonce += workers
				}
			}
		}(w)
	}

	found := <-resultChan
	close(stopChan)
	return found.nonce, found.hash
}
