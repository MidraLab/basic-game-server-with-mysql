// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.21
// +build go1.21

package cpu

import (
	"reflect"
	"testing"
	_ "unsafe"
)

func Test_runtime_getAuxv(t *testing.T) {
	tests := []struct {
		name string
		want []uintptr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runtime_getAuxv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runtime_getAuxv() = %v, want %v", got, tt.want)
			}
		})
	}
}
