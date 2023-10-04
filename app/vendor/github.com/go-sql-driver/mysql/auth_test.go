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
	"crypto/rsa"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestRegisterServerPubKey(t *testing.T) {
	type args struct {
		name   string
		pubKey *rsa.PublicKey
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterServerPubKey(tt.args.name, tt.args.pubKey)
		})
	}
}

func TestDeregisterServerPubKey(t *testing.T) {
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
			DeregisterServerPubKey(tt.args.name)
		})
	}
}

func Test_getServerPubKey(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantPubKey *rsa.PublicKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPubKey := getServerPubKey(tt.args.name); !reflect.DeepEqual(gotPubKey, tt.wantPubKey) {
				t.Errorf("getServerPubKey() = %v, want %v", gotPubKey, tt.wantPubKey)
			}
		})
	}
}

func Test_newMyRnd(t *testing.T) {
	type args struct {
		seed1 uint32
		seed2 uint32
	}
	tests := []struct {
		name string
		args args
		want *myRnd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMyRnd(tt.args.seed1, tt.args.seed2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMyRnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myRnd_NextByte(t *testing.T) {
	type fields struct {
		seed1 uint32
		seed2 uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &myRnd{
				seed1: tt.fields.seed1,
				seed2: tt.fields.seed2,
			}
			if got := r.NextByte(); got != tt.want {
				t.Errorf("myRnd.NextByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pwHash(t *testing.T) {
	type args struct {
		password []byte
	}
	tests := []struct {
		name       string
		args       args
		wantResult [2]uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := pwHash(tt.args.password); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("pwHash() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_scrambleOldPassword(t *testing.T) {
	type args struct {
		scramble []byte
		password string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrambleOldPassword(tt.args.scramble, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrambleOldPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scramblePassword(t *testing.T) {
	type args struct {
		scramble []byte
		password string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scramblePassword(tt.args.scramble, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scramblePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scrambleSHA256Password(t *testing.T) {
	type args struct {
		scramble []byte
		password string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scrambleSHA256Password(tt.args.scramble, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrambleSHA256Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encryptPassword(t *testing.T) {
	type args struct {
		password string
		seed     []byte
		pub      *rsa.PublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encryptPassword(tt.args.password, tt.args.seed, tt.args.pub)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encryptPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_sendEncryptedPassword(t *testing.T) {
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
		seed []byte
		pub  *rsa.PublicKey
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
			if err := mc.sendEncryptedPassword(tt.args.seed, tt.args.pub); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.sendEncryptedPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlConn_auth(t *testing.T) {
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
		plugin   string
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
			got, err := mc.auth(tt.args.authData, tt.args.plugin)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlConn.auth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlConn_handleAuthResult(t *testing.T) {
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
		oldAuthData []byte
		plugin      string
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
			if err := mc.handleAuthResult(tt.args.oldAuthData, tt.args.plugin); (err != nil) != tt.wantErr {
				t.Errorf("mysqlConn.handleAuthResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
