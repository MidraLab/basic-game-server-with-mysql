package bunrouter

import (
	"reflect"
	"testing"
)

func Test_node_addRoute(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	type args struct {
		route string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			got, got1 := n.addRoute(tt.args.route)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.addRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("node.addRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_node_addPart(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			if got := n.addPart(tt.args.part); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.addPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_findRoute(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	type args struct {
		meth string
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  *routeHandler
		want2  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			got, got1, got2 := n.findRoute(tt.args.meth, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.findRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("node.findRoute() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("node.findRoute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_node__findRoute(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	type args struct {
		meth string
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  *routeHandler
		want2  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			got, got1, got2 := n._findRoute(tt.args.meth, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node._findRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("node._findRoute() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("node._findRoute() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_node_indexNodes(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			n.indexNodes()
		})
	}
}

func Test_node__indexNodes(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			n._indexNodes()
		})
	}
}

func Test_node_setHandler(t *testing.T) {
	type fields struct {
		route      string
		part       string
		handlerMap *handlerMap
		parent     *node
		colon      *node
		isWC       bool
		nodes      []*node
		index      struct {
			table   []uint8
			minChar byte
			maxChar byte
		}
	}
	type args struct {
		verb    string
		handler *routeHandler
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
			n := &node{
				route:      tt.fields.route,
				part:       tt.fields.part,
				handlerMap: tt.fields.handlerMap,
				parent:     tt.fields.parent,
				colon:      tt.fields.colon,
				isWC:       tt.fields.isWC,
				nodes:      tt.fields.nodes,
				index:      tt.fields.index,
			}
			n.setHandler(tt.args.verb, tt.args.handler)
		})
	}
}

func Test_routeParser_valid(t *testing.T) {
	type fields struct {
		segments []string
		i        int
		acc      []string
		parts    []string
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
			p := &routeParser{
				segments: tt.fields.segments,
				i:        tt.fields.i,
				acc:      tt.fields.acc,
				parts:    tt.fields.parts,
			}
			if got := p.valid(); got != tt.want {
				t.Errorf("routeParser.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_routeParser_next(t *testing.T) {
	type fields struct {
		segments []string
		i        int
		acc      []string
		parts    []string
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
			p := &routeParser{
				segments: tt.fields.segments,
				i:        tt.fields.i,
				acc:      tt.fields.acc,
				parts:    tt.fields.parts,
			}
			if got := p.next(); got != tt.want {
				t.Errorf("routeParser.next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_routeParser_accumulate(t *testing.T) {
	type fields struct {
		segments []string
		i        int
		acc      []string
		parts    []string
	}
	type args struct {
		s string
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
			p := &routeParser{
				segments: tt.fields.segments,
				i:        tt.fields.i,
				acc:      tt.fields.acc,
				parts:    tt.fields.parts,
			}
			p.accumulate(tt.args.s)
		})
	}
}

func Test_routeParser_finalizePart(t *testing.T) {
	type fields struct {
		segments []string
		i        int
		acc      []string
		parts    []string
	}
	type args struct {
		withSlash bool
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
			p := &routeParser{
				segments: tt.fields.segments,
				i:        tt.fields.i,
				acc:      tt.fields.acc,
				parts:    tt.fields.parts,
			}
			p.finalizePart(tt.args.withSlash)
		})
	}
}

func Test_join(t *testing.T) {
	type args struct {
		ss        []string
		withSlash bool
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
			if got := join(tt.args.ss, tt.args.withSlash); got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitRoute(t *testing.T) {
	type args struct {
		route string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitRoute(tt.args.route)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_paramMap(t *testing.T) {
	type args struct {
		route  string
		params []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paramMap(tt.args.route, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paramMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHandlerMap(t *testing.T) {
	tests := []struct {
		name string
		want *handlerMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHandlerMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHandlerMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handlerMap_Get(t *testing.T) {
	type fields struct {
		get        *routeHandler
		post       *routeHandler
		put        *routeHandler
		delete     *routeHandler
		head       *routeHandler
		options    *routeHandler
		patch      *routeHandler
		notAllowed *routeHandler
	}
	type args struct {
		meth string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *routeHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerMap{
				get:        tt.fields.get,
				post:       tt.fields.post,
				put:        tt.fields.put,
				delete:     tt.fields.delete,
				head:       tt.fields.head,
				options:    tt.fields.options,
				patch:      tt.fields.patch,
				notAllowed: tt.fields.notAllowed,
			}
			if got := h.Get(tt.args.meth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handlerMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handlerMap_Set(t *testing.T) {
	type fields struct {
		get        *routeHandler
		post       *routeHandler
		put        *routeHandler
		delete     *routeHandler
		head       *routeHandler
		options    *routeHandler
		patch      *routeHandler
		notAllowed *routeHandler
	}
	type args struct {
		meth    string
		handler *routeHandler
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
			h := &handlerMap{
				get:        tt.fields.get,
				post:       tt.fields.post,
				put:        tt.fields.put,
				delete:     tt.fields.delete,
				head:       tt.fields.head,
				options:    tt.fields.options,
				patch:      tt.fields.patch,
				notAllowed: tt.fields.notAllowed,
			}
			h.Set(tt.args.meth, tt.args.handler)
		})
	}
}
