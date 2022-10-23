package main

import (
	"fmt"
	"github.com/forgoer/openssl"
)

func main() {
	dst, err := openssl.AesECBEncrypt([]byte("qwdqedfq131"), []byte("ze-ho~&ng!w+a*ng"), openssl.PKCS7_PADDING)
	fmt.Println(dst, err)
}
