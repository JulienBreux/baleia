package template

import (
	"bytes"
	"fmt"
	gotemplate "text/template"

	"github.com/Masterminds/sprig"
	"github.com/julienbreux/baleia/internal/config"
	"github.com/julienbreux/baleia/internal/template/files"
	"github.com/julienbreux/baleia/pkg/file"
)

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
		Option("missingkey=zero").
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
