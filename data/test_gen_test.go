package data

import "testing"

func TestGenTestData(t *testing.T) {
	GenTestData("./test.plain", 16)
	GenTestData("./test.key", 16)
}
