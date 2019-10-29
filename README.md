# How to test
- `cd to/this/repo && code .`
- edit the linter codebase or the files under `.vscode/extension/src`
- click on `Debug > Start Debugging`

This will open a new vscode window in which the extension will be automatically running.

# Linter rules
Add them under `checks.go`. 

To capture a lint error use `lint.Log` and pass the field as the first argument. This way the lint log will contain the field's line and column numbers.

For example:
```Go
// title
func(s step.YML) {
	if s.Title.Value == "" {
		lint.Log(s.Title, "title is required")
	}
},
```