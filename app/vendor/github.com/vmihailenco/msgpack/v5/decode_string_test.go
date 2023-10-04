package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func TestDecoder_bytesLen(t *testing.T) {
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
			got, err := d.bytesLen(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.bytesLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.bytesLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeString(t *testing.T) {
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
		want    string
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
			got, err := d.DecodeString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_string(t *testing.T) {
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
		want    string
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
			got, err := d.string(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.string() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.string() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_stringWithLen(t *testing.T) {
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
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
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
			got, err := d.stringWithLen(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.stringWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.stringWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeStringValue(t *testing.T) {
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
			if err := decodeStringValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_DecodeBytesLen(t *testing.T) {
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
			got, err := d.DecodeBytesLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeBytesLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.DecodeBytesLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeBytes(t *testing.T) {
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
		want    []byte
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
			got, err := d.DecodeBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_bytes(t *testing.T) {
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
		b []byte
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
			got, err := d.bytes(tt.args.c, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeStringTemp(t *testing.T) {
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
		want    string
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
			got, err := d.decodeStringTemp()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeStringTemp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.decodeStringTemp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeBytesPtr(t *testing.T) {
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
		ptr *[]byte
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
			if err := d.decodeBytesPtr(tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeBytesPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_bytesPtr(t *testing.T) {
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
		c   byte
		ptr *[]byte
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
			if err := d.bytesPtr(tt.args.c, tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.bytesPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_skipBytes(t *testing.T) {
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
			if err := d.skipBytes(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.skipBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeBytesValue(t *testing.T) {
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
			if err := decodeBytesValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeBytesValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeByteArrayValue(t *testing.T) {
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
			if err := decodeByteArrayValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeByteArrayValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
