package json

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseJson(s []byte) (any, error) {
	var json any = nil
	var err error = nil

	for i := 0; i < len(s); i++ {
		c := s[i]

		if canTryParseAsString(c) {
			json, err = parseString(s, &i)
		} else if canTryParseAsInt(c) {
			json, err = parseFloat(s, &i)

			if err != nil {
				json, err = parseInt(s, &i)
			}
		} else if canTryParseAsNull(c) {
			json, err = parseNull(s, &i)
		} else if canTryParseAsFalse(c) {
			json, err = parseFalse(s, &i)
		} else if canTryParseAsTrue(c) {
			json, err = parseTrue(s, &i)
		}

		if err != nil {
			return json, err	
		}
	}

	if len(s) == 0 {
		json = ""
	}
	
	return json, nil
}

func canTryParseAsString(c byte) bool {
	return c == '"'
}

func parseString(s []byte, i *int) (string, error) {
	// skip start "
	*i += 1
	result := ""

	for *i < len(s) && s[*i] != '"' {
		result += string(s[*i])
		*i += 1
	}

	if s[*i] != '"' {
		return "", errors.New("invalid []byte: expected end \"")
	}

	// skip end "
	*i += 1

	return result, nil
}

func canTryParseAsTrue(c byte) bool {
	return c == 't'
}

func parseTrue(s []byte, i *int) (bool, error) {
	if (*i+4 <= len(s) && s[*i] == 't' && s[*i+1] == 'r' && s[*i+2] == 'u' && s[*i+3] == 'e') {
		return true, nil
	}

	return false, errors.New("invalid true")
}

func canTryParseAsFalse(c byte) bool {
	return c == 'f'
}

func parseFalse(s []byte, i *int) (bool, error) {
	if (*i+5 <= len(s) && s[*i] == 'f' && s[*i+1] == 'a' && s[*i+2] == 'l' && s[*i+3] == 's' && s[*i+4] == 'e') {
		return false, nil
	}

	return false, errors.New("invalid false")
}

func canTryParseAsNull(c byte) bool {
	return c == 'n'
}

func parseNull(s []byte, i *int) (any, error) {
	if (*i+4 <= len(s) && s[*i] == 'n' && s[*i+1] == 'u' && s[*i+2] == 'l' && s[*i+3] == 'l') {
		return nil, nil
	}

	return nil, errors.New("invalid null")
}

func canTryParseAsInt(c byte) bool {
	return c == '-' || ('0' <= c && c <= '9')
}

func parseInt(s []byte, i *int) (int, error) {
	num := ""

	for *i < len(s) && canTryParseAsInt(s[*i]) {
		num += string(s[*i])
		*i += 1
	}

	result, err := strconv.ParseInt(num, 10, 64)

	return int(result), err
}

func parseFloat(s []byte, i *int) (float64, error) {
	decimalPart := ""
	fractionPart := ""
	initialIndex := *i

	for *i < len(s) && (canTryParseAsInt(s[*i]) || s[*i] == '.') {
		if s[*i] == '.' {
			fractionPart += string(s[*i])
		} else if len(fractionPart) > 0 {
			fractionPart += string(s[*i])
		} else {
			decimalPart += string(s[*i])
		}

		*i += 1
	}

	if len(fractionPart) == 0 {
		*i = initialIndex
		return 0, errors.New("invalid float: expected a fraction part")
	}

	result, err := strconv.ParseFloat(fmt.Sprintf("%s%s", decimalPart, fractionPart), 64)

	return result, err
}