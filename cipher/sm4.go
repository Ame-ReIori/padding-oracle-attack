package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"padding_oracle_attack/utils"
)

var (
	ErrInvalidBlockSize = errors.New("invalid block size")
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
	BadPadding = errors.New("bad padding")
)

func Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	plaintext, _ = PKCS7Pad(plaintext, aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize + len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	plaintext := make([]byte, len(ciphertext) - aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	plaintext, err = PKCS7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, BadPadding
	} else {
		return plaintext, nil
	}
}

func DecryptWithOutUnpad(key []byte, ciphertext []byte) {
	plaintext := make([]byte, len(ciphertext) - aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	utils.TypePrint(plaintext, utils.HEX)
}

func PKCS7Pad(msg []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if msg == nil || len(msg) == 0 {
		return nil ,ErrInvalidPKCS7Data
	}
	n := blockSize - (len(msg) % blockSize)
	paddedMsg := make([]byte, len(msg) + n)
	copy(paddedMsg, msg)
	copy(paddedMsg[len(msg):], bytes.Repeat([]byte{byte(n)}, n))
	return paddedMsg, nil
}

func PKCS7Unpad(msg []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if msg == nil || len(msg) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(msg) % blockSize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	b := msg[len(msg)-1]
	n := int(b)
	if n == 0 || n > blockSize {
		return nil ,ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if msg[len(msg) - n + i] != b {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return msg[:len(msg) - n], nil
}
