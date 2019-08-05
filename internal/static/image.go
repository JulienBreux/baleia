package static

// ImageTemplate used by init command.
// nolint: gochecknoglobals
const ImageTemplate = `# This is an example of Dockerfile.tmpl.
FROM {{ .baseImage }}

{{- if .maintainers }}

# Maintainers.
{{- range $maintainer := .maintainers }}
LABEL maintainer="{{ $maintainer }}"
{{- end }}
{{- end }}

{{- if .arguments }}

# Arguments.
{{- range $argument := .arguments }}
ARG {{ $argument }}
{{- end }}
{{- end }}

{{- if .labels -}}
# Labels.
{{- range $key, $val := .labels }}
LABEL {{ $key }}="{{ $val }}"
{{- end }}
{{- end }}

`
