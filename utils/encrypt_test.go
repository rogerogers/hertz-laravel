package utils

import (
	"os"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestAes256CbcEncrypt(t *testing.T) {
}

func TestAes256CbcDecrypt(t *testing.T) {
	tokenMap, err := safeDecode("eyJpdiI6IllubnlkdVI1aEZVWWlUbk1SMzBHd1E9PSIsInZhbHVlIjoiQnBQdlRmbVErbU5IWHFIZXMwLzJkdkcyMXNuUHR1WmNtNUVmcTc2Q2Q3aW9oQkNuZzlPbnhzVWJqNW13dmJNc0VLMFV1N2c4QUJ5M1RzMWhZQytsR1RITjllZ0RzcTZBR3pybmRsSm1EVjlMaENJeFZUenFwcytLTFRqTTNWUFkiLCJtYWMiOiI0ZDc1YWRlOWUzYWE5OTk1NDc2NGMwNTJhNWI1OGIwOTVjNzBmYTYxNDI3MjVlMGZiZjNiYzkyNWZmZGE0MzI3IiwidGFnIjoiIn0%3D")
	assert.Nil(t, err)
	key, err := B64Decode(os.Getenv("AUTH_KEY"))
	assert.Nil(t, err)
	iv, err := B64Decode(tokenMap["iv"].(string))
	assert.Nil(t, err)
	a, err := Aes256CbcDecrypt(tokenMap["value"].(string), key, iv)
	assert.Nil(t, err)
	t.Logf("%s", a)
}

func safeDecode(token string) (map[string]any, error) {
	tokenMap := make(map[string]any)
	tokenStr, err := B64SafeDecode(token)
	if err != nil {
		return nil, err
	}
	err = sonic.Unmarshal(tokenStr, &tokenMap)
	if err != nil {
		return nil, err
	}
	return tokenMap, nil
}
