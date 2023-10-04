package swag

import "testing"

func TestSpec_ReadDoc(t *testing.T) {
	type fields struct {
		Version          string
		Host             string
		BasePath         string
		Schemes          []string
		Title            string
		Description      string
		InfoInstanceName string
		SwaggerTemplate  string
		LeftDelim        string
		RightDelim       string
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
			i := &Spec{
				Version:          tt.fields.Version,
				Host:             tt.fields.Host,
				BasePath:         tt.fields.BasePath,
				Schemes:          tt.fields.Schemes,
				Title:            tt.fields.Title,
				Description:      tt.fields.Description,
				InfoInstanceName: tt.fields.InfoInstanceName,
				SwaggerTemplate:  tt.fields.SwaggerTemplate,
				LeftDelim:        tt.fields.LeftDelim,
				RightDelim:       tt.fields.RightDelim,
			}
			if got := i.ReadDoc(); got != tt.want {
				t.Errorf("Spec.ReadDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpec_InstanceName(t *testing.T) {
	type fields struct {
		Version          string
		Host             string
		BasePath         string
		Schemes          []string
		Title            string
		Description      string
		InfoInstanceName string
		SwaggerTemplate  string
		LeftDelim        string
		RightDelim       string
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
			i := &Spec{
				Version:          tt.fields.Version,
				Host:             tt.fields.Host,
				BasePath:         tt.fields.BasePath,
				Schemes:          tt.fields.Schemes,
				Title:            tt.fields.Title,
				Description:      tt.fields.Description,
				InfoInstanceName: tt.fields.InfoInstanceName,
				SwaggerTemplate:  tt.fields.SwaggerTemplate,
				LeftDelim:        tt.fields.LeftDelim,
				RightDelim:       tt.fields.RightDelim,
			}
			if got := i.InstanceName(); got != tt.want {
				t.Errorf("Spec.InstanceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
