package util

import "strings"

//client_id=test_client_1&redirect_uri=https://www.example.com&response_type=code&scope=all&state=dsdsd

func ParseUrlQuery(query string) *map[string]string {
	result := make(map[string]string)
	strArr := strings.Split(query, "&")
	for _, v := range strArr {
		vs := strings.Split(v, "=")
		if len(vs) == 2 {
			result[vs[0]] = vs[1]
		}
	}
	return &result
}
