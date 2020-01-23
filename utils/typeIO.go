package utils

import (
	"crypto/aes"
	"fmt"
)

type IOTYPE int

const (
	HEX = iota
	CHAR
	INTEGER
)

func TypePrint(msg []byte, TYPE IOTYPE) {
	switch TYPE {
	case HEX:
		printHex(msg)
	case CHAR:
		printChar(msg)
	case INTEGER:
		printInteger(msg)
	}
}

func printHex(msg []byte) {
	for i, item := range msg {
		fmt.Printf("%02X ", item)
		if (i + 1) % aes.BlockSize == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func printChar(msg []byte) {
	fmt.Printf("%s\n", string(msg))
}

func printInteger(msg []byte) {
	for i, item := range msg {
		fmt.Printf("%03d ", item)
		if (i + 1) % aes.BlockSize == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}
