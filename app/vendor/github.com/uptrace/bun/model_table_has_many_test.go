package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

func Test_newHasManyModel(t *testing.T) {
	type args struct {
		j *relationJoin
	}
	tests := []struct {
		name string
		args args
		want *hasManyModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHasManyModel(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHasManyModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasManyModel_ScanRows(t *testing.T) {
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
			m := &hasManyModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("hasManyModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hasManyModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasManyModel_Scan(t *testing.T) {
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
			m := &hasManyModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("hasManyModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hasManyModel_parkStruct(t *testing.T) {
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
			m := &hasManyModel{
				sliceTableModel: tt.fields.sliceTableModel,
				baseTable:       tt.fields.baseTable,
				rel:             tt.fields.rel,
				baseValues:      tt.fields.baseValues,
				structKey:       tt.fields.structKey,
			}
			if err := m.parkStruct(); (err != nil) != tt.wantErr {
				t.Errorf("hasManyModel.parkStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_baseValues(t *testing.T) {
	type args struct {
		model  TableModel
		fields []*schema.Field
	}
	tests := []struct {
		name string
		args args
		want map[internal.MapKey][]reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseValues(tt.args.model, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modelKey(t *testing.T) {
	type args struct {
		key    []interface{}
		strct  reflect.Value
		fields []*schema.Field
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modelKey(tt.args.key, tt.args.strct, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modelKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
