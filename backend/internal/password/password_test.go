package password

import "testing"

func TestHashAndVerify(t *testing.T) {
	plain := "super-secret-123"

	hashed, err := Hash(plain)
	if err != nil {
		t.Fatalf("Hash return error: %v", err)
	}

	if !Verify(plain, hashed) {
		t.Error("expected Verify to return true for correct password")
	}

	if Verify("wrong password", hashed) {
		t.Error("expected Verify to return false for incorrect password")
	}
}
