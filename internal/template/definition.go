package template

import (
	"io"
)

// Template represents the interface template
type Template interface {
	// Parse parses files
	Parse() (err error)

	// Write writes files
	Write() (err error)

	// Print prints files changes
	Print(w io.Writer, diff bool)
}
