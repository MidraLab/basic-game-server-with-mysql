// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import "testing"

func Test_mysqlTx_Commit(t *testing.T) {
	type fields struct {
		mc *mysqlConn
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
			tx := &mysqlTx{
				mc: tt.fields.mc,
			}
			if err := tx.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlTx.Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlTx_Rollback(t *testing.T) {
	type fields struct {
		mc *mysqlConn
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
			tx := &mysqlTx{
				mc: tt.fields.mc,
			}
			if err := tx.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlTx.Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
