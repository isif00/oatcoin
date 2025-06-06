package crypto

import (
	"crypto/sha256"

	"github.com/btcsuite/btcutil/base58"
)

func Base58CheckEncode(payload []byte) string {
	checksum := checkSum(payload)
	full := append(payload, checksum...)
	return base58.Encode(full)
}

func checkSum(payload []byte) []byte {
	first := sha256.Sum256(payload)
	second := sha256.Sum256(first[:])
	return second[:4]
}
