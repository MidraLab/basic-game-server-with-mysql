package schema

import (
	"reflect"
	"testing"
)

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

func Test_fieldByIndex(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
	}
	tests := []struct {
		name   string
		args   args
		want   reflect.Value
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := fieldByIndex(tt.args.v, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldByIndex() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("fieldByIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
