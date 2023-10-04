// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2016 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"reflect"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Clone(t *testing.T) {
	type fields struct {
		User                     string
		Passwd                   string
		Net                      string
		Addr                     string
		DBName                   string
		Params                   map[string]string
		Collation                string
		Loc                      *time.Location
		MaxAllowedPacket         int
		ServerPubKey             string
		pubKey                   *rsa.PublicKey
		TLSConfig                string
		TLS                      *tls.Config
		Timeout                  time.Duration
		ReadTimeout              time.Duration
		WriteTimeout             time.Duration
		AllowAllFiles            bool
		AllowCleartextPasswords  bool
		AllowFallbackToPlaintext bool
		AllowNativePasswords     bool
		AllowOldPasswords        bool
		CheckConnLiveness        bool
		ClientFoundRows          bool
		ColumnsWithAlias         bool
		InterpolateParams        bool
		MultiStatements          bool
		ParseTime                bool
		RejectReadOnly           bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				User:                     tt.fields.User,
				Passwd:                   tt.fields.Passwd,
				Net:                      tt.fields.Net,
				Addr:                     tt.fields.Addr,
				DBName:                   tt.fields.DBName,
				Params:                   tt.fields.Params,
				Collation:                tt.fields.Collation,
				Loc:                      tt.fields.Loc,
				MaxAllowedPacket:         tt.fields.MaxAllowedPacket,
				ServerPubKey:             tt.fields.ServerPubKey,
				pubKey:                   tt.fields.pubKey,
				TLSConfig:                tt.fields.TLSConfig,
				TLS:                      tt.fields.TLS,
				Timeout:                  tt.fields.Timeout,
				ReadTimeout:              tt.fields.ReadTimeout,
				WriteTimeout:             tt.fields.WriteTimeout,
				AllowAllFiles:            tt.fields.AllowAllFiles,
				AllowCleartextPasswords:  tt.fields.AllowCleartextPasswords,
				AllowFallbackToPlaintext: tt.fields.AllowFallbackToPlaintext,
				AllowNativePasswords:     tt.fields.AllowNativePasswords,
				AllowOldPasswords:        tt.fields.AllowOldPasswords,
				CheckConnLiveness:        tt.fields.CheckConnLiveness,
				ClientFoundRows:          tt.fields.ClientFoundRows,
				ColumnsWithAlias:         tt.fields.ColumnsWithAlias,
				InterpolateParams:        tt.fields.InterpolateParams,
				MultiStatements:          tt.fields.MultiStatements,
				ParseTime:                tt.fields.ParseTime,
				RejectReadOnly:           tt.fields.RejectReadOnly,
			}
			if got := cfg.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_normalize(t *testing.T) {
	type fields struct {
		User                     string
		Passwd                   string
		Net                      string
		Addr                     string
		DBName                   string
		Params                   map[string]string
		Collation                string
		Loc                      *time.Location
		MaxAllowedPacket         int
		ServerPubKey             string
		pubKey                   *rsa.PublicKey
		TLSConfig                string
		TLS                      *tls.Config
		Timeout                  time.Duration
		ReadTimeout              time.Duration
		WriteTimeout             time.Duration
		AllowAllFiles            bool
		AllowCleartextPasswords  bool
		AllowFallbackToPlaintext bool
		AllowNativePasswords     bool
		AllowOldPasswords        bool
		CheckConnLiveness        bool
		ClientFoundRows          bool
		ColumnsWithAlias         bool
		InterpolateParams        bool
		MultiStatements          bool
		ParseTime                bool
		RejectReadOnly           bool
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
			cfg := &Config{
				User:                     tt.fields.User,
				Passwd:                   tt.fields.Passwd,
				Net:                      tt.fields.Net,
				Addr:                     tt.fields.Addr,
				DBName:                   tt.fields.DBName,
				Params:                   tt.fields.Params,
				Collation:                tt.fields.Collation,
				Loc:                      tt.fields.Loc,
				MaxAllowedPacket:         tt.fields.MaxAllowedPacket,
				ServerPubKey:             tt.fields.ServerPubKey,
				pubKey:                   tt.fields.pubKey,
				TLSConfig:                tt.fields.TLSConfig,
				TLS:                      tt.fields.TLS,
				Timeout:                  tt.fields.Timeout,
				ReadTimeout:              tt.fields.ReadTimeout,
				WriteTimeout:             tt.fields.WriteTimeout,
				AllowAllFiles:            tt.fields.AllowAllFiles,
				AllowCleartextPasswords:  tt.fields.AllowCleartextPasswords,
				AllowFallbackToPlaintext: tt.fields.AllowFallbackToPlaintext,
				AllowNativePasswords:     tt.fields.AllowNativePasswords,
				AllowOldPasswords:        tt.fields.AllowOldPasswords,
				CheckConnLiveness:        tt.fields.CheckConnLiveness,
				ClientFoundRows:          tt.fields.ClientFoundRows,
				ColumnsWithAlias:         tt.fields.ColumnsWithAlias,
				InterpolateParams:        tt.fields.InterpolateParams,
				MultiStatements:          tt.fields.MultiStatements,
				ParseTime:                tt.fields.ParseTime,
				RejectReadOnly:           tt.fields.RejectReadOnly,
			}
			if err := cfg.normalize(); (err != nil) != tt.wantErr {
				t.Errorf("Config.normalize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_writeDSNParam(t *testing.T) {
	type args struct {
		buf      *bytes.Buffer
		hasParam *bool
		name     string
		value    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeDSNParam(tt.args.buf, tt.args.hasParam, tt.args.name, tt.args.value)
		})
	}
}

func TestConfig_FormatDSN(t *testing.T) {
	type fields struct {
		User                     string
		Passwd                   string
		Net                      string
		Addr                     string
		DBName                   string
		Params                   map[string]string
		Collation                string
		Loc                      *time.Location
		MaxAllowedPacket         int
		ServerPubKey             string
		pubKey                   *rsa.PublicKey
		TLSConfig                string
		TLS                      *tls.Config
		Timeout                  time.Duration
		ReadTimeout              time.Duration
		WriteTimeout             time.Duration
		AllowAllFiles            bool
		AllowCleartextPasswords  bool
		AllowFallbackToPlaintext bool
		AllowNativePasswords     bool
		AllowOldPasswords        bool
		CheckConnLiveness        bool
		ClientFoundRows          bool
		ColumnsWithAlias         bool
		InterpolateParams        bool
		MultiStatements          bool
		ParseTime                bool
		RejectReadOnly           bool
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
			cfg := &Config{
				User:                     tt.fields.User,
				Passwd:                   tt.fields.Passwd,
				Net:                      tt.fields.Net,
				Addr:                     tt.fields.Addr,
				DBName:                   tt.fields.DBName,
				Params:                   tt.fields.Params,
				Collation:                tt.fields.Collation,
				Loc:                      tt.fields.Loc,
				MaxAllowedPacket:         tt.fields.MaxAllowedPacket,
				ServerPubKey:             tt.fields.ServerPubKey,
				pubKey:                   tt.fields.pubKey,
				TLSConfig:                tt.fields.TLSConfig,
				TLS:                      tt.fields.TLS,
				Timeout:                  tt.fields.Timeout,
				ReadTimeout:              tt.fields.ReadTimeout,
				WriteTimeout:             tt.fields.WriteTimeout,
				AllowAllFiles:            tt.fields.AllowAllFiles,
				AllowCleartextPasswords:  tt.fields.AllowCleartextPasswords,
				AllowFallbackToPlaintext: tt.fields.AllowFallbackToPlaintext,
				AllowNativePasswords:     tt.fields.AllowNativePasswords,
				AllowOldPasswords:        tt.fields.AllowOldPasswords,
				CheckConnLiveness:        tt.fields.CheckConnLiveness,
				ClientFoundRows:          tt.fields.ClientFoundRows,
				ColumnsWithAlias:         tt.fields.ColumnsWithAlias,
				InterpolateParams:        tt.fields.InterpolateParams,
				MultiStatements:          tt.fields.MultiStatements,
				ParseTime:                tt.fields.ParseTime,
				RejectReadOnly:           tt.fields.RejectReadOnly,
			}
			if got := cfg.FormatDSN(); got != tt.want {
				t.Errorf("Config.FormatDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDSN(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *Config
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := ParseDSN(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("ParseDSN() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func Test_parseDSNParams(t *testing.T) {
	type args struct {
		cfg    *Config
		params string
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
			if err := parseDSNParams(tt.args.cfg, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("parseDSNParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ensureHavePort(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ensureHavePort(tt.args.addr); got != tt.want {
				t.Errorf("ensureHavePort() = %v, want %v", got, tt.want)
			}
		})
	}
}
