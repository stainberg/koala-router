package koala

import (
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var  (
	hs = md5.New()
)

func SecurityMD5(src string) string {
	hs.Reset()
	hs.Write([]byte(src))
	return strings.ToLower(hex.EncodeToString(hs.Sum(nil)))
}

func Base64Encode(src string) string {
	bs := base64.URLEncoding.EncodeToString([]byte(src))
	dst := strings.Replace(string(bs), "/", "_", -1)
	dst = strings.Replace(dst, "+", "-", -1)
	dst = strings.Replace(dst, "=", "", -1)
	return dst
}

func Base64Decode(src string) (string) {
	var missing = (4 - len(src)%4) % 4
	src += strings.Repeat("=", missing)
	db, err := base64.URLEncoding.DecodeString(src)
	if err != nil {
		println(err.Error())
	}
	return string(db)
}