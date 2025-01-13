package utils

import (
	"testing"
)

func TestGetVars(t *testing.T) {
	vars := map[string]string{
		"id":       "5ee2371d-3065-4f69-8e3f-d1a31df2ef74",
		"category": "91f43fe2-4d8a-4a21-865a-fe9dc49aa3c7",
	}

	tests := []struct {
		name   string
		url    string
		path   string
		fields []string
	}{
		{
			name:   "have id",
			url:    "/products/5ee2371d-3065-4f69-8e3f-d1a31df2ef74",
			path:   "/{id}",
			fields: []string{"id"},
		},
		{
			name:   "have id and category",
			url:    "/products/5ee2371d-3065-4f69-8e3f-d1a31df2ef74/91f43fe2-4d8a-4a21-865a-fe9dc49aa3c7",
			path:   "/{id}/{category}",
			fields: []string{"id", "category"},
		},
		// {
		// 	name:   "don't have id",
		// 	url:    "/products",
		// 	path:   "/{id}",
		// 	fields: []string{"id"},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetVars(tt.url, tt.path)
			for _, field := range tt.fields {
				exp, ok := vars[field]
				if ok {
					if exp != got[field] {
						t.Errorf("Incorrect result, expected: %s, got: %s", exp, got)
					}
				} else {
					// if exp == got[field] {
					// 	t.Errorf("Incorrect result, expected: %s, got: %s", exp, got)
					// }
				}
			}
		})
	}
}
