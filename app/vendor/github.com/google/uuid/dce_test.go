// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"reflect"
	"testing"
)

func TestNewDCESecurity(t *testing.T) {
	type args struct {
		domain Domain
		id     uint32
	}
	tests := []struct {
		name    string
		args    args
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDCESecurity(tt.args.domain, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDCESecurity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDCESecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDCEPerson(t *testing.T) {
	tests := []struct {
		name    string
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDCEPerson()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDCEPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDCEPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDCEGroup(t *testing.T) {
	tests := []struct {
		name    string
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDCEGroup()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDCEGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDCEGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Domain(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want Domain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.Domain(); got != tt.want {
				t.Errorf("UUID.Domain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_ID(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.ID(); got != tt.want {
				t.Errorf("UUID.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomain_String(t *testing.T) {
	tests := []struct {
		name string
		d    Domain
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Domain.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
