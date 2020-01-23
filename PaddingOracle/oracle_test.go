package PaddingOracle

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"padding_oracle_attack/utils"
	"testing"
)

func TestNew(t *testing.T) {
	plaintext := []byte("aaa")
	key, _ := hex.DecodeString("0123456789abcdeffedcba9876543210")
	oracle := New(key, plaintext)
	utils.TypePrint(oracle.key, utils.HEX)
	utils.TypePrint(oracle.banMsg, utils.HEX)
}

func TestPaddingOracle_Query(t *testing.T) {
	plaintext := []byte("abcdefghijklmnopqrstuvwxyz")
	key, _ := hex.DecodeString("0123456789abcdeffedbca9876543210")
	oracle := New(key, plaintext)
	utils.TypePrint(oracle.banMsg, utils.HEX)
	utils.TypePrint(oracle.banMsg, utils.INTEGER)
	challenge := oracle.banMsg
	for i := 0; i < len(challenge); i++ {
		fmt.Printf("%03d ", challenge[i])
		if (i + 1) % aes.BlockSize == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
	for j := 16; j < 32; j++ {
		challenge[j] = uint8(115)
		isSuccess := oracle.Query(challenge)
		fmt.Println(isSuccess)
	}
}
