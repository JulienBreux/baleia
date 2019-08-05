package config

// GetMaintainers returns maintainers
func (i schemaImage) GetMaintainers() []string {
	return i.Maintainers
}

// GetName returns name
func (i schemaImage) GetName() string {
	return i.Name
}

// GetLabels returns labels
func (i schemaImage) GetLabels() map[string]string {
	return i.Labels
}

// GetBaseImage returns base image
func (i schemaImage) GetBaseImage() string {
	return i.BaseImage
}

// GetImageTag returns image tag
func (i schemaImage) GetImageTag() string {
	return i.ImageTag
}

// GetVars returns variables
func (i schemaImage) GetVars() map[string]string {
	return i.Vars
}

// GetOutput returns output
func (i schemaImage) GetOutput() string {
	return i.Output
}

// GetArguments returns arguments
func (i schemaImage) GetArguments() []string {
	return i.Arguments
}
