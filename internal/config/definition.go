package config

// Config represents a configuration interface
type Config interface {
	GetVersion() string
	GetTemplate() string
	LenImages() int
	GetImages() []Image
}

// Image represents a configuration image interface
type Image interface {
	GetVars() map[string]string
	GetMaintainers() []string
	GetName() string
	GetLabels() map[string]string
	GetBaseImage() string
	GetImageTag() string
	GetOutput() string
	GetArguments() []string
}

// config represents the internal configuration
type config struct {
	schema *schema
}

// schema represents the configuration schema
type schema struct {
	Version  string        `yaml:"version,omitempty"`
	Template string        `yaml:"template,omitempty"`
	Images   []schemaImage `yaml:"images"`

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

// schemaImage represents the configuration schema images
type schemaImage struct {
	Vars        map[string]string `yaml:"vars,omitempty"`
	Maintainers []string          `yaml:"maintainers,omitempty"`
	Name        string            `yaml:"name"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	BaseImage   string            `yaml:"baseImage,omitempty"`
	ImageTag    string            `yaml:"imageTag,omitempty"`
	Output      string            `yaml:"output,omitempty"`
	Arguments   []string          `yaml:"arguments,omitempty"`
}
