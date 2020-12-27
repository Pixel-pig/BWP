package utils

import "encoding/base64"

func Base64EnConding(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
