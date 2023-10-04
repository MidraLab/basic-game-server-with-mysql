package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewDropIndexQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDropIndexQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDropIndexQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Concurrently(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Concurrently(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Concurrently() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_IfExists(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.IfExists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.IfExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Cascade(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Cascade(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Cascade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Restrict(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Restrict(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Restrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Index(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Index(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
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
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("DropIndexQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
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
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropIndexQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropIndexQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		concurrently bool
		ifExists     bool
		index        schema.QueryWithArgs
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
			q := &DropIndexQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				concurrently: tt.fields.concurrently,
				ifExists:     tt.fields.ifExists,
				index:        tt.fields.index,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropIndexQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropIndexQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
