package msgpcode

import "testing"

func TestIsFixedNum(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedNum(tt.args.c); got != tt.want {
				t.Errorf("IsFixedNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFixedMap(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedMap(tt.args.c); got != tt.want {
				t.Errorf("IsFixedMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFixedArray(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedArray(tt.args.c); got != tt.want {
				t.Errorf("IsFixedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFixedString(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedString(tt.args.c); got != tt.want {
				t.Errorf("IsFixedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsString(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsString(tt.args.c); got != tt.want {
				t.Errorf("IsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBin(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBin(tt.args.c); got != tt.want {
				t.Errorf("IsBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFixedExt(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedExt(tt.args.c); got != tt.want {
				t.Errorf("IsFixedExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsExt(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExt(tt.args.c); got != tt.want {
				t.Errorf("IsExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
