package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword
// uses bcrypt to hash an incoming string password for storage
// This is a simple demo implementation
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

// PasswordsMatch
// compares an incoming password hash against the stored hash to validate login requests
func PasswordsMatch(hashedPassword, incomingPassword string) bool {
	byteHash := []byte(hashedPassword)
	byteIncoming := []byte(incomingPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, byteIncoming)
	return err == nil
}
