// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !aix && !linux && (ppc64 || ppc64le)
// +build !aix
// +build !linux
// +build ppc64 ppc64le

package cpu

import "testing"

func Test_archInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			archInit()
		})
	}
}
