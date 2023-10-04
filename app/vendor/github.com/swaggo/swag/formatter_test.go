package swag

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestNewFormatter(t *testing.T) {
	tests := []struct {
		name string
		want *Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatter_Format(t *testing.T) {
	type fields struct {
		debug Debugger
	}
	type args struct {
		fileName string
		contents []byte
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
			f := &Formatter{
				debug: tt.fields.debug,
			}
			got, err := f.Format(tt.args.fileName, tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("Formatter.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Formatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_edits_apply(t *testing.T) {
	type args struct {
		contents []byte
	}
	tests := []struct {
		name  string
		edits edits
		args  args
		want  []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.edits.apply(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("edits.apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatFuncDoc(t *testing.T) {
	type args struct {
		fileSet     *token.FileSet
		commentList []*ast.Comment
		edits       *edits
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatFuncDoc(tt.args.fileSet, tt.args.commentList, tt.args.edits)
		})
	}
}

func Test_splitComment2(t *testing.T) {
	type args struct {
		attr string
		body string
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
			if got := splitComment2(tt.args.attr, tt.args.body); got != tt.want {
				t.Errorf("splitComment2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceRange(t *testing.T) {
	type args struct {
		s     string
		start int
		end   int
		new   string
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
			if got := replaceRange(tt.args.s, tt.args.start, tt.args.end, tt.args.new); got != tt.want {
				t.Errorf("replaceRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swagComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := swagComment(tt.args.comment)
			if got != tt.want {
				t.Errorf("swagComment() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swagComment() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("swagComment() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
