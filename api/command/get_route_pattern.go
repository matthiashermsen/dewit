package command

import "strings"

func GetRoutePattern(parts ...string) string {
	route := strings.Join(parts, "/")

	return "POST /" + route
}
