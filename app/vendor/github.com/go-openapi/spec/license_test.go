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

package spec

import (
	"reflect"
	"testing"
)

func TestLicense_UnmarshalJSON(t *testing.T) {
	type fields struct {
		LicenseProps     LicenseProps
		VendorExtensible VendorExtensible
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
			l := &License{
				LicenseProps:     tt.fields.LicenseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if err := l.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("License.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLicense_MarshalJSON(t *testing.T) {
	type fields struct {
		LicenseProps     LicenseProps
		VendorExtensible VendorExtensible
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
			l := License{
				LicenseProps:     tt.fields.LicenseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			got, err := l.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("License.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("License.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
