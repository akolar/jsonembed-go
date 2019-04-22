package jsonembed

import (
	"errors"
	"fmt"
	"strings"
)

type arbitraryType interface{}
type naiveMap map[string]arbitraryType

type mapSerializerFunc func(map[string]arbitraryType) (string, bool)
type listSerializerFunc func([]arbitraryType) (string, bool)

type naiveMapSerializer struct {
	mapType string
	parser  mapParserFunc
}

const (
	mapBaseSize       = 25
	mapEstimatePerRow = 15
)

const (
	float64Type = "float64"
	intType     = "int"
	boolType    = "bool"
	stringType  = "string"
	mixedType   = "interface{}"
)

var (
	SerializationFailedError = errors.New("serialization failed")
)

var (
	mapSerializers = []naiveMapSerializer{
		{mapType: intType, parser: parseIntMap},
		{mapType: boolType, parser: parseBoolMap},
		{mapType: float64Type, parser: parseFloatMap},
		{mapType: stringType, parser: parseStringMap},
		{mapType: mixedType, parser: parseMixedMap},
	}
	listSerializers = []listSerializerFunc{}
)

func serializeNaive(obj naiveMap) (string, error) {
	for _, s := range mapSerializers {
		if data, ok := s.parser(obj); ok {
			return naiveMapString(s.mapType, data), nil
		}
	}

	return "", SerializationFailedError
}

func naiveMapString(mapType string, data naiveMap) string {
	var sb strings.Builder
	sb.Grow(mapBaseSize + len(data)*mapEstimatePerRow)

	// map[string]type {\n
	sb.WriteString("map[string]")
	sb.WriteString(mapType)
	sb.WriteString(" {\n")

	// "key": value,\n or "key": "string"\n
	for k, v := range data {
		_, vIsString := v.(string)

		sb.WriteString(`"`)
		sb.WriteString(k)
		sb.WriteString(`": `)

		if vIsString {
			sb.WriteString(`"`)
		}
		sb.WriteString(fmt.Sprint(v))
		if vIsString {
			sb.WriteString(`"`)
		}

		sb.WriteString(",\n")
	}

	// }
	sb.WriteString("}")

	return sb.String()
}
