package cmd

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

const (
	app     = "Baleia"
	author  = "Julien BREUX <julien.breux@gmail.com>"
	website = "https://julienbreux.github.io/baleia/"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version/build info",
	Long:  "Show the version and build information",
	Run:   versionRun,
}

// versionRun represents the run command
func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println(aurora.White(app))
	fmt.Println()
	fmt.Println(aurora.Cyan("Version:  "), aurora.White(version))
	fmt.Println(aurora.Cyan("Commit:   "), aurora.White(commit))
	fmt.Println(aurora.Cyan("Date:     "), aurora.White(date))
	fmt.Println(aurora.Cyan("Author:   "), aurora.White(author))
	fmt.Println()
	fmt.Println(aurora.Cyan("Website:  "), aurora.White(website))

	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
