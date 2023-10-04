// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && arm64
// +build linux,arm64

package cpu

import "testing"

func Test_readLinuxProcCPUInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readLinuxProcCPUInfo(); (err != nil) != tt.wantErr {
				t.Errorf("readLinuxProcCPUInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
