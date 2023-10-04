package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun/schema"
)

func TestQueryEvent_Operation(t *testing.T) {
	type fields struct {
		DB            *DB
		QueryAppender schema.QueryAppender
		IQuery        Query
		Query         string
		QueryTemplate string
		QueryArgs     []interface{}
		Model         Model
		StartTime     time.Time
		Result        sql.Result
		Err           error
		Stash         map[interface{}]interface{}
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
			e := &QueryEvent{
				DB:            tt.fields.DB,
				QueryAppender: tt.fields.QueryAppender,
				IQuery:        tt.fields.IQuery,
				Query:         tt.fields.Query,
				QueryTemplate: tt.fields.QueryTemplate,
				QueryArgs:     tt.fields.QueryArgs,
				Model:         tt.fields.Model,
				StartTime:     tt.fields.StartTime,
				Result:        tt.fields.Result,
				Err:           tt.fields.Err,
				Stash:         tt.fields.Stash,
			}
			if got := e.Operation(); got != tt.want {
				t.Errorf("QueryEvent.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryOperation(t *testing.T) {
	type args struct {
		query string
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
			if got := queryOperation(tt.args.query); got != tt.want {
				t.Errorf("queryOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_beforeQuery(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx           context.Context
		iquery        Query
		queryTemplate string
		queryArgs     []interface{}
		query         string
		model         Model
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
		want1  *QueryEvent
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, got1 := db.beforeQuery(tt.args.ctx, tt.args.iquery, tt.args.queryTemplate, tt.args.queryArgs, tt.args.query, tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.beforeQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DB.beforeQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDB_afterQuery(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx   context.Context
		event *QueryEvent
		res   sql.Result
		err   error
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			db.afterQuery(tt.args.ctx, tt.args.event, tt.args.res, tt.args.err)
		})
	}
}

func TestDB_afterQueryFromIndex(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx       context.Context
		event     *QueryEvent
		hookIndex int
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			db.afterQueryFromIndex(tt.args.ctx, tt.args.event, tt.args.hookIndex)
		})
	}
}
