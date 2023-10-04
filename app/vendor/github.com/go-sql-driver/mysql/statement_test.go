// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
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

func Test_mysqlStmt_Close(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
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
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			if err := stmt.Close(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlStmt_NumInput(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			if got := stmt.NumInput(); got != tt.want {
				t.Errorf("mysqlStmt.NumInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_ColumnConverter(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   driver.ValueConverter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			if got := stmt.ColumnConverter(tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.ColumnConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_CheckNamedValue(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		nv *driver.NamedValue
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
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			if err := stmt.CheckNamedValue(tt.args.nv); (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.CheckNamedValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlStmt_Exec(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		args []driver.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			got, err := stmt.Exec(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_Query(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		args []driver.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			got, err := stmt.Query(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_query(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		args []driver.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *binaryRows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := &mysqlStmt{
				mc:         tt.fields.mc,
				id:         tt.fields.id,
				paramCount: tt.fields.paramCount,
			}
			got, err := stmt.query(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_converter_ConvertValue(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		c       converter
		args    args
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := converter{}
			got, err := c.ConvertValue(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("converter.ConvertValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("converter.ConvertValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callValuerValue(t *testing.T) {
	type args struct {
		vr driver.Valuer
	}
	tests := []struct {
		name    string
		args    args
		wantV   driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, err := callValuerValue(tt.args.vr)
			if (err != nil) != tt.wantErr {
				t.Errorf("callValuerValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("callValuerValue() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
