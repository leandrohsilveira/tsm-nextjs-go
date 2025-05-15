package crypto_test

import (
	"testing"
	"tsm/crypto"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "123456"

	hash, err := crypto.HashPassword(password)

	assert.Nil(t, err, "password hash should not fail")
	assert.NotEqual(t, "", hash, "password hash should not be an empty string")

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.Nil(t, err, "hashed password should br verifiable")
}

func TestVerifyPasswordMatch(t *testing.T) {
	password := "123456"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	assert.Nil(t, err, "password generation should not fail")

	matched, err := crypto.VerifyPassword(password, string(hash))

	assert.Nil(t, err, "password verification should not fail")
	assert.True(t, matched, "password verification should match")
}

func TestVerifyPasswordMismatch(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("987654321"), bcrypt.MinCost)

	assert.Nil(t, err, "password generation should not fail")

	matched, err := crypto.VerifyPassword("123456", string(hash))

	assert.Nil(t, err, "password verification should not fail")
	assert.False(t, matched, "password verification should not match")
}
