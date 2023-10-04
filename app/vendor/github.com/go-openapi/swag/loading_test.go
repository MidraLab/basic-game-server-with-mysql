// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swag

import (
	"reflect"
	"testing"
	"time"
)

func TestLoadFromFileOrHTTP(t *testing.T) {
	type args struct {
		path string
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
			got, err := LoadFromFileOrHTTP(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromFileOrHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFromFileOrHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadFromFileOrHTTPWithTimeout(t *testing.T) {
	type args struct {
		path    string
		timeout time.Duration
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
			got, err := LoadFromFileOrHTTPWithTimeout(tt.args.path, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromFileOrHTTPWithTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFromFileOrHTTPWithTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadStrategy(t *testing.T) {
	type args struct {
		path   string
		local  func(string) ([]byte, error)
		remote func(string) ([]byte, error)
	}
	tests := []struct {
		name string
		args args
		want func(string) ([]byte, error)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadStrategy(tt.args.path, tt.args.local, tt.args.remote); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadHTTPBytes(t *testing.T) {
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want func(path string) ([]byte, error)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadHTTPBytes(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadHTTPBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
