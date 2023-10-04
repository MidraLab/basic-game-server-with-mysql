// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2013 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func Test_newBuffer(t *testing.T) {
	type args struct {
		nc net.Conn
	}
	tests := []struct {
		name string
		args args
		want buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBuffer(tt.args.nc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffer_flip(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			b.flip()
		})
	}
}

func Test_buffer_fill(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	type args struct {
		need int
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			if err := b.fill(tt.args.need); (err != nil) != tt.wantErr {
				t.Errorf("buffer.fill() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buffer_readNext(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	type args struct {
		need int
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			got, err := b.readNext(tt.args.need)
			if (err != nil) != tt.wantErr {
				t.Errorf("buffer.readNext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buffer.readNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffer_takeBuffer(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	type args struct {
		length int
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			got, err := b.takeBuffer(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("buffer.takeBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buffer.takeBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffer_takeSmallBuffer(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	type args struct {
		length int
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			got, err := b.takeSmallBuffer(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("buffer.takeSmallBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buffer.takeSmallBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffer_takeCompleteBuffer(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			got, err := b.takeCompleteBuffer()
			if (err != nil) != tt.wantErr {
				t.Errorf("buffer.takeCompleteBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buffer.takeCompleteBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buffer_store(t *testing.T) {
	type fields struct {
		buf     []byte
		nc      net.Conn
		idx     int
		length  int
		timeout time.Duration
		dbuf    [2][]byte
		flipcnt uint
	}
	type args struct {
		buf []byte
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
			b := &buffer{
				buf:     tt.fields.buf,
				nc:      tt.fields.nc,
				idx:     tt.fields.idx,
				length:  tt.fields.length,
				timeout: tt.fields.timeout,
				dbuf:    tt.fields.dbuf,
				flipcnt: tt.fields.flipcnt,
			}
			if err := b.store(tt.args.buf); (err != nil) != tt.wantErr {
				t.Errorf("buffer.store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
