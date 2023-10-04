package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewUpdateQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Conn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Model(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Err(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Apply(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fn func(*UpdateQuery) *UpdateQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_With(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.With(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WithRecursive(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WithRecursive(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WithRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Table(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_TableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Column(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_ExcludeColumn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.ExcludeColumn(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.ExcludeColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Set(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Set(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_SetColumn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		column string
		query  string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.SetColumn(tt.args.column, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.SetColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Value(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		column string
		query  string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Value(tt.args.column, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_OmitZero(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.OmitZero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.OmitZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WherePK(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		cols []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Where(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WhereOr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WhereGroup(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		sep string
		fn  func(*UpdateQuery) *UpdateQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WhereDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Returning(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Returning(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Returning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Operation(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("UpdateQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_AppendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_mustAppendSet(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.mustAppendSet(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.mustAppendSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.mustAppendSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_appendSetStruct(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		model *structTableModel
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.appendSetStruct(tt.args.fmter, tt.args.b, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.appendSetStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.appendSetStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_appendOtherTables(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.appendOtherTables(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.appendOtherTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.appendOtherTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Bulk(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.Bulk(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Bulk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_updateSliceSet(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fmter schema.Formatter
		model *sliceTableModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.updateSliceSet(tt.args.fmter, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.updateSliceSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateQuery.updateSliceSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_updateSliceWhere(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fmter schema.Formatter
		model *sliceTableModel
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.updateSliceWhere(tt.args.fmter, tt.args.model); got != tt.want {
				t.Errorf("UpdateQuery.updateSliceWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_Scan(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateQuery_Exec(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_scanOrExec(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			got, err := q.scanOrExec(tt.args.ctx, tt.args.dest, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.scanOrExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.scanOrExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_beforeUpdateHook(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if err := q.beforeUpdateHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.beforeUpdateHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateQuery_afterUpdateHook(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if err := q.afterUpdateHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateQuery.afterUpdateHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateQuery_FQN(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		column string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Ident
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.FQN(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.FQN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_hasTableAlias(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fmter schema.Formatter
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.hasTableAlias(tt.args.fmter); got != tt.want {
				t.Errorf("UpdateQuery.hasTableAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_String(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("UpdateQuery.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_QueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
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
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.QueryBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.QueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_ApplyQueryBuilder(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		fn func(QueryBuilder) QueryBuilder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.ApplyQueryBuilder(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.ApplyQueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_WhereGroup(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.WhereGroup(tt.args.sep, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.WhereGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_Where(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_WhereOr(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_WhereDeleted(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.WhereDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.WhereDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_WhereAllWithDeleted(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.WhereAllWithDeleted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.WhereAllWithDeleted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_WherePK(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.WherePK(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.WherePK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateQueryBuilder_Unwrap(t *testing.T) {
	type fields struct {
		UpdateQuery *UpdateQuery
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
			q := &updateQueryBuilder{
				UpdateQuery: tt.fields.UpdateQuery,
			}
			if got := q.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateQueryBuilder.Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_UseIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.UseIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.UseIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_IgnoreIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.IgnoreIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.IgnoreIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateQuery_ForceIndex(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		setQuery         setQuery
		idxHintsQuery    idxHintsQuery
		omitZero         bool
	}
	type args struct {
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UpdateQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				setQuery:         tt.fields.setQuery,
				idxHintsQuery:    tt.fields.idxHintsQuery,
				omitZero:         tt.fields.omitZero,
			}
			if got := q.ForceIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateQuery.ForceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
