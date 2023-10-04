package internal

import "testing"

func TestFlag_Has(t *testing.T) {
	type args struct {
		other Flag
	}
	tests := []struct {
		name string
		flag Flag
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.flag.Has(tt.args.other); got != tt.want {
				t.Errorf("Flag.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlag_Set(t *testing.T) {
	type args struct {
		other Flag
	}
	tests := []struct {
		name string
		flag Flag
		args args
		want Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.flag.Set(tt.args.other); got != tt.want {
				t.Errorf("Flag.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlag_Remove(t *testing.T) {
	type args struct {
		other Flag
	}
	tests := []struct {
		name string
		flag Flag
		args args
		want Flag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.flag.Remove(tt.args.other); got != tt.want {
				t.Errorf("Flag.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
