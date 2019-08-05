package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/julienbreux/baleia/internal/config"
	"github.com/julienbreux/baleia/internal/template"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	defaultFile = ".baleia.yaml"
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

	debug = false
	file  = defaultFile
	diff  = false
	dry   = false
)

// generateRun represents the run command
func generateRun(cmd *cobra.Command, args []string) {
	c, err := config.New(file)
	if err != nil {
		printError(os.Stderr, fmt.Sprintf("Unable to read config file '%s'", file), nil)
	}

	// Open template
	t, err := template.New(c)
	if err != nil {
		printError(os.Stderr, err.Error(), nil)
	}

	// Parse templates
	if err := t.Parse(); err != nil {
		printError(os.Stderr, fmt.Sprintf("Unable to parse template '%s'", c.GetTemplate()), err)
	}

	// Print output
	t.Print(os.Stdout, diff)

	// Write templates changed
	if !dry {
		if err = t.Write(); err != nil {
			printError(os.Stdout, "Unable to write files", nil)
		}
	}

	os.Exit(0)
}

// printError prints an error O_o
func printError(w io.Writer, str string, err error) {
	fmt.Fprintln(w, aurora.Red(str))
	if err != nil && debug {
		fmt.Fprintln(w, aurora.Red(fmt.Sprintf(" ▹ %+v", err)))
	}
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&file, "file", "f", defaultFile, "File used to work")
	generateCmd.Flags().BoolVarP(&dry, "dry", "y", false, "Dry mode to display changes only")
	generateCmd.Flags().BoolVarP(&diff, "diff", "i", false, "Display diff")
	generateCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
}
