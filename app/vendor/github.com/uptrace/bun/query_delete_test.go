package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun/schema"
)

func TestNewDeleteQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeleteQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeleteQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Conn(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Model(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Err(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Apply(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		fn func(*DeleteQuery) *DeleteQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_With(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.With(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WithRecursive(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WithRecursive(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WithRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Table(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_TableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WherePK(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		cols []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Where(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WhereOr(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WhereGroup(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		sep string
		fn  func(*DeleteQuery) *DeleteQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WhereDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_ForceDelete(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.ForceDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.ForceDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Returning(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Returning(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Returning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Operation(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("DeleteQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_AppendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_isSoftDelete(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.isSoftDelete(); got != tt.want {
				t.Errorf("DeleteQuery.isSoftDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_softDeleteSet(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		fmter schema.Formatter
		tm    time.Time
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.softDeleteSet(tt.args.fmter, tt.args.tm); got != tt.want {
				t.Errorf("DeleteQuery.softDeleteSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_Scan(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteQuery_Exec(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_scanOrExec(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			got, err := q.scanOrExec(tt.args.ctx, tt.args.dest, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.scanOrExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.scanOrExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_beforeDeleteHook(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if err := q.beforeDeleteHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.beforeDeleteHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteQuery_afterDeleteHook(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if err := q.afterDeleteHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DeleteQuery.afterDeleteHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteQuery_String(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("DeleteQuery.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_QueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
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
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.QueryBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.QueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteQuery_ApplyQueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery whereBaseQuery
		returningQuery returningQuery
	}
	type args struct {
		fn func(QueryBuilder) QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DeleteQuery{
				whereBaseQuery: tt.fields.whereBaseQuery,
				returningQuery: tt.fields.returningQuery,
			}
			if got := q.ApplyQueryBuilder(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteQuery.ApplyQueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_WhereGroup(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_Where(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_WhereOr(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_WhereDeleted(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_WherePK(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteQueryBuilder_Unwrap(t *testing.T) {
	type fields struct {
		DeleteQuery *DeleteQuery
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
			q := &deleteQueryBuilder{
				DeleteQuery: tt.fields.DeleteQuery,
			}
			if got := q.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteQueryBuilder.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
