// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestNullTime_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		nt      *NullTime
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nt := &NullTime{}
			if err := nt.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("NullTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNullTime_Value(t *testing.T) {
	tests := []struct {
		name    string
		nt      NullTime
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nt := NullTime{}
			got, err := nt.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("NullTime.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullTime.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
