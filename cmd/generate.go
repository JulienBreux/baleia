package cmd

import (
	"fmt"
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
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate command is used to generate files",
	Long: `Generate command is used to generate files.

It's recommended to user the --dry option to display your changes.`,
	Run: generateRun,
}

// generateRun represents the run command
func generateRun(cmd *cobra.Command, args []string) {
	debug, _ := cmd.Flags().GetBool("debug")
	f, _ := cmd.Flags().GetString("file")
	diff, _ := cmd.Flags().GetBool("diff")

	c, err := config.New(f)
	if err != nil {
		fmt.Println(aurora.Red(fmt.Sprintf("Unable to read config file '%s'", f)))
		os.Exit(1)
	}

	// if err := c.Validate(); err != nil {
	// 	fmt.Println(aurora.Red(fmt.Sprintf("Config is not valid '%s'", err)))
	// 	os.Exit(1)
	// }

	// Open template
	t, err := template.New(c)
	if err != nil {
		fmt.Println(aurora.Red(err))
		os.Exit(1)
	}

	// Parse templates
	if err := t.Parse(); err != nil {
		fmt.Println(aurora.Red(fmt.Sprintf("Unable to parse template '%s'", c.GetTemplate())))
		if err != nil && debug {
			fmt.Println(aurora.Red(fmt.Sprintf(" ▹ %+v", err)))
		}
		os.Exit(1)
	}

	// Print output
	t.Print(os.Stdout, diff)

	// Exit if only dry
	if d, err := cmd.Flags().GetBool("dry"); d && err == nil {
		os.Exit(0)
	}

	// Write templates changed
	if err = t.Write(); err != nil {
		fmt.Println(aurora.Red(fmt.Sprint("Unable to write files")))
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("file", "f", defaultFile, "File used to work")
	generateCmd.Flags().BoolP("dry", "y", false, "Dry mode to display changes only")
	generateCmd.Flags().BoolP("diff", "i", false, "Display diff")
	generateCmd.Flags().BoolP("debug", "d", false, "Debug mode")
}
