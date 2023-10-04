// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package astutil contains common utilities for working with the Go AST.
package astutil

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestAddImport(t *testing.T) {
	type args struct {
		fset *token.FileSet
		f    *ast.File
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantAdded bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdded := AddImport(tt.args.fset, tt.args.f, tt.args.path); gotAdded != tt.wantAdded {
				t.Errorf("AddImport() = %v, want %v", gotAdded, tt.wantAdded)
			}
		})
	}
}

func TestAddNamedImport(t *testing.T) {
	type args struct {
		fset *token.FileSet
		f    *ast.File
		name string
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantAdded bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdded := AddNamedImport(tt.args.fset, tt.args.f, tt.args.name, tt.args.path); gotAdded != tt.wantAdded {
				t.Errorf("AddNamedImport() = %v, want %v", gotAdded, tt.wantAdded)
			}
		})
	}
}

func Test_isThirdParty(t *testing.T) {
	type args struct {
		importPath string
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
			if got := isThirdParty(tt.args.importPath); got != tt.want {
				t.Errorf("isThirdParty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteImport(t *testing.T) {
	type args struct {
		fset *token.FileSet
		f    *ast.File
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantDeleted bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDeleted := DeleteImport(tt.args.fset, tt.args.f, tt.args.path); gotDeleted != tt.wantDeleted {
				t.Errorf("DeleteImport() = %v, want %v", gotDeleted, tt.wantDeleted)
			}
		})
	}
}

func TestDeleteNamedImport(t *testing.T) {
	type args struct {
		fset *token.FileSet
		f    *ast.File
		name string
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantDeleted bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDeleted := DeleteNamedImport(tt.args.fset, tt.args.f, tt.args.name, tt.args.path); gotDeleted != tt.wantDeleted {
				t.Errorf("DeleteNamedImport() = %v, want %v", gotDeleted, tt.wantDeleted)
			}
		})
	}
}

func TestRewriteImport(t *testing.T) {
	type args struct {
		fset    *token.FileSet
		f       *ast.File
		oldPath string
		newPath string
	}
	tests := []struct {
		name        string
		args        args
		wantRewrote bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRewrote := RewriteImport(tt.args.fset, tt.args.f, tt.args.oldPath, tt.args.newPath); gotRewrote != tt.wantRewrote {
				t.Errorf("RewriteImport() = %v, want %v", gotRewrote, tt.wantRewrote)
			}
		})
	}
}

func TestUsesImport(t *testing.T) {
	type args struct {
		f    *ast.File
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantUsed bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUsed := UsesImport(tt.args.f, tt.args.path); gotUsed != tt.wantUsed {
				t.Errorf("UsesImport() = %v, want %v", gotUsed, tt.wantUsed)
			}
		})
	}
}

func Test_visitFn_Visit(t *testing.T) {
	type args struct {
		node ast.Node
	}
	tests := []struct {
		name string
		fn   visitFn
		args args
		want ast.Visitor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn.Visit(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visitFn.Visit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imports(t *testing.T) {
	type args struct {
		f    *ast.File
		name string
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
			if got := imports(tt.args.f, tt.args.name, tt.args.path); got != tt.want {
				t.Errorf("imports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importSpec(t *testing.T) {
	type args struct {
		f    *ast.File
		path string
	}
	tests := []struct {
		name string
		args args
		want *ast.ImportSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := importSpec(tt.args.f, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importName(t *testing.T) {
	type args struct {
		s *ast.ImportSpec
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
			if got := importName(tt.args.s); got != tt.want {
				t.Errorf("importName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importPath(t *testing.T) {
	type args struct {
		s *ast.ImportSpec
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
			if got := importPath(tt.args.s); got != tt.want {
				t.Errorf("importPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_declImports(t *testing.T) {
	type args struct {
		gen  *ast.GenDecl
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
			if got := declImports(tt.args.gen, tt.args.path); got != tt.want {
				t.Errorf("declImports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchLen(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchLen(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("matchLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTopName(t *testing.T) {
	type args struct {
		n    ast.Expr
		name string
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
			if got := isTopName(tt.args.n, tt.args.name); got != tt.want {
				t.Errorf("isTopName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImports(t *testing.T) {
	type args struct {
		fset *token.FileSet
		f    *ast.File
	}
	tests := []struct {
		name string
		args args
		want [][]*ast.ImportSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Imports(tt.args.fset, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Imports() = %v, want %v", got, tt.want)
			}
		})
	}
}
