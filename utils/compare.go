package utils

import (
	"bytes"
	"io/ioutil"
)

func CompareFiles(f1 string, f2 string) bool {
	content1, err := ioutil.ReadFile(f1)
	if err != nil {
		panic(err)
	}
	content2, err := ioutil.ReadFile(f2)
	if err != nil {
		panic(err)
	}

	return bytes.Equal(content1, content2)
}