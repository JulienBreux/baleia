package cmd

import (
	"fmt"
	"os"

	"github.com/julienbreux/baleia/internal/config"
	"github.com/julienbreux/baleia/internal/template"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate command is used to generate files",
		Long: `Generate command is used to generate files.

It's recommended to user the --dry option to display your changes.`,
		Run: generateRun,
	}

	diff = false
	dry  = false
)

// generateRun represents the run command
func generateRun(cmd *cobra.Command, args []string) {
	c, err := config.New(configFile)
	if err != nil {
		printError(fmt.Sprintf("Unable to read config file '%s'", configFile), nil)
	}

	// Open template
	t, err := template.New(c)
	if err != nil {
		printError(err.Error(), nil)
	}

	// Parse templates
	if err := t.Parse(); err != nil {
		printError(fmt.Sprintf("Unable to parse template '%s'", c.GetTemplate()), err)
	}

	// Print output
	t.Print(diff)

	// Write templates changed
	if !dry {
		if err = t.Write(); err != nil {
			printError("Unable to write files", nil)
		}
	}

	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolVarP(&dry, "dry", "y", false, "Dry mode to display changes only")
	generateCmd.Flags().BoolVarP(&diff, "diff", "i", false, "Display diff")
}
