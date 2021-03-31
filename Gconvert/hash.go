package Gconvert

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Md5(data string) string {
	ret := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	return ret
}
func Sha1(data string) string {
	ret := fmt.Sprintf("%x", sha1.Sum([]byte(data)))
	return ret
}
func Sha256(data string) string {
	ret := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	return ret
}
func Sha512(data string) string {
	ret := fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
	return ret
}
