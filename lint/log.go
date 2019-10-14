package lint

import "fmt"

// Log prints linter formatted log output
func Log(file string, line, column int, f string, args ...interface{}) {
	fmt.Printf("%s:%d:%d: %s\n", file, line, column, fmt.Sprintf(f, args...))
}
