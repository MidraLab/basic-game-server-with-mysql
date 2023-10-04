// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutil

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {
	type args struct {
		root ast.Node
		pre  ApplyFunc
		post ApplyFunc
	}
	tests := []struct {
		name       string
		args       args
		wantResult ast.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Apply(tt.args.root, tt.args.pre, tt.args.post); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Apply() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCursor_Node(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   ast.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			if got := c.Node(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cursor.Node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursor_Parent(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   ast.Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			if got := c.Parent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cursor.Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursor_Name(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Cursor.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursor_Index(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			if got := c.Index(); got != tt.want {
				t.Errorf("Cursor.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursor_field(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			if got := c.field(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cursor.field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCursor_Replace(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			c.Replace(tt.args.n)
		})
	}
}

func TestCursor_Delete(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			c.Delete()
		})
	}
}

func TestCursor_InsertAfter(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			c.InsertAfter(tt.args.n)
		})
	}
}

func TestCursor_InsertBefore(t *testing.T) {
	type fields struct {
		parent ast.Node
		name   string
		iter   *iterator
		node   ast.Node
	}
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				parent: tt.fields.parent,
				name:   tt.fields.name,
				iter:   tt.fields.iter,
				node:   tt.fields.node,
			}
			c.InsertBefore(tt.args.n)
		})
	}
}

func Test_application_apply(t *testing.T) {
	type fields struct {
		pre    ApplyFunc
		post   ApplyFunc
		cursor Cursor
		iter   iterator
	}
	type args struct {
		parent ast.Node
		name   string
		iter   *iterator
		n      ast.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &application{
				pre:    tt.fields.pre,
				post:   tt.fields.post,
				cursor: tt.fields.cursor,
				iter:   tt.fields.iter,
			}
			a.apply(tt.args.parent, tt.args.name, tt.args.iter, tt.args.n)
		})
	}
}

func Test_application_applyList(t *testing.T) {
	type fields struct {
		pre    ApplyFunc
		post   ApplyFunc
		cursor Cursor
		iter   iterator
	}
	type args struct {
		parent ast.Node
		name   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &application{
				pre:    tt.fields.pre,
				post:   tt.fields.post,
				cursor: tt.fields.cursor,
				iter:   tt.fields.iter,
			}
			a.applyList(tt.args.parent, tt.args.name)
		})
	}
}
