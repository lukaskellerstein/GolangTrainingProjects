package main

import (
	"io"
	"io/ioutil"
	"regexp"
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

var specialChars = regexp.MustCompile(`[|;:]`)

func replace(text string) string {
	return specialChars.ReplaceAllString(text, "-")
}
