package main

import (
	"strings"
	"testing"
)

func Test_ReadAndReplace(t *testing.T) {

	cases := []struct{ text, result string }{
		{"asdf|asdf|asdf|asfd|", "asdf-asdf-asdf-asfd-"},
		{"asdf;asdf;asdf;asfd;", "asdf-asdf-asdf-asfd-"},
		{"asdf:asdf:asdf:asfd:", "asdf-asdf-asdf-asfd-"},
	}

	for _, c := range cases {
		result := ReadAndReplace(strings.NewReader(c.text))

		if result != c.result {
			t.Log("result should be "+string(c.result)+", but got ", result)
			t.Fail()
		}
	}

}
