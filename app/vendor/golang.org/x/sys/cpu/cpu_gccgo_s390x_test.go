// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gccgo
// +build gccgo

package cpu

import (
	"reflect"
	"testing"
)

func Test_haveAsmFunctions(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haveAsmFunctions(); got != tt.want {
				t.Errorf("haveAsmFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stfle(t *testing.T) {
	tests := []struct {
		name string
		want facilityList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stfle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stfle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmcQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmcQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmcQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmctrQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmctrQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmctrQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmaQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmaQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmaQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kimdQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kimdQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kimdQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_klmdQuery(t *testing.T) {
	tests := []struct {
		name string
		want queryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := klmdQuery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("klmdQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
