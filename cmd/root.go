package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	defaultConfigFile = ".baleia.yaml"
)

var (
	debug      = false
	configFile = defaultConfigFile

	version = "dev"
	commit  = "dev"
	date    = "n/a"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "baleia",
	Short: "Baleia is a tiny CLI to generate docker images.",
	Long: `Baleia is a tiny CLI to generate docker images.

It's very useful to generate a big repository of images.`,
	Version: version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", defaultConfigFile, "Configuration file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// printError prints an error O_o
func printError(w io.Writer, str string, err error) {
	fmt.Fprintln(w, aurora.Red(str))
	if err != nil && debug {
		fmt.Fprintln(w, aurora.Red(fmt.Sprintf(" ▹ %+v", err)))
	}
	os.Exit(1)
}
