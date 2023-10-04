// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "testing"

func TestTime_UnixTime(t *testing.T) {
	tests := []struct {
		name     string
		tr       Time
		wantSec  int64
		wantNsec int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSec, gotNsec := tt.tr.UnixTime()
			if gotSec != tt.wantSec {
				t.Errorf("Time.UnixTime() gotSec = %v, want %v", gotSec, tt.wantSec)
			}
			if gotNsec != tt.wantNsec {
				t.Errorf("Time.UnixTime() gotNsec = %v, want %v", gotNsec, tt.wantNsec)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	tests := []struct {
		name    string
		want    Time
		want1   uint16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getTime(t *testing.T) {
	tests := []struct {
		name    string
		want    Time
		want1   uint16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClockSequence(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClockSequence(); got != tt.want {
				t.Errorf("ClockSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clockSequence(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clockSequence(); got != tt.want {
				t.Errorf("clockSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetClockSequence(t *testing.T) {
	type args struct {
		seq int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClockSequence(tt.args.seq)
		})
	}
}

func Test_setClockSequence(t *testing.T) {
	type args struct {
		seq int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setClockSequence(tt.args.seq)
		})
	}
}

func TestUUID_Time(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.Time(); got != tt.want {
				t.Errorf("UUID.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_ClockSequence(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.ClockSequence(); got != tt.want {
				t.Errorf("UUID.ClockSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
