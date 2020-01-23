package utils

import (
	"fmt"
	"testing"
)

func TestTypePrint(t *testing.T) {
	msg := []byte("dfaiosdufniasdufasdufcsdfucmads")
	TypePrint(msg, HEX)
	TypePrint(msg, CHAR)
	TypePrint(msg, INTEGER)
}

func TestCompareFiles(t *testing.T) {
	file1 := "./test/data1.txt"
	file2 := "./test/data2.txt"
	file3 := "./test/data3.txt"
	fmt.Println(CompareFiles(file1, file2))
	fmt.Println(CompareFiles(file1, file3))
}