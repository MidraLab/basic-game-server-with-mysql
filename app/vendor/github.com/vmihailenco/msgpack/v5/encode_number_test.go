package msgpack

import (
	"reflect"
	"testing"
)

func TestEncoder_EncodeUint8(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint8
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
			if err := e.EncodeUint8(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeUint8() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeUint8Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint8
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
			if err := e.encodeUint8Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeUint8Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeUint16(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint16
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
			if err := e.EncodeUint16(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeUint16() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeUint16Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint16
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
			if err := e.encodeUint16Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeUint16Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeUint32(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint32
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
			if err := e.EncodeUint32(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeUint32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeUint32Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint32
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
			if err := e.encodeUint32Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeUint32Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeUint64(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint64
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
			if err := e.EncodeUint64(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeUint64Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint64
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
			if err := e.encodeUint64Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeUint64Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeInt8(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int8
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
			if err := e.EncodeInt8(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeInt8() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInt8Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int8
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
			if err := e.encodeInt8Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInt8Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeInt16(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int16
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
			if err := e.EncodeInt16(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeInt16() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInt16Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int16
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
			if err := e.encodeInt16Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInt16Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeInt32(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int32
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
			if err := e.EncodeInt32(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeInt32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInt32Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int32
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
			if err := e.encodeInt32Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInt32Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeInt64(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int64
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
			if err := e.EncodeInt64(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeInt64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeInt64Cond(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int64
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
			if err := e.encodeInt64Cond(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.encodeInt64Cond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeUint(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n uint64
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
			if err := e.EncodeUint(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeUint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeInt(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n int64
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
			if err := e.EncodeInt(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeInt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeFloat32(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n float32
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
			if err := e.EncodeFloat32(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeFloat32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeFloat64(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		n float64
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
			if err := e.EncodeFloat64(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_write1(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		code byte
		n    uint8
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
			if err := e.write1(tt.args.code, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.write1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_write2(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		code byte
		n    uint16
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
			if err := e.write2(tt.args.code, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.write2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_write4(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		code byte
		n    uint32
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
			if err := e.write4(tt.args.code, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.write4() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_write8(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		code byte
		n    uint64
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
			if err := e.write8(tt.args.code, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.write8() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUintValue(t *testing.T) {
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
			if err := encodeUintValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUintValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeIntValue(t *testing.T) {
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
			if err := encodeIntValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeIntValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUint8CondValue(t *testing.T) {
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
			if err := encodeUint8CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUint8CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUint16CondValue(t *testing.T) {
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
			if err := encodeUint16CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUint16CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUint32CondValue(t *testing.T) {
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
			if err := encodeUint32CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUint32CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUint64CondValue(t *testing.T) {
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
			if err := encodeUint64CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUint64CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInt8CondValue(t *testing.T) {
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
			if err := encodeInt8CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInt8CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInt16CondValue(t *testing.T) {
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
			if err := encodeInt16CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInt16CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInt32CondValue(t *testing.T) {
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
			if err := encodeInt32CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInt32CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInt64CondValue(t *testing.T) {
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
			if err := encodeInt64CondValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInt64CondValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeFloat32Value(t *testing.T) {
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
			if err := encodeFloat32Value(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeFloat32Value() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeFloat64Value(t *testing.T) {
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
			if err := encodeFloat64Value(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeFloat64Value() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
