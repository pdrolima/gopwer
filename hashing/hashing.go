package hashing

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/webmasterdro/gopwer/config"
	"io"
)

func HashPassword(username string, password string) (string, error) {
	switch config.Config("PASSWORD_HASH") {
	case "md5":
		return hashPasswordMD5(username, password)
	case "base64md5":
		return hashPasswordBase64MD5(username, password)
	}
	return "", nil
}

func hashPasswordMD5(username string, password string) (string, error) {
	hash := md5.Sum([]byte(username + password))
	return "0x" + hex.EncodeToString(hash[:]), nil
}

func hashPasswordBase64MD5(username string, password string) (string, error) {
	hash := md5.New()
	io.WriteString(hash, username+password)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

func CheckPassword(username string, password string, hashedPassword string) (bool, error) {
	switch config.Config("PASSWORD_HASH") {
	case "md5":
		return checkPasswordMD5(username, password, hashedPassword), nil
	case "base64md5":
		return checkPasswordBase64MD5(username, password, hashedPassword), nil
	}
	return false, nil
}

func checkPasswordBase64MD5(username string, password string, hashedPassword string) bool {
	hash, _ := hashPasswordBase64MD5(username, password)
	return hash == hashedPassword
}

func checkPasswordMD5(username string, password string, hashedPassword string) bool {
	hash, _ := hashPasswordMD5(username, password)
	return hash == hashedPassword
}
