package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func Aes256CbcEncrypt(plaintext string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedPlaintext := pkcs7Padding([]byte(plaintext), block.BlockSize())

	ciphertext := make([]byte, len(paddedPlaintext))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Aes256CbcDecrypt(ciphertext string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(decodedCiphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, decodedCiphertext)

	unPaddedPlaintext, err := pkcs7UnPadding(decrypted)
	if err != nil {
		return "", err
	}

	return string(unPaddedPlaintext), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padData...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	unPadding := int(data[length-1])

	if unPadding < 1 || unPadding > 16 {
		return nil, errors.New("invalid padding")
	}

	return data[:(length - unPadding)], nil
}
