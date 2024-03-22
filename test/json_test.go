package test

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	json "github.com/thiago-figueredo/json/src"
)

func TestParsePositiveInt(t *testing.T) {
	n := rand.Intn(math.MaxInt)
	result, err := json.ParseJson([]byte(fmt.Sprintf("%d", n)))

	assert.Nil(t, err)
	assert.Equal(t, n, result)
}

func TestParseNegativeInt(t *testing.T) {
	n := rand.Intn(math.MaxInt)
	result, err := json.ParseJson([]byte(fmt.Sprintf("-%d", n)))

	assert.Nil(t, err)
	assert.Equal(t, -n, result)
}

func TestParsePositiveFloat(t *testing.T) {
	n := rand.Float64()
	result, err := json.ParseJson([]byte(fmt.Sprintf("%f", n)))

	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseNegativeFloat(t *testing.T) {
	n := rand.Float64()
	result, err := json.ParseJson([]byte(fmt.Sprintf("%f", -n)))

	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, -expected, result)
}

func TestParseNull(t *testing.T) {
	result, err := json.ParseJson([]byte("null"))

	assert.Nil(t, err)
	assert.Equal(t, nil, result)
}

func TestParseFalse(t *testing.T) {
	result, err := json.ParseJson([]byte("false"))

	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func TestParseTrue(t *testing.T) {
	result, err := json.ParseJson([]byte("true"))

	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func TestParseEmptyString(t *testing.T) {
	result, err := json.ParseJson([]byte(""))

	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

func TestParseNotEmptyString(t *testing.T) {
	s := `"Hello, World!"`
	result, err := json.ParseJson([]byte(s))
	
	assert.Nil(t, err)

	str, err := strconv.Unquote(s)

	assert.Nil(t, err)
	assert.Equal(t, str, result)
}

func TestParseEmptyArray(t *testing.T) {
	result, err := json.ParseJson([]byte("[]"))

	assert.Nil(t, err)
	assert.Equal(t, []any{}, result)
}

func TestParseArrayWithOneNull(t *testing.T) {
	str := "[null]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{nil}, result)
}

func TestParseArrayWithOneFalse(t *testing.T) {
	str := "[false]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{false}, result)
}

func TestParseArrayWithOneTrue(t *testing.T) {
	str := "[true]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{true}, result)
}

func TestParseArrayWithOnePositiveInt(t *testing.T) {
	str := "["

	n := rand.Intn(math.MaxInt)
	str += fmt.Sprintf("%d", n)

	str += "]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{n}, result)
}

func TestParseArrayWithOneNegativeInt(t *testing.T) {
	str := "["

	n := -rand.Intn(math.MaxInt)
	str += fmt.Sprintf("%d", n)

	str += "]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{n}, result)
}

func TestParseArrayWithOnePositiveFloat(t *testing.T) {
	str := "["

	n := rand.Float64()
	str += fmt.Sprintf("%f", n)

	str += "]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, []any{expected}, result)
}

func TestParseArrayWithOneNegativeFloat(t *testing.T) {
	str := "["

	n := -rand.Float64()
	str += fmt.Sprintf("%f", n)

	str += "]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, []any{expected}, result)
}

func TestParseArrayWithOneEmptyString(t *testing.T) {
	str := `[""]`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{""}, result)
}

func TestParseArrayWithOneNonEmptyString(t *testing.T) {
	str := `["hello world"]`
	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{"hello world"}, result)
}

func TestParseArrayWithTwoNulls(t *testing.T) {
	str := "[null, null]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{nil, nil}, result)
}

func TestParseArrayWithTwoBooleans(t *testing.T) {
	str := "[true, false]"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{true, false}, result)
}

func TestParseArrayWithTwoInts(t *testing.T) {
	str := "["
	n1 := rand.Intn(math.MaxInt)
	n2 := rand.Intn(math.MaxInt)
	str += fmt.Sprintf("%d, %d]", n1, n2)

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{n1, n2}, result)
}

func TestParseArrayWithTwoFloats(t *testing.T) {
	str := "["
	n1 := rand.Float64()
	n2 := rand.Float64()
	str += fmt.Sprintf("%f, %f]", n1, n2)

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)

	expected1, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n1), 64)

	assert.Nil(t, err)

	expected2, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n2), 64)

	assert.Nil(t, err)

	assert.Equal(t, []any{expected1, expected2}, result)
}

func TestParseArrayWithTwoStrings(t *testing.T) {
	str := `["hello", "world"]`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{"hello", "world"}, result)
}

func TestParseArrayWithTwoMixedValues(t *testing.T) {
	str := `[null, true, 1, 1.0, "hello"]`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{nil, true, 1, 1.0, "hello"}, result)
}

