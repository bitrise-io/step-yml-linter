package lint

import "fmt"

// Field ...
type Field interface {
	Line() int
	Column() int
}

var filePath string

// SetFilePath ...
func SetFilePath(path string) {
	filePath = path
}

// Log ...
func Log(f Field, message string) {
	fmt.Printf("%s:%d:%d: %s\n", filePath, f.Line(), f.Column(), message)
}
