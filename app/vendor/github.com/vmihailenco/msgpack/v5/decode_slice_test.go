package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func TestDecoder_DecodeArrayLen(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			got, err := d.DecodeArrayLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeArrayLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.DecodeArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_arrayLen(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	type args struct {
		c byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			got, err := d.arrayLen(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.arrayLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.arrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeStringSliceValue(t *testing.T) {
	type args struct {
		d *Decoder
		v reflect.Value
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
			if err := decodeStringSliceValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeStringSliceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeStringSlicePtr(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	type args struct {
		ptr *[]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			if err := d.decodeStringSlicePtr(tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeStringSlicePtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_makeStrings(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStrings(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeSliceValue(t *testing.T) {
	type args struct {
		d *Decoder
		v reflect.Value
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
			if err := decodeSliceValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeSliceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_growSliceValue(t *testing.T) {
	type args struct {
		v reflect.Value
		n int
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := growSliceValue(tt.args.v, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("growSliceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeArrayValue(t *testing.T) {
	type args struct {
		d *Decoder
		v reflect.Value
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
			if err := decodeArrayValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeArrayValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_DecodeSlice(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			got, err := d.DecodeSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeSlice(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	type args struct {
		c byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			got, err := d.decodeSlice(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.decodeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_skipSlice(t *testing.T) {
	type fields struct {
		r          io.Reader
		s          io.ByteScanner
		buf        []byte
		rec        []byte
		dict       []string
		flags      uint32
		structTag  string
		mapDecoder func(*Decoder) (interface{}, error)
	}
	type args struct {
		c byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Decoder{
				r:          tt.fields.r,
				s:          tt.fields.s,
				buf:        tt.fields.buf,
				rec:        tt.fields.rec,
				dict:       tt.fields.dict,
				flags:      tt.fields.flags,
				structTag:  tt.fields.structTag,
				mapDecoder: tt.fields.mapDecoder,
			}
			if err := d.skipSlice(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.skipSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
