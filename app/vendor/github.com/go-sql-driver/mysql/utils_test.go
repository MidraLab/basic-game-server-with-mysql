// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"crypto/tls"
	"database/sql/driver"
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

func TestRegisterTLSConfig(t *testing.T) {
	type args struct {
		key    string
		config *tls.Config
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
			if err := RegisterTLSConfig(tt.args.key, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("RegisterTLSConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeregisterTLSConfig(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeregisterTLSConfig(tt.args.key)
		})
	}
}

func Test_getTLSConfigClone(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig *tls.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConfig := getTLSConfigClone(tt.args.key); !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("getTLSConfigClone() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

func Test_readBool(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
		wantValid bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotValid := readBool(tt.args.input)
			if gotValue != tt.wantValue {
				t.Errorf("readBool() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotValid != tt.wantValid {
				t.Errorf("readBool() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_parseDateTime(t *testing.T) {
	type args struct {
		b   []byte
		loc *time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDateTime(tt.args.b, tt.args.loc)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseByteYear(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseByteYear(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseByteYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseByteYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseByte2Digits(t *testing.T) {
	type args struct {
		b1 byte
		b2 byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseByte2Digits(tt.args.b1, tt.args.b2)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseByte2Digits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseByte2Digits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseByteNanoSec(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseByteNanoSec(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseByteNanoSec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseByteNanoSec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bToi(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bToi(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("bToi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bToi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBinaryDateTime(t *testing.T) {
	type args struct {
		num  uint64
		data []byte
		loc  *time.Location
	}
	tests := []struct {
		name    string
		args    args
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBinaryDateTime(tt.args.num, tt.args.data, tt.args.loc)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseBinaryDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBinaryDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendDateTime(t *testing.T) {
	type args struct {
		buf []byte
		t   time.Time
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
			got, err := appendDateTime(tt.args.buf, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("appendDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendMicrosecs(t *testing.T) {
	type args struct {
		dst      []byte
		src      []byte
		decimals int
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
			if got := appendMicrosecs(tt.args.dst, tt.args.src, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendMicrosecs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatBinaryDateTime(t *testing.T) {
	type args struct {
		src    []byte
		length uint8
	}
	tests := []struct {
		name    string
		args    args
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatBinaryDateTime(tt.args.src, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatBinaryDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatBinaryDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatBinaryTime(t *testing.T) {
	type args struct {
		src    []byte
		length uint8
	}
	tests := []struct {
		name    string
		args    args
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatBinaryTime(tt.args.src, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatBinaryTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatBinaryTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint64ToBytes(t *testing.T) {
	type args struct {
		n uint64
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
			if got := uint64ToBytes(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uint64ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uint64ToString(t *testing.T) {
	type args struct {
		n uint64
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
			if got := uint64ToString(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uint64ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringToInt(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToInt(tt.args.b); got != tt.want {
				t.Errorf("stringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readLengthEncodedString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   bool
		want2   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := readLengthEncodedString(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLengthEncodedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLengthEncodedString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readLengthEncodedString() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("readLengthEncodedString() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_skipLengthEncodedString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := skipLengthEncodedString(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("skipLengthEncodedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("skipLengthEncodedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readLengthEncodedInteger(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
		want2 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := readLengthEncodedInteger(tt.args.b)
			if got != tt.want {
				t.Errorf("readLengthEncodedInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readLengthEncodedInteger() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("readLengthEncodedInteger() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_appendLengthEncodedInteger(t *testing.T) {
	type args struct {
		b []byte
		n uint64
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
			if got := appendLengthEncodedInteger(tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendLengthEncodedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reserveBuffer(t *testing.T) {
	type args struct {
		buf        []byte
		appendSize int
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
			if got := reserveBuffer(tt.args.buf, tt.args.appendSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reserveBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escapeBytesBackslash(t *testing.T) {
	type args struct {
		buf []byte
		v   []byte
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
			if got := escapeBytesBackslash(tt.args.buf, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escapeBytesBackslash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escapeStringBackslash(t *testing.T) {
	type args struct {
		buf []byte
		v   string
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
			if got := escapeStringBackslash(tt.args.buf, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escapeStringBackslash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escapeBytesQuotes(t *testing.T) {
	type args struct {
		buf []byte
		v   []byte
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
			if got := escapeBytesQuotes(tt.args.buf, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escapeBytesQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escapeStringQuotes(t *testing.T) {
	type args struct {
		buf []byte
		v   string
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
			if got := escapeStringQuotes(tt.args.buf, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("escapeStringQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noCopy_Lock(t *testing.T) {
	tests := []struct {
		name string
		n    *noCopy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &noCopy{}
			n.Lock()
		})
	}
}

func Test_noCopy_Unlock(t *testing.T) {
	tests := []struct {
		name string
		n    *noCopy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &noCopy{}
			n.Unlock()
		})
	}
}

func Test_atomicError_Set(t *testing.T) {
	type fields struct {
		noCopy noCopy
		value  atomic.Value
	}
	type args struct {
		value error
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
			ae := &atomicError{
				_:     tt.fields.noCopy,
				value: tt.fields.value,
			}
			ae.Set(tt.args.value)
		})
	}
}

func Test_atomicError_Value(t *testing.T) {
	type fields struct {
		noCopy noCopy
		value  atomic.Value
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
			ae := &atomicError{
				_:     tt.fields.noCopy,
				value: tt.fields.value,
			}
			if err := ae.Value(); (err != nil) != tt.wantErr {
				t.Errorf("atomicError.Value() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_namedValueToValue(t *testing.T) {
	type args struct {
		named []driver.NamedValue
	}
	tests := []struct {
		name    string
		args    args
		want    []driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := namedValueToValue(tt.args.named)
			if (err != nil) != tt.wantErr {
				t.Errorf("namedValueToValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("namedValueToValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapIsolationLevel(t *testing.T) {
	type args struct {
		level driver.IsolationLevel
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapIsolationLevel(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapIsolationLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mapIsolationLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
