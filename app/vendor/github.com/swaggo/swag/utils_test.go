package swag

import (
	"reflect"
	"testing"
)

func TestFieldsFunc(t *testing.T) {
	type args struct {
		s string
		f func(rune2 rune) bool
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldsFunc(tt.args.s, tt.args.f, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldsByAnySpace(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldsByAnySpace(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldsByAnySpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
