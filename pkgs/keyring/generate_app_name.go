package keyring

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os/user"
)

func GenerateAppName(appName string, user *user.User) string {
	hashingString := fmt.Sprintf("%s-%s",appName, user.Uid)
	hash := sha256.Sum256([]byte(hashingString))
	hashedString := hex.EncodeToString(hash[:])
	return hashedString
}
