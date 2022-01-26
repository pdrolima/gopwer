package hashing_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/webmasterdro/gopwer/entity"
	"github.com/webmasterdro/gopwer/hashing"
	"testing"
)

func TestHashPasswordMD5(t *testing.T) {

	u := entity.User{
		Name:   "test",
		Passwd: "test",
	}

	hash, err := hashing.HashPasswordMD5(u.Name, u.Passwd)
	assert.Equal(t, err, nil)
	assert.Equal(t, hash, "0x05a671c66aefea124cc08b76ea6d30bb")
}

func TestCheckPasswordBase64MD5(t *testing.T) {

	u := entity.User{
		Name:   "test",
		Passwd: "test",
	}

	hash, err := hashing.HashPasswordBase64MD5(u.Name, u.Passwd)
	assert.Equal(t, err, nil)
	assert.Equal(t, hash, "BaZxxmrv6hJMwIt26m0wuw==")
}
