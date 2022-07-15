package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	nByte := []byte{}

	for i := len(sd); i > 0; i-- {
		nByte = append(nByte, sd[i-1])
	}

	whatIsIt = string(nByte)
	fmt.Println(whatIsIt)
}
