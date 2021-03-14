package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// EncryptPassword .
func EncryptPassword(password string, salt string) (encryptPassword string, err error) {
	s := sha256.New()
	_, err = s.Write([]byte(salt + password + salt))
	if err != nil {
		return
	}
	return hex.EncodeToString(s.Sum(nil)), nil
}
