// Package jlexer contains a JSON lexer implementation.
//
// It is expected that it is mostly used with generated parser code, so the interface is tuned
// for a parser that knows what kind of data is expected.
package jlexer

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLexer_FetchToken(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.FetchToken()
		})
	}
}

func Test_isTokenEnd(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTokenEnd(tt.args.c); got != tt.want {
				t.Errorf("isTokenEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_fetchNull(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.fetchNull()
		})
	}
}

func TestLexer_fetchTrue(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.fetchTrue()
		})
	}
}

func TestLexer_fetchFalse(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.fetchFalse()
		})
	}
}

func TestLexer_fetchNumber(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.fetchNumber()
		})
	}
}

func Test_findStringLen(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
		wantLength  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, gotLength := findStringLen(tt.args.data)
			if gotIsValid != tt.wantIsValid {
				t.Errorf("findStringLen() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
			if gotLength != tt.wantLength {
				t.Errorf("findStringLen() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func TestLexer_unescapeStringToken(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if err := r.unescapeStringToken(); (err != nil) != tt.wantErr {
				t.Errorf("Lexer.unescapeStringToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getu4(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getu4(tt.args.s); got != tt.want {
				t.Errorf("getu4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeEscape(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name               string
		args               args
		wantDecoded        rune
		wantBytesProcessed int
		wantErr            bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDecoded, gotBytesProcessed, err := decodeEscape(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeEscape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDecoded != tt.wantDecoded {
				t.Errorf("decodeEscape() gotDecoded = %v, want %v", gotDecoded, tt.wantDecoded)
			}
			if gotBytesProcessed != tt.wantBytesProcessed {
				t.Errorf("decodeEscape() gotBytesProcessed = %v, want %v", gotBytesProcessed, tt.wantBytesProcessed)
			}
		})
	}
}

func TestLexer_fetchString(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.fetchString()
		})
	}
}

func TestLexer_scanToken(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.scanToken()
		})
	}
}

func TestLexer_consume(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.consume()
		})
	}
}

func TestLexer_Ok(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Ok(); got != tt.want {
				t.Errorf("Lexer.Ok() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_errParse(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		what string
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.errParse(tt.args.what)
		})
	}
}

func TestLexer_errSyntax(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.errSyntax()
		})
	}
}

func TestLexer_errInvalidToken(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		expected string
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.errInvalidToken(tt.args.expected)
		})
	}
}

func TestLexer_GetPos(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.GetPos(); got != tt.want {
				t.Errorf("Lexer.GetPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Delim(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		c byte
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.Delim(tt.args.c)
		})
	}
}

func TestLexer_IsDelim(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		c byte
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.IsDelim(tt.args.c); got != tt.want {
				t.Errorf("Lexer.IsDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Null(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.Null()
		})
	}
}

func TestLexer_IsNull(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.IsNull(); got != tt.want {
				t.Errorf("Lexer.IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Skip(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.Skip()
		})
	}
}

func TestLexer_SkipRecursive(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.SkipRecursive()
		})
	}
}

