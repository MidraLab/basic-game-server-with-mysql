// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "testing"

func Test_parseRelease(t *testing.T) {
	type args struct {
		rel string
	}
	tests := []struct {
		name      string
		args      args
		wantMajor int
		wantMinor int
		wantPatch int
		wantOk    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMajor, gotMinor, gotPatch, gotOk := parseRelease(tt.args.rel)
			if gotMajor != tt.wantMajor {
				t.Errorf("parseRelease() gotMajor = %v, want %v", gotMajor, tt.wantMajor)
			}
			if gotMinor != tt.wantMinor {
				t.Errorf("parseRelease() gotMinor = %v, want %v", gotMinor, tt.wantMinor)
			}
			if gotPatch != tt.wantPatch {
				t.Errorf("parseRelease() gotPatch = %v, want %v", gotPatch, tt.wantPatch)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseRelease() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
