package ext_test

import (
	"fmt"
	"github.com/speedyhoon/ext"
)

func ExampleEqualFold() {
	fmt.Println(ext.EqualFold("my/path/index.htM", ext.HTM))
	// Output: true
}

func ExampleIs() {
	fmt.Println(ext.Is("my/path/index.htm", ext.HTM))
	// Output: true
}

func ExampleIsGo() {
	fmt.Println(ext.IsGo("my/path/main.go"))
	// Output: true
}

func ExampleIsGoTest() {
	fmt.Println(ext.IsGoTest("my/path/val_test.go"))
	// Output: true
}

func ExampleReplace() {
	fmt.Println(ext.Replace("my/path/val_test.go", ext.GoTest, ext.Go))
	// Output: my/path/val.go
}

func ExampleReplaceFold() {
	fmt.Println(ext.ReplaceFold("my/path/val_TEST.GO", ext.GoTest, ext.Go))
	// Output: my/path/val.go
}
