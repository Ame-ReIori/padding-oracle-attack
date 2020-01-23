package PaddingOracle

import (
	"bytes"
	"crypto/aes"
)

func GetPad(oracle PaddingOracle, ciphertext []byte) int {
	length := len(ciphertext)
	blockNumber := length/ aes.BlockSize
	processMsg := make([]byte, length)
	copy(processMsg, ciphertext)
	prevBlock := processMsg[aes.BlockSize * (blockNumber - 2): aes.BlockSize *(blockNumber - 1)]

	for i := 0; i < aes.BlockSize; i++ {
		prevBlock[i] = 0
		isSuccess := oracle.Query(processMsg)
		if !isSuccess {
			return aes.BlockSize - i
		}
	}
	return 0
}

func ProcessBlock(block []byte, pad int) {
	delta := make([]byte, aes.BlockSize)
	copy(delta[aes.BlockSize - pad:], bytes.Repeat([]byte{byte((pad + 1) ^ pad)}, pad))
	for i := aes.BlockSize - pad; i < aes.BlockSize; i++ {
		block[i] ^= delta[i]
	}
}

func AttackBlock(oracle PaddingOracle, ciphertext []byte, pad int) []byte {
	out := make([]byte, aes.BlockSize)

	length := len(ciphertext)
	blockNumber := length/ aes.BlockSize
	processMsg := make([]byte, length)
	copy(processMsg, ciphertext)
	prevBlock := processMsg[aes.BlockSize * (blockNumber - 2): aes.BlockSize *(blockNumber - 1)]

	for i := aes.BlockSize - pad - 1; i >= 0; i-- {
		ProcessBlock(prevBlock, aes.BlockSize - 1 - i)
		for j := 0; j != 255; j++ {
			prevBlock[i] ^= byte(j)
			isSuccess := oracle.Query(processMsg)
			if isSuccess {
				out[i] = byte(j ^ (aes.BlockSize - i))
				break
			}
			prevBlock[i] ^= byte(j)
		}
	}
	return out
}

func Attack(oracle PaddingOracle, ciphertext []byte) []byte {
	pad := GetPad(oracle, ciphertext)
	length := len(ciphertext) - pad - aes.BlockSize
	blockNumber := length / aes.BlockSize
	out := make([]byte, length)

	lastBlock := AttackBlock(oracle, ciphertext, pad)
	copy(out[aes.BlockSize * blockNumber:], lastBlock)

	for i := blockNumber - 1; i >= 0; i-- {
		block := AttackBlock(oracle, ciphertext[:aes.BlockSize * (i + 2)], 0)
		copy(out[aes.BlockSize * i:aes.BlockSize * (i + 1)], block)
	}
	return out
}


