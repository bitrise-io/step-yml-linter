package lint

import "fmt"

// Field ...
type Field interface {
	Line() int
	Column() int
	IsEmpty() bool
	ParentField() Field
	FieldName() string
}

var filePath string

// SetFilePath ...
func SetFilePath(path string) {
	filePath = path
}

// Log ...
func Log(f Field, message string) {
	sss := ""
	for {
		sss = f.FieldName() + ">" + sss
		f = f.ParentField()
		if f == nil {
			break
		}
		if !f.IsEmpty() {
			fmt.Printf("%s:%d:%d: %s %s\n", filePath, f.Line(), f.Column(), sss, message)
		}
	}
}
