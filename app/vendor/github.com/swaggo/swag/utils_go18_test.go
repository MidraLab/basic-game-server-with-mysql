//go:build go1.18
// +build go1.18

package swag

import (
	"reflect"
	"testing"
)

func TestAppendUtf8Rune(t *testing.T) {
	type args struct {
		p []byte
		r rune
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
			if got := AppendUtf8Rune(tt.args.p, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUtf8Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanIntegerValue_CanInt(t *testing.T) {
	type fields struct {
		Value reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CanIntegerValue{
				Value: tt.fields.Value,
			}
			if got := v.CanInt(); got != tt.want {
				t.Errorf("CanIntegerValue.CanInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanIntegerValue_CanUint(t *testing.T) {
	type fields struct {
		Value reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CanIntegerValue{
				Value: tt.fields.Value,
			}
			if got := v.CanUint(); got != tt.want {
				t.Errorf("CanIntegerValue.CanUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
