package bunrouter

import (
	"reflect"
	"testing"
)

func Test_option_apply(t *testing.T) {
	type args struct {
		cfg *config
	}
	tests := []struct {
		name string
		fn   option
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn.apply(tt.args.cfg)
		})
	}
}

func TestWithNotFoundHandler(t *testing.T) {
	type args struct {
		handler HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithNotFoundHandler(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithNotFoundHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMethodNotAllowedHandler(t *testing.T) {
	type args struct {
		handler HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMethodNotAllowedHandler(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMethodNotAllowedHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupOption_apply(t *testing.T) {
	type args struct {
		cfg *config
	}
	tests := []struct {
		name string
		fn   groupOption
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn.apply(tt.args.cfg)
		})
	}
}

func Test_groupOption_groupOption(t *testing.T) {
	tests := []struct {
		name string
		fn   groupOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn.groupOption()
		})
	}
}

func TestWithGroup(t *testing.T) {
	type args struct {
		fn func(g *Group)
	}
	tests := []struct {
		name string
		args args
		want GroupOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithGroup(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMiddleware(t *testing.T) {
	type args struct {
		fns []MiddlewareFunc
	}
	tests := []struct {
		name string
		args args
		want GroupOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMiddleware(tt.args.fns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUse(t *testing.T) {
	type args struct {
		fns []MiddlewareFunc
	}
	tests := []struct {
		name string
		args args
		want GroupOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Use(tt.args.fns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Use() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithHandler(t *testing.T) {
	type args struct {
		fn HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want GroupOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithHandler(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
