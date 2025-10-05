# ext

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/ext.svg)](https://pkg.go.dev/github.com/speedyhoon/ext)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/ext)](https://goreportcard.com/report/github.com/speedyhoon/ext)

Go string constants for commonly used file extension and file suffix constants.

```go
package main

import "github.com/speedyhoon/ext"

func main() {
	print("favicon" + ext.ICO)
}
```
Prints:<br>
`favicon.ico`
