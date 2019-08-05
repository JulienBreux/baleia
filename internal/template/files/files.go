package files

// files represents the internal struct of files manager
type files struct {
	files map[State][]File
}

// New creates an instance of Files manager
func New() Files {
	return &files{
		files: make(map[State][]File),
	}
}

// Add add a file to a state collection
func (f *files) Add(path string, s State, c []byte) {
	f.files[s] = append(f.files[s], &file{
		path: path,
		s:    s,
		c:    c,
	})
}

// List returns the list of files of a state collection
func (f *files) List(s State) []File {
	return f.files[s]
}

// Len returns the length of a state collection
func (f *files) Len(s State) int {
	return len(f.files[s])
}

// LenAll returns the lenght of all state collections
func (f *files) LenAll() (i int) {
	for _, s := range States {
		i = i + f.Len(s)
	}
	return
}
