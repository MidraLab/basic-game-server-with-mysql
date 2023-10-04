package bun

import (
	"context"
	"reflect"
	"testing"

	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

func Test_relationJoin_applyTo(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		q *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			j.applyTo(tt.args.q)
		})
	}
}

func Test_relationJoin_Select(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		ctx context.Context
		q   *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if err := j.Select(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("relationJoin.Select() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_relationJoin_selectMany(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		ctx context.Context
		q   *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if err := j.selectMany(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("relationJoin.selectMany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_relationJoin_manyQuery(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		q *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.manyQuery(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.manyQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_manyQueryCompositeIn(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		where []byte
		q     *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.manyQueryCompositeIn(tt.args.where, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.manyQueryCompositeIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_manyQueryMulti(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		where []byte
		q     *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.manyQueryMulti(tt.args.where, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.manyQueryMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_hasManyColumns(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		q *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.hasManyColumns(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.hasManyColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_selectM2M(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		ctx context.Context
		q   *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if err := j.selectM2M(tt.args.ctx, tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("relationJoin.selectM2M() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_relationJoin_m2mQuery(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		q *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.m2mQuery(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.m2mQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_hasParent(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.hasParent(); got != tt.want {
				t.Errorf("relationJoin.hasParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_appendAlias(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.appendAlias(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.appendAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_appendAliasColumn(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		column string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.appendAliasColumn(tt.args.fmter, tt.args.b, tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.appendAliasColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_appendBaseAlias(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.appendBaseAlias(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.appendBaseAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_appendSoftDelete(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		flags internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			if got := j.appendSoftDelete(tt.args.fmter, tt.args.b, tt.args.flags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.appendSoftDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendAlias(t *testing.T) {
	type args struct {
		b []byte
		j *relationJoin
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendAlias(tt.args.b, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relationJoin_appendHasOneJoin(t *testing.T) {
	type fields struct {
		Parent    *relationJoin
		BaseModel TableModel
		JoinModel TableModel
		Relation  *schema.Relation
		apply     func(*SelectQuery) *SelectQuery
		columns   []schema.QueryWithArgs
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		q     *SelectQuery
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
			j := &relationJoin{
				Parent:    tt.fields.Parent,
				BaseModel: tt.fields.BaseModel,
				JoinModel: tt.fields.JoinModel,
				Relation:  tt.fields.Relation,
				apply:     tt.fields.apply,
				columns:   tt.fields.columns,
			}
			got, err := j.appendHasOneJoin(tt.args.fmter, tt.args.b, tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("relationJoin.appendHasOneJoin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relationJoin.appendHasOneJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendChildValues(t *testing.T) {
	type args struct {
		fmter  schema.Formatter
		b      []byte
		v      reflect.Value
		index  []int
		fields []*schema.Field
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendChildValues(tt.args.fmter, tt.args.b, tt.args.v, tt.args.index, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendChildValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendMultiValues(t *testing.T) {
	type args struct {
		fmter      schema.Formatter
		b          []byte
		v          reflect.Value
		index      []int
		baseFields []*schema.Field
		joinFields []*schema.Field
		joinTable  schema.Safe
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendMultiValues(tt.args.fmter, tt.args.b, tt.args.v, tt.args.index, tt.args.baseFields, tt.args.joinFields, tt.args.joinTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendMultiValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
