package hash

import (
	"crypto/sha1"
	"math/big"
)

func Hash(key string) *big.Int {
	sha := sha1.New()
	sha.Write([]byte(key))
	id := new(big.Int).SetBytes(sha.Sum(nil))
	// TODO: Do not use a % 100 (this is just for local testing)
	return id.Mod(id, big.NewInt(100))
}
