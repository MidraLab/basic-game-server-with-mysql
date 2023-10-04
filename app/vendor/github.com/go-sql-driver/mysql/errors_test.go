// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import "testing"

func TestSetLogger(t *testing.T) {
	type args struct {
		logger Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetLogger(tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("SetLogger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMySQLError_Error(t *testing.T) {
	type fields struct {
		Number   uint16
		SQLState [5]byte
		Message  string
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
			me := &MySQLError{
				Number:   tt.fields.Number,
				SQLState: tt.fields.SQLState,
				Message:  tt.fields.Message,
			}
			if got := me.Error(); got != tt.want {
				t.Errorf("MySQLError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQLError_Is(t *testing.T) {
	type fields struct {
		Number   uint16
		SQLState [5]byte
		Message  string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &MySQLError{
				Number:   tt.fields.Number,
				SQLState: tt.fields.SQLState,
				Message:  tt.fields.Message,
			}
			if got := me.Is(tt.args.err); got != tt.want {
				t.Errorf("MySQLError.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
