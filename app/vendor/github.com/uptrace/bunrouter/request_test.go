package bunrouter

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestParamsFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamsFromContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamsFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contextWithParams(t *testing.T) {
	type args struct {
		ctx    context.Context
		params Params
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contextWithParams(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contextWithParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPHandler(t *testing.T) {
	type args struct {
		handler http.Handler
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
			if got := HTTPHandler(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPHandlerFunc(t *testing.T) {
	type args struct {
		handler http.HandlerFunc
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
			if got := HTTPHandlerFunc(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPHandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerFunc_ServeHTTP(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		h    HandlerFunc
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ServeHTTP(tt.args.w, tt.args.req)
		})
	}
}

func TestNewRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRequestParams(t *testing.T) {
	type args struct {
		req    *http.Request
		params Params
	}
	tests := []struct {
		name string
		args args
		want Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRequestParams(tt.args.req, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRequestParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_WithContext(t *testing.T) {
	type fields struct {
		Request *http.Request
		params  Params
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Request: tt.fields.Request,
				params:  tt.fields.params,
			}
			if got := req.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Params(t *testing.T) {
	type fields struct {
		Request *http.Request
		params  Params
	}
	tests := []struct {
		name   string
		fields fields
		want   Params
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Request: tt.fields.Request,
				params:  tt.fields.params,
			}
			if got := req.Params(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Params() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Param(t *testing.T) {
	type fields struct {
		Request *http.Request
		params  Params
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Request: tt.fields.Request,
				params:  tt.fields.params,
			}
			if got := req.Param(tt.args.key); got != tt.want {
				t.Errorf("Request.Param() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Route(t *testing.T) {
	type fields struct {
		Request *http.Request
		params  Params
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := Request{
				Request: tt.fields.Request,
				params:  tt.fields.params,
			}
			if got := req.Route(); got != tt.want {
				t.Errorf("Request.Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_IsZero(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			if got := ps.IsZero(); got != tt.want {
				t.Errorf("Params.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Route(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			if got := ps.Route(); got != tt.want {
				t.Errorf("Params.Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Get(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, got1 := ps.Get(tt.args.name)
			if got != tt.want {
				t.Errorf("Params.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Params.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParams_findParam(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		paramIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, got1 := ps.findParam(tt.args.paramIndex)
			if got != tt.want {
				t.Errorf("Params.findParam() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Params.findParam() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParams_ByName(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			if got := ps.ByName(tt.args.name); got != tt.want {
				t.Errorf("Params.ByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Uint32(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, err := ps.Uint32(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Params.Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Params.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Uint64(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, err := ps.Uint64(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Params.Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Params.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Int32(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, err := ps.Int32(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Params.Int32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Params.Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Int64(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			got, err := ps.Int64(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Params.Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Params.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Map(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			if got := ps.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Params.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_Slice(t *testing.T) {
	type fields struct {
		path        string
		node        *node
		handler     *routeHandler
		wildcardLen uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   []Param
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := Params{
				path:        tt.fields.path,
				node:        tt.fields.node,
				handler:     tt.fields.handler,
				wildcardLen: tt.fields.wildcardLen,
			}
			if got := ps.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Params.Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON(t *testing.T) {
	type args struct {
		w     http.ResponseWriter
		value interface{}
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
			if err := JSON(tt.args.w, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("JSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
