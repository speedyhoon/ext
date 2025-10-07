package ext

import "strings"

// EqualFold determines if the path's extension is the same as ext, ignoring case.
func EqualFold(path, ext string) bool {
	return ext != "" && strings.HasSuffix(strings.ToLower(path), strings.ToLower(ext))
}

// Is returns true if the path ends with a certain file extension or filename suffix.
func Is(path, ext string) bool {
	return ext != "" && strings.HasSuffix(path, ext)
}

// IsGo performs a case-sensitive check if a filepath ends with a `.go` file extension.
// IsGo is shorthand for ext.Is(path, ext.Go)
func IsGo(path string) bool {
	return Is(path, Go)
}

// IsGoTest performs a case-sensitive check if a filepath ends with an `_test.go` file suffix.
// IsGoTest is shorthand for ext.Is(path, ext.GoTest)
func IsGoTest(path string) bool {
	return Is(path, GoTest)
}

// Replace performs a case-sensitive swap of path's file extension.
func Replace(path, old, new string) string {
	if Is(path, old) {
		return path[:len(path)-len(old)] + new
	}

	return path
}

// ReplaceFold swaps a path's file extension ignoring case.
func ReplaceFold(path, old, new string) string {
	if EqualFold(path, old) {
		return path[:len(path)-len(old)] + new
	}

	return path
}
