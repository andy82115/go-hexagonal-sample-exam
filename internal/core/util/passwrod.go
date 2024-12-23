package util

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes input password using bcrypt
// HashPassword は、入力パスワードを bcrypt を使ってハッシュする。
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares input password with hashed password
// ComparePassword は入力パスワードとハッシュ化されたパスワードを比較する。
func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
