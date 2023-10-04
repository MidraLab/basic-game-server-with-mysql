// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_newResolverContext(t *testing.T) {
	type args struct {
		options *ExpandOptions
	}
	tests := []struct {
		name string
		args args
		want *resolverContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newResolverContext(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResolverContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schemaLoader_transitiveResolver(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		basePath string
		ref      Ref
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *schemaLoader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if got := r.transitiveResolver(tt.args.basePath, tt.args.ref); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("schemaLoader.transitiveResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schemaLoader_updateBasePath(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		transitive *schemaLoader
		basePath   string
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
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if got := r.updateBasePath(tt.args.transitive, tt.args.basePath); got != tt.want {
				t.Errorf("schemaLoader.updateBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schemaLoader_resolveRef(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		ref      *Ref
		target   interface{}
		basePath string
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
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if err := r.resolveRef(tt.args.ref, tt.args.target, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("schemaLoader.resolveRef() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_schemaLoader_load(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		refURL *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		want1   url.URL
		want2   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			got, got1, got2, err := r.load(tt.args.refURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("schemaLoader.load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("schemaLoader.load() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("schemaLoader.load() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("schemaLoader.load() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_schemaLoader_isCircular(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		ref        *Ref
		basePath   string
		parentRefs []string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantFoundCycle bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if gotFoundCycle := r.isCircular(tt.args.ref, tt.args.basePath, tt.args.parentRefs...); gotFoundCycle != tt.wantFoundCycle {
				t.Errorf("schemaLoader.isCircular() = %v, want %v", gotFoundCycle, tt.wantFoundCycle)
			}
		})
	}
}

func Test_schemaLoader_Resolve(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		ref      *Ref
		target   interface{}
		basePath string
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
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if err := r.Resolve(tt.args.ref, tt.args.target, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("schemaLoader.Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_schemaLoader_deref(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		input      interface{}
		parentRefs []string
		basePath   string
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
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if err := r.deref(tt.args.input, tt.args.parentRefs, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("schemaLoader.deref() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_schemaLoader_shouldStopOnError(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			if got := r.shouldStopOnError(tt.args.err); got != tt.want {
				t.Errorf("schemaLoader.shouldStopOnError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_schemaLoader_setSchemaID(t *testing.T) {
	type fields struct {
		root    interface{}
		options *ExpandOptions
		cache   ResolutionCache
		context *resolverContext
	}
	type args struct {
		target   interface{}
		id       string
		basePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &schemaLoader{
				root:    tt.fields.root,
				options: tt.fields.options,
				cache:   tt.fields.cache,
				context: tt.fields.context,
			}
			got, got1 := r.setSchemaID(tt.args.target, tt.args.id, tt.args.basePath)
			if got != tt.want {
				t.Errorf("schemaLoader.setSchemaID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("schemaLoader.setSchemaID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_defaultSchemaLoader(t *testing.T) {
	type args struct {
		root          interface{}
		expandOptions *ExpandOptions
		cache         ResolutionCache
		context       *resolverContext
	}
	tests := []struct {
		name string
		args args
		want *schemaLoader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultSchemaLoader(tt.args.root, tt.args.expandOptions, tt.args.cache, tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultSchemaLoader() = %v, want %v", got, tt.want)
			}
		})
	}
}
