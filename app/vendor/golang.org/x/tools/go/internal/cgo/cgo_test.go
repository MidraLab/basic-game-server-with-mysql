// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cgo handles cgo preprocessing of files containing `import "C"`.
//
// DESIGN
//
// The approach taken is to run the cgo processor on the package's
// CgoFiles and parse the output, faking the filenames of the
// resulting ASTs so that the synthetic file containing the C types is
// called "C" (e.g. "~/go/src/net/C") and the preprocessed files
// have their original names (e.g. "~/go/src/net/cgo_unix.go"),
// not the names of the actual temporary files.
//
// The advantage of this approach is its fidelity to 'go build'.  The
// downside is that the token.Position.Offset for each AST node is
// incorrect, being an offset within the temporary file.  Line numbers
// should still be correct because of the //line comments.
//
// The logic of this file is mostly plundered from the 'go build'
// tool, which also invokes the cgo preprocessor.
//
//
// REJECTED ALTERNATIVE
//
// An alternative approach that we explored is to extend go/types'
// Importer mechanism to provide the identity of the importing package
// so that each time `import "C"` appears it resolves to a different
// synthetic package containing just the objects needed in that case.
// The loader would invoke cgo but parse only the cgo_types.go file
// defining the package-level objects, discarding the other files
// resulting from preprocessing.
//
// The benefit of this approach would have been that source-level
// syntax information would correspond exactly to the original cgo
// file, with no preprocessing involved, making source tools like
// godoc, guru, and eg happy.  However, the approach was rejected
// due to the additional complexity it would impose on go/types.  (It
// made for a beautiful demo, though.)
//
// cgo files, despite their *.go extension, are not legal Go source
// files per the specification since they may refer to unexported
// members of package "C" such as C.int.  Also, a function such as
// C.getpwent has in effect two types, one matching its C type and one
// which additionally returns (errno C.int).  The cgo preprocessor
// uses name mangling to distinguish these two functions in the
// processed code, but go/types would need to duplicate this logic in
// its handling of function calls, analogous to the treatment of map
// lookups in which y=m[k] and y,ok=m[k] are both legal.

package cgo

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestProcessFiles(t *testing.T) {
	type args struct {
		bp          *build.Package
		fset        *token.FileSet
		DisplayPath func(path string) string
		mode        parser.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    []*ast.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessFiles(tt.args.bp, tt.args.fset, tt.args.DisplayPath, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		bp     *build.Package
		pkgdir string
		tmpdir string
		useabs bool
	}
	tests := []struct {
		name             string
		args             args
		wantFiles        []string
		wantDisplayFiles []string
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, gotDisplayFiles, err := Run(tt.args.bp, tt.args.pkgdir, tt.args.tmpdir, tt.args.useabs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("Run() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
			if !reflect.DeepEqual(gotDisplayFiles, tt.wantDisplayFiles) {
				t.Errorf("Run() gotDisplayFiles = %v, want %v", gotDisplayFiles, tt.wantDisplayFiles)
			}
		})
	}
}

func Test_cflags(t *testing.T) {
	type args struct {
		p   *build.Package
		def bool
	}
	tests := []struct {
		name         string
		args         args
		wantCppflags []string
		wantCflags   []string
		wantCxxflags []string
		wantLdflags  []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCppflags, gotCflags, gotCxxflags, gotLdflags := cflags(tt.args.p, tt.args.def)
			if !reflect.DeepEqual(gotCppflags, tt.wantCppflags) {
				t.Errorf("cflags() gotCppflags = %v, want %v", gotCppflags, tt.wantCppflags)
			}
			if !reflect.DeepEqual(gotCflags, tt.wantCflags) {
				t.Errorf("cflags() gotCflags = %v, want %v", gotCflags, tt.wantCflags)
			}
			if !reflect.DeepEqual(gotCxxflags, tt.wantCxxflags) {
				t.Errorf("cflags() gotCxxflags = %v, want %v", gotCxxflags, tt.wantCxxflags)
			}
			if !reflect.DeepEqual(gotLdflags, tt.wantLdflags) {
				t.Errorf("cflags() gotLdflags = %v, want %v", gotLdflags, tt.wantLdflags)
			}
		})
	}
}

func Test_envList(t *testing.T) {
	type args struct {
		key string
		def string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := envList(tt.args.key, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("envList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringList(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringList(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringList() = %v, want %v", got, tt.want)
			}
		})
	}
}
