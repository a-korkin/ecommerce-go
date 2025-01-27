package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetQueryParams(t *testing.T) {
	params := map[string]string{
		"page":     "1",
		"limit":    "100",
		"offset":   "20",
		"category": "91f43fe2-4d8a-4a21-865a-fe9dc49aa3c7",
	}
	tests := []struct {
		name     string
		positive bool
		url      string
		keys     []string
	}{
		{
			name:     "have page and offset",
			positive: true,
			url:      "page=1&offset=20",
			keys:     []string{"page", "offset"},
		},
		{
			name:     "dont' have page",
			positive: false,
			url:      "limit=100&offset=20",
			keys:     []string{"page"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetQueryParams(tt.url)
			if tt.positive {
				for _, key := range tt.keys {
					exp := params[key]
					got := result[key]
					if exp != got {
						t.Errorf("Incorrect result, expected: %s, got: %s", exp, got)
					}
				}
			} else {
				for _, key := range tt.keys {
					_, ok := result[key]
					if ok {
						t.Errorf("Incorrect result, can't have a : %s in params", key)
					}
				}
			}
		})
	}
}

func TestGetResource(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "is products",
			url:      "/products",
			expected: "/products",
		},
		{
			name:     "is categories",
			url:      "/categories/91f43fe2-4d8a-4a21-865a-fe9dc49aa3c7",
			expected: "/categories",
		},
		{
			name:     "is users",
			url:      "/users?page=2&limit=20",
			expected: "/users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetResouce(tt.url)
			if tt.expected != result {
				t.Errorf("Incorrect result, expected: %s, got: %s", tt.expected, result)
			}
		})
	}
}

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetVars(tt.url, tt.path)
			for _, field := range tt.fields {
				exp := vars[field]
				if exp != got[field] {
					t.Errorf("Incorrect result, expected: %s, got: %s", exp, got)
				}
			}
		})
	}
}

type User struct {
	ID        string `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

func TestUnmarshallingFromFile(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Failed to get working dir: %s", err)
	}
	filePath := filepath.Join(dir, "../../test", "users.json")
	data := make([]*User, 3)
	UnmarshallingFromFile(filePath, &data)
	if data[0] == nil {
		t.Errorf(
			"Wrong result unsmarshalling, expected: not nil, got: %v", data[0])
	}
}
