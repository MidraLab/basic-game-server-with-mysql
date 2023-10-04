package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewMergeQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMergeQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMergeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Apply(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		fn func(*MergeQuery) *MergeQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_With(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.With(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_WithRecursive(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.WithRecursive(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.WithRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Returning(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Returning(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Returning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Using(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Using(tt.args.s, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Using() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_On(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.On(tt.args.s, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.On() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_WhenInsert(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		expr string
		fn   func(q *InsertQuery) *InsertQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.WhenInsert(tt.args.expr, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.WhenInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_WhenUpdate(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		expr string
		fn   func(q *UpdateQuery) *UpdateQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.WhenUpdate(tt.args.expr, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.WhenUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_WhenDelete(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		expr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.WhenDelete(tt.args.expr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.WhenDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_When(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		expr string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.When(tt.args.expr, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.When() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("MergeQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_Scan(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("MergeQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMergeQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_scanOrExec(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
	}
	type args struct {
		ctx     context.Context
		dest    []interface{}
		hasDest bool
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			got, err := q.scanOrExec(tt.args.ctx, tt.args.dest, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("MergeQuery.scanOrExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeQuery.scanOrExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeQuery_String(t *testing.T) {
	type fields struct {
		baseQuery      baseQuery
		returningQuery returningQuery
		using          schema.QueryWithArgs
		on             schema.QueryWithArgs
		when           []schema.QueryAppender
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
			q := &MergeQuery{
				baseQuery:      tt.fields.baseQuery,
				returningQuery: tt.fields.returningQuery,
				using:          tt.fields.using,
				on:             tt.fields.on,
				when:           tt.fields.when,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("MergeQuery.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whenInsert_AppendQuery(t *testing.T) {
	type fields struct {
		expr  string
		query *InsertQuery
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
			w := &whenInsert{
				expr:  tt.fields.expr,
				query: tt.fields.query,
			}
			got, err := w.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("whenInsert.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whenInsert.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whenUpdate_AppendQuery(t *testing.T) {
	type fields struct {
		expr  string
		query *UpdateQuery
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
			w := &whenUpdate{
				expr:  tt.fields.expr,
				query: tt.fields.query,
			}
			got, err := w.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("whenUpdate.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whenUpdate.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whenDelete_AppendQuery(t *testing.T) {
	type fields struct {
		expr string
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
			w := &whenDelete{
				expr: tt.fields.expr,
			}
			got, err := w.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("whenDelete.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whenDelete.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
