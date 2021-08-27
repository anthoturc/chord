package hash

import (
	"math/big"
	"testing"
)

func TestHash(t *testing.T) {
	id := Hash("SomeRandomKey")
	if id.Cmp(big.NewInt(100)) > 0 || id.Cmp(big.NewInt(0)) < 0 {
		t.Errorf("Hashed key should result in an 0 <= 0 < 100")
	}
}
