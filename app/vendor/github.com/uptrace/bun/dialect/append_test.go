package dialect

import (
	"reflect"
	"testing"
)

func TestAppendError(t *testing.T) {
	type args struct {
		b   []byte
		err error
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
			if got := AppendError(tt.args.b, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendNull(t *testing.T) {
	type args struct {
		b []byte
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
			if got := AppendNull(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendBool(t *testing.T) {
	type args struct {
		b []byte
		v bool
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
			if got := AppendBool(tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFloat32(t *testing.T) {
	type args struct {
		b []byte
		v float32
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
			if got := AppendFloat32(tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFloat64(t *testing.T) {
	type args struct {
		b []byte
		v float64
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
			if got := AppendFloat64(tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendFloat(t *testing.T) {
	type args struct {
		b       []byte
		v       float64
		bitSize int
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
			if got := appendFloat(tt.args.b, tt.args.v, tt.args.bitSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendName(t *testing.T) {
	type args struct {
		b     []byte
		ident string
		quote byte
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
			if got := AppendName(tt.args.b, tt.args.ident, tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendName(t *testing.T) {
	type args struct {
		b     []byte
		ident []byte
		quote byte
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
			if got := appendName(tt.args.b, tt.args.ident, tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendIdent(t *testing.T) {
	type args struct {
		b     []byte
		name  string
		quote byte
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
			if got := AppendIdent(tt.args.b, tt.args.name, tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendIdent(t *testing.T) {
	type args struct {
		b     []byte
		name  []byte
		quote byte
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
			if got := appendIdent(tt.args.b, tt.args.name, tt.args.quote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}
