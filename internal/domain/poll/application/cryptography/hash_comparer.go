package cryptography

type HashComparer interface {
	Compare(text string, hash string) bool
}
