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

func Test_mysqlRows_Columns(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			if got := rows.Columns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRows.Columns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRows_ColumnTypeDatabaseTypeName(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			if got := rows.ColumnTypeDatabaseTypeName(tt.args.i); got != tt.want {
				t.Errorf("mysqlRows.ColumnTypeDatabaseTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRows_ColumnTypeNullable(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	type args struct {
		i int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNullable bool
		wantOk       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			gotNullable, gotOk := rows.ColumnTypeNullable(tt.args.i)
			if gotNullable != tt.wantNullable {
				t.Errorf("mysqlRows.ColumnTypeNullable() gotNullable = %v, want %v", gotNullable, tt.wantNullable)
			}
			if gotOk != tt.wantOk {
				t.Errorf("mysqlRows.ColumnTypeNullable() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_mysqlRows_ColumnTypePrecisionScale(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
		want1  int64
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			got, got1, got2 := rows.ColumnTypePrecisionScale(tt.args.i)
			if got != tt.want {
				t.Errorf("mysqlRows.ColumnTypePrecisionScale() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("mysqlRows.ColumnTypePrecisionScale() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("mysqlRows.ColumnTypePrecisionScale() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_mysqlRows_ColumnTypeScanType(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			if got := rows.ColumnTypeScanType(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRows.ColumnTypeScanType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRows_Close(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
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
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			if err := rows.Close(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlRows.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlRows_HasNextResultSet(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	tests := []struct {
		name   string
		fields fields
		wantB  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			if gotB := rows.HasNextResultSet(); gotB != tt.wantB {
				t.Errorf("mysqlRows.HasNextResultSet() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_mysqlRows_nextResultSet(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			got, err := rows.nextResultSet()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRows.nextResultSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlRows.nextResultSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRows_nextNotEmptyResultSet(t *testing.T) {
	type fields struct {
		mc     *mysqlConn
		rs     resultSet
		finish func()
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := &mysqlRows{
				mc:     tt.fields.mc,
				rs:     tt.fields.rs,
				finish: tt.fields.finish,
			}
			got, err := rows.nextNotEmptyResultSet()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRows.nextNotEmptyResultSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlRows.nextNotEmptyResultSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryRows_NextResultSet(t *testing.T) {
	type fields struct {
		mysqlRows mysqlRows
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
			rows := &binaryRows{
				mysqlRows: tt.fields.mysqlRows,
			}
			if err := rows.NextResultSet(); (err != nil) != tt.wantErr {
				t.Errorf("binaryRows.NextResultSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_binaryRows_Next(t *testing.T) {
	type fields struct {
		mysqlRows mysqlRows
	}
	type args struct {
		dest []driver.Value
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
			rows := &binaryRows{
				mysqlRows: tt.fields.mysqlRows,
			}
			if err := rows.Next(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("binaryRows.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_textRows_NextResultSet(t *testing.T) {
	type fields struct {
		mysqlRows mysqlRows
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
			rows := &textRows{
				mysqlRows: tt.fields.mysqlRows,
			}
			if err := rows.NextResultSet(); (err != nil) != tt.wantErr {
				t.Errorf("textRows.NextResultSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_textRows_Next(t *testing.T) {
	type fields struct {
		mysqlRows mysqlRows
	}
	type args struct {
		dest []driver.Value
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
			rows := &textRows{
				mysqlRows: tt.fields.mysqlRows,
			}
			if err := rows.Next(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("textRows.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
