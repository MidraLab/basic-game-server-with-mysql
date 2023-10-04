// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildutil

import (
	"go/build"
	"io"
	"reflect"
	"testing"
)

func TestOverlayContext(t *testing.T) {
	type args struct {
		orig    *build.Context
		overlay map[string][]byte
	}
	tests := []struct {
		name string
		args args
		want *build.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OverlayContext(tt.args.orig, tt.args.overlay); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OverlayContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseOverlayArchive(t *testing.T) {
	type args struct {
		archive io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string][]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOverlayArchive(tt.args.archive)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOverlayArchive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOverlayArchive() = %v, want %v", got, tt.want)
			}
		})
	}
}
