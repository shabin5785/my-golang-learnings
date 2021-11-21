package gobyex

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func Encode() {

	s := "hash this string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)

	base64Enc()

}

func base64Enc() {

	data := "asd28asdk23@dwasd"
	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)

	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(string(dec))

}
