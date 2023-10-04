package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewSelectQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSelectQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSelectQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Conn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Model(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Err(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Apply(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fn func(*SelectQuery) *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_With(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.With(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WithRecursive(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WithRecursive(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WithRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Distinct(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Distinct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_DistinctOn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.DistinctOn(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.DistinctOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Table(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_TableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Column(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ExcludeColumn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ExcludeColumn(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ExcludeColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WherePK(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		cols []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Where(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WhereOr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WhereGroup(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		sep string
		fn  func(*SelectQuery) *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WhereDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_UseIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.UseIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.UseIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_UseIndexForJoin(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.UseIndexForJoin(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.UseIndexForJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_UseIndexForOrderBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.UseIndexForOrderBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.UseIndexForOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_UseIndexForGroupBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.UseIndexForGroupBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.UseIndexForGroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_IgnoreIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.IgnoreIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.IgnoreIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_IgnoreIndexForJoin(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.IgnoreIndexForJoin(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.IgnoreIndexForJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_IgnoreIndexForOrderBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.IgnoreIndexForOrderBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.IgnoreIndexForOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_IgnoreIndexForGroupBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.IgnoreIndexForGroupBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.IgnoreIndexForGroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ForceIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ForceIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ForceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ForceIndexForJoin(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ForceIndexForJoin(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ForceIndexForJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ForceIndexForOrderBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ForceIndexForOrderBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ForceIndexForOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ForceIndexForGroupBy(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ForceIndexForGroupBy(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ForceIndexForGroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Group(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Group(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_GroupExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		group string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.GroupExpr(tt.args.group, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.GroupExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Having(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		having string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Having(tt.args.having, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Having() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Order(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		orders []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Order(tt.args.orders...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_OrderExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.OrderExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.OrderExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Limit(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Limit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Offset(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Offset(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_For(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.For(tt.args.s, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.For() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Union(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Union(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_UnionAll(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.UnionAll(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.UnionAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Intersect(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Intersect(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_IntersectAll(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.IntersectAll(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.IntersectAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Except(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Except(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Except() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ExceptAll(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ExceptAll(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ExceptAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_addUnion(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		expr  string
		other *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.addUnion(tt.args.expr, tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.addUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Join(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		join string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Join(tt.args.join, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_JoinOn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		cond string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.JoinOn(tt.args.cond, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.JoinOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_JoinOnOr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		cond string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.JoinOnOr(tt.args.cond, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.JoinOnOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_joinOn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		cond string
		args []interface{}
		sep  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.joinOn(tt.args.cond, tt.args.args, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.joinOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Relation(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		name  string
		apply []func(*SelectQuery) *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Relation(tt.args.name, tt.args.apply...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Relation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_forEachInlineRelJoin(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fn func(*relationJoin) error
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q.forEachInlineRelJoin(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.forEachInlineRelJoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery__forEachInlineRelJoin(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fn    func(*relationJoin) error
		joins []relationJoin
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q._forEachInlineRelJoin(tt.args.fn, tt.args.joins); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery._forEachInlineRelJoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery_selectJoins(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx   context.Context
		joins []relationJoin
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q.selectJoins(tt.args.ctx, tt.args.joins); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.selectJoins() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery_Operation(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("SelectQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_AppendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_appendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		count bool
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.appendQuery(tt.args.fmter, tt.args.b, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.appendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.appendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_appendColumns(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.appendColumns(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.appendColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.appendColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_appendInlineRelColumns(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		join  *relationJoin
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.appendInlineRelColumns(tt.args.fmter, tt.args.b, tt.args.join)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.appendInlineRelColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.appendInlineRelColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_appendTables(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.appendTables(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.appendTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.appendTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_appendOrder(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.appendOrder(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.appendOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.appendOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Rows(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.Rows(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.Rows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Exec(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			gotRes, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SelectQuery.Exec() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestSelectQuery_Scan(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery_beforeSelectHook(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q.beforeSelectHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.beforeSelectHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery_afterSelectHook(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if err := q.afterSelectHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.afterSelectHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectQuery_Count(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.Count(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ScanAndCount(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.ScanAndCount(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.ScanAndCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.ScanAndCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_scanAndCountConc(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.scanAndCountConc(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.scanAndCountConc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.scanAndCountConc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_scanAndCountSeq(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.scanAndCountSeq(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.scanAndCountSeq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.scanAndCountSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_Exists(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.Exists(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_selectExists(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.selectExists(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.selectExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.selectExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_whereExists(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			got, err := q.whereExists(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectQuery.whereExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SelectQuery.whereExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_String(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
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
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("SelectQuery.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_QueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	tests := []struct {
		name   string
		fields fields
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.QueryBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.QueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectQuery_ApplyQueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		idxHintsQuery  idxHintsQuery
		distinctOn     []schema.QueryWithArgs
		joins          []joinQuery
		group          []schema.QueryWithArgs
		having         []schema.QueryWithArgs
		order          []schema.QueryWithArgs
		limit          int32
		offset         int32
		selFor         schema.QueryWithArgs
		union          []union
	}
	type args struct {
		fn func(QueryBuilder) QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SelectQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				idxHintsQuery:  tt.fields.idxHintsQuery,
				distinctOn:     tt.fields.distinctOn,
				joins:          tt.fields.joins,
				group:          tt.fields.group,
				having:         tt.fields.having,
				order:          tt.fields.order,
				limit:          tt.fields.limit,
				offset:         tt.fields.offset,
				selFor:         tt.fields.selFor,
				union:          tt.fields.union,
			}
			if got := q.ApplyQueryBuilder(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectQuery.ApplyQueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_WhereGroup(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	type args struct {
		sep string
		fn  func(QueryBuilder) QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_Where(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_WhereOr(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_WhereDeleted(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_WherePK(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	type args struct {
		cols []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectQueryBuilder_Unwrap(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &selectQueryBuilder{
				SelectQuery: tt.fields.SelectQuery,
			}
			if got := q.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectQueryBuilder.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinQuery_AppendQuery(t *testing.T) {
	type fields struct {
		join schema.QueryWithArgs
		on   []schema.QueryWithSep
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
			j := &joinQuery{
				join: tt.fields.join,
				on:   tt.fields.on,
			}
			got, err := j.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("joinQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joinQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countQuery_AppendQuery(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
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
			q := countQuery{
				SelectQuery: tt.fields.SelectQuery,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("countQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectExistsQuery_AppendQuery(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
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
			q := selectExistsQuery{
				SelectQuery: tt.fields.SelectQuery,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("selectExistsQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectExistsQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereExistsQuery_AppendQuery(t *testing.T) {
	type fields struct {
		SelectQuery *SelectQuery
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
			q := whereExistsQuery{
				SelectQuery: tt.fields.SelectQuery,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereExistsQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereExistsQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
