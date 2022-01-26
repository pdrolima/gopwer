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
		return HashPasswordMD5(username, password)
	case "base64md5":
		return HashPasswordBase64MD5(username, password)
	}
	return "", nil
}

func HashPasswordMD5(username string, password string) (string, error) {
	hash := md5.Sum([]byte(username + password))
	return "0x" + hex.EncodeToString(hash[:]), nil
}

func HashPasswordBase64MD5(username string, password string) (string, error) {
	hash := md5.New()
	io.WriteString(hash, username+password)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

func CheckPassword(username string, password string, hashedPassword string) (bool, error) {
	switch config.Config("PASSWORD_HASH") {
	case "md5":
		return CheckPasswordMD5(username, password, hashedPassword), nil
	case "base64md5":
		return CheckPasswordBase64MD5(username, password, hashedPassword), nil
	}
	return false, nil
}

func CheckPasswordBase64MD5(username string, password string, hashedPassword string) bool {
	hash, _ := HashPasswordBase64MD5(username, password)
	return hash == hashedPassword
}

func CheckPasswordMD5(username string, password string, hashedPassword string) bool {
	hash, _ := HashPasswordMD5(username, password)
	return hash == hashedPassword
}
