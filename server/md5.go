package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func main() {
	h := md5.New()
	h.Write([]byte("film_discuss_"))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
}
