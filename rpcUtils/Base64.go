package rpcUtils

import "encoding/base64"

func Base64Encode(data []byte) []byte {
	encoding := base64.StdEncoding
	dst := make([]byte, encoding.EncodedLen(len(data)))
	encoding.Encode(dst, data)
	return dst
}

/**
 * base64编码
 */
func Base64Decode(data []byte) []byte {
	encoding := base64.StdEncoding
	dst := make([]byte, encoding.DecodedLen(len(data)))
	n, _ :=encoding.Decode(dst, data)
	return dst[0:n]
}

