package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	gotemplate "text/template"

	"github.com/julienbreux/baleia/internal/config"
	"github.com/julienbreux/baleia/internal/template/files"
	"github.com/julienbreux/baleia/pkg/file"
	"github.com/logrusorgru/aurora"
)

// template represents the internal template
type template struct {
	template []byte
	files    files.Files
	config   config.Config
}

// New creates a new template instance
func New(c config.Config) (Template, error) {
	if !file.Exists(c.GetTemplate()) {
		return nil, fmt.Errorf("Unable to open the template file '%s'", c.GetTemplate())
	}

	cts, err := ioutil.ReadFile(c.GetTemplate())
	if err != nil {
		return nil, fmt.Errorf("Unable to read the template file '%s'", c.GetTemplate())
	}

	return &template{
		template: cts,
		config:   c,
		files:    files.New(),
	}, nil
}

// Parse parses files
func (t *template) Parse() (err error) {
	for _, img := range t.config.GetImages() {
		c, err := t.computeContent(img)
		if err != nil {
			return err
		}

		t.files.Add(
			files.NewFile(
				img.GetOutput(),
				t.getFileState(img, c),
				c,
			),
		)
	}

	return
}

// getFileState returns the file state
func (t *template) getFileState(img config.Image, c []byte) files.State {
	if file.Exists(img.GetOutput()) {
		if d, _ := file.Compare(img.GetOutput(), c); d {
			return files.StateChanged
		}
		return files.StateUnchanged
	}

	return files.StateCreated
}

// computeContent computes file content
func (t *template) computeContent(img config.Image) ([]byte, error) {
	vars := make(map[string]interface{})
	// TODO: Generate assign
	// Direct variables
	vars["maintainers"] = img.GetMaintainers()
	vars["name"] = img.GetName()
	vars["labels"] = img.GetLabels()
	vars["baseImage"] = img.GetBaseImage()
	vars["imageTag"] = img.GetImageTag()
	vars["arguments"] = img.GetArguments()
	vars["vars"] = img.GetVars()

	outputTpl, err := gotemplate.New("output").Parse(string(t.template))
	if err != nil {
		return []byte{}, err
	}

	var buf bytes.Buffer
	if err := outputTpl.Execute(&buf, vars); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

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

// Print prints files
func (t *template) Print(diff bool) {
	// Changes print
	first := true
	for stateName, state := range files.States {
		if !first {
			fmt.Println("")
		}
		t.printFile(state, stateName, diff)
		if first {
			first = false
		}
	}
}

// printFile prints a file line
func (t *template) printFile(state files.State, stateName string, diff bool) {
	if t.files.Len(state) == 0 {
		return
	}

	fmt.Println(aurora.Yellow(fmt.Sprintf("%s file(s): %d", stateName, t.files.Len(state))))
	for _, f := range t.files.List(state) {
		fmt.Println(aurora.White(fmt.Sprintf("▹ %s", f.Path())))
		if diff && state == files.StateChanged {
			d, _ := f.Diff()
			fmt.Print(aurora.Gray(12, d))
		}
	}
}
