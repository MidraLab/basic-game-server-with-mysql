// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import (
	"reflect"
	"testing"
)

func Test_getAuxv(t *testing.T) {
	tests := []struct {
		name string
		want []uintptr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAuxv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAuxv() = %v, want %v", got, tt.want)
			}
		})
	}
}
