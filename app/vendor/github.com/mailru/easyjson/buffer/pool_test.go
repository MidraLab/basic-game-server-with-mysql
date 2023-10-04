// Package buffer implements a buffer for serialization, consisting of a chain of []byte-s to
// reduce copying and to allow reuse of individual chunks.
package buffer

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_initBuffers(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initBuffers()
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		cfg PoolConfig
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.cfg)
		})
	}
}

func Test_putBuf(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			putBuf(tt.args.buf)
		})
	}
}

func Test_getBuf(t *testing.T) {
	type args struct {
		size int
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
			if got := getBuf(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBuf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_EnsureSpace(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		s int
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.EnsureSpace(tt.args.s)
		})
	}
}

func TestBuffer_ensureSpaceSlow(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		s int
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.ensureSpaceSlow(tt.args.s)
		})
	}
}

func TestBuffer_AppendByte(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		data byte
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.AppendByte(tt.args.data)
		})
	}
}

func TestBuffer_AppendBytes(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.AppendBytes(tt.args.data)
		})
	}
}

func TestBuffer_appendBytesSlow(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.appendBytesSlow(tt.args.data)
		})
	}
}

func TestBuffer_AppendString(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		data string
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.AppendString(tt.args.data)
		})
	}
}

func TestBuffer_appendStringSlow(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		data string
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			b.appendStringSlow(tt.args.data)
		})
	}
}

func TestBuffer_Size(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			if got := b.Size(); got != tt.want {
				t.Errorf("Buffer.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuffer_DumpTo(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	tests := []struct {
		name        string
		fields      fields
		wantWritten int
		wantW       string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			w := &bytes.Buffer{}
			gotWritten, err := b.DumpTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buffer.DumpTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWritten != tt.wantWritten {
				t.Errorf("Buffer.DumpTo() = %v, want %v", gotWritten, tt.wantWritten)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Buffer.DumpTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestBuffer_BuildBytes(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	type args struct {
		reuse [][]byte
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
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			if got := b.BuildBytes(tt.args.reuse...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buffer.BuildBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCloser_Read(t *testing.T) {
	type fields struct {
		offset int
		bufs   [][]byte
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &readCloser{
				offset: tt.fields.offset,
				bufs:   tt.fields.bufs,
			}
			gotN, err := r.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCloser.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("readCloser.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_readCloser_Close(t *testing.T) {
	type fields struct {
		offset int
		bufs   [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &readCloser{
				offset: tt.fields.offset,
				bufs:   tt.fields.bufs,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("readCloser.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuffer_ReadCloser(t *testing.T) {
	type fields struct {
		Buf    []byte
		toPool []byte
		bufs   [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   io.ReadCloser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				Buf:    tt.fields.Buf,
				toPool: tt.fields.toPool,
				bufs:   tt.fields.bufs,
			}
			if got := b.ReadCloser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buffer.ReadCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
