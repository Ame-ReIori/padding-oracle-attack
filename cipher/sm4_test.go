package cipher

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"padding_oracle_attack/utils"
	"testing"
)

func TestPKCS7Pad(t *testing.T) {
	msg := []byte("abcdefghijklmnop")
	paddedMsg, err := PKCS7Pad(msg, aes.BlockSize)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(paddedMsg)
	}
}

func TestPKCS7Unpad(t *testing.T) {
	msg, _ := hex.DecodeString("6162636465666768696a6b6c6d6e6f70710f0f0f0f0f0f0f0f0f0f0f0f0f0f0f")
	UnpaddedMsg, err := PKCS7Unpad(msg, aes.BlockSize)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(UnpaddedMsg)
	}
}

func TestEncrypt(t *testing.T) {
	plaintext := []byte("gyw")
	key, err := hex.DecodeString("0123456789abcdeffedcba9876543210")
	if err != nil {
		panic(err)
	} else {
		ciphertext, _ := Encrypt(key, plaintext)
		utils.TypePrint(ciphertext, utils.CHAR)
	}
}

func TestDecrypt(t *testing.T) {
	plaintext, _ := ioutil.ReadFile("../data/test.plain")
	key, err := ioutil.ReadFile("../data/test.key")
	if err != nil {
		panic(err)
	}
	ciphertext, _ := Encrypt(key, plaintext)
	decText, _ := Decrypt(key, ciphertext)
	isEqual := bytes.Equal(decText, plaintext)
	fmt.Println(isEqual)
}