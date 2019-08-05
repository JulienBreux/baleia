package files

import (
	filepkg "github.com/julienbreux/baleia/pkg/file"
)

// file represents the struct of a file
type file struct {
	path string
	s    State
	c    []byte
}

// Path returns the file's path
func (f *file) Path() string {
	return f.path
}

// Diff returns the file's diff content
func (f *file) Diff() (string, error) {
	return filepkg.Diff(f.path, f.c)
}

// Content returns the file's content
func (f *file) Content() []byte {
	return f.c
}
