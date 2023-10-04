package msgpack

import (
	"reflect"
	"testing"
)

func Test_encodeMapValue(t *testing.T) {
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
			if err := encodeMapValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeMapValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeMapStringStringValue(t *testing.T) {
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
			if err := encodeMapStringStringValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeMapStringStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeMapStringInterfaceValue(t *testing.T) {
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
			if err := encodeMapStringInterfaceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeMapStringInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeMap(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		m map[string]interface{}
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
			if err := e.EncodeMap(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeMapSorted(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		m map[string]interface{}
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
			if err := e.EncodeMapSorted(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeMapSorted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeSortedMapStringString(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		m map[string]string
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
			if err := e.encodeSortedMapStringString(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeSortedMapStringString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeMapLen(t *testing.T) {
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
			if err := e.EncodeMapLen(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeMapLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeStructValue(t *testing.T) {
	type args struct {
		e     *Encoder
		strct reflect.Value
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
			if err := encodeStructValue(tt.args.e, tt.args.strct); (err != nil) != tt.wantErr {
				t.Errorf("encodeStructValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeStructValueAsArray(t *testing.T) {
	type args struct {
		e      *Encoder
		strct  reflect.Value
		fields []*field
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
			if err := encodeStructValueAsArray(tt.args.e, tt.args.strct, tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("encodeStructValueAsArray() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
