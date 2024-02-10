package cryptography_test

import (
	"fmt"
	"strings"
)

type FakeHasher struct{}

func (f *FakeHasher) Hash(text string) (string, error) {
	text = fmt.Sprintf("hashed:" + text)
	return text, nil
}

func (f *FakeHasher) Compare(text string, hash string) bool {
	return strings.Contains(hash, "hashed:")
}
