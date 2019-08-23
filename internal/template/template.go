package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	gotemplate "text/template"

	"github.com/Masterminds/sprig"
	"github.com/logrusorgru/aurora"

	"github.com/julienbreux/baleia/internal/config"
	"github.com/julienbreux/baleia/internal/template/files"
	"github.com/julienbreux/baleia/pkg/file"
)

// template represents the internal template
type template struct {
	config    config.Config
	templates map[string][]byte
	files     files.Files
}

// New creates a new template instance
func New(c config.Config) (Template, error) {
	templates := make(map[string][]byte)

	t := &template{
		config:    c,
		templates: templates,
		files:     files.New(),
	}

	if err := t.loadTemplates(); err != nil {
		return nil, err
	}

	return t, nil
}

// loadTemplates loads the templates content
func (t *template) loadTemplates() error {
	for name, tpl := range t.config.GetTemplates() {
		cts, err := t.loadTemplateContent(tpl)
		if err != nil {
			return err
		}
		t.templates[name] = cts
	}

	return nil
}

// loadTemplateContent loads the template content
func (t *template) loadTemplateContent(template string) ([]byte, error) {
	if !file.Exists(template) {
		return nil, fmt.Errorf("Unable to open the template file '%s'", template)
	}

	cts, err := ioutil.ReadFile(template)
	if err != nil {
		return nil, fmt.Errorf("Unable to read the template file '%s'", template)
	}

	return cts, nil
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

	return t.buildImageTemplate(img, vars)
}

// checkTemplateRef checks based template
func (t *template) checkTemplateRef(img config.Image) (template string, err error) {
	template = img.GetTemplate()
	if img.GetTemplate() == "" {
		template = config.DefaultTemplateRef
	}

	if _, ok := t.config.GetTemplates()[template]; !ok {
		return "", fmt.Errorf("reference \"%s\" doesn't exists", template)
	}

	return
}

// buildImageTemplate builds the image template
func (t *template) buildImageTemplate(img config.Image, vars map[string]interface{}) ([]byte, error) {
	tplRef, err := t.checkTemplateRef(img)
	if err != nil {
		return nil, err
	}

	outputTpl, err := gotemplate.
		New("output").
		Funcs(sprig.GenericFuncMap()).
		Parse(string(t.templates[tplRef]))
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
	output := ""
	for stateName, state := range files.SortedStates() {
		outputFile := t.printFile(state, stateName, diff)
		if outputFile != "" {
			output = fmt.Sprintf("%s%s\n", output, outputFile)
		}
	}
	fmt.Print(output[0 : len(output)-1])
}

// printFile prints a file line
func (t *template) printFile(state files.State, stateName string, diff bool) (output string) {
	if t.files.Len(state) == 0 {
		return
	}

	output = fmt.Sprintln(aurora.Yellow(fmt.Sprintf("%s file(s): %d", stateName, t.files.Len(state))))
	for _, f := range t.files.List(state) {
		output = fmt.Sprintf("%s%s\n", output, aurora.White(fmt.Sprintf("▹ %s", f.Path())))
		if diff && state == files.StateChanged {
			d, _ := f.Diff()
			output = fmt.Sprintf("%s%s", output, aurora.Gray(12, d))
		}
	}

	return
}
