package msgpack

import (
	"io"
	"reflect"
	"testing"
)

func Test_getDecoder(t *testing.T) {
	type args struct {
		typ reflect.Type
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
			if got := getDecoder(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDecoder(t *testing.T) {
	type args struct {
		typ reflect.Type
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
			if got := _getDecoder(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_getDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ptrValueDecoder(t *testing.T) {
	type args struct {
		typ reflect.Type
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
			if got := ptrValueDecoder(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ptrValueDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addrDecoder(t *testing.T) {
	type args struct {
		fn decoderFunc
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
			if got := addrDecoder(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addrDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nilAwareDecoder(t *testing.T) {
	type args struct {
		typ reflect.Type
		fn  decoderFunc
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
			if got := nilAwareDecoder(tt.args.typ, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nilAwareDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeBoolValue(t *testing.T) {
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
			if err := decodeBoolValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeBoolValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeInterfaceValue(t *testing.T) {
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
			if err := decodeInterfaceValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecoder_interfaceValue(t *testing.T) {
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
			if err := d.interfaceValue(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.interfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeUnsupportedValue(t *testing.T) {
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
			if err := decodeUnsupportedValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeUnsupportedValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_decodeCustomValue(t *testing.T) {
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
			if err := decodeCustomValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("decodeCustomValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unmarshalValue(t *testing.T) {
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
			if err := unmarshalValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unmarshalBinaryValue(t *testing.T) {
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
			if err := unmarshalBinaryValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalBinaryValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unmarshalTextValue(t *testing.T) {
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
			if err := unmarshalTextValue(tt.args.d, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalTextValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
