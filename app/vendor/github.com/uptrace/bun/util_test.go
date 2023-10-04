package bun

import (
	"reflect"
	"testing"
)

func Test_indirect(t *testing.T) {
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
			if got := indirect(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walk(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
		fn    func(reflect.Value)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			walk(tt.args.v, tt.args.index, tt.args.fn)
		})
	}
}

func Test_visitField(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
		fn    func(reflect.Value)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			visitField(tt.args.v, tt.args.index, tt.args.fn)
		})
	}
}

func Test_typeByIndex(t *testing.T) {
	type args struct {
		t     reflect.Type
		index []int
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeByIndex(tt.args.t, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("typeByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indirectType(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indirectType(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceElemType(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceElemType(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceElemType() = %v, want %v", got, tt.want)
			}
		})
	}
}
