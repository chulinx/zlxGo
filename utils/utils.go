package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func MakeUrl(url string) string {
	if strings.Contains(url, "http://") || strings.Contains(url, "https://") {
		return url
	}
	return fmt.Sprintf("http://%s", url)
}

func Md5(input string) string {
	c := md5.New()
	c.Write([]byte(input))
	bytes := c.Sum(nil)
	return hex.EncodeToString(bytes)
}
