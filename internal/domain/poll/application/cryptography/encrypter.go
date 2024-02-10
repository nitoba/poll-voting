package cryptography

type Encrypter interface {
	Encrypt(payload map[string]interface{}) string
	Verify(token string) (map[string]interface{}, error)
}
