// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"context"
	"database/sql/driver"
	"net"
	"reflect"
	"testing"
	"time"
)

func Test_mysqlConn_handleParams(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
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
			if err := mc.handleParams(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.handleParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_markBadConn(t *testing.T) {
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
		err error
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
			if err := mc.markBadConn(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.markBadConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_Begin(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		want    driver.Tx
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
			got, err := mc.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_begin(t *testing.T) {
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
		readOnly bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Tx
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
			got, err := mc.begin(tt.args.readOnly)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_Close(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
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
			if err := mc.Close(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_cleanup(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			mc.cleanup()
		})
	}
}

func Test_mysqlConn_error(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
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
			if err := mc.error(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_Prepare(t *testing.T) {
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
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Stmt
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
			got, err := mc.Prepare(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_interpolateParams(t *testing.T) {
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
		query string
		args  []driver.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
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
			got, err := mc.interpolateParams(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.interpolateParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlConn.interpolateParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_Exec(t *testing.T) {
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
		query string
		args  []driver.Value
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
			got, err := mc.Exec(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_exec(t *testing.T) {
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
		query string
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
			if err := mc.exec(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_Query(t *testing.T) {
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
		query string
		args  []driver.Value
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
			got, err := mc.Query(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_query(t *testing.T) {
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
		query string
		args  []driver.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *textRows
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
			got, err := mc.query(tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_getSystemVar(t *testing.T) {
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
		want    []byte
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
			got, err := mc.getSystemVar(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.getSystemVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.getSystemVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_cancel(t *testing.T) {
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
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			mc.cancel(tt.args.err)
		})
	}
}

func Test_mysqlConn_finish(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			mc.finish()
		})
	}
}

func Test_mysqlConn_Ping(t *testing.T) {
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
		ctx context.Context
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
			if err := mc.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_BeginTx(t *testing.T) {
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
		ctx  context.Context
		opts driver.TxOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Tx
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
			got, err := mc.BeginTx(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.BeginTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.BeginTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_QueryContext(t *testing.T) {
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
		ctx   context.Context
		query string
		args  []driver.NamedValue
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
			got, err := mc.QueryContext(tt.args.ctx, tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.QueryContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_ExecContext(t *testing.T) {
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
		ctx   context.Context
		query string
		args  []driver.NamedValue
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
			got, err := mc.ExecContext(tt.args.ctx, tt.args.query, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.ExecContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_PrepareContext(t *testing.T) {
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
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    driver.Stmt
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
			got, err := mc.PrepareContext(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.PrepareContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.PrepareContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_QueryContext(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		ctx  context.Context
		args []driver.NamedValue
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
			got, err := stmt.QueryContext(tt.args.ctx, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.QueryContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_ExecContext(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		ctx  context.Context
		args []driver.NamedValue
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
			got, err := stmt.ExecContext(tt.args.ctx, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlStmt.ExecContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_watchCancel(t *testing.T) {
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
		ctx context.Context
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
			if err := mc.watchCancel(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.watchCancel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_startWatcher(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			mc.startWatcher()
		})
	}
}

func Test_mysqlConn_CheckNamedValue(t *testing.T) {
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
			if err := mc.CheckNamedValue(tt.args.nv); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.CheckNamedValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_ResetSession(t *testing.T) {
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
		ctx context.Context
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
			if err := mc.ResetSession(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.ResetSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_IsValid(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
		want   bool
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
			if got := mc.IsValid(); got != tt.want {
				t.Errorf("mysqlConn.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
