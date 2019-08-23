package files

import "sort"

var (
	// States represents the list of states
	States = map[string]State{"Created": StateCreated, "Changed": StateChanged, "Unchanged": StateUnchanged}
)

// SortedStates returns sorted states
func SortedStates() map[string]State {
	names := make([]string, 0, len(States))
	for name := range States {
		names = append(names, name)
	}

	sort.Strings(names)

	ss := make(map[string]State, len(States))
	for _, name := range names {
		ss[name] = States[name]
	}

	return ss
}
