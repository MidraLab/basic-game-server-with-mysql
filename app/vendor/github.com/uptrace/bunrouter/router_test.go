package bunrouter

import (
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_ServeHTTP(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
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
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			r.ServeHTTP(tt.args.w, tt.args.req)
		})
	}
}

func TestRouter_ServeHTTPError(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
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
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			if err := r.ServeHTTPError(tt.args.w, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Router.ServeHTTPError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRouter_lookup(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HandlerFunc
		want1  Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			got, got1 := r.lookup(tt.args.w, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.lookup() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Router.lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRouter_redir(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			if got := r.redir(tt.args.method, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.redir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_Compat(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	tests := []struct {
		name   string
		fields fields
		want   *CompatRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			if got := r.Compat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.Compat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouter_Verbose(t *testing.T) {
	type fields struct {
		config config
		Group  Group
		mu     sync.Mutex
		tree   node
	}
	tests := []struct {
		name   string
		fields fields
		want   *VerboseRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Router{
				config: tt.fields.config,
				Group:  tt.fields.Group,
				mu:     tt.fields.mu,
				tree:   tt.fields.tree,
			}
			if got := r.Verbose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router.Verbose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redirectHandler(t *testing.T) {
	type args struct {
		newPath string
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := redirectHandler(tt.args.newPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redirectHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_methodNotAllowedHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := methodNotAllowedHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("methodNotAllowedHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_notFoundHandler(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := notFoundHandler(tt.args.w, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("notFoundHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
