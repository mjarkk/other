package other

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

// NewErr adds extra information to a error
// This is handy if you want to add a bit of a context to your errors
func NewErr(prefix string, actualError error) error {
	return errors.New(prefix + ": " + actualError.Error())
}

// GetSha1 returns the sha1 hash of the inputted bytes as string
func GetSha1(input []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(input))
}
