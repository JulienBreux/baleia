package template

import (
	"fmt"
	"io/ioutil"

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
