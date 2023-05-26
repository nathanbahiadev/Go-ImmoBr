package entity

import (
	"errors"
)

type Password struct {
	Value string `json:"password"`
}

type PasswordHasher interface {
	Hash(raw string) (hash string, err error)
	Verify(hash, raw string) error
}

func (p *Password) Validate() error {
	if len(p.Value) < 8 {
		return errors.New("the password must contain at least 8 characters")
	}
	return nil
}

func (p *Password) Hash(hasher PasswordHasher) error {
	hashedPassword, err := hasher.Hash(p.Value)
	if err != nil {
		return err
	}

	p.Value = string(hashedPassword)
	return nil
}

func (p *Password) Verify(hasher PasswordHasher, hash, raw string) error {
	if err := hasher.Verify(hash, raw); err != nil {
		return err
	}
	return nil
}
