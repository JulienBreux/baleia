package static

// ConfigTemplate used by the init command.
// nolint: gochecknoglobals
const ConfigTemplate = `# This is an example of .baleia.yaml.
# Make sure to check the documentation at https://julienbreux.github.io/baleia/

version: "1"                                                                    # Optional - Version of Beleia file used

template: Dockerfile.tmpl                                                       # Optional - Dockerfile template used
maintainers:                                                                    # Optional - List of maintainers
  - Julien BREUX <julien.breux@gmail.com>
name: go                                                                        # Required - Name of the image

baseImage: "{{.name}}:{{.goVersion}}-{{.distrib}}{{.distribVersion}}"           # Required - Based image
imageTag: "{{.goVersion}}-{{.distrib}}{{.distribVersion'}}-{{.Version}}"        # Optional - Image tag
output: "{{.name}}/{{.goVersion}}/{{.distrib}}/{{.distribVersion}}/Dockerfile"  # Required - Output file

vars:                                                                           # Optional - Global variables
  distrib: alpine
  distribVersion: 3.10

images:                                                                         # Required - Images generated
  # Go 1.11 / Alpine 3.10
  - vars:
      goVersion: 1.11

   # Go 1.12 / Alpine 3.10
  - vars:
      goVersion: 1.12
`
