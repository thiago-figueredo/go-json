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
	assert.Equal(t, float64(expected), result)
}

func TestParseNegativeFloat(t *testing.T) {
	n := rand.Float64()
	result, err := json.ParseJson([]byte(fmt.Sprintf("%f", -n)))

	assert.Nil(t, err)

	expected, err := strconv.ParseFloat(fmt.Sprintf("%.6f", n), 64)

	assert.Nil(t, err)
	assert.Equal(t, -float64(expected), result)
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

func TestEmptyString(t *testing.T) {
	result, err := json.ParseJson([]byte(""))

	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

func TestNotEmptyString(t *testing.T) {
	str := `"Hello, World!"`
	result, err := json.ParseJson([]byte(str))
	
	assert.Nil(t, err)

	s, err := strconv.Unquote(str)

	assert.Nil(t, err)
	assert.Equal(t, s, result)
}