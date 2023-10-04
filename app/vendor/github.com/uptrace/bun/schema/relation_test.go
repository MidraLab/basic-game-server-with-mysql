package schema

import "testing"

func TestRelation_String(t *testing.T) {
	type fields struct {
		Type             int
		Field            *Field
		JoinTable        *Table
		BaseFields       []*Field
		JoinFields       []*Field
		OnUpdate         string
		OnDelete         string
		Condition        []string
		PolymorphicField *Field
		PolymorphicValue string
		M2MTable         *Table
		M2MBaseFields    []*Field
		M2MJoinFields    []*Field
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
			r := &Relation{
				Type:             tt.fields.Type,
				Field:            tt.fields.Field,
				JoinTable:        tt.fields.JoinTable,
				BaseFields:       tt.fields.BaseFields,
				JoinFields:       tt.fields.JoinFields,
				OnUpdate:         tt.fields.OnUpdate,
				OnDelete:         tt.fields.OnDelete,
				Condition:        tt.fields.Condition,
				PolymorphicField: tt.fields.PolymorphicField,
				PolymorphicValue: tt.fields.PolymorphicValue,
				M2MTable:         tt.fields.M2MTable,
				M2MBaseFields:    tt.fields.M2MBaseFields,
				M2MJoinFields:    tt.fields.M2MJoinFields,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Relation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
