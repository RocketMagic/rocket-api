package util

import (
	"fmt"
	"hash"
	"log"
)

func Hash(s string, h hash.Hash) string {
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Fatalln("hash fail", err.Error())
	}

	b := h.Sum([]byte(""))

	return fmt.Sprintf("%x", b) // %x - 接受一个数字并将其转化为十六进制数格式,使用小写字母
}
