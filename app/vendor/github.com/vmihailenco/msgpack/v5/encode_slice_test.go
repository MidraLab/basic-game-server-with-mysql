package msgpack

import (
	"reflect"
	"testing"
)

func Test_encodeStringValue(t *testing.T) {
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
			if err := encodeStringValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeStringValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeByteSliceValue(t *testing.T) {
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
			if err := encodeByteSliceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeByteSliceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeByteArrayValue(t *testing.T) {
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
			if err := encodeByteArrayValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeByteArrayValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_grow(t *testing.T) {
	type args struct {
		b []byte
		n int
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
			if got := grow(tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoder_EncodeBytesLen(t *testing.T) {
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
			if err := e.EncodeBytesLen(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeBytesLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeStringLen(t *testing.T) {
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
			if err := e.encodeStringLen(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeStringLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeString(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		v string
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
			if err := e.EncodeString(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeNormalString(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		v string
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
			if err := e.encodeNormalString(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeNormalString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeBytes(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		v []byte
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
			if err := e.EncodeBytes(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeArrayLen(t *testing.T) {
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
			if err := e.EncodeArrayLen(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeArrayLen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeStringSliceValue(t *testing.T) {
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
			if err := encodeStringSliceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeStringSliceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeStringSlice(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		s []string
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
			if err := e.encodeStringSlice(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeStringSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeSliceValue(t *testing.T) {
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
			if err := encodeSliceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeSliceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeArrayValue(t *testing.T) {
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
			if err := encodeArrayValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeArrayValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
