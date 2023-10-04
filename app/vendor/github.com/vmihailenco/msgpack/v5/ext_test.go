package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func TestRegisterExt(t *testing.T) {
	type args struct {
		extID int8
		value MarshalerUnmarshaler
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterExt(tt.args.extID, tt.args.value)
		})
	}
}

func TestUnregisterExt(t *testing.T) {
	type args struct {
		extID int8
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UnregisterExt(tt.args.extID)
		})
	}
}

func TestRegisterExtEncoder(t *testing.T) {
	type args struct {
		extID   int8
		value   interface{}
		encoder func(enc *Encoder, v reflect.Value) ([]byte, error)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterExtEncoder(tt.args.extID, tt.args.value, tt.args.encoder)
		})
	}
}

func Test_unregisterExtEncoder(t *testing.T) {
	type args struct {
		extID int8
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unregisterExtEncoder(tt.args.extID)
		})
	}
}

func Test_makeExtEncoder(t *testing.T) {
	type args struct {
		extID   int8
		typ     reflect.Type
		encoder func(enc *Encoder, v reflect.Value) ([]byte, error)
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeExtEncoder(tt.args.extID, tt.args.typ, tt.args.encoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeExtEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeExtEncoderAddr(t *testing.T) {
	type args struct {
		extEncoder encoderFunc
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeExtEncoderAddr(tt.args.extEncoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeExtEncoderAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterExtDecoder(t *testing.T) {
	type args struct {
		extID   int8
		value   interface{}
		decoder func(dec *Decoder, v reflect.Value, extLen int) error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterExtDecoder(tt.args.extID, tt.args.value, tt.args.decoder)
		})
	}
}

func Test_unregisterExtDecoder(t *testing.T) {
	type args struct {
		extID int8
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unregisterExtDecoder(tt.args.extID)
		})
	}
}

func Test_makeExtDecoder(t *testing.T) {
	type args struct {
		wantedExtID int8
		typ         reflect.Type
		decoder     func(d *Decoder, v reflect.Value, extLen int) error
	}
	tests := []struct {
		name string
		args args
		want decoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeExtDecoder(tt.args.wantedExtID, tt.args.typ, tt.args.decoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeExtDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeExtDecoderAddr(t *testing.T) {
	type args struct {
		extDecoder decoderFunc
	}
	tests := []struct {
		name string
		args args
		want decoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeExtDecoderAddr(tt.args.extDecoder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeExtDecoderAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_EncodeExtHeader(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		extID  int8
		extLen int
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
			if err := e.EncodeExtHeader(tt.args.extID, tt.args.extLen); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeExtHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeExtLen(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		l int
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
			if err := e.encodeExtLen(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeExtLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_DecodeExtHeader(t *testing.T) {
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
		name       string
		fields     fields
		wantExtID  int8
		wantExtLen int
		wantErr    bool
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
			gotExtID, gotExtLen, err := d.DecodeExtHeader()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeExtHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExtID != tt.wantExtID {
				t.Errorf("Decoder.DecodeExtHeader() gotExtID = %v, want %v", gotExtID, tt.wantExtID)
			}
			if gotExtLen != tt.wantExtLen {
				t.Errorf("Decoder.DecodeExtHeader() gotExtLen = %v, want %v", gotExtLen, tt.wantExtLen)
			}
		})
	}
}

func TestDecoder_extHeader(t *testing.T) {
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
		want    int8
		want1   int
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
			got, got1, err := d.extHeader(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.extHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.extHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Decoder.extHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDecoder_parseExtLen(t *testing.T) {
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
			got, err := d.parseExtLen(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.parseExtLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.parseExtLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeInterfaceExt(t *testing.T) {
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
		want    interface{}
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
			got, err := d.decodeInterfaceExt(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeInterfaceExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.decodeInterfaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_skipExt(t *testing.T) {
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
			if err := d.skipExt(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.skipExt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_skipExtHeader(t *testing.T) {
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
			if err := d.skipExtHeader(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.skipExtHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_extHeaderLen(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extHeaderLen(tt.args.c); got != tt.want {
				t.Errorf("extHeaderLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
