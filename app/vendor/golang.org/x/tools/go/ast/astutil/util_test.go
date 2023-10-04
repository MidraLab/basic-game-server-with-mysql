// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutil

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestUnparen(t *testing.T) {
	type args struct {
		e ast.Expr
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
			if got := Unparen(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unparen() = %v, want %v", got, tt.want)
			}
		})
	}
}
