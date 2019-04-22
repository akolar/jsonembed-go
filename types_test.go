package jsonembed

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypesTestSuite struct {
	suite.Suite
}

func (suite *TypesTestSuite) TestParseIntMap() {
	tests := []struct {
		input   naiveMap
		success bool
	}{
		{naiveMap{"a": 1.0}, true},
		{naiveMap{"a": 1.1}, false},
		{naiveMap{"a": true}, false},
		{naiveMap{"a": "123"}, false},
		{naiveMap{"a": 1.0, "b": 2.0}, true},
		{naiveMap{"a": 1.0, "b": 2.0 + 1e-15}, false},
	}

	for i, t := range tests {
		res, ok := parseIntMap(t.input)
		suite.Equal(t.success, ok, fmt.Sprintf("failed test %d", i))

		for k, v := range res {
			_, ok := v.(int)
			suite.True(ok, fmt.Sprintf("failed test %d: %s:%v is not int", i, k, v))
		}
	}
}

func (suite *TypesTestSuite) TestParseFloatMap() {
	tests := []struct {
		input   naiveMap
		success bool
	}{
		{naiveMap{"a": 1.0}, true},
		{naiveMap{"a": 1.1}, true},
		{naiveMap{"a": true}, false},
		{naiveMap{"a": "123"}, false},
		{naiveMap{"a": 1.0, "b": 2.0}, true},
		{naiveMap{"a": 1.0, "b": 2.0 + 1e-15}, true},
	}

	for i, t := range tests {
		res, ok := parseFloatMap(t.input)
		suite.Equal(t.success, ok, fmt.Sprintf("failed test %d", i))

		if !ok {
			continue
		}

		suite.Equal(len(t.input), len(res))
		for k, v := range res {
			suite.Equal(t.input[k], v)
		}
	}
}

func (suite *TypesTestSuite) TestParseBoolMap() {
	tests := []struct {
		input   naiveMap
		success bool
	}{
		{naiveMap{"a": 1.0}, false},
		{naiveMap{"a": 1.1}, false},
		{naiveMap{"a": true}, true},
		{naiveMap{"a": "123"}, false},
		{naiveMap{"a": false, "b": true}, true},
	}

	for i, t := range tests {
		res, ok := parseBoolMap(t.input)
		suite.Equal(t.success, ok, fmt.Sprintf("failed test %d", i))

		if !ok {
			continue
		}

		suite.Equal(len(t.input), len(res))
		for k, v := range res {
			suite.Equal(t.input[k], v)
		}
	}
}

func (suite *TypesTestSuite) TestParseStringMap() {
	tests := []struct {
		input   naiveMap
		success bool
	}{
		{naiveMap{"a": 1.0}, false},
		{naiveMap{"a": 1.1}, false},
		{naiveMap{"a": true}, false},
		{naiveMap{"a": "123"}, true},
		{naiveMap{"a": "1", "b": "false"}, true},
	}

	for i, t := range tests {
		res, ok := parseStringMap(t.input)
		suite.Equal(t.success, ok, fmt.Sprintf("failed test %d", i))

		if !ok {
			continue
		}

		suite.Equal(len(t.input), len(res))
		for k, v := range res {
			suite.Equal(t.input[k], v)
		}
	}
}

func (suite *TypesTestSuite) TestParseMixedMap() {
	in := naiveMap{"a": "1", "b": false, "c": 1.0}

	res, ok := parseMixedMap(in)
	suite.True(ok)

	suite.Equal(len(in), len(res))
	for k, v := range res {
		suite.Equal(in[k], v)
	}
}

func (suite *TypesTestSuite) TestIsValidType() {
	tests := []struct {
		input arbitraryType
		valid bool
	}{
		{float64(1.0), true},
		{false, true},
		{"string", true},
		{struct{}{}, false},
	}

	for i, t := range tests {
		suite.Equal(t.valid, isValidType(t.input), fmt.Sprintf("failed test %d", i))
	}
}

func TestTypesTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}
