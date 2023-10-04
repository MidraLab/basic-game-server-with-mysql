// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import "testing"

func Test_mysqlResult_LastInsertId(t *testing.T) {
	type fields struct {
		affectedRows int64
		insertId     int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mysqlResult{
				affectedRows: tt.fields.affectedRows,
				insertId:     tt.fields.insertId,
			}
			got, err := res.LastInsertId()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlResult.LastInsertId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlResult.LastInsertId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlResult_RowsAffected(t *testing.T) {
	type fields struct {
		affectedRows int64
		insertId     int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mysqlResult{
				affectedRows: tt.fields.affectedRows,
				insertId:     tt.fields.insertId,
			}
			got, err := res.RowsAffected()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlResult.RowsAffected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlResult.RowsAffected() = %v, want %v", got, tt.want)
			}
		})
	}
}
