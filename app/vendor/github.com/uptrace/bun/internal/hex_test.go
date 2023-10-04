package internal

import (
	"reflect"
	"testing"
)

func TestNewHexEncoder(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want *HexEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexEncoder(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexEncoder_Bytes(t *testing.T) {
	type fields struct {
		b       []byte
		written bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &HexEncoder{
				b:       tt.fields.b,
				written: tt.fields.written,
			}
			if got := enc.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexEncoder.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexEncoder_Write(t *testing.T) {
	type fields struct {
		b       []byte
		written bool
	}
	type args struct {
		b []byte
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
			enc := &HexEncoder{
				b:       tt.fields.b,
				written: tt.fields.written,
			}
			got, err := enc.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexEncoder.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HexEncoder.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexEncoder_Close(t *testing.T) {
	type fields struct {
		b       []byte
		written bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := &HexEncoder{
				b:       tt.fields.b,
				written: tt.fields.written,
			}
			if err := enc.Close(); (err != nil) != tt.wantErr {
				t.Errorf("HexEncoder.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
