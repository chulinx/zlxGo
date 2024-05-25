package endecryption

import (
	"fmt"
	"github.com/UnderTreeTech/waterdrop/pkg/utils/xcrypto"
)

// Decrypt Encrypt aes 加解密
func Decrypt(src []byte, secretKey string) (result string, err error) {
	iv := fmt.Sprintf("%x", xcrypto.Hash(secretKey, xcrypto.MD5))
	iv = iv[:16]
	cbc, err := xcrypto.NewAesCbcCrypto(secretKey, []byte(iv)...)
	if err != nil {
		return
	}
	src, err = xcrypto.DecodeString(string(src), xcrypto.Base64)
	if err != nil {
		return
	}
	dst, err := cbc.Decrypt(src, []byte(iv))
	if err != nil {
		return
	}
	result = string(dst)
	return
}

func Encrypt(src []byte, secretKey string) (result string, err error) {

	iv := fmt.Sprintf("%x", xcrypto.Hash(secretKey, xcrypto.MD5))
	iv = iv[:16]
	cbc, err := xcrypto.NewAesCbcCrypto(secretKey, []byte(iv)...)
	if err != nil {
		return
	}
	dst, err := cbc.Encrypt(src, []byte(iv))
	if err != nil {
		return
	}
	result, _ = xcrypto.EncodeToString(dst, xcrypto.Base64)
	return
}
