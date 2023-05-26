package utils

import "golang.org/x/crypto/bcrypt"

type PasswordHasher struct{}

func (p PasswordHasher) Hash(raw string) (hash string, err error) {
	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}

	hashedPassword := string(hashedPasswordInBytes)
	return hashedPassword, nil
}

func (p PasswordHasher) Verify(hash, raw string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw)); err != nil {
		return err
	}
	return nil
}
