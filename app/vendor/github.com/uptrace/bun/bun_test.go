package bun

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

func TestSetLogger(t *testing.T) {
	type args struct {
		logger internal.Logging
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogger(tt.args.logger)
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want schema.QueryAppender
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullZero(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want schema.QueryAppender
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullZero(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
