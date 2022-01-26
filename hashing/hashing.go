package hashing

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func HashPassword(username string, password string, method string) (string, error) {
	switch method {
	case "md5":
		return hashPasswordMD5(username, password)
	case "base64md5":
		return hashPasswordBase64MD5(username, password)
	}
	return "", nil
}

func hashPasswordMD5(username string, password string) (string, error) {
	hash := md5.Sum([]byte(username + password))
	return hex.EncodeToString(hash[:]), nil
}

func hashPasswordBase64MD5(username string, password string) (string, error) {
	hash := md5.Sum([]byte(username + password))
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}
