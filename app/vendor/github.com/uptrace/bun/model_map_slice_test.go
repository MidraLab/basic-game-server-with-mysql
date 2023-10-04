package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func Test_newMapSliceModel(t *testing.T) {
	type args struct {
		db   *DB
		dest *[]map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *mapSliceModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMapSliceModel(tt.args.db, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMapSliceModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSliceModel_Value(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
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
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			if got := m.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapSliceModel.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSliceModel_SetCap(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
	}
	type args struct {
		cap int
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
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			m.SetCap(tt.args.cap)
		})
	}
}

func Test_mapSliceModel_ScanRows(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
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
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapSliceModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mapSliceModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSliceModel_appendColumns(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
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
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			got, err := m.appendColumns(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapSliceModel.appendColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapSliceModel.appendColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSliceModel_appendValues(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
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
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			got, err := m.appendValues(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapSliceModel.appendValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapSliceModel.appendValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapSliceModel_initKeys(t *testing.T) {
	type fields struct {
		mapModel mapModel
		dest     *[]map[string]interface{}
		keys     []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapSliceModel{
				mapModel: tt.fields.mapModel,
				dest:     tt.fields.dest,
				keys:     tt.fields.keys,
			}
			if err := m.initKeys(); (err != nil) != tt.wantErr {
				t.Errorf("mapSliceModel.initKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
