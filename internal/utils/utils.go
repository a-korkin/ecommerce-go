package utils

import (
	"fmt"
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

func GetResouce(url string) string {
	path := strings.Split(url, "/")
	if len(path) == 1 {
		return "/"
	}
	return fmt.Sprintf("/%s", path[1])
}

func getPathTokens(url string) []string {
	tokens := strings.Split(url, "/")
	return tokens[2:]
}

func zip(pathPattern []string, tokens []string) map[string]string {
	tokensLen := len(tokens)
	result := make(map[string]string, 0)
	for i := 0; i < len(pathPattern); i++ {
		if tokensLen > i {
			value := tokens[i]
			key := strings.Replace(
				strings.Replace(pathPattern[i], "{", "", 1), "}", "", 1)
			result[key] = value
		}
	}
	return result
}

func GetVars(url string, path string) map[string]string {
	tokens := getPathTokens(url)
	patternPath := strings.Split(path, "/")[1:]
	return zip(patternPath, tokens)
}
