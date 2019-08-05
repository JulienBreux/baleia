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

// NewFile creates a new file instance
func NewFile(path string, state State, content []byte) File {
	return &file{
		path: path,
		s:    state,
		c:    content,
	}
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

// State returns the files's state
func (f *file) State() State {
	return f.s
}
