package internal

import (
	"reflect"
	"testing"
)

func TestMakeSliceNextElemFunc(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want func() reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeSliceNextElemFunc(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeSliceNextElemFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnwrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unwrap(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldByIndexAlloc(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldByIndexAlloc(tt.args.v, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldByIndexAlloc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indirectNil(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indirectNil(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
