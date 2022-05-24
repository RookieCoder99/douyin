package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func EncoderSha256(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return string(s)
}
func Md5(data string) string {
	w := md5.New()
	io.WriteString(w, data)
	bydate := w.Sum(nil)
	result := fmt.Sprintf("%x", bydate)
	return result
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
