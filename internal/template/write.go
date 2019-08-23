package template

import (
	"github.com/julienbreux/baleia/internal/template/files"
	"github.com/julienbreux/baleia/pkg/file"
)

// Write writes the files
func (t *template) Write() (err error) {
	for _, s := range []files.State{files.StateCreated, files.StateChanged} {
		if err := t.writeFilesByState(s); err != nil {
			return err
		}
	}

	return nil
}

// writeFilesByStates writes all the files by state
func (t *template) writeFilesByState(s files.State) error {
	for _, f := range t.files.List(s) {
		if _, err := file.Write(f.Path(), f.Content()); err != nil {
			return err
		}
	}

	return nil
}
