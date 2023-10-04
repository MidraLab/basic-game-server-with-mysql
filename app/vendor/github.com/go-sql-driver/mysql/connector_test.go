// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2018 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"context"
	"database/sql/driver"
	"reflect"
	"testing"
)

func Test_connector_Connect(t *testing.T) {
	type fields struct {
		cfg *Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connector{
				cfg: tt.fields.cfg,
			}
			got, err := c.Connect(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("connector.Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connector.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_connector_Driver(t *testing.T) {
	type fields struct {
		cfg *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   driver.Driver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connector{
				cfg: tt.fields.cfg,
			}
			if got := c.Driver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connector.Driver() = %v, want %v", got, tt.want)
			}
		})
	}
}
