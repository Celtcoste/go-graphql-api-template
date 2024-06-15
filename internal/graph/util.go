package graph

import "github.com/celtcoste/go-graphql-api-template/pkg/util/gqlutil"

func ConvertArrayUUIDToString(uuids []gqlutil.UUID) []string {
	var arr []string
	for _, val := range uuids {
		arr = append(arr, val.String())
	}
	return arr
}
