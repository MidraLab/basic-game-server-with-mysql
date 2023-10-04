package schema

import (
	"reflect"
	"testing"
)

func TestFieldAppender(t *testing.T) {
	type args struct {
		dialect Dialect
		field   *Field
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldAppender(tt.args.dialect, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldAppender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppender(t *testing.T) {
	type args struct {
		dialect Dialect
		typ     reflect.Type
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Appender(tt.args.dialect, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Appender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appender(t *testing.T) {
	type args struct {
		dialect Dialect
		typ     reflect.Type
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appender(tt.args.dialect, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ifaceAppenderFunc(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := ifaceAppenderFunc(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ifaceAppenderFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nilAwareAppender(t *testing.T) {
	type args struct {
		fn AppenderFunc
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nilAwareAppender(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nilAwareAppender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPtrAppender(t *testing.T) {
	type args struct {
		fn AppenderFunc
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PtrAppender(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PtrAppender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendBoolValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendBoolValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendBoolValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendIntValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendIntValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendIntValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendUintValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendUintValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendUintValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendUint32Value(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendUint32Value(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendUint32Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendUint64Value(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendUint64Value(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendUint64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFloat32Value(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendFloat32Value(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendFloat32Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendFloat64Value(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendFloat64Value(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendFloat64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendBytesValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendBytesValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendBytesValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendArrayBytesValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendArrayBytesValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendArrayBytesValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendStringValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendStringValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendJSONValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := AppendJSONValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendJSONValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendTimeValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendTimeValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendTimeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendIPValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendIPValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendIPValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendIPNetValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendIPNetValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendIPNetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendJSONRawMessageValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendJSONRawMessageValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendJSONRawMessageValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendQueryAppenderValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendQueryAppenderValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendQueryAppenderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendDriverValue(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendDriverValue(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendDriverValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addrAppender(t *testing.T) {
	type args struct {
		fn AppenderFunc
	}
	tests := []struct {
		name string
		args args
		want AppenderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addrAppender(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addrAppender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendMsgpack(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     reflect.Value
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
			if got := appendMsgpack(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendMsgpack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendQueryAppender(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		app   QueryAppender
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
			if got := AppendQueryAppender(tt.args.fmter, tt.args.b, tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendQueryAppender() = %v, want %v", got, tt.want)
			}
		})
	}
}
