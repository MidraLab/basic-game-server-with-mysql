package schema

import (
	"reflect"
	"testing"
)

func TestFieldScanner(t *testing.T) {
	type args struct {
		dialect Dialect
		field   *Field
	}
	tests := []struct {
		name string
		args args
		want ScannerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldScanner(tt.args.dialect, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanner(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want ScannerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scanner(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanner(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want ScannerFunc
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

func Test_scanBool(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanBool(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanBool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanInt64(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanInt64(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanUint64(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanUint64(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanFloat64(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanFloat64(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanString(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanString(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanBytes(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanBytes(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanTime(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanTime(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanScanner(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanScanner(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanScanner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanMsgpack(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanMsgpack(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanMsgpack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanJSON(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanJSON(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanJSONUseNumber(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanJSONUseNumber(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanJSONUseNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanIP(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanIP(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanIP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanIPNet(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanIPNet(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanIPNet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addrScanner(t *testing.T) {
	type args struct {
		fn ScannerFunc
	}
	tests := []struct {
		name string
		args args
		want ScannerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addrScanner(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addrScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBytes(t *testing.T) {
	type args struct {
		src interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytes(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrScanner(t *testing.T) {
	type args struct {
		fn ScannerFunc
	}
	tests := []struct {
		name string
		args args
		want ScannerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrScanner(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrScanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanNull(t *testing.T) {
	type args struct {
		dest reflect.Value
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
			if err := scanNull(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("scanNull() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanJSONIntoInterface(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanJSONIntoInterface(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanJSONIntoInterface() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanInterface(t *testing.T) {
	type args struct {
		dest reflect.Value
		src  interface{}
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
			if err := scanInterface(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanInterface() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nilable(t *testing.T) {
	type args struct {
		kind reflect.Kind
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
			if got := nilable(tt.args.kind); got != tt.want {
				t.Errorf("nilable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanError(t *testing.T) {
	type args struct {
		dest reflect.Type
		src  interface{}
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
			if err := scanError(tt.args.dest, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
