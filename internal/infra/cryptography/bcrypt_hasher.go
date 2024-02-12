package cryptography

import "golang.org/x/crypto/bcrypt"

type BCryptHasher struct{}

func (*BCryptHasher) Hash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (*BCryptHasher) Compare(text string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

func CreateBCryptHasher() *BCryptHasher {
	return &BCryptHasher{}
}
