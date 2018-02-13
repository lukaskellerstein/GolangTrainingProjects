package main

import (
	"io"
	"io/ioutil"
)

//BAD WAY - use file
// func ReadAndReplace(f *io.File) string {
// 	bs, _ := ioutil.ReadAll(f)
// 	DO REPLACE LOGIC
// 	return result
// }

//BETTER WAY - use reader
func ReadAndReplace(rdr io.Reader) string {

	bs, _ := ioutil.ReadAll(rdr)

	// DO REPLACE LOGIC
	result := replace(string(bs))

	return result
}

func replace(text string) string {
	newStr := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		switch c := text[i]; c {
		case '|', ';', ':':
			newStr[i] = '-'
		default:
			newStr[i] = c
		}
	}
	return string(newStr)
}