func TestLexer_Raw(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Raw(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.Raw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_IsStart(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.IsStart(); got != tt.want {
				t.Errorf("Lexer.IsStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Consumed(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.Consumed()
		})
	}
}

func TestLexer_unsafeString(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		skipUnescape bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			got, got1 := r.unsafeString(tt.args.skipUnescape)
			if got != tt.want {
				t.Errorf("Lexer.unsafeString() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Lexer.unsafeString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLexer_UnsafeString(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.UnsafeString(); got != tt.want {
				t.Errorf("Lexer.UnsafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_UnsafeBytes(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.UnsafeBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.UnsafeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_UnsafeFieldName(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		skipUnescape bool
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.UnsafeFieldName(tt.args.skipUnescape); got != tt.want {
				t.Errorf("Lexer.UnsafeFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_String(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Lexer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_StringIntern(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.StringIntern(); got != tt.want {
				t.Errorf("Lexer.StringIntern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Bytes(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Bool(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Bool(); got != tt.want {
				t.Errorf("Lexer.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_number(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.number(); got != tt.want {
				t.Errorf("Lexer.number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint8(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint8(); got != tt.want {
				t.Errorf("Lexer.Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint16(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint16(); got != tt.want {
				t.Errorf("Lexer.Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint32(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint32(); got != tt.want {
				t.Errorf("Lexer.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint64(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint64(); got != tt.want {
				t.Errorf("Lexer.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint(); got != tt.want {
				t.Errorf("Lexer.Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int8(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int8(); got != tt.want {
				t.Errorf("Lexer.Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int16(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int16(); got != tt.want {
				t.Errorf("Lexer.Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int32(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int32(); got != tt.want {
				t.Errorf("Lexer.Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int64(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int64(); got != tt.want {
				t.Errorf("Lexer.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int(); got != tt.want {
				t.Errorf("Lexer.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint8Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint8Str(); got != tt.want {
				t.Errorf("Lexer.Uint8Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint16Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint16Str(); got != tt.want {
				t.Errorf("Lexer.Uint16Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint32Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint32Str(); got != tt.want {
				t.Errorf("Lexer.Uint32Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Uint64Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Uint64Str(); got != tt.want {
				t.Errorf("Lexer.Uint64Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_UintStr(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.UintStr(); got != tt.want {
				t.Errorf("Lexer.UintStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_UintptrStr(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   uintptr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.UintptrStr(); got != tt.want {
				t.Errorf("Lexer.UintptrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int8Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int8Str(); got != tt.want {
				t.Errorf("Lexer.Int8Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int16Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int16Str(); got != tt.want {
				t.Errorf("Lexer.Int16Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int32Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int32Str(); got != tt.want {
				t.Errorf("Lexer.Int32Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Int64Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Int64Str(); got != tt.want {
				t.Errorf("Lexer.Int64Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_IntStr(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.IntStr(); got != tt.want {
				t.Errorf("Lexer.IntStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Float32(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Float32(); got != tt.want {
				t.Errorf("Lexer.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Float32Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Float32Str(); got != tt.want {
				t.Errorf("Lexer.Float32Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Float64(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Float64(); got != tt.want {
				t.Errorf("Lexer.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Float64Str(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Float64Str(); got != tt.want {
				t.Errorf("Lexer.Float64Str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Error(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if err := r.Error(); (err != nil) != tt.wantErr {
				t.Errorf("Lexer.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLexer_AddError(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		e error
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.AddError(tt.args.e)
		})
	}
}

func TestLexer_AddNonFatalError(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		e error
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.AddNonFatalError(tt.args.e)
		})
	}
}

func TestLexer_addNonfatalError(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	type args struct {
		err *LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.addNonfatalError(tt.args.err)
		})
	}
}

func TestLexer_GetNonFatalErrors(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   []*LexerError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.GetNonFatalErrors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.GetNonFatalErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_JsonNumber(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
		want   json.Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.JsonNumber(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.JsonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_Interface(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
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
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			if got := r.Interface(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lexer.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLexer_WantComma(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.WantComma()
		})
	}
}

func TestLexer_WantColon(t *testing.T) {
	type fields struct {
		Data              []byte
		start             int
		pos               int
		token             token
		firstElement      bool
		wantSep           byte
		UseMultipleErrors bool
		fatalError        error
		multipleErrors    []*LexerError
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Lexer{
				Data:              tt.fields.Data,
				start:             tt.fields.start,
				pos:               tt.fields.pos,
				token:             tt.fields.token,
				firstElement:      tt.fields.firstElement,
				wantSep:           tt.fields.wantSep,
				UseMultipleErrors: tt.fields.UseMultipleErrors,
				fatalError:        tt.fields.fatalError,
				multipleErrors:    tt.fields.multipleErrors,
			}
			r.WantColon()
		})
	}
}
