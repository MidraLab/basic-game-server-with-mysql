package swag

import (
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		name    string
		swagger Swagger
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.name, tt.args.swagger)
		})
	}
}

func TestGetSwagger(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Swagger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSwagger(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSwagger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDoc(t *testing.T) {
	type args struct {
		optionalName []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDoc(tt.args.optionalName...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
