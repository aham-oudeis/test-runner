package value_interpolator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func InterpolateValue(value string, context interface{}) ( result interface{}, err error) {
	var path = SplitStringByRegex(value, "\\.|\\[|\\]");
	path = filter(path, func(s string) bool { return s != "" })
	//the first key is part of the struct and hence is converted to Title case when converting to struct
	//the rest of the nested keys are simply part of map and hence preserve their original case
	path[0] = cases.Title(language.English).String(path[0])

	result = context;

	for _, key := range path {
		fmt.Printf("result%+v:", result)
		if result == nil {
			//that means the path cannot be followed
			return nil, errors.New("invalid reference in the assertion")
		} else {
			result = FollowPath(result, key)
			//this needs to be added so that there is no pointer dereference error
			if result == nil {
				//that means the path cannot be followed
				return nil, errors.New("invalid reference in the assertion")
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
		case []string:
			idx, err := strconv.Atoi(key)

			if err != nil {
				return nil
			}

			if idx < 0 || idx >= len(v) {
				return nil
			}

			context = v[idx]
		case map[string]interface{}:
			context = v[key] 
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