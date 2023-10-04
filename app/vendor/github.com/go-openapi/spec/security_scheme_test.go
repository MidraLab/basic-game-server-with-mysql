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

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		name string
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BasicAuth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIKeyAuth(t *testing.T) {
	type args struct {
		fieldName   string
		valueSource string
	}
	tests := []struct {
		name string
		args args
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := APIKeyAuth(tt.args.fieldName, tt.args.valueSource); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIKeyAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuth2Implicit(t *testing.T) {
	type args struct {
		authorizationURL string
	}
	tests := []struct {
		name string
		args args
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OAuth2Implicit(tt.args.authorizationURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuth2Implicit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuth2Password(t *testing.T) {
	type args struct {
		tokenURL string
	}
	tests := []struct {
		name string
		args args
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OAuth2Password(tt.args.tokenURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuth2Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuth2Application(t *testing.T) {
	type args struct {
		tokenURL string
	}
	tests := []struct {
		name string
		args args
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OAuth2Application(tt.args.tokenURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuth2Application() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOAuth2AccessToken(t *testing.T) {
	type args struct {
		authorizationURL string
		tokenURL         string
	}
	tests := []struct {
		name string
		args args
		want *SecurityScheme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OAuth2AccessToken(tt.args.authorizationURL, tt.args.tokenURL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OAuth2AccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecuritySchemeProps_AddScope(t *testing.T) {
	type fields struct {
		Description      string
		Type             string
		Name             string
		In               string
		Flow             string
		AuthorizationURL string
		TokenURL         string
		Scopes           map[string]string
	}
	type args struct {
		scope       string
		description string
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
			s := &SecuritySchemeProps{
				Description:      tt.fields.Description,
				Type:             tt.fields.Type,
				Name:             tt.fields.Name,
				In:               tt.fields.In,
				Flow:             tt.fields.Flow,
				AuthorizationURL: tt.fields.AuthorizationURL,
				TokenURL:         tt.fields.TokenURL,
				Scopes:           tt.fields.Scopes,
			}
			s.AddScope(tt.args.scope, tt.args.description)
		})
	}
}

func TestSecurityScheme_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible    VendorExtensible
		SecuritySchemeProps SecuritySchemeProps
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SecurityScheme{
				VendorExtensible:    tt.fields.VendorExtensible,
				SecuritySchemeProps: tt.fields.SecuritySchemeProps,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecurityScheme.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SecurityScheme.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityScheme_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible    VendorExtensible
		SecuritySchemeProps SecuritySchemeProps
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
			s := SecurityScheme{
				VendorExtensible:    tt.fields.VendorExtensible,
				SecuritySchemeProps: tt.fields.SecuritySchemeProps,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SecurityScheme.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SecurityScheme.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityScheme_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible    VendorExtensible
		SecuritySchemeProps SecuritySchemeProps
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
			s := &SecurityScheme{
				VendorExtensible:    tt.fields.VendorExtensible,
				SecuritySchemeProps: tt.fields.SecuritySchemeProps,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SecurityScheme.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
