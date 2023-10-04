// Copyright 2017 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js
// +build js

package uuid

import (
	"reflect"
	"testing"
)

func Test_getHardwareInterface(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getHardwareInterface(tt.args.name)
			if got != tt.want {
				t.Errorf("getHardwareInterface() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getHardwareInterface() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
