package config

import (
	"bytes"
	"text/template"
)

const (
	// DefaultTemplateRef returns the default template reference
	DefaultTemplateRef = "default"

	defaultTemplate = "Dockerfile.tmpl"
)

// config represents the internal configuration
type config struct {
	schema *schema
}

// schema represents the configuration schema
type schema struct {
	Version   string            `yaml:"version,omitempty"`
	Images    []schemaImage     `yaml:"images"`
	Templates map[string]string `yaml:"templates,omitempty"`

	// copy: schemaImage
	Vars        map[string]string `yaml:"vars,omitempty"`
	Maintainers []string          `yaml:"maintainers,omitempty"`
	Name        string            `yaml:"name"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	BaseImage   string            `yaml:"baseImage,omitempty"`
	ImageTag    string            `yaml:"imageTag,omitempty"`
	Output      string            `yaml:"output,omitempty"`
	Arguments   []string          `yaml:"arguments,omitempty"`
}

// New creates a config instance
func New(f string) (Config, error) {
	s, err := fileToSchema(f)
	if err != nil {
		return nil, err
	}

	// Default values
	if len(s.Templates) == 0 {
		s.Templates[DefaultTemplateRef] = defaultTemplate
	}

	return config{
		schema: s,
	}, nil
}

// GetVersion returns the Version
func (c config) GetVersion() string {
	return c.schema.Version
}

// GetTemplates returns the templates filenames
func (c config) GetTemplates() map[string]string {
	return c.schema.Templates
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
	// Variables
	vars := img.SetDefaultVars(c.schema)

	// Default fields
	img.Maintainers = c.defaultStringSlice(img.Maintainers, c.schema.Maintainers)
	img.Labels = c.defaultStringMap(img.Labels, c.schema.Labels)
	img.Arguments = c.defaultStringSlice(img.Arguments, c.schema.Arguments)
	img.Name = c.defaultStringWithVars(img.Name, c.schema.Name, vars)
	img.BaseImage = c.defaultStringWithVars(img.BaseImage, c.schema.BaseImage, vars)
	img.ImageTag = c.defaultStringWithVars(img.ImageTag, c.schema.ImageTag, vars)
	img.Output = c.defaultStringWithVars(img.Output, c.schema.Output, vars)

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

// defaultStringWithVars returns a string or default
func (c config) defaultStringWithVars(i, d string, vars map[string]interface{}) string {
	if i == "" {
		i = d
	}

	v, err := c.computeValue(vars, i)
	if err != nil {
		return i
	}

	return v
}

// defaultStringSlice returns a string slice or default
func (c config) defaultStringSlice(i, d []string) []string {
	if len(i) == 0 {
		return d
	}

	return i
}

// defaultStringMap returns a string map or default
func (c config) defaultStringMap(i, d map[string]string) map[string]string {
	if len(i) == 0 {
		return d
	}

	return i
}
