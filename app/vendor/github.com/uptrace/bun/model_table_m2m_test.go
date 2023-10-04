package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

func Test_newM2MModel(t *testing.T) {
	type args struct {
		j *relationJoin
	}
	tests := []struct {
		name string
		args args
		want *m2mModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newM2MModel(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newM2MModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_m2mModel_ScanRows(t *testing.T) {
	type fields struct {
		sliceTableModel *sliceTableModel
		baseTable       *schema.Table
		rel             *schema.Relation
		baseValues      map[internal.MapKey][]reflect.Value
		structKey       []interface{}
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
			m := &m2mModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("m2mModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("m2mModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_m2mModel_Scan(t *testing.T) {
	type fields struct {
		sliceTableModel *sliceTableModel
		baseTable       *schema.Table
		rel             *schema.Relation
		baseValues      map[internal.MapKey][]reflect.Value
		structKey       []interface{}
	}
	type args struct {
		src interface{}
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
			m := &m2mModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("m2mModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_m2mModel_scanM2MColumn(t *testing.T) {
	type fields struct {
		sliceTableModel *sliceTableModel
		baseTable       *schema.Table
		rel             *schema.Relation
		baseValues      map[internal.MapKey][]reflect.Value
		structKey       []interface{}
	}
	type args struct {
		column string
		src    interface{}
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
			m := &m2mModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			if err := m.scanM2MColumn(tt.args.column, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("m2mModel.scanM2MColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_m2mModel_parkStruct(t *testing.T) {
	type fields struct {
		sliceTableModel *sliceTableModel
		baseTable       *schema.Table
		rel             *schema.Relation
		baseValues      map[internal.MapKey][]reflect.Value
		structKey       []interface{}
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
			m := &m2mModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			if err := m.parkStruct(); (err != nil) != tt.wantErr {
				t.Errorf("m2mModel.parkStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
