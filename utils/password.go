package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// SaltSize defines the size of the salt
const SaltSize = 16

// HashPassword generates a salted hash of the password
func HashPassword(password string) (string, error) {
	// Generate a random salt
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Combine password and salt
	passwordBytes := []byte(password)
	passwordWithSalt := append(passwordBytes, salt...)

	// Calculate the hash
	hash := sha256.Sum256(passwordWithSalt)

	// Encode the hash and salt to base64
	hashSalt := append(hash[:], salt...)
	hashSaltBase64 := base64.StdEncoding.EncodeToString(hashSalt)

	return hashSaltBase64, nil
}

func VerifyPassword(password, hashedPassword string) bool {
	// Decode the hashed password from base64
	hashSalt, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	// Extract the salt from the hash
	salt := hashSalt[len(hashSalt)-SaltSize:]

	// Generate the hash using the provided password and extracted salt
	passwordBytes := []byte(password)
	passwordWithSalt := append(passwordBytes, salt...)
	hash := sha256.Sum256(passwordWithSalt)

	// Concatenate the hash with the extracted salt
	hashWithSalt := append(hash[:], salt...)

	// Encode the concatenated hash and salt to base64
	calculatedHashSalt := base64.StdEncoding.EncodeToString(hashWithSalt)

	// Compare the generated hash with the stored hash
	return calculatedHashSalt == hashedPassword
}
