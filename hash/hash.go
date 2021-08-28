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

// Helper method to see if a given id is between start and end while accounting
// for identifiers that wrap around the chord ring
func IsBetween(start, id, end *big.Int, inclusive bool) bool {
	if end.Cmp(start) > 0 { // Ids can wrap around the ring
		return (start.Cmp(id) < 0 && id.Cmp(end) < 0) || (inclusive && id.Cmp(end) == 0)
	} else {
		return start.Cmp(id) < 0 || id.Cmp(end) < 0 || (inclusive && id.Cmp(end) == 0)
	}
}
