package schema

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		v     interface{}
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
			if got := Append(tt.args.fmter, tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
		want QueryAppender
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inValues_AppendQuery(t *testing.T) {
	type fields struct {
		slice reflect.Value
		err   error
	}
	type args struct {
		fmter Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &inValues{
				slice: tt.fields.slice,
				err:   tt.fields.err,
			}
			got, err := in.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("inValues.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inValues.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendIn(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
		slice reflect.Value
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
			if got := appendIn(tt.args.fmter, tt.args.b, tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullZero(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want QueryAppender
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullZero(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nullZero_AppendQuery(t *testing.T) {
	type fields struct {
		value interface{}
	}
	type args struct {
		fmter Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nz := nullZero{
				value: tt.fields.value,
			}
			got, err := nz.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("nullZero.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nullZero.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
