package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func Test_decodeInternedStringExt(t *testing.T) {
	type args struct {
		d      *Decoder
		v      reflect.Value
		extLen int
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
			if err := decodeInternedStringExt(tt.args.d, tt.args.v, tt.args.extLen); (err != nil) != tt.wantErr {
				t.Errorf("decodeInternedStringExt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInternedInterfaceValue(t *testing.T) {
	type args struct {
		e *Encoder
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
			if err := encodeInternedInterfaceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInternedInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInternedStringValue(t *testing.T) {
	type args struct {
		e *Encoder
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
			if err := encodeInternedStringValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInternedStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInternedString(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		s      string
		intern bool
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
			e := &Encoder{
				w:         tt.fields.w,
				buf:       tt.fields.buf,
				timeBuf:   tt.fields.timeBuf,
				dict:      tt.fields.dict,
				flags:     tt.fields.flags,
				structTag: tt.fields.structTag,
			}
			if err := e.encodeInternedString(tt.args.s, tt.args.intern); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInternedString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInternedStringIndex(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		idx int
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
			e := &Encoder{
				w:         tt.fields.w,
				buf:       tt.fields.buf,
				timeBuf:   tt.fields.timeBuf,
				dict:      tt.fields.dict,
				flags:     tt.fields.flags,
				structTag: tt.fields.structTag,
			}
			if err := e.encodeInternedStringIndex(tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInternedStringIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeInternedInterfaceValue(t *testing.T) {
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
			if err := decodeInternedInterfaceValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeInternedInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeInternedStringValue(t *testing.T) {
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
			if err := decodeInternedStringValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeInternedStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeInternedString(t *testing.T) {
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
		intern bool
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
			got, err := d.decodeInternedString(tt.args.intern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeInternedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.decodeInternedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeInternedStringIndex(t *testing.T) {
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
		extLen int
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
			got, err := d.decodeInternedStringIndex(tt.args.extLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeInternedStringIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.decodeInternedStringIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_internedStringAtIndex(t *testing.T) {
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
		idx int
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
			got, err := d.internedStringAtIndex(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.internedStringAtIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.internedStringAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeInternedStringWithLen(t *testing.T) {
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
		n      int
		intern bool
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
			got, err := d.decodeInternedStringWithLen(tt.args.n, tt.args.intern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeInternedStringWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.decodeInternedStringWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
