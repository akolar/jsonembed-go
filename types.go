package jsonembed

type mapParserFunc func(naiveMap) (naiveMap, bool)

func parseIntMap(obj naiveMap) (naiveMap, bool) {
	intMap := make(naiveMap)

	for k, v := range obj {
		num, ok := v.(float64)
		if !ok {
			return nil, false
		}

		integer := int(num)

		if num != float64(integer) {
			return nil, false
		}

		intMap[k] = integer
	}

	return intMap, true
}

func parseFloatMap(obj naiveMap) (naiveMap, bool) {
	for _, v := range obj {
		_, ok := v.(float64)
		if !ok {
			return nil, false
		}
	}

	return obj, true
}

func parseBoolMap(obj naiveMap) (naiveMap, bool) {
	for _, v := range obj {
		_, ok := v.(bool)
		if !ok {
			return nil, false
		}
	}

	return obj, true
}

func parseStringMap(obj naiveMap) (naiveMap, bool) {
	for _, v := range obj {
		_, ok := v.(string)
		if !ok {
			return nil, false
		}
	}

	return obj, true
}

func parseMixedMap(obj naiveMap) (naiveMap, bool) {
	for _, v := range obj {
		ok := isValidType(v)
		if !ok {
			return nil, false
		}
	}

	return obj, true
}

func isValidType(v arbitraryType) bool {
	switch v.(type) {
	case float64, bool, string:
		return true
	}

	return false
}
