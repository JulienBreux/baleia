package config

// Config represents a configuration interface
type Config interface {
	GetVersion() string
	GetTemplates() map[string]string
	LenImages() int
	GetImages() []Image
}

// Image represents a configuration image interface
type Image interface {
	GetVars() map[string]string
	GetTemplate() string
	GetMaintainers() []string
	GetName() string
	GetLabels() map[string]string
	GetBaseImage() string
	GetImageTag() string
	GetOutput() string
	GetArguments() []string
}
