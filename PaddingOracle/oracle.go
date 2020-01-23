package PaddingOracle

import (
	"bytes"
	"padding_oracle_attack/cipher"
)

type PaddingOracle struct {
	key    []byte
	banMsg []byte
}

func New(key []byte, plaintext []byte) PaddingOracle {
	ciphertext, _ := cipher.Encrypt(key, plaintext)
	return PaddingOracle{
		key:    key,
		banMsg: ciphertext,
	}
}

func (oracle *PaddingOracle) Query(msg []byte) bool {
	if bytes.Equal(msg, oracle.banMsg) {
		panic("Cannot query the secret.")
	}
	_, err := cipher.Decrypt(oracle.key, msg)
	if err != nil {
		return false
	} else {
		return true
	}
}
