//
// Copyright (c) 2011-2019 Canonical Ltd
// Copyright (c) 2006-2010 Kirill Simonov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package yaml

import "testing"

func Test_is_alpha(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_alpha(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_alpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_digit(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_digit(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_digit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_as_digit(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := as_digit(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("as_digit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_hex(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_hex(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_as_hex(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := as_hex(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("as_hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_ascii(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_ascii(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_ascii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_printable(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_printable(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_printable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_z(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_z(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_z() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_bom(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_bom(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_bom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_space(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_space(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_space() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_tab(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_tab(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_tab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_blank(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_blank(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_blank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_break(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_break(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_break() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_crlf(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_crlf(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_crlf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_breakz(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_breakz(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_breakz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_spacez(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_spacez(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_spacez() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_blankz(t *testing.T) {
	type args struct {
		b []byte
		i int
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
			if got := is_blankz(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("is_blankz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_width(t *testing.T) {
	type args struct {
		b byte
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
			if got := width(tt.args.b); got != tt.want {
				t.Errorf("width() = %v, want %v", got, tt.want)
			}
		})
	}
}
