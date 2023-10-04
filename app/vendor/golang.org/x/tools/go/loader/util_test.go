// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loader

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func Test_parseFiles(t *testing.T) {
	type args struct {
		fset        *token.FileSet
		ctxt        *build.Context
		displayPath func(string) string
		dir         string
		files       []string
		mode        parser.Mode
	}
	tests := []struct {
		name  string
		args  args
		want  []*ast.File
		want1 []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseFiles(tt.args.fset, tt.args.ctxt, tt.args.displayPath, tt.args.dir, tt.args.files, tt.args.mode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFiles() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseFiles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_scanImports(t *testing.T) {
	type args struct {
		files []*ast.File
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanImports(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanImports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenFileContainsPos(t *testing.T) {
	type args struct {
		f   *token.File
		pos token.Pos
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenFileContainsPos(tt.args.f, tt.args.pos); got != tt.want {
				t.Errorf("tokenFileContainsPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
