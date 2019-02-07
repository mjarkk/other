package other

import (
	"errors"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErr(t *testing.T) {
	err := NewErr("prefix", errors.New("test error message"))
	assert.Equal(t, "prefix: test error message", err.Error())
}

func TestGetSha1(t *testing.T) {
	assert.Equal(
		t,
		"57162c5f0c9d65d5968484699f0b9858935be699",
		GetSha1([]byte("a very long string")),
	)
}

func TestPathDoesNotExist(t *testing.T) {
	dir, err := user.Current()
	if err == nil {
		assert.False(t, PathDoesNotExist(dir.HomeDir))
	}
	assert.True(t, PathDoesNotExist("/some/long/path/that/probebly/does/not/exsists"))
}

func TestMatch(t *testing.T) {
	assert.True(t, Match(`[a-z]{3}`, "123abc321"))
	assert.False(t, Match(`[a-z]{3}`, "123b321"))
	assert.False(t, Match(`[`, "123b321"))
}

func TestFullMatch(t *testing.T) {
	assert.False(t, FullMatch(`[a-z]{3}`, "123abc321"))
	assert.False(t, FullMatch(`[`, ""))
	assert.True(t, FullMatch(`[0-9]{3}[a-z]{3}[0-9]{3}`, "123abc321"))
}

func TestFindMatch(t *testing.T) {
	matched, err := FindMatch(`[a-z]{3}`, "abc def ghi", 0)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	assert.Equal(t, "abc", matched)

	_, err = FindMatch(`[a-z]{3}`, "", 0)
	if err == nil {
		assert.FailNow(t, "Expected a error")
	}

	_, err = FindMatch(`[a-z]{3}`, "abc", 1)
	if err == nil {
		assert.FailNow(t, "Expected out of range error")
	}
}

func TestFindAllMatches(t *testing.T) {
	assert.Len(t, FindAllMatches(`[a-z]{3}`, "abc def ghi"), 3)
	assert.Len(t, FindAllMatches(`[0-9]{3}`, "abc def ghi"), 0)
}

func TestReplace(t *testing.T) {
	assert.Equal(t, "foo baz", Replace(`bar`, "foo bar", "baz"))
}
