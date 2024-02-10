package cryptography

type HashGenerator interface {
	Hash(text string) (string, error)
}
