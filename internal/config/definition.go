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
