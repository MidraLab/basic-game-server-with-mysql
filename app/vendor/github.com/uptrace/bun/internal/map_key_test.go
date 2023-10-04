package internal

import (
	"reflect"
	"testing"
)

func TestNewMapKey(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want MapKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMapKey(tt.args.is); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newMapKey(t *testing.T) {
	type args struct {
		is []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMapKey(tt.args.is); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
