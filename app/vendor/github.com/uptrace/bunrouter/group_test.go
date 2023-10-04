package bunrouter

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGroup_NewGroup(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path string
		opts []GroupOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.NewGroup(tt.args.path, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_cloneStack(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   []MiddlewareFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.cloneStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.cloneStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Use(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		middlewares []MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.Use(tt.args.middlewares...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.Use() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_WithMiddleware(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		middleware MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Group
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.WithMiddleware(tt.args.middleware); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.WithMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_WithGroup(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path string
		fn   func(g *Group)
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.WithGroup(tt.args.path, tt.args.fn)
		})
	}
}

func TestGroup_Handle(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		meth    string
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.Handle(tt.args.meth, tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_GET(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.GET(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_POST(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.POST(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_PUT(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.PUT(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_DELETE(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.DELETE(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_PATCH(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.PATCH(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_HEAD(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.HEAD(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_OPTIONS(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		path    string
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			g.OPTIONS(tt.args.path, tt.args.handler)
		})
	}
}

func TestGroup_wrap(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	type args struct {
		handler HandlerFunc
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
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.wrap(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Compat(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *CompatGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.Compat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.Compat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Verbose(t *testing.T) {
	type fields struct {
		router *Router
		path   string
		stack  []MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *VerboseGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				router: tt.fields.router,
				path:   tt.fields.path,
				stack:  tt.fields.stack,
			}
			if got := g.Verbose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.Verbose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompatGroup_NewGroup(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path string
		opts []GroupOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CompatGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CompatGroup{
				group: tt.fields.group,
			}
			if got := g.NewGroup(tt.args.path, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompatGroup.NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompatGroup_WithMiddleware(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		middleware MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CompatGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CompatGroup{
				group: tt.fields.group,
			}
			if got := g.WithMiddleware(tt.args.middleware); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompatGroup.WithMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompatGroup_WithGroup(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path string
		fn   func(g *CompatGroup)
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.WithGroup(tt.args.path, tt.args.fn)
		})
	}
}

func TestCompatGroup_Handle(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		method  string
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.Handle(tt.args.method, tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_GET(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.GET(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_POST(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.POST(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_PUT(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.PUT(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_DELETE(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.DELETE(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_PATCH(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.PATCH(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_HEAD(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.HEAD(tt.args.path, tt.args.handler)
		})
	}
}

func TestCompatGroup_OPTIONS(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler http.HandlerFunc
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
			g := CompatGroup{
				group: tt.fields.group,
			}
			g.OPTIONS(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_NewGroup(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path string
		opts []GroupOption
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VerboseGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := VerboseGroup{
				group: tt.fields.group,
			}
			if got := g.NewGroup(tt.args.path, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerboseGroup.NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerboseGroup_WithMiddleware(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		middleware MiddlewareFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VerboseGroup
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := VerboseGroup{
				group: tt.fields.group,
			}
			if got := g.WithMiddleware(tt.args.middleware); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerboseGroup.WithMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerboseGroup_WithGroup(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path string
		fn   func(g *VerboseGroup)
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.WithGroup(tt.args.path, tt.args.fn)
		})
	}
}

func TestVerboseGroup_Handle(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		method  string
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.Handle(tt.args.method, tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_GET(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.GET(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_POST(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.POST(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_PUT(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.PUT(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_DELETE(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.DELETE(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_PATCH(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.PATCH(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_HEAD(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.HEAD(tt.args.path, tt.args.handler)
		})
	}
}

func TestVerboseGroup_OPTIONS(t *testing.T) {
	type fields struct {
		group *Group
	}
	type args struct {
		path    string
		handler VerboseHandlerFunc
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
			g := VerboseGroup{
				group: tt.fields.group,
			}
			g.OPTIONS(tt.args.path, tt.args.handler)
		})
	}
}

func Test_joinPath(t *testing.T) {
	type args struct {
		base string
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinPath(tt.args.base, tt.args.path); got != tt.want {
				t.Errorf("joinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkPath(tt.args.path)
		})
	}
}
