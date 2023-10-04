// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "testing"

func Test_readHWCAP(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readHWCAP(); (err != nil) != tt.wantErr {
				t.Errorf("readHWCAP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
