package PaddingOracle

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"padding_oracle_attack/cipher"
	"padding_oracle_attack/utils"
	"testing"
)

func TestGetPad(t *testing.T) {
	plaintext := []byte("abcdefghijklmnopqrstuvwxyzabcdef")
	key, _ := hex.DecodeString("0123456789abcdeffedcba9876543210")
	ciphertext, _ := cipher.Encrypt(key, plaintext)

	oracle := New(key, plaintext)
	pad := GetPad(oracle, ciphertext)
	fmt.Println(pad)

	utils.TypePrint(ciphertext, utils.HEX)
}

func TestProcessBlock(t *testing.T) {
	block := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 15, 15, 15, 15, 15, 15, 15}
	ProcessBlock(block, 7)
	utils.TypePrint(block, utils.HEX)
}

func TestAttack(t *testing.T) {
	plaintext, _ := ioutil.ReadFile("../data/test.plain")
	key, _ := ioutil.ReadFile("../data/test.key")
	ciphertextFile, _ := os.Create("../data/test.enc")
	ciphertext, _ := cipher.Encrypt(key, plaintext)
	_, _ = ciphertextFile.Write(ciphertext)

	oracle := New(key, plaintext)
	decText := Attack(oracle, ciphertext)
	decTextFile, _ := os.Create("../data/test.dec")
	_, _ = decTextFile.Write(decText)

	isEqual := bytes.Equal(plaintext, decText)
	fmt.Println(isEqual)
}

func TestAttackVisible(t *testing.T) {
	plaintext := []byte("sdfoasdinfhsajdfb32p98oybyq97ocb7qova7icacyba7co3qcuai3cyachuiiuhiuhuihiuhuiohiuygytgi679gf867gr765e54se5$%F*GHOY&Gf9775d56kc37qkc")
	key, _ := ioutil.ReadFile("../data/test.key")
	ciphertext, _ := cipher.Encrypt(key, plaintext)

	oracle := New(key, plaintext)
	decText := Attack(oracle, ciphertext)

	isEqual := bytes.Equal(plaintext, decText)
	fmt.Println(isEqual)
}