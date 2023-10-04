package schema

import (
	"reflect"
	"sync"
	"testing"
)

func Test_newTableInProgress(t *testing.T) {
	type args struct {
		table *Table
	}
	tests := []struct {
		name string
		args args
		want *tableInProgress
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTableInProgress(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTableInProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tableInProgress_init1(t *testing.T) {
	type fields struct {
		table     *Table
		init1Once sync.Once
		init2Once sync.Once
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
			inp := &tableInProgress{
				table:     tt.fields.table,
				init1Once: tt.fields.init1Once,
				init2Once: tt.fields.init2Once,
			}
			if got := inp.init1(); got != tt.want {
				t.Errorf("tableInProgress.init1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tableInProgress_init2(t *testing.T) {
	type fields struct {
		table     *Table
		init1Once sync.Once
		init2Once sync.Once
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
			inp := &tableInProgress{
				table:     tt.fields.table,
				init1Once: tt.fields.init1Once,
				init2Once: tt.fields.init2Once,
			}
			if got := inp.init2(); got != tt.want {
				t.Errorf("tableInProgress.init2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTables(t *testing.T) {
	type args struct {
		dialect Dialect
	}
	tests := []struct {
		name string
		args args
		want *Tables
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTables(tt.args.dialect); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTables_Register(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		models []interface{}
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
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			tr.Register(tt.args.models...)
		})
	}
}

func TestTables_Get(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			if got := tr.Get(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tables.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTables_Ref(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			if got := tr.Ref(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tables.Ref() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTables_table(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		typ             reflect.Type
		allowInProgress bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			if got := tr.table(tt.args.typ, tt.args.allowInProgress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tables.table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTables_ByModel(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			if got := tr.ByModel(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tables.ByModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTables_ByName(t *testing.T) {
	type fields struct {
		dialect    Dialect
		tables     sync.Map
		mu         sync.RWMutex
		inProgress map[reflect.Type]*tableInProgress
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tables{
				dialect:    tt.fields.dialect,
				tables:     tt.fields.tables,
				mu:         tt.fields.mu,
				inProgress: tt.fields.inProgress,
			}
			if got := tr.ByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tables.ByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
