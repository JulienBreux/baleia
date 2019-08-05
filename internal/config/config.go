package config

import (
	"bytes"
	"text/template"
)

const (
	defaultTemplate = "Dockerfile.tmpl"
)

// New creates a config instance
func New(f string) (Config, error) {
	s, err := fileToSchema(f)
	if err != nil {
		return nil, err
	}

	// Default values
	if s.Template == "" {
		s.Template = defaultTemplate
	}

	return config{
		schema: s,
	}, nil
}

// GetVersion returns the Version
func (c config) GetVersion() string {
	return c.schema.Version
}

// GetTemplate returns the template filename
func (c config) GetTemplate() string {
	return c.schema.Template
}

// LenImages counts the number of images
func (c config) LenImages() int {
	return len(c.schema.Images)
}

// GetImages returns the images
func (c config) GetImages() (imgs []Image) {
	for _, img := range c.schema.Images {
		imgs = append(imgs, c.computeImageValues(img))
	}
	return
}

// computeImageValues computes image values
// FIXME: Before major release, return errors
func (c config) computeImageValues(img schemaImage) Image {
	var v string

	// Field: Maintainers
	if len(img.Maintainers) == 0 {
		img.Maintainers = c.schema.Maintainers
	}

	// Field: Labels
	if len(img.Labels) == 0 {
		img.Labels = c.schema.Labels
	}

	// Field: Arguments
	if len(img.Arguments) == 0 {
		img.Arguments = c.schema.Arguments
	}

	// Variables
	vars := img.SetDefaultVars(c.schema)

	// Field: Name
	if img.Name == "" {
		img.Name = c.schema.Name
	}
	v, _ = c.computeValue(vars, img.Name)
	img.Name = v

	// Field: Base image
	if img.BaseImage == "" {
		img.BaseImage = c.schema.BaseImage
	}
	v, _ = c.computeValue(vars, img.BaseImage)
	img.BaseImage = v

	// Field: Image tag
	if img.ImageTag == "" {
		img.ImageTag = c.schema.ImageTag
	}
	v, _ = c.computeValue(vars, img.ImageTag)
	img.ImageTag = v
	// Field: Output
	if img.Output == "" {
		img.Output = c.schema.Output
	}
	v, _ = c.computeValue(vars, img.Output)
	img.Output = v

	return img
}

// computeValue computes a value with variables
func (c config) computeValue(vars map[string]interface{}, value string) (string, error) {
	outputTpl, err := template.New("output").Parse(value)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := outputTpl.Execute(&buf, vars); err != nil {
		return "", err
	}
	return buf.String(), nil
}
