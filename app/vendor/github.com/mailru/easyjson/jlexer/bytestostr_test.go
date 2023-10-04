// This file will only be included to the build if neither
// easyjson_nounsafe nor appengine build tag is set. See README notes
// for more details.

//go:build !easyjson_nounsafe && !appengine
// +build !easyjson_nounsafe,!appengine

package jlexer

import "testing"

func Test_bytesToStr(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bytesToStr(tt.args.data); got != tt.want {
				t.Errorf("bytesToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
