package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hashed == "" {
		t.Errorf("expected hashed to be not empty")
	}

	if hashed == password {
		t.Errorf("expected hashed to be different from password")
	}
}


func TestComparePassword(t *testing.T) {
	password := "hereisthepassword"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	t.Run("correct password should match the hashed password", func(t *testing.T) {
		if !ComparePassword(hashed, []byte(password)) {
			t.Errorf("expected password to match hashed")
		}	
	})
	
	t.Run("incorrect password should not match the hashed password", func(t *testing.T) {
		if ComparePassword(hashed, []byte("notcorrectpassword")) {
			t.Errorf("expected password to not match hashed")
		}
	})
}