// Package jwriter contains a JSON writer.
package jwriter

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/mailru/easyjson/buffer"
)

func TestWriter_Size(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			if got := w.Size(); got != tt.want {
				t.Errorf("Writer.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_DumpTo(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantWritten int
		wantOut     string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			out := &bytes.Buffer{}
			gotWritten, err := w.DumpTo(out)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.DumpTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWritten != tt.wantWritten {
				t.Errorf("Writer.DumpTo() = %v, want %v", gotWritten, tt.wantWritten)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Writer.DumpTo() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestWriter_BuildBytes(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		reuse [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			got, err := w.BuildBytes(tt.args.reuse...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.BuildBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Writer.BuildBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_ReadCloser(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			got, err := w.ReadCloser()
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.ReadCloser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Writer.ReadCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_RawByte(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		c byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.RawByte(tt.args.c)
		})
	}
}

func TestWriter_RawString(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.RawString(tt.args.s)
		})
	}
}

func TestWriter_Raw(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		data []byte
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Raw(tt.args.data, tt.args.err)
		})
	}
}

func TestWriter_RawText(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		data []byte
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.RawText(tt.args.data, tt.args.err)
		})
	}
}

func TestWriter_Base64Bytes(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Base64Bytes(tt.args.data)
		})
	}
}

func TestWriter_Uint8(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint8(tt.args.n)
		})
	}
}

func TestWriter_Uint16(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint16(tt.args.n)
		})
	}
}

func TestWriter_Uint32(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint32(tt.args.n)
		})
	}
}

func TestWriter_Uint(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint(tt.args.n)
		})
	}
}

func TestWriter_Uint64(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint64(tt.args.n)
		})
	}
}

func TestWriter_Int8(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int8(tt.args.n)
		})
	}
}

func TestWriter_Int16(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int16(tt.args.n)
		})
	}
}

func TestWriter_Int32(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int32(tt.args.n)
		})
	}
}

func TestWriter_Int(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int(tt.args.n)
		})
	}
}

func TestWriter_Int64(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int64(tt.args.n)
		})
	}
}

func TestWriter_Uint8Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint8Str(tt.args.n)
		})
	}
}

func TestWriter_Uint16Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint16Str(tt.args.n)
		})
	}
}

func TestWriter_Uint32Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint32Str(tt.args.n)
		})
	}
}

func TestWriter_UintStr(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.UintStr(tt.args.n)
		})
	}
}

func TestWriter_Uint64Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Uint64Str(tt.args.n)
		})
	}
}

func TestWriter_UintptrStr(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n uintptr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.UintptrStr(tt.args.n)
		})
	}
}

func TestWriter_Int8Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int8Str(tt.args.n)
		})
	}
}

func TestWriter_Int16Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int16Str(tt.args.n)
		})
	}
}

func TestWriter_Int32Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int32Str(tt.args.n)
		})
	}
}

func TestWriter_IntStr(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.IntStr(tt.args.n)
		})
	}
}

func TestWriter_Int64Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Int64Str(tt.args.n)
		})
	}
}

func TestWriter_Float32(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Float32(tt.args.n)
		})
	}
}

func TestWriter_Float32Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Float32Str(tt.args.n)
		})
	}
}

func TestWriter_Float64(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Float64(tt.args.n)
		})
	}
}

func TestWriter_Float64Str(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Float64Str(tt.args.n)
		})
	}
}

func TestWriter_Bool(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.Bool(tt.args.v)
		})
	}
}

func Test_getTable(t *testing.T) {
	type args struct {
		falseValues []int
	}
	tests := []struct {
		name string
		args args
		want [128]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTable(tt.args.falseValues...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_String(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.String(tt.args.s)
		})
	}
}

func TestWriter_base64(t *testing.T) {
	type fields struct {
		Flags        Flags
		Error        error
		Buffer       buffer.Buffer
		NoEscapeHTML bool
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Flags:        tt.fields.Flags,
				Error:        tt.fields.Error,
				Buffer:       tt.fields.Buffer,
				NoEscapeHTML: tt.fields.NoEscapeHTML,
			}
			w.base64(tt.args.in)
		})
	}
}
