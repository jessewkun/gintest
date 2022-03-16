package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"gintest/config"
)

type AesCbc struct {
	Key string
	Iv  string
}

// md5
func Md5X(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// cbc加密
func (ac *AesCbc) Encode(data string) (string, *APIException) {
	_data := []byte(data)
	_key := []byte(ac.Key)
	_iv := []byte(ac.Iv)

	_data = ac.PKCS7Padding(_data)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", NewAPIException(config.ERROR_NEWCIPHER_FAIL, err)
	}
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	return base64.StdEncoding.EncodeToString(_data), nil
}

// cbc解密
func (ac *AesCbc) Decode(data string) (string, *APIException) {
	_data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", NewAPIException(config.ERROR_DECOCDE_FAIL, err)
	}
	_key := []byte(ac.Key)
	_iv := []byte(ac.Iv)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", NewAPIException(config.ERROR_NEWCIPHER_FAIL, err)
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	_data = ac.PKCS7UnPadding(_data)

	return string(_data), nil
}

func (ac *AesCbc) PKCS7Padding(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func (ac *AesCbc) PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
