package files

const (
	// StateCreated represents the state created
	StateCreated State = iota
	// StateChanged represents the state changed
	StateChanged
	// StateUnchanged represents the state unchanged
	StateUnchanged
)

// State represents a file type
type State int

// Files represents the interface of files manager
type Files interface {
	Add(f File)
	List(s State) []File
	Len(s State) int
	LenAll() int
}

// File represent the interface of a file
type File interface {
	Path() string
	Diff() (string, error)
	Content() []byte
	State() State
}
