// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutil

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestPathEnclosingInterval(t *testing.T) {
	type args struct {
		root  *ast.File
		start token.Pos
		end   token.Pos
	}
	tests := []struct {
		name      string
		args      args
		wantPath  []ast.Node
		wantExact bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotExact := PathEnclosingInterval(tt.args.root, tt.args.start, tt.args.end)
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("PathEnclosingInterval() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotExact != tt.wantExact {
				t.Errorf("PathEnclosingInterval() gotExact = %v, want %v", gotExact, tt.wantExact)
			}
		})
	}
}

func Test_tokenNode_Pos(t *testing.T) {
	type fields struct {
		pos token.Pos
		end token.Pos
	}
	tests := []struct {
		name   string
		fields fields
		want   token.Pos
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tokenNode{
				pos: tt.fields.pos,
				end: tt.fields.end,
			}
			if got := n.Pos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenNode.Pos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenNode_End(t *testing.T) {
	type fields struct {
		pos token.Pos
		end token.Pos
	}
	tests := []struct {
		name   string
		fields fields
		want   token.Pos
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tokenNode{
				pos: tt.fields.pos,
				end: tt.fields.end,
			}
			if got := n.End(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenNode.End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tok(t *testing.T) {
	type args struct {
		pos token.Pos
		len int
	}
	tests := []struct {
		name string
		args args
		want ast.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tok(tt.args.pos, tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tok() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_childrenOf(t *testing.T) {
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name string
		args args
		want []ast.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := childrenOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("childrenOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byPos_Len(t *testing.T) {
	tests := []struct {
		name string
		sl   byPos
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.Len(); got != tt.want {
				t.Errorf("byPos.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byPos_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		sl   byPos
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sl.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byPos.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byPos_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		sl   byPos
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sl.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestNodeDescription(t *testing.T) {
	type args struct {
		n ast.Node
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
			if got := NodeDescription(tt.args.n); got != tt.want {
				t.Errorf("NodeDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
