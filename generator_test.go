package jsonembed

import (
	"fmt"
	goparser "go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GeneratorTestSuite struct {
	suite.Suite
}

func (suite *GeneratorTestSuite) TestSerializeFail() {
	mapSerializers = []naiveMapSerializer{
		{mapType: intType, parser: parseIntMap},
	}

	_, err := serializeNaive(naiveMap{"a": false})
	suite.Equal(SerializationFailedError, err)
}

func (suite *GeneratorTestSuite) TestSerialize() {
	tests := []struct {
		input   naiveMap
		mapType string
	}{
		{naiveMap{"a": 1.0, "b": 2.0}, "map[string]int"},
		{naiveMap{"a": 1.1, "b": 2.2}, "map[string]float64"},
		{naiveMap{"a": true, "b": false}, "map[string]bool"},
		{naiveMap{"a": "a", "b": "b"}, "map[string]string"},
		{naiveMap{"a": "a", "b": 1.0}, "map[string]interface{}"},
	}

	for i, t := range tests {
		res, err := serializeNaive(t.input)
		suite.Nil(err, fmt.Sprintf("failed test %d", i))
		suite.Contains(res, t.mapType, fmt.Sprintf("failed test %d", i))
	}
}

func (suite *GeneratorTestSuite) TestNaiveMapString() {
	tests := []struct {
		input   naiveMap
		mapType string
		output  string
	}{
		{naiveMap{"a": true}, "bool", "map[string]bool {\n\"a\": true,\n}"},
		{naiveMap{"a": 1}, "int", "map[string]int {\n\"a\": 1,\n}"},
		{naiveMap{"a": 1.1}, "float64", "map[string]float64 {\n\"a\": 1.1,\n}"},
		{naiveMap{"a": "string"}, "string", "map[string]string {\n\"a\": \"string\",\n}"},
	}

	for i, t := range tests {
		res := naiveMapString(t.mapType, t.input)
		suite.Equal(t.output, res, fmt.Sprintf("failed test %d", i))

		_, err := goparser.ParseExpr(res)
		suite.Nil(err, fmt.Sprintf("failed test %d", i))
	}
}

func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}
