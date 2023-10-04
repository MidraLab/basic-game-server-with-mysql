package schema

import (
	"reflect"
	"testing"
)

func TestSafe_AppendQuery(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		s       Safe
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Safe.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Safe.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName_AppendQuery(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		s       Name
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Name.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdent_AppendQuery(t *testing.T) {
	type args struct {
		fmter Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		s       Ident
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ident.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ident.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeQuery(t *testing.T) {
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want QueryWithArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeQuery(tt.args.query, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeIdent(t *testing.T) {
	type args struct {
		ident string
	}
	tests := []struct {
		name string
		args args
		want QueryWithArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsafeIdent(tt.args.ident); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsafeIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryWithArgs_IsZero(t *testing.T) {
	type fields struct {
		Query string
		Args  []interface{}
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
			q := QueryWithArgs{
				Query: tt.fields.Query,
				Args:  tt.fields.Args,
			}
			if got := q.IsZero(); got != tt.want {
				t.Errorf("QueryWithArgs.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryWithArgs_AppendQuery(t *testing.T) {
	type fields struct {
		Query string
		Args  []interface{}
	}
	type args struct {
		fmter Formatter
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
			q := QueryWithArgs{
				Query: tt.fields.Query,
				Args:  tt.fields.Args,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryWithArgs.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryWithArgs.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeQueryWithSep(t *testing.T) {
	type args struct {
		query string
		args  []interface{}
		sep   string
	}
	tests := []struct {
		name string
		args args
		want QueryWithSep
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeQueryWithSep(tt.args.query, tt.args.args, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeQueryWithSep() = %v, want %v", got, tt.want)
			}
		})
	}
}
