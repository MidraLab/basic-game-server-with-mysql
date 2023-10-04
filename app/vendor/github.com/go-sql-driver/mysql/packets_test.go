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

func Test_mysqlConn_readPacket(t *testing.T) {
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
			got, err := mc.readPacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.readPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_writePacket(t *testing.T) {
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
		data []byte
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
			if err := mc.writePacket(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writePacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_readHandshakePacket(t *testing.T) {
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
		name       string
		fields     fields
		wantData   []byte
		wantPlugin string
		wantErr    bool
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
			gotData, gotPlugin, err := mc.readHandshakePacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readHandshakePacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("mysqlConn.readHandshakePacket() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotPlugin != tt.wantPlugin {
				t.Errorf("mysqlConn.readHandshakePacket() gotPlugin = %v, want %v", gotPlugin, tt.wantPlugin)
			}
		})
	}
}

func Test_mysqlConn_writeHandshakeResponsePacket(t *testing.T) {
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
		authResp []byte
		plugin   string
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
			if err := mc.writeHandshakeResponsePacket(tt.args.authResp, tt.args.plugin); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writeHandshakeResponsePacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_writeAuthSwitchPacket(t *testing.T) {
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
		authData []byte
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
			if err := mc.writeAuthSwitchPacket(tt.args.authData); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writeAuthSwitchPacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_writeCommandPacket(t *testing.T) {
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
		command byte
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
			if err := mc.writeCommandPacket(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writeCommandPacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_writeCommandPacketStr(t *testing.T) {
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
		command byte
		arg     string
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
			if err := mc.writeCommandPacketStr(tt.args.command, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writeCommandPacketStr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_writeCommandPacketUint32(t *testing.T) {
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
		command byte
		arg     uint32
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
			if err := mc.writeCommandPacketUint32(tt.args.command, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.writeCommandPacketUint32() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_readAuthResult(t *testing.T) {
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
		want    []byte
		want1   string
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
			got, got1, err := mc.readAuthResult()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readAuthResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.readAuthResult() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("mysqlConn.readAuthResult() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_mysqlConn_readResultOK(t *testing.T) {
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
			if err := mc.readResultOK(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readResultOK() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_readResultSetHeaderPacket(t *testing.T) {
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
		want    int
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
			got, err := mc.readResultSetHeaderPacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readResultSetHeaderPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlConn.readResultSetHeaderPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_handleErrorPacket(t *testing.T) {
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
		data []byte
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
			if err := mc.handleErrorPacket(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.handleErrorPacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readStatus(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want statusFlag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readStatus(tt.args.b); got != tt.want {
				t.Errorf("readStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_handleOkPacket(t *testing.T) {
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
		data []byte
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
			if err := mc.handleOkPacket(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.handleOkPacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_readColumns(t *testing.T) {
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
		count int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []mysqlField
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
			got, err := mc.readColumns(tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.readColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_textRows_readRow(t *testing.T) {
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
			if err := rows.readRow(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("textRows.readRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_readUntilEOF(t *testing.T) {
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
			if err := mc.readUntilEOF(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.readUntilEOF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlStmt_readPrepareResultPacket(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint16
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
			got, err := stmt.readPrepareResultPacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.readPrepareResultPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlStmt.readPrepareResultPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlStmt_writeCommandLongData(t *testing.T) {
	type fields struct {
		mc         *mysqlConn
		id         uint32
		paramCount int
	}
	type args struct {
		paramID int
		arg     []byte
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
			if err := stmt.writeCommandLongData(tt.args.paramID, tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.writeCommandLongData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlStmt_writeExecutePacket(t *testing.T) {
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
			if err := stmt.writeExecutePacket(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("mysqlStmt.writeExecutePacket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_discardResults(t *testing.T) {
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
			if err := mc.discardResults(); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.discardResults() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_binaryRows_readRow(t *testing.T) {
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
			if err := rows.readRow(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("binaryRows.readRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
