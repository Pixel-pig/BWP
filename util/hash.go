package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Hash(data []byte) []byte {
	//、对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
func MD5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	psswordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(psswordBytes)
}