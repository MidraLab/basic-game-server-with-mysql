// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"reflect"
	"testing"
)

func TestUUID_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uuid.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_UnmarshalText(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		uuid    *UUID
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uuid.UnmarshalText(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uuid.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_UnmarshalBinary(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		uuid    *UUID
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uuid.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
