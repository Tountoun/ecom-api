package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	secret := []byte("secretkey")

	token, err := CreateJWT(secret, 3)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("expected token to be not empty")
	}
}