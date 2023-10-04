// Copyright 2021 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestNullUUID_Scan(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
	}
	type args struct {
		value interface{}
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
			nu := &NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			if err := nu.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullUUID_Value(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			got, err := nu.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullUUID.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullUUID_MarshalBinary(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
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
			nu := NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			got, err := nu.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullUUID.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullUUID_UnmarshalBinary(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
	}
	type args struct {
		data []byte
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
			nu := &NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			if err := nu.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullUUID_MarshalText(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
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
			nu := NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			got, err := nu.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullUUID.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullUUID_UnmarshalText(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
	}
	type args struct {
		data []byte
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
			nu := &NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			if err := nu.UnmarshalText(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullUUID_MarshalJSON(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
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
			nu := NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			got, err := nu.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullUUID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullUUID_UnmarshalJSON(t *testing.T) {
	type fields struct {
		UUID  UUID
		Valid bool
	}
	type args struct {
		data []byte
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
			nu := &NullUUID{
				UUID:  tt.fields.UUID,
				Valid: tt.fields.Valid,
			}
			if err := nu.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("NullUUID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
