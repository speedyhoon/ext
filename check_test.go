package ext_test

import (
	"fmt"
	"github.com/speedyhoon/ext"
	"testing"
)

func Test_EqualFold(t *testing.T) {
	var tests = []struct {
		s, ext  string
		want    bool
		swapped bool
	}{
		{"abc", "abc", true, true},
		{"ABcd", "ABcd", true, true},
		{"123abc", "123ABC", true, true},
		{"αβδ", "ΑΒΔ", false, false},
		{"abc", "xyz", false, false},
		{"abc", "XYZ", false, false},
		{"abcdefghijk", "abcdefghijX", false, false},
		{"abcdefghijk", "abcdefghij\u212A", false, false},
		{"abcdefghijK", "abcdefghij\u212A", false, false},
		{"abcdefghijkz", "abcdefghij\u212Ay", false, false},
		{"abcdefghijKz", "abcdefghij\u212Ay", false, false},
		{"1", "2", false, false},
		{"utf-8", "US-ASCII", false, false},
		{"😀✏️😜⚙️↘️🔗😆🏴‍🏁☠️🫥☹️", "☠️🫥☹️", true, false},
		{"🏎️🏍️🚗🛻🚲🛴🛹", "🏎️🏍️🚗🛻🚲🛴🛹", true, true},
		{"112211", "211", true, false},
	}
	for _, tt := range tests {
		if out := ext.EqualFold(tt.s, tt.ext); out != tt.want {
			t.Errorf("EqualFold(%#q, %#q) = %v, want: %v", tt.s, tt.ext, out, tt.want)
		}
		if out := ext.EqualFold(tt.ext, tt.s); out != tt.swapped {
			t.Errorf("EqualFold(%#q, %#q) = %v, swapped: %v", tt.ext, tt.s, out, tt.swapped)
		}
	}
}

