package crypto

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func VerifyPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return false, err
	}
	return err == nil, nil
}
