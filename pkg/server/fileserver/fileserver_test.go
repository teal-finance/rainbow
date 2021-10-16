// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package fileserver

import (
	"testing"
)

func Test_extIndex(t *testing.T) {
	cases := []struct {
		name string
		path string
		ext  string
	}{
		{"regular folder and filename", "folder/file.ext", "ext"},
		{"without folder", "file.ext", "ext"},
		{"filename without extension", "folder/file", ""},
		{"empty path has no extension", "", ""},
		{"valid folder but empty filename", "folder/", ""},
		{"ignore dot in folder", "folder.ext/file", ""},
		{"ignore dot in folder even when no file", "folder.ext/", ""},
		{"filename ending with a dot has no extension", "ending-dot.", ""},
		{"filename ending with a double dot has no extension", "double-dot..", ""},
		{"only consider the last dot", "a..b.c..ext", "ext"},
		{"filename starting with a dot has an extension", ".gitignore", "gitignore"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			extPos := extIndex(c.path)
			max := len(c.path)
			if extPos < 0 || extPos > max {
				t.Errorf("extIndex() = %v out of range [0..%v]", extPos, max)
			}

			got := c.path[extPos:]
			if got != c.ext {
				t.Errorf("extIndex() = %v, want %v", got, c.ext)
			}
		})
	}
}
