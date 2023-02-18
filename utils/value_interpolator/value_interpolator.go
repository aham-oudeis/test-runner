package value_interpolator

import (
	"errors"
	"regexp"
	"strconv"
)

func InterpolateValue(value string, context interface{}) ( result interface{}, err error) {
	var path = SplitStringByRegex(value, "\\.|\\[|\\]");
	path = filter(path, func(s string) bool { return s != "" })

	result = context;

	for _, key := range path {
		if result == nil {
			//that means the path cannot be followed
			return nil, errors.New("invalid path")
		} else {
			result = FollowPath(result, key)
			//this needs to be added so that there is no pointer dereference error
			if result == nil {
				//that means the path cannot be followed
				return nil, errors.New("invalid path")
			}
		}
	}

	return result, nil
}

func SplitStringByRegex(str string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.Split(str, -1)
}

func FollowPath(context interface{}, key string) interface{} {
	switch v := context.(type) {
		case map[string]interface{}:
			context = v[key]
		case []string:
			idx, err := strconv.Atoi(key)

			if err != nil {
				return nil
			}

			if idx < 0 || idx >= len(v) {
				return nil
			}

			context = v[idx]
		default:
			context = nil
	}

	return context
}

func filter[T any](s []T, fn func(T) bool) []T {
	var r []T
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}