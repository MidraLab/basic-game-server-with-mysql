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

func TestContactInfo_UnmarshalJSON(t *testing.T) {
	type fields struct {
		ContactInfoProps ContactInfoProps
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
			c := &ContactInfo{
				ContactInfoProps: tt.fields.ContactInfoProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if err := c.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ContactInfo.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContactInfo_MarshalJSON(t *testing.T) {
	type fields struct {
		ContactInfoProps ContactInfoProps
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
			c := ContactInfo{
				ContactInfoProps: tt.fields.ContactInfoProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			got, err := c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContactInfo.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactInfo.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
