// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"reflect"
	"testing"
)

func TestNewUUID(t *testing.T) {
	tests := []struct {
		name    string
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
