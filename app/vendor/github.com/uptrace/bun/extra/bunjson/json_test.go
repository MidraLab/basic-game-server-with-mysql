package bunjson

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestStdProvider_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		s       StdProvider
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StdProvider{}
			got, err := s.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StdProvider.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdProvider.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStdProvider_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		s       StdProvider
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StdProvider{}
			if err := s.Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("StdProvider.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStdProvider_NewEncoder(t *testing.T) {
	tests := []struct {
		name  string
		s     StdProvider
		want  Encoder
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StdProvider{}
			w := &bytes.Buffer{}
			if got := s.NewEncoder(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdProvider.NewEncoder() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("StdProvider.NewEncoder() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestStdProvider_NewDecoder(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		s    StdProvider
		args args
		want Decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StdProvider{}
			if got := s.NewDecoder(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StdProvider.NewDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
