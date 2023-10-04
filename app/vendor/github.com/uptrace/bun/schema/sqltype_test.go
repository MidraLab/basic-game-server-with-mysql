package schema

import (
	"reflect"
	"testing"
	"time"
)

func TestDiscoverSQLType(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiscoverSQLType(tt.args.typ); got != tt.want {
				t.Errorf("DiscoverSQLType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := NullTime{
				Time: tt.fields.Time,
			}
			got, err := tm.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		b []byte
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
			tm := &NullTime{
				Time: tt.fields.Time,
			}
			if err := tm.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("NullTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullTime_AppendQuery(t *testing.T) {
	type fields struct {
		Time time.Time
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
			tm := NullTime{
				Time: tt.fields.Time,
			}
			got, err := tm.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("NullTime.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullTime.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullTime_Scan(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		src interface{}
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
			tm := &NullTime{
				Time: tt.fields.Time,
			}
			if err := tm.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("NullTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
