package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewCreateIndexQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateIndexQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateIndexQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Conn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Model(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Err(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Unique(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Unique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Concurrently(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Concurrently(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Concurrently() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_IfNotExists(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.IfNotExists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.IfNotExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Index(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Index(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_IndexExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.IndexExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.IndexExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Table(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_TableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Using(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Using(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Using() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Column(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_ExcludeColumn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.ExcludeColumn(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.ExcludeColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Include(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Include(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Include() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_IncludeExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.IncludeExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.IncludeExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Where(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_WhereOr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Operation(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
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
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("CreateIndexQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_AppendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateIndexQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateIndexQuery_Exec(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		unique         bool
		fulltext       bool
		spatial        bool
		concurrently   bool
		ifNotExists    bool
		index          schema.QueryWithArgs
		using          schema.QueryWithArgs
		include        []schema.QueryWithArgs
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateIndexQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				unique:         tt.fields.unique,
				fulltext:       tt.fields.fulltext,
				spatial:        tt.fields.spatial,
				concurrently:   tt.fields.concurrently,
				ifNotExists:    tt.fields.ifNotExists,
				index:          tt.fields.index,
				using:          tt.fields.using,
				include:        tt.fields.include,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateIndexQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateIndexQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
