package ext_test

import (
	"github.com/speedyhoon/ext"
	"strings"
	"testing"
)

func TestExtensions(t *testing.T) {
	list := [][2]string{
		// {Constant, Expected},
		{ext.BMP, ".bmp"},
		{ext.CSS, ".css"},
		{ext.Go, ".go"},
		{ext.GoHTML, ".gohtml"},
		{ext.GoTest, "_test.go"},
		{ext.HTM, ".htm"},
		{ext.HTML, ".html"},
		{ext.ICO, ".ico"},
		{ext.INI, ".ini"},
		{ext.ISS, ".iss"},
		{ext.JPG, ".jpg"},
		{ext.JS, ".js"},
		{ext.JSON, ".json"},
		{ext.Log, ".log"},
		{ext.MD, ".md"},
		{ext.PDF, ".pdf"},
		{ext.PNG, ".png"},
		{ext.STYL, ".styl"},
		{ext.SVG, ".svg"},
		{ext.TXT, ".txt"},
		{ext.XHTML, ".xhtml"},
		{ext.YML, ".yml"},
	}
	for i, item := range list {
		val, expected := item[0], item[1]
		t.Run(expected, func(t *testing.T) {
			if val != expected {
				t.Errorf("expected extension[%d]: `%s`, got: `%s`", i, expected, val)
			}
			if val != strings.ToLower(val) {
				t.Errorf("expected extension[%d] to be lowercase: `%s`", i, val)
			}
		})
	}
}
