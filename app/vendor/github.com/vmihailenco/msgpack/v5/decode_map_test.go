package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func Test_decodeMapValue(t *testing.T) {
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
			if err := decodeMapValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeMapValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeMapDefault(t *testing.T) {
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
			got, err := d.decodeMapDefault()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeMapDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.decodeMapDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeMapLen(t *testing.T) {
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
			got, err := d.DecodeMapLen()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeMapLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.DecodeMapLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_mapLen(t *testing.T) {
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
			got, err := d.mapLen(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.mapLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decoder.mapLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeMapStringStringValue(t *testing.T) {
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
			if err := decodeMapStringStringValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeMapStringStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeMapStringStringPtr(t *testing.T) {
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
		ptr *map[string]string
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
			if err := d.decodeMapStringStringPtr(tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeMapStringStringPtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeMapStringInterfaceValue(t *testing.T) {
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
			if err := decodeMapStringInterfaceValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeMapStringInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeMapStringInterfacePtr(t *testing.T) {
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
		ptr *map[string]interface{}
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
			if err := d.decodeMapStringInterfacePtr(tt.args.ptr); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeMapStringInterfacePtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_DecodeMap(t *testing.T) {
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
		want    map[string]interface{}
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
			got, err := d.DecodeMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeUntypedMap(t *testing.T) {
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
		want    map[interface{}]interface{}
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
			got, err := d.DecodeUntypedMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeUntypedMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeUntypedMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeTypedMap(t *testing.T) {
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
			got, err := d.DecodeTypedMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeTypedMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeTypedMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeTypedMapValue(t *testing.T) {
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
		v reflect.Value
		n int
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
			if err := d.decodeTypedMapValue(tt.args.v, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeTypedMapValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_skipMap(t *testing.T) {
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
			if err := d.skipMap(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.skipMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeStructValue(t *testing.T) {
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
			if err := decodeStructValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeStructValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_decodeStruct(t *testing.T) {
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
		v reflect.Value
		n int
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
			if err := d.decodeStruct(tt.args.v, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
