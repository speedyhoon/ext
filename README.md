# ext

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/ext.svg)](https://pkg.go.dev/github.com/speedyhoon/ext)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/ext)](https://goreportcard.com/report/github.com/speedyhoon/ext)

Helpful functions for checking and modifying file extensions and file suffixes. Plus a list of string constants for commonly used file extensions and file suffixes.

## Example
```go
package main

import "github.com/speedyhoon/ext"

func main() {
	println("favicon" + ext.ICO)
	// Output: favicon.ico
	
	println(ext.EqualFold("index.htM", ext.HTM))
	// Output: true
	
	println(ext.Is("index.htm", ext.HTM))
	// Output: true
	
	println(ext.IsGo("main.go"))
	// Output: true
	
	println(ext.IsGoTest("val_test.go"))
	// Output: true
	
	println(ext.Replace("val_test.go", ext.GoTest, ext.Go))
	// Output: val.go
	
	println(ext.ReplaceFold("val_TEST.GO", ext.GoTest, ext.Go))
	// Output: val.go
}
```
