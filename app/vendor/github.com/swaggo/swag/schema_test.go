package swag

import (
	"reflect"
	"testing"

	"github.com/go-openapi/spec"
)

func TestCheckSchemaType(t *testing.T) {
	type args struct {
		typeName string
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
			if err := CheckSchemaType(tt.args.typeName); (err != nil) != tt.wantErr {
				t.Errorf("CheckSchemaType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsSimplePrimitiveType(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSimplePrimitiveType(tt.args.typeName); got != tt.want {
				t.Errorf("IsSimplePrimitiveType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrimitiveType(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrimitiveType(tt.args.typeName); got != tt.want {
				t.Errorf("IsPrimitiveType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInterfaceLike(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterfaceLike(tt.args.typeName); got != tt.want {
				t.Errorf("IsInterfaceLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumericType(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumericType(tt.args.typeName); got != tt.want {
				t.Errorf("IsNumericType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransToValidSchemeType(t *testing.T) {
	type args struct {
		typeName string
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
			if got := TransToValidSchemeType(tt.args.typeName); got != tt.want {
				t.Errorf("TransToValidSchemeType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGolangPrimitiveType(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGolangPrimitiveType(tt.args.typeName); got != tt.want {
				t.Errorf("IsGolangPrimitiveType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransToValidCollectionFormat(t *testing.T) {
	type args struct {
		format string
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
			if got := TransToValidCollectionFormat(tt.args.format); got != tt.want {
				t.Errorf("TransToValidCollectionFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ignoreNameOverride(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ignoreNameOverride(tt.args.name); got != tt.want {
				t.Errorf("ignoreNameOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsComplexSchema(t *testing.T) {
	type args struct {
		schema *spec.Schema
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsComplexSchema(tt.args.schema); got != tt.want {
				t.Errorf("IsComplexSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRefSchema(t *testing.T) {
	type args struct {
		schema *spec.Schema
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRefSchema(tt.args.schema); got != tt.want {
				t.Errorf("IsRefSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefSchema(t *testing.T) {
	type args struct {
		refType string
	}
	tests := []struct {
		name string
		args args
		want *spec.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefSchema(tt.args.refType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimitiveSchema(t *testing.T) {
	type args struct {
		refType string
	}
	tests := []struct {
		name string
		args args
		want *spec.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimitiveSchema(tt.args.refType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimitiveSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildCustomSchema(t *testing.T) {
	type args struct {
		types []string
	}
	tests := []struct {
		name    string
		args    args
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildCustomSchema(tt.args.types)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildCustomSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCustomSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSchema(t *testing.T) {
	type args struct {
		dst *spec.Schema
		src *spec.Schema
	}
	tests := []struct {
		name string
		args args
		want *spec.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSchema(tt.args.dst, tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
