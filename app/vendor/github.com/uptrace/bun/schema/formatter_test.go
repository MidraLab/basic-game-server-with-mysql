package schema

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/dialect/feature"
	"github.com/uptrace/bun/internal/parser"
)

func TestNewFormatter(t *testing.T) {
	type args struct {
		dialect Dialect
	}
	tests := []struct {
		name string
		args args
		want Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFormatter(tt.args.dialect); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNopFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNopFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNopFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_IsNop(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.IsNop(); got != tt.want {
				t.Errorf("Formatter.IsNop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_Dialect(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	tests := []struct {
		name   string
		fields fields
		want   Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_IdentQuote(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.IdentQuote(); got != tt.want {
				t.Errorf("Formatter.IdentQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_AppendName(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		b    []byte
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.AppendName(tt.args.b, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.AppendName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_AppendIdent(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		b     []byte
		ident string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.AppendIdent(tt.args.b, tt.args.ident); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.AppendIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_AppendValue(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		b []byte
		v reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.AppendValue(tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.AppendValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_HasFeature(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		feature feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.HasFeature(tt.args.feature); got != tt.want {
				t.Errorf("Formatter.HasFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_WithArg(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		arg NamedArgAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.WithArg(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.WithArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_WithNamedArg(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.WithNamedArg(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.WithNamedArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_FormatQuery(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.FormatQuery(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("Formatter.FormatQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_AppendQuery(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		dst   []byte
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.AppendQuery(tt.args.dst, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_append(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		dst  []byte
		p    *parser.Parser
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.append(tt.args.dst, tt.args.p, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_appendArg(t *testing.T) {
	type fields struct {
		dialect Dialect
		args    *namedArgList
	}
	type args struct {
		b   []byte
		arg interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formatter{
				dialect: tt.fields.dialect,
				args:    tt.fields.args,
			}
			if got := f.appendArg(tt.args.b, tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.appendArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_namedArgList_WithArg(t *testing.T) {
	type fields struct {
		arg  NamedArgAppender
		next *namedArgList
	}
	type args struct {
		arg NamedArgAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *namedArgList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &namedArgList{
				arg:  tt.fields.arg,
				next: tt.fields.next,
			}
			if got := l.WithArg(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("namedArgList.WithArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_namedArgList_AppendNamedArg(t *testing.T) {
	type fields struct {
		arg  NamedArgAppender
		next *namedArgList
	}
	type args struct {
		fmter Formatter
		b     []byte
		name  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &namedArgList{
				arg:  tt.fields.arg,
				next: tt.fields.next,
			}
			got, got1 := l.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("namedArgList.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("namedArgList.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_namedArg_AppendNamedArg(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
	}
	type args struct {
		fmter Formatter
		b     []byte
		name  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &namedArg{
				name:  tt.fields.name,
				value: tt.fields.value,
			}
			got, got1 := a.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("namedArg.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("namedArg.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_newStructArgs(t *testing.T) {
	type args struct {
		fmter Formatter
		strct interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  *structArgs
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := newStructArgs(tt.args.fmter, tt.args.strct)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStructArgs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("newStructArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_structArgs_AppendNamedArg(t *testing.T) {
	type fields struct {
		table *Table
		strct reflect.Value
	}
	type args struct {
		fmter Formatter
		b     []byte
		name  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structArgs{
				table: tt.fields.table,
				strct: tt.fields.strct,
			}
			got, got1 := m.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structArgs.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("structArgs.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
