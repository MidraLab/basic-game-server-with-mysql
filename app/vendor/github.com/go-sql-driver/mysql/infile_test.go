// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"context"
	"io"
	"net"
	"testing"
	"time"
)

func TestRegisterLocalFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterLocalFile(tt.args.filePath)
		})
	}
}

func TestDeregisterLocalFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeregisterLocalFile(tt.args.filePath)
		})
	}
}

func TestRegisterReaderHandler(t *testing.T) {
	type args struct {
		name    string
		handler func() io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterReaderHandler(tt.args.name, tt.args.handler)
		})
	}
}

func TestDeregisterReaderHandler(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeregisterReaderHandler(tt.args.name)
		})
	}
}

func Test_deferredClose(t *testing.T) {
	type args struct {
		err    *error
		closer io.Closer
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deferredClose(tt.args.err, tt.args.closer)
		})
	}
}

func Test_mysqlConn_handleInFileRequest(t *testing.T) {
	type fields struct {
		buf              buffer
		netConn          net.Conn
		rawConn          net.Conn
		affectedRows     uint64
		insertId         uint64
		cfg              *Config
		maxAllowedPacket int
		maxWriteSize     int
		writeTimeout     time.Duration
		flags            clientFlag
		status           statusFlag
		sequence         uint8
		parseTime        bool
		reset            bool
		watching         bool
		watcher          chan<- context.Context
		closech          chan struct{}
		finished         chan<- struct{}
		canceled         atomicError
		closed           atomicBool
	}
	type args struct {
		name string
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
			mc := &mysqlConn{
				buf:              tt.fields.buf,
				netConn:          tt.fields.netConn,
				rawConn:          tt.fields.rawConn,
				affectedRows:     tt.fields.affectedRows,
				insertId:         tt.fields.insertId,
				cfg:              tt.fields.cfg,
				maxAllowedPacket: tt.fields.maxAllowedPacket,
				maxWriteSize:     tt.fields.maxWriteSize,
				writeTimeout:     tt.fields.writeTimeout,
				flags:            tt.fields.flags,
				status:           tt.fields.status,
				sequence:         tt.fields.sequence,
				parseTime:        tt.fields.parseTime,
				reset:            tt.fields.reset,
				watching:         tt.fields.watching,
				watcher:          tt.fields.watcher,
				closech:          tt.fields.closech,
				finished:         tt.fields.finished,
				canceled:         tt.fields.canceled,
				closed:           tt.fields.closed,
			}
			if err := mc.handleInFileRequest(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.handleInFileRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
