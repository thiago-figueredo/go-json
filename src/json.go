package json

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseJson(s []byte) (any, error) {
	json, _, err := parseJson(s)
	return json, err
}

func parseJson(s []byte) (any, int, error) {
	var json any = nil
	var err error = nil
	var offset int = 0
	var i int = 0

	for i < len(s) {
		if isWhiteSpace(s[i]) {
			i += 1
			continue
		}

		json, offset, err = parseJsonElement(s, i)

		if err != nil {
			return json, 0, err	
		}

		i += offset
	}

	if len(s) == 0 {
		json = ""
	}
	
	return json, offset, nil
}

func isWhiteSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}

func parseJsonElement(s []byte, i int) (any, int, error) {
	var json any
	var err error
	var offset int = 0
	var c byte = s[i]

	if canTryParseAsObject(c) {
		return parseObject(s, i)
	} else if canTryParseAsArray(c) {
		return parseArray(s, i)
	} else if canTryParseAsString(c) {
		return parseString(s, i)
	} else if canTryParseAsInt(c) {
		json, offset, err = parseFloat(s, i)

		if err != nil {
			return parseInt(s, i)
		}

		return json, offset, err
	} else if canTryParseAsNull(c) {
		return parseNull(s, i)
	} else if canTryParseAsFalse(c) {
		return parseFalse(s, i)
	} else if canTryParseAsTrue(c) {
		return parseTrue(s, i)
	} 
		
	return nil, 0, fmt.Errorf("invalid character: `%s` at index %d in `%s`", string(c), i, string(s))
}

func canTryParseAsString(c byte) bool {
	return c == '"'
}

func parseString(s []byte, i int) (string, int, error) {
	i += 1 // skip start "
	result := ""

	for i < len(s) && s[i] != '"' {
		result += string(s[i])
		i += 1
	}

	if i < len(s) && s[i] != '"' {
		return "", 0, fmt.Errorf("invalid string `%s`: expected end \"", s)
	}

	i += 1 // skip end "

	return result, i, nil
}

func canTryParseAsTrue(c byte) bool {
	return c == 't'
}

func parseTrue(s []byte, i int) (bool, int, error) {
	if (i+4 <= len(s) && s[i] == 't' && s[i+1] == 'r' && s[i+2] == 'u' && s[i+3] == 'e') {
		return true, i+4, nil
	}

	return false, 0, errors.New("invalid true")
}

func canTryParseAsFalse(c byte) bool {
	return c == 'f'
}

func parseFalse(s []byte, i int) (bool, int, error) {
	if (i+5 <= len(s) && s[i] == 'f' && s[i+1] == 'a' && s[i+2] == 'l' && s[i+3] == 's' && s[i+4] == 'e') {
		return false, i+5, nil
	}

	return false, 0, errors.New("invalid false")
}

func canTryParseAsNull(c byte) bool {
	return c == 'n'
}

func parseNull(s []byte, i int) (any, int, error) {
	if (i+4 <= len(s) && s[i] == 'n' && s[i+1] == 'u' && s[i+2] == 'l' && s[i+3] == 'l') {
		return nil, i + 4, nil
	}

	return nil, 0,errors.New("invalid null")
}

func canTryParseAsInt(c byte) bool {
	return c == '-' || ('0' <= c && c <= '9')
}

func parseInt(s []byte, i int) (int, int, error) {
	num := ""

	for i < len(s) && canTryParseAsInt(s[i]) {
		num += string(s[i])
		i += 1
	}

	result, err := strconv.ParseInt(num, 10, 64)

	return int(result), i, err
}

func parseFloat(s []byte, i int) (float64, int, error) {
	decimalPart := ""
	fractionPart := ""

	for i < len(s) && (canTryParseAsInt(s[i]) || s[i] == '.') {
		if s[i] == '.' {
			fractionPart += string(s[i])
		} else if len(fractionPart) > 0 {
			fractionPart += string(s[i])
		} else {
			decimalPart += string(s[i])
		}

		i += 1
	}

	if len(fractionPart) == 0 {
		return 0, i, errors.New("invalid float: expected a fraction part")
	}

	result, err := strconv.ParseFloat(fmt.Sprintf("%s%s", decimalPart, fractionPart), 64)

	return result, i, err
}

func canTryParseAsArray(c byte) bool {
	return c == '['
}

func parseArray(s []byte, i int) ([]any, int, error) {
	// skip [ and white spaces
	i = skipWhiteSpaces(s, i + 1)

	arr := []any{}

	for i < len(s) && s[i] != ']' {
		if s[i] == ',' || isWhiteSpace((s[i])) {
			i += 1
			continue
		}

		element, offset, err := parseJsonElement(s, i)

		if err != nil {
			return arr, offset, err
		}

		arr = append(arr, element)
		i = offset
	}

	if (i < len(s) && s[i] != ']') {
		return nil, 0, errors.New("invalid array: expected end ]")
	}

	i += 1 // skip ]

	return arr, i, nil
}

func canTryParseAsObject(c byte) bool {
	return c == '{'
}	

func skipWhiteSpaces(s []byte, i int) int {
	for i < len(s) && isWhiteSpace(s[i]) {
		i += 1
	}

	return i
}

func parseObject(s []byte, i int) (map[string]any, int, error) {
	i = skipWhiteSpaces(s, i)
	i += 1 // skip {
	result := map[string]any{}

	i = skipWhiteSpaces(s, i)

	if s[i] != '"' && s[i] != '}'{
		return result, 0, fmt.Errorf("invalid object `%s`: expected start \"", string(s[i]))
	}

	for i < len(s) && s[i] != '}' {
		if s[i] == ',' || isWhiteSpace(s[i]) {
			i += 1
			continue
		}

		key, offset, err := parseString(s, i)

		if err != nil {
			return result, offset, err
		}

		i = skipWhiteSpaces(s, offset)

		if s[i] != ':' {
			return result, 0, fmt.Errorf("invalid character `%s`: expected `:`", string(s[i]))
		}

		// skip : and any white spaces after it
		i = skipWhiteSpaces(s, i+1)

		value, offset, err := parseJsonElement(s, i)

		if err != nil {
			return result, offset, err
		}

		i = offset
		result[key] = value
	}

	if i < len(s) && s[i] != '}' {
		return nil, 0, errors.New("invalid object: expected end }")
	}

	i += 1 // skip }

	return result, i, nil
}
