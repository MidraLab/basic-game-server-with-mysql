package msgpack

import (
	"io"
	"reflect"
	"testing"
	"time"
)

func Test_timeEncoder(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := timeEncoder(tt.args.e, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("timeEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeDecoder(t *testing.T) {
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
			if err := timeDecoder(tt.args.d, tt.args.v, tt.args.extLen); (err != nil) != tt.wantErr {
				t.Errorf("timeDecoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_EncodeTime(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		tm time.Time
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
			if err := e.EncodeTime(tt.args.tm); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.EncodeTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_encodeTime(t *testing.T) {
	type fields struct {
		w         writer
		buf       []byte
		timeBuf   []byte
		dict      map[string]int
		flags     uint32
		structTag string
	}
	type args struct {
		tm time.Time
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
			e := &Encoder{
				w:         tt.fields.w,
				buf:       tt.fields.buf,
				timeBuf:   tt.fields.timeBuf,
				dict:      tt.fields.dict,
				flags:     tt.fields.flags,
				structTag: tt.fields.structTag,
			}
			if got := e.encodeTime(tt.args.tm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder.encodeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_DecodeTime(t *testing.T) {
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
		want    time.Time
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
			got, err := d.DecodeTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.DecodeTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.DecodeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_decodeTime(t *testing.T) {
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
		want    time.Time
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
			got, err := d.decodeTime(tt.args.extLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.decodeTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.decodeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
