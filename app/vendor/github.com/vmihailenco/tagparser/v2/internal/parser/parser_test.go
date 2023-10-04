package parser

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Bytes(t *testing.T) {
	type fields struct {
		b []byte
		i int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Valid(t *testing.T) {
	type fields struct {
		b []byte
		i int
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.Valid(); got != tt.want {
				t.Errorf("Parser.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Read(t *testing.T) {
	type fields struct {
		b []byte
		i int
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.Read(); got != tt.want {
				t.Errorf("Parser.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Peek(t *testing.T) {
	type fields struct {
		b []byte
		i int
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.Peek(); got != tt.want {
				t.Errorf("Parser.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Advance(t *testing.T) {
	type fields struct {
		b []byte
		i int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			p.Advance()
		})
	}
}

func TestParser_Skip(t *testing.T) {
	type fields struct {
		b []byte
		i int
	}
	type args struct {
		skip byte
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.Skip(tt.args.skip); got != tt.want {
				t.Errorf("Parser.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_SkipBytes(t *testing.T) {
	type fields struct {
		b []byte
		i int
	}
	type args struct {
		skip []byte
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			if got := p.SkipBytes(tt.args.skip); got != tt.want {
				t.Errorf("Parser.SkipBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ReadSep(t *testing.T) {
	type fields struct {
		b []byte
		i int
	}
	type args struct {
		sep byte
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
			p := &Parser{
				b: tt.fields.b,
				i: tt.fields.i,
			}
			got, got1 := p.ReadSep(tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ReadSep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parser.ReadSep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
