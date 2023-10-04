// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package mysql provides a MySQL driver for Go's database/sql package.
//
// The driver should be used via the database/sql package:
//
//	import "database/sql"
//	import _ "github.com/go-sql-driver/mysql"
//
//	db, err := sql.Open("mysql", "user:password@/dbname")
//
// See https://github.com/go-sql-driver/mysql#usage for details
package mysql

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestRegisterDialContext(t *testing.T) {
	type args struct {
		net  string
		dial DialContextFunc
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterDialContext(tt.args.net, tt.args.dial)
		})
	}
}

func TestRegisterDial(t *testing.T) {
	type args struct {
		network string
		dial    DialFunc
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterDial(tt.args.network, tt.args.dial)
		})
	}
}

func TestMySQLDriver_Open(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		d       MySQLDriver
		args    args
		want    driver.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MySQLDriver{}
			got, err := d.Open(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLDriver.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLDriver.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConnector(t *testing.T) {
	type args struct {
		cfg *Config
	}
	tests := []struct {
		name    string
		args    args
		want    driver.Connector
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConnector(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConnector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQLDriver_OpenConnector(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		d       MySQLDriver
		args    args
		want    driver.Connector
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := MySQLDriver{}
			got, err := d.OpenConnector(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLDriver.OpenConnector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQLDriver.OpenConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}
