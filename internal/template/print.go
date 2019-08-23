package template

import (
	"fmt"

	"github.com/logrusorgru/aurora"

	"github.com/julienbreux/baleia/internal/template/files"
)

// Print prints files
func (t *template) Print(diff bool) {
	output := ""
	for stateName, state := range files.SortedStates() {
		outputFile := t.printFile(state, stateName, diff)
		if outputFile != "" {
			output = fmt.Sprintf("%s%s\n", output, outputFile)
		}
	}
	fmt.Print(output[0 : len(output)-1])
}

// printFile prints a file line
func (t *template) printFile(state files.State, stateName string, diff bool) (output string) {
	if t.files.Len(state) == 0 {
		return
	}

	output = fmt.Sprintln(aurora.Yellow(fmt.Sprintf("%s file(s): %d", stateName, t.files.Len(state))))
	for _, f := range t.files.List(state) {
		output = fmt.Sprintf("%s%s\n", output, aurora.White(fmt.Sprintf("▹ %s", f.Path())))
		if diff && state == files.StateChanged {
			d, _ := f.Diff()
			output = fmt.Sprintf("%s%s", output, aurora.Gray(12, d))
		}
	}

	return
}
