package mysqldialect

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func Test_scanner(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want schema.ScannerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanner(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
