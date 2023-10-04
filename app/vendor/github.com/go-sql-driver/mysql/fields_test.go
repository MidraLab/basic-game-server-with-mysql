// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2017 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"reflect"
	"testing"
)

func Test_mysqlField_typeDatabaseName(t *testing.T) {
	type fields struct {
		tableName string
		name      string
		length    uint32
		flags     fieldFlag
		fieldType fieldType
		decimals  byte
		charSet   uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := &mysqlField{
				tableName: tt.fields.tableName,
				name:      tt.fields.name,
				length:    tt.fields.length,
				flags:     tt.fields.flags,
				fieldType: tt.fields.fieldType,
				decimals:  tt.fields.decimals,
				charSet:   tt.fields.charSet,
			}
			if got := mf.typeDatabaseName(); got != tt.want {
				t.Errorf("mysqlField.typeDatabaseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlField_scanType(t *testing.T) {
	type fields struct {
		tableName string
		name      string
		length    uint32
		flags     fieldFlag
		fieldType fieldType
		decimals  byte
		charSet   uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := &mysqlField{
				tableName: tt.fields.tableName,
				name:      tt.fields.name,
				length:    tt.fields.length,
				flags:     tt.fields.flags,
				fieldType: tt.fields.fieldType,
				decimals:  tt.fields.decimals,
				charSet:   tt.fields.charSet,
			}
			if got := mf.scanType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlField.scanType() = %v, want %v", got, tt.want)
			}
		})
	}
}
