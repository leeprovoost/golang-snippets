package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func runMd5() {
	fmt.Println(GetMD5Hash("Hello World"))
}

// Create an MD5 hash based on a string
// See: http://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang
func GetMD5Hash(text string) string {
	hashText := md5.New()
	hashText.Write([]byte(text))
	return hex.EncodeToString(hashText.Sum(nil))
}
