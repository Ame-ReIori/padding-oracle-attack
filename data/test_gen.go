package data

import (
	"io/ioutil"
	"math/rand"
)

func GenTestData(filename string, length int) {
	str := make([]byte, length)
	rand.Read(str)
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		panic(err)
	}
}
