// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package typeparams contains common utilities for writing tools that interact
// with generic Go code, as introduced with Go 1.18.
//
// Many of the types and functions in this package are proxies for the new APIs
// introduced in the standard library with Go 1.18. For example, the
// typeparams.Union type is an alias for go/types.Union, and the ForTypeSpec
// function returns the value of the go/ast.TypeSpec.TypeParams field. At Go
// versions older than 1.18 these helpers are implemented as stubs, allowing
// users of this package to write code that handles generic constructs inline,
// even if the Go version being used to compile does not support generics.
//
// Additionally, this package contains common utilities for working with the
// new generic constructs, to supplement the standard library APIs. Notably,
// the StructuralTerms API computes a minimal representation of the structural
// restrictions on a type parameter.
//
// An external version of these APIs is available in the
// golang.org/x/exp/typeparams module.
package typeparams

import (
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"testing"
)

func TestUnpackIndexExpr(t *testing.T) {
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name        string
		args        args
		wantX       ast.Expr
		wantLbrack  token.Pos
		wantIndices []ast.Expr
		wantRbrack  token.Pos
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotLbrack, gotIndices, gotRbrack := UnpackIndexExpr(tt.args.n)
			if !reflect.DeepEqual(gotX, tt.wantX) {
				t.Errorf("UnpackIndexExpr() gotX = %v, want %v", gotX, tt.wantX)
			}
			if !reflect.DeepEqual(gotLbrack, tt.wantLbrack) {
				t.Errorf("UnpackIndexExpr() gotLbrack = %v, want %v", gotLbrack, tt.wantLbrack)
			}
			if !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("UnpackIndexExpr() gotIndices = %v, want %v", gotIndices, tt.wantIndices)
			}
			if !reflect.DeepEqual(gotRbrack, tt.wantRbrack) {
				t.Errorf("UnpackIndexExpr() gotRbrack = %v, want %v", gotRbrack, tt.wantRbrack)
			}
		})
	}
}

func TestPackIndexExpr(t *testing.T) {
	type args struct {
		x       ast.Expr
		lbrack  token.Pos
		indices []ast.Expr
		rbrack  token.Pos
	}
	tests := []struct {
		name string
		args args
		want ast.Expr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PackIndexExpr(tt.args.x, tt.args.lbrack, tt.args.indices, tt.args.rbrack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackIndexExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTypeParam(t *testing.T) {
	type args struct {
		t types.Type
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
			if got := IsTypeParam(tt.args.t); got != tt.want {
				t.Errorf("IsTypeParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOriginMethod(t *testing.T) {
	type args struct {
		fn *types.Func
	}
	tests := []struct {
		name string
		args args
		want *types.Func
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OriginMethod(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OriginMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericAssignableTo(t *testing.T) {
	type args struct {
		ctxt *Context
		V    types.Type
		T    types.Type
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
			if got := GenericAssignableTo(tt.args.ctxt, tt.args.V, tt.args.T); got != tt.want {
				t.Errorf("GenericAssignableTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
