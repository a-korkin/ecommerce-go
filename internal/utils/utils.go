package utils

import (
	"strings"
)

func GetQueryParams(url string) map[string]string {
	params := make(map[string]string, 0)
	if url == "" {
		return params
	}
	for _, query := range strings.Split(url, "&") {
		q := strings.Split(query, "=")
		params[q[0]] = q[1]
	}

	return params
}
