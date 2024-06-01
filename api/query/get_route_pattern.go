package query

import "strings"

func GetRoutePattern(parts ...string) string {
	route := strings.Join(parts, "/")

	return "GET /" + route
}
