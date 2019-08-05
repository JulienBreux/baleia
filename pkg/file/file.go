package file

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pmezard/go-difflib/difflib"
)

const (
	// ModeUserPerm represents the permission user
	ModeUserPerm os.FileMode = 0755
)

// Exists returns if file exists
func Exists(f string) bool {
	if _, err := os.Stat(f); err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

// Compare compares file content to buffer content
func Compare(f string, c []byte) (bool, error) {
	if !Exists(f) {
		return true, fmt.Errorf("File '%s' does not exists", f)
	}

	cts, err := ioutil.ReadFile(f)
	if err != nil {
		return true, fmt.Errorf("Unable to read file '%s'", f)
	}

	return bytes.Compare(c, cts) != 0, nil
}

// Diff compares file content to buffer content and exporte diff
func Diff(f string, c []byte) (string, error) {
	if !Exists(f) {
		return "", fmt.Errorf("File '%s' does not exists", f)
	}

	cts, err := ioutil.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("Unable to read file '%s'", f)
	}

	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(cts)),
		B:        difflib.SplitLines(string(c)),
		FromFile: "Original",
		ToFile:   "Current",
	}
	return difflib.GetUnifiedDiffString(diff)
}

// Write writes and creates file and path
func Write(f string, c []byte) (bool, error) {
	path := filepath.Dir(f)
	// panic(path)
	if !Exists(path) {
		if err := os.MkdirAll(path, ModeUserPerm); err != nil {
			return false, err
		}
	}

	if err := ioutil.WriteFile(f, c, ModeUserPerm); err != nil {
		return false, err
	}

	return true, nil
}
