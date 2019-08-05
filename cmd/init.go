package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/julienbreux/baleia/internal/static"
	filepkg "github.com/julienbreux/baleia/pkg/file"
)

const (
	defaultImageFile = "Dockerfile.tmpl"
)

var (
	// initCmd represents the init command
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize templates",
		Long:  "This command initialize templates",
		Run:   initRun,
	}

	initImageFile  = true
	initConfigFile = true
	imageFile      = defaultImageFile
)

// initRun represents the run command
func initRun(cmd *cobra.Command, args []string) {
	generateFile(initImageFile, configFile, []byte(static.ConfigTemplate))
	generateFile(initConfigFile, imageFile, []byte(static.ImageTemplate))

	os.Exit(0)
}

// Generate file
func generateFile(i bool, f string, c []byte) {
	if i {
		if filepkg.Exists(f) {
			printError(os.Stderr, fmt.Sprintf("File '%s' already exists", f), nil)
		}

		if _, err := filepkg.Write(f, c); err != nil {
			printError(os.Stderr, fmt.Sprintf("Unable to create '%s' file", f), err)
		}

		printSuccess(os.Stdout, fmt.Sprintf("File '%s' successfully created", f))
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&configFile, "image-file", "m", defaultConfigFile, "Image file")
	initCmd.Flags().StringVarP(&imageFile, "config-file", "o", defaultImageFile, "Config file")
	initCmd.Flags().BoolVarP(&initImageFile, "init-image-file", "a", true, "Initialize image file")
	initCmd.Flags().BoolVarP(&initConfigFile, "init-config-file", "g", true, "Initialize config file")
}