func TestIs(t *testing.T) {
	tests := map[string]struct {
		path string
		ext  string
		want bool
	}{
		"empty":       {path: "", ext: "", want: false},
		"no path":     {path: "", ext: ext.JSON, want: false},
		"no ext":      {path: "index.htm", ext: "", want: false},
		"only ext":    {path: ".htm", ext: ext.HTM, want: true},
		"missing ext": {path: "README", ext: ext.MD, want: false},
		"readme":      {path: "README.md", ext: ext.MD, want: true},
		"filename":    {path: "main.css", ext: ext.CSS, want: true},
		"file suffix": {path: "hello_test", ext: ext.Go, want: false},
		"path":        {path: "C:\\Users\\Foo\\favicon", ext: ext.Go, want: false},
		"filepath":    {path: "C:\\Users\\Foo\\favicon.ico", ext: ext.ICO, want: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ext.Is(tt.path, tt.ext); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAny(t *testing.T) {
	tests := []struct {
		path       string
		extensions []string
		want       bool
	}{
		{path: "", extensions: nil, want: false},
		{path: "", extensions: []string{ext.Go}, want: false},
		{path: "name.go", extensions: []string{ext.Go}, want: true},
		{path: "/name.go", extensions: []string{ext.Go}, want: true},
		{path: "dir/name.go", extensions: []string{ext.Go}, want: true},
		{path: "name.GO", extensions: []string{ext.Go}, want: false},
		{path: "/name.GO", extensions: []string{ext.Go}, want: false},
		{path: "dir/name.GO", extensions: []string{ext.Go}, want: false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			if got := ext.IsAny(tt.path, tt.extensions...); got != tt.want {
				t.Errorf("IsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnyFold(t *testing.T) {
	tests := []struct {
		path       string
		extensions []string
		want       bool
	}{
		{path: "", extensions: nil, want: false},
		{path: "", extensions: []string{ext.Go}, want: false},
		{path: "name.go", extensions: []string{ext.Go}, want: true},
		{path: "/name.go", extensions: []string{ext.Go}, want: true},
		{path: "dir/name.go", extensions: []string{ext.Go}, want: true},
		{path: "name.GO", extensions: []string{ext.Go}, want: true},
		{path: "/name.GO", extensions: []string{ext.Go}, want: true},
		{path: "dir/name.GO", extensions: []string{ext.Go}, want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			if got := ext.IsAnyFold(tt.path, tt.extensions...); got != tt.want {
				t.Errorf("IsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGo(t *testing.T) {
	tests := map[string]struct {
		path string
		ext  string
		want bool
	}{
		"empty":       {path: "", want: false},
		"only ext":    {path: ".go", want: true},
		"suffix":      {path: "_test.go", want: true},
		"no ext":      {path: "hello", want: false},
		"missing ext": {path: "hello_test", want: false},
		"filename":    {path: "hello.go", want: true},
		"path":        {path: "C:\\Users\\Foo\\hello", want: false},
		"filepath":    {path: "C:\\Users\\Foo\\hello.go", want: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ext.IsGo(tt.path); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGoTest(t *testing.T) {
	tests := map[string]struct {
		path string
		want bool
	}{
		"file suffix":     {path: "hello_test.go", want: true},
		"path suffix":     {path: "C:\\Users\\Foo\\hello_test.go", want: true},
		"filepath suffix": {path: "C:\\Users\\Foo\\hello_test.go", want: true},
		"only suffix":     {path: "_test.go", want: true},
		"no suffix":       {path: "hello", want: false},

		"empty":       {path: "", want: false},
		"only ext":    {path: ".go", want: false},
		"no ext":      {path: "hello", want: false},
		"missing ext": {path: "hello", want: false},
		"filename":    {path: "hello.go", want: false},
		"path":        {path: "C:\\Users\\Foo\\hello", want: false},
		"filepath":    {path: "C:\\Users\\Foo\\hello.go", want: false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ext.IsGoTest(tt.path); got != tt.want {
				t.Errorf("IsGoTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tests := map[string]struct {
		path string
		old  string
		new  string
		want string
	}{
		"empty":      {path: "", old: "", new: "", want: ""},
		"file":       {path: "file.go", old: ext.Go, new: ext.GoTest, want: "file_test.go"},
		"base":       {path: "file.go", old: ext.Go, new: "", want: "file"},
		"none":       {path: "file.go", old: "", new: ext.HTM, want: "file.go"},
		"file upper": {path: "FILE.GO", old: ext.Go, new: ext.GoTest, want: "FILE.GO"},
		"ext upper":  {path: "file.GO", old: ext.Go, new: ext.GoTest, want: "file.GO"},
		"ext cap":    {path: "file.Go", old: ext.Go, new: ext.GoTest, want: "file.Go"},
		"ext mixed":  {path: "file.gO", old: ext.Go, new: ext.GoTest, want: "file.gO"},
		"path":       {path: "/project/file.go", old: ext.Go, new: ext.GoTest, want: "/project/file_test.go"},
		"path cap":   {path: "/project/file.Go", old: ext.Go, new: ext.GoTest, want: "/project/file.Go"},
		"path mixed": {path: "/project/file.gO", old: ext.Go, new: ext.GoTest, want: "/project/file.gO"},
		"path upper": {path: "/PROJECT/FILE.GO", old: ext.Go, new: ext.GoTest, want: "/PROJECT/FILE.GO"},
		// File suffixes.
		"x-file":       {path: "file_test.go", old: ext.GoTest, new: ext.Go, want: "file.go"},
		"x-base":       {path: "file_test.go", old: ext.GoTest, new: "", want: "file"},
		"x-none":       {path: "file_test.go", old: "", new: ext.HTM, want: "file_test.go"},
		"x-file upper": {path: "FILE_TEST.GO", old: ext.GoTest, new: ext.Go, want: "FILE_TEST.GO"},
		"x-ext upper":  {path: "file_TEST.GO", old: ext.GoTest, new: ext.Go, want: "file_TEST.GO"},
		"x-ext cap":    {path: "file_Test.Go", old: ext.GoTest, new: ext.Go, want: "file_Test.Go"},
		"x-ext mixed":  {path: "file_tEsT.gO", old: ext.GoTest, new: ext.Go, want: "file_tEsT.gO"},
		"x-path":       {path: "/project/file_test.go", old: ext.GoTest, new: ext.Go, want: "/project/file.go"},
		"x-path cap":   {path: "/project/file_Test.Go", old: ext.GoTest, new: ext.Go, want: "/project/file_Test.Go"},
		"x-path mixed": {path: "/project/file_tEsT.gO", old: ext.GoTest, new: ext.Go, want: "/project/file_tEsT.gO"},
		"x-path upper": {path: "/PROJECT/FILE_TEST.GO", old: ext.GoTest, new: ext.Go, want: "/PROJECT/FILE_TEST.GO"},
		"x-path semi":  {path: "/PROJECT/FILE_test.GO", old: ext.GoTest, new: ext.Go, want: "/PROJECT/FILE_test.GO"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ext.Replace(tt.path, tt.old, tt.new); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceFold(t *testing.T) {
	tests := map[string]struct {
		path string
		old  string
		new  string
		want string
	}{
		"empty":      {path: "", old: "", new: "", want: ""},
		"file":       {path: "file.go", old: ext.Go, new: ext.GoTest, want: "file_test.go"},
		"base":       {path: "file.go", old: ext.Go, new: "", want: "file"},
		"none":       {path: "file.go", old: "", new: ext.HTM, want: "file.go"},
		"file upper": {path: "FILE.GO", old: ext.Go, new: ext.GoTest, want: "FILE_test.go"},
		"ext upper":  {path: "file.GO", old: ext.Go, new: ext.GoTest, want: "file_test.go"},
		"ext cap":    {path: "file.Go", old: ext.Go, new: ext.GoTest, want: "file_test.go"},
		"ext mixed":  {path: "file.gO", old: ext.Go, new: ext.GoTest, want: "file_test.go"},
		"path":       {path: "/project/file.go", old: ext.Go, new: ext.GoTest, want: "/project/file_test.go"},
		"path cap":   {path: "/project/file.Go", old: ext.Go, new: ext.GoTest, want: "/project/file_test.go"},
		"path mixed": {path: "/project/file.gO", old: ext.Go, new: ext.GoTest, want: "/project/file_test.go"},
		"path upper": {path: "/PROJECT/FILE.GO", old: ext.Go, new: ext.GoTest, want: "/PROJECT/FILE_test.go"},
		// File suffixes.
		"x-file":       {path: "file_test.go", old: ext.GoTest, new: ext.Go, want: "file.go"},
		"x-base":       {path: "file_test.go", old: ext.GoTest, new: "", want: "file"},
		"x-none":       {path: "file_test.go", old: "", new: ext.HTM, want: "file_test.go"},
		"x-file upper": {path: "FILE_TEST.GO", old: ext.GoTest, new: ext.Go, want: "FILE.go"},
		"x-ext upper":  {path: "file_TEST.GO", old: ext.GoTest, new: ext.Go, want: "file.go"},
		"x-ext cap":    {path: "file_Test.Go", old: ext.GoTest, new: ext.Go, want: "file.go"},
		"x-ext mixed":  {path: "file_tEsT.gO", old: ext.GoTest, new: ext.Go, want: "file.go"},
		"x-path":       {path: "/project/file_test.go", old: ext.GoTest, new: ext.Go, want: "/project/file.go"},
		"x-path cap":   {path: "/project/file_Test.Go", old: ext.GoTest, new: ext.Go, want: "/project/file.go"},
		"x-path mixed": {path: "/project/file_tEsT.gO", old: ext.GoTest, new: ext.Go, want: "/project/file.go"},
		"x-path upper": {path: "/PROJECT/FILE_TEST.GO", old: ext.GoTest, new: ext.Go, want: "/PROJECT/FILE.go"},
		"x-path semi":  {path: "/PROJECT/FILE_test.GO", old: ext.GoTest, new: ext.Go, want: "/PROJECT/FILE.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ext.ReplaceFold(tt.path, tt.old, tt.new); got != tt.want {
				t.Errorf("ReplaceFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
