// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildutil

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	type args struct {
		fset        *token.FileSet
		ctxt        *build.Context
		displayPath func(string) string
		dir         string
		file        string
		mode        parser.Mode
	}
	tests := []struct {
		name    string
		args    args
		want    *ast.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFile(tt.args.fset, tt.args.ctxt, tt.args.displayPath, tt.args.dir, tt.args.file, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainingPackage(t *testing.T) {
	type args struct {
		ctxt     *build.Context
		dir      string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *build.Package
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ContainingPackage(tt.args.ctxt, tt.args.dir, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContainingPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContainingPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSubdir(t *testing.T) {
	type args struct {
		ctxt *build.Context
		root string
		dir  string
	}
	tests := []struct {
		name    string
		args    args
		wantRel string
		wantOk  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRel, gotOk := HasSubdir(tt.args.ctxt, tt.args.root, tt.args.dir)
			if gotRel != tt.wantRel {
				t.Errorf("HasSubdir() gotRel = %v, want %v", gotRel, tt.wantRel)
			}
			if gotOk != tt.wantOk {
				t.Errorf("HasSubdir() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_hasSubdir(t *testing.T) {
	type args struct {
		root string
		dir  string
	}
	tests := []struct {
		name    string
		args    args
		wantRel string
		wantOk  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRel, gotOk := hasSubdir(tt.args.root, tt.args.dir)
			if gotRel != tt.wantRel {
				t.Errorf("hasSubdir() gotRel = %v, want %v", gotRel, tt.wantRel)
			}
			if gotOk != tt.wantOk {
				t.Errorf("hasSubdir() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path string
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
			if got := FileExists(tt.args.ctxt, tt.args.path); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenFile(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenFile(tt.args.ctxt, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAbsPath(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path string
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
			if got := IsAbsPath(tt.args.ctxt, tt.args.path); got != tt.want {
				t.Errorf("IsAbsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPath(tt.args.ctxt, tt.args.path...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path string
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
			if got := IsDir(tt.args.ctxt, tt.args.path); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDir(t *testing.T) {
	type args struct {
		ctxt *build.Context
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.ctxt, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPathList(t *testing.T) {
	type args struct {
		ctxt *build.Context
		s    string
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
			if got := SplitPathList(tt.args.ctxt, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPathList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameFile(t *testing.T) {
	type args struct {
		x string
		y string
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
			if got := sameFile(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("sameFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
