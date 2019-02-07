package other

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"regexp"
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

// PathDoesNotExist returns true if the path to a folder is not found
func PathDoesNotExist(path string) bool {
	stats, err := os.Stat(path)
	if err != nil || !stats.IsDir() {
		return true
	}
	return false
}

// Match returns true if a match is found
func Match(regx string, toMatch string) bool {
	matched, err := regexp.MatchString(regx, toMatch)
	if err != nil {
		return false
	}
	return matched
}

// FullMatch returns true if the full input strings gets matched
// This function is mostly the same as other.Match but
// this one adds a ^( and )$ to the begin and end of a string
func FullMatch(regx string, toMatch string) bool {
	matched, err := regexp.MatchString("^("+regx+")$", toMatch)
	if err != nil {
		return false
	}
	return matched
}

// FindMatch finds a specific match and returns that specific
// match based on a selector number
func FindMatch(regx string, toMatch string, toSelect int) (string, error) {
	re := regexp.MustCompile(regx)
	out := re.FindStringSubmatch(string(toMatch))
	if len(out) > toSelect {
		return out[toSelect], nil
	}
	if len(out) == 0 {
		return "", errors.New("Nothing matched on toMatch")
	}
	return "", errors.New("Selector is more than the amoud of matches found")
}

// FindAllMatches finds matches in a string and returns them
func FindAllMatches(regx string, toMatch string) [][]string {
	re := regexp.MustCompile(regx)
	return re.FindAllStringSubmatch(string(toMatch), -1)
}

// Replace a string with a regex
func Replace(regx string, toReplace string, Replaceval string) string {
	re := regexp.MustCompile(regx)
	return re.ReplaceAllString(toReplace, Replaceval)
}