func TestParseArrayOfArrays(t *testing.T) {
	str := `[[1, false, "lorem"], [true, null, []], [7.0, -8, -923]]`
	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, []any{ []any{1, false, "lorem"}, []any{true, nil, []any{}}, []any{7.0, -8, -923} }, result)
}

func TestParseEmptyObject(t *testing.T) {
	str := "{}"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}

func TestParseEmptyObjectWithNewLines(t *testing.T) {
	str := "{\n\r\n}"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}

func TestParseEmptyObjectWithNewLinesAndWhiteSpaces(t *testing.T) {
	str := "{   \n  \r            \n}"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}


func TestParseMultiLineEmptyObject(t *testing.T) {
	str := `
	
	  {}

	`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}

func TestParseEmptyObjectWithWhiteSpaces(t *testing.T) {
	str := "{         }"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}

func TestParseEmptyObjectWithTabs(t *testing.T) {
	str := "{\t\t}"

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{}, result)
}

func TestParseObjectWithOneKeyAndNullValue(t *testing.T) {
	str := `{"key": null}`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"key": nil}, result)
}

func TestParseObjectWithOneKeyAndFalseValue(t *testing.T) {
	str := `{"key": false}`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"key": false}, result)
}

func TestParseObjectWithOneKeyAndTrueValue(t *testing.T) {
	str := `{"key": true}`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"key": true}, result)
}

func TestParseObjectWithOneKeyAndPositiveIntValue(t *testing.T) {
	str := `{   "key": `
	n := rand.Intn(math.MaxInt)
	str += fmt.Sprintf("%d}", n)

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"key": n}, result)
}

func TestParseObjectWithOneKeyAndNegativeIntValue(t *testing.T) {
	str := `{"bar": `
	n := -rand.Intn(math.MaxInt)
	str += fmt.Sprintf("\r\n  %d\t}", n)

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"bar": n}, result)
}

func TestParseObjectWithOneKeyAndPositiveFloatValue(t *testing.T) {
	str := `{"foo": `

	n := rand.Float64()

	str += fmt.Sprintf("%f}", n)

	result, err := json.ParseJson([]byte(str))
	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"foo": expected}, result)
}

func TestParseObjectWithOneKeyAndNegativeFloatValue(t *testing.T) {
	str := `{"foo": `

	n := -rand.Float64()

	str += fmt.Sprintf("%f}", n)

	result, err := json.ParseJson([]byte(str))
	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"foo": expected}, result)
}

func TestParseObjectWithOneKeyAndEmptyStringValue(t *testing.T) {
	str := `{"foo"   :	""}`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"foo": ""}, result)
}

func TestParseObjectWithOneKeyAndNonEmptyStringValue(t *testing.T) {
	str := `{
		"foo": "lorem ipsum"
	}`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"foo": "lorem ipsum"}, result)
}

func TestParseObjectWithMultipleKeyAndValues(t *testing.T) {
	str := `{
		"foo": "lorem ipsum",
		"bar": 123,
		"baz": 123.456,
		"qux": true,
		"quux": false,
		"corge": null,
		"grault": [],
		"garply": {},
		"waldo": [1, 2, 3],
		"fred": [
			4.01,   
			5.1203, 
			6.7819
		],
		"plugh": ["foo", "bar", "baz"],
		"nested": {
			"foo": "lorem ipsum",
			"foo2": {
				"foo3": [
					"lorem ipsum",
					{
						"foo4": "lorem ipsum"
						"foo5": [
							{
								"foo6": "lorem ipsum"
								"foo7": []
							}
						]
					}
				]
			}
		}
		"xyzzy": [true, false, null]
	}
	`

	result, err := json.ParseJson([]byte(str))

	assert.Nil(t, err)
	assert.Equal(t, map[string]any{ 
		"foo": "lorem ipsum",
		"bar": 123,
		"baz": 123.456,
		"qux": true,
		"quux": false,
		"corge": nil,
		"grault": []any{},
		"garply": map[string]any{},
		"waldo": []any{1, 2, 3},
		"fred": []any{4.01, 5.1203, 6.7819},
		"plugh": []any{"foo", "bar", "baz"},
		"xyzzy": []any{true, false, nil},
		"nested": map[string]interface{}{
			"foo": "lorem ipsum",
			"foo2": map[string]interface{}{
				"foo3": []interface{}{
					"lorem ipsum",
					map[string]interface{}{
						"foo4": "lorem ipsum",
						"foo5": []interface{}{
							map[string]interface{}{
								"foo6": "lorem ipsum",
								"foo7": []interface{}{},
							},
						},
					},
				},
			},
		},
	}, result)
}