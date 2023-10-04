/*
Package inflection pluralizes and singularizes English nouns.

		inflection.Plural("person") => "people"
		inflection.Plural("Person") => "People"
		inflection.Plural("PERSON") => "PEOPLE"

		inflection.Singular("people") => "person"
		inflection.Singular("People") => "Person"
		inflection.Singular("PEOPLE") => "PERSON"

		inflection.Plural("FancyPerson") => "FancydPeople"
		inflection.Singular("FancyPeople") => "FancydPerson"

Standard rules are from Rails's ActiveSupport (https://github.com/rails/rails/blob/master/activesupport/lib/active_support/inflections.rb)

If you want to register more rules, follow:

		inflection.AddUncountable("fish")
		inflection.AddIrregular("person", "people")
		inflection.AddPlural("(bu)s$", "${1}ses") # "bus" => "buses" / "BUS" => "BUSES" / "Bus" => "Buses"
		inflection.AddSingular("(bus)(es)?$", "${1}") # "buses" => "bus" / "Buses" => "Bus" / "BUSES" => "BUS"
*/
package inflection

import (
	"reflect"
	"testing"
)

func Test_compile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compile()
		})
	}
}

func TestAddPlural(t *testing.T) {
	type args struct {
		find    string
		replace string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddPlural(tt.args.find, tt.args.replace)
		})
	}
}

func TestAddSingular(t *testing.T) {
	type args struct {
		find    string
		replace string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddSingular(tt.args.find, tt.args.replace)
		})
	}
}

func TestAddIrregular(t *testing.T) {
	type args struct {
		singular string
		plural   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddIrregular(tt.args.singular, tt.args.plural)
		})
	}
}

func TestAddUncountable(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUncountable(tt.args.values...)
		})
	}
}

func TestGetPlural(t *testing.T) {
	tests := []struct {
		name string
		want RegularSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPlural(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlural() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSingular(t *testing.T) {
	tests := []struct {
		name string
		want RegularSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSingular(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSingular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIrregular(t *testing.T) {
	tests := []struct {
		name string
		want IrregularSlice
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIrregular(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIrregular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUncountable(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUncountable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUncountable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPlural(t *testing.T) {
	type args struct {
		inflections RegularSlice
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPlural(tt.args.inflections)
		})
	}
}

func TestSetSingular(t *testing.T) {
	type args struct {
		inflections RegularSlice
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSingular(tt.args.inflections)
		})
	}
}

func TestSetIrregular(t *testing.T) {
	type args struct {
		inflections IrregularSlice
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetIrregular(tt.args.inflections)
		})
	}
}

func TestSetUncountable(t *testing.T) {
	type args struct {
		inflections []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUncountable(tt.args.inflections)
		})
	}
}

func TestPlural(t *testing.T) {
	type args struct {
		str string
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
			if got := Plural(tt.args.str); got != tt.want {
				t.Errorf("Plural() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingular(t *testing.T) {
	type args struct {
		str string
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
			if got := Singular(tt.args.str); got != tt.want {
				t.Errorf("Singular() = %v, want %v", got, tt.want)
			}
		})
	}
}
