package cmd

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
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
	fmt.Print(aurora.Cyan("Version:  %s"), aurora.White(version))
	fmt.Print(aurora.Cyan("Commit:   %s"), aurora.White(commit))
	fmt.Print(aurora.Cyan("Date:     %s"), aurora.White(date))

	os.Exit(0)
}
