package utils

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestAes256CbcEncrypt(t *testing.T) {
}

func TestAes256CbcDecrypt(t *testing.T) {
	key, err := B64Decode("aj5F8ziDtRkdSShsM30vsWNNK1UqJcDDbuUsBlGkz2I=")
	assert.Nil(t, err)
	iv, err := B64Decode("2d1u0X0yK+dwjYmUZkt8zA==")
	assert.Nil(t, err)
	a, err := Aes256CbcDecrypt("lGJuvE0FvRmJiEkX9RL0zictoB31AGoanvLcNMImks45qaHcIkw9KWTjNMnpWEdPMjqGjsyF/iUqJ2KXuXHpKgr3NrpMkpRxEzE+/GAXt2Q7fHo9dgdpYjs7IUpl6JfH", key, iv)
	assert.Nil(t, err)
	t.Logf("%s", a)
}
