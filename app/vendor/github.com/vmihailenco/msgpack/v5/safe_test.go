//go:build appengine
// +build appengine

package msgpack

import (
	"reflect"
	"testing"
)

func Test_bytesToString(t *testing.T) {
	type args struct {
		b []byte
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
			if got := bytesToString(tt.args.b); got != tt.want {
				t.Errorf("bytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
