package cryptography_test

import (
	"bytes"
	"fmt"
)

type FakeEncrypter struct{}

func (f *FakeEncrypter) Encrypt(payload map[string]interface{}) string {
	return createKeyValuePairs(payload)
}

func (f *FakeEncrypter) Verify(token string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func createKeyValuePairs(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "\"%s\":\"%s\"\n", key, value)
	}
	return b.String()
}
