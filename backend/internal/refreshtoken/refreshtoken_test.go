package refreshtoken

import "testing"

func TestGenerateAndHash(t *testing.T) {
	tokenA, err := Generate()
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	tokenB, err := Generate()
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	if tokenA == tokenB {
		t.Error("expected two generated tokens to be different")
	}

	hashA1 := Hash(tokenA)
	hashA2 := Hash(tokenA)

	if hashA1 != hashA2 {
		t.Error("expected hashing the same token twice to produce the same hash")
	}

	if hashA1 == tokenA {
		t.Error("hash should not equal the raw token")
	}
}
