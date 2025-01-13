package utils

import "testing"

func TestGetVars(t *testing.T) {
	url := "/products/5ee2371d-3065-4f69-8e3f-d1a31df2ef74"
	path := "/{id}"
	vars := map[string]string{
		"id": "5ee2371d-3065-4f69-8e3f-d1a31df2ef74",
	}
	if vars["id"] != GetVars(url, path)["id"] {
		t.Errorf("Result was incorrect, expected: %s, got: %s", vars["id"], GetVars(url, path)["id"])
	}
}
