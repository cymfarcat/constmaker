// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"fmt"
	"hash/crc32"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// ////////////////////////////////////////////////////////////////////////////
func HashStr(str string) string {
	hash := crc32.ChecksumIEEE([]byte(str))
	return fmt.Sprintf("%08x", hash)
}

// match like: [user_id]、"user_id"、'user_id'
var identRegexp = regexp.MustCompile(`[\[\"'](.+?)[\"'\]]`)

func UnwrapIdent(input string) string {
	if identRegexp.MatchString(input) {
		var ret = identRegexp.FindStringSubmatch(input)
		return ret[1]
	}
	return input
}

func UnquoteStr(str string) string {
	result, err := strconv.Unquote(str)
	if err != nil {
		result = str
	}
	return result
}

func UpperCamelCase(ident string) string {
	// split by UnderScore="_"
	parts := strings.Split(strings.ToLower(ident), UnderScore)

	var result string
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		// upper part of first char
		r := []rune(part)
		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}
		result += string(r)
	}

	return result
}

func LowerCamelCase(ident string) string {
	// split by UnderScore="_"
	parts := strings.Split(strings.ToLower(ident), UnderScore)

	var result string
	for idx, part := range parts {
		if len(part) == 0 {
			continue
		}

		if idx == 0 {
			result += part
			continue
		}

		// upper part of first char
		r := []rune(part)
		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}
		result += string(r)
	}

	return result
}

// calc like (1<<2)
func CalcShift(expr string) string {
	expr = strings.ReplaceAll(expr, "(", "")
	expr = strings.ReplaceAll(expr, ")", "")
	expr = strings.ReplaceAll(expr, " ", "")

	parts := strings.Split(expr, "<<")
	if len(parts) != 2 {
		return expr
	}

	base, err := strconv.Atoi(parts[0])
	if err != nil {
		return expr
	}
	shift, err := strconv.Atoi(parts[1])
	if err != nil {
		return expr
	}

	return "0x" + strconv.FormatInt(int64(base<<shift), 16)
}

func CleanCommentDoc(comment string, sign string) string {
	comment = strings.Trim(comment, " \t\n\r")
	comment = strings.TrimPrefix(comment, "/*")
	comment = strings.TrimSuffix(comment, "*/")

	ret := ""
	lines := strings.Split(comment, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " *\t\n\r")
		if len(line) > 0 {
			if len(ret) > 0 {
				ret = sign + " " + line
			} else {
				ret = line
			}
		}
	}
	return ret
}

func ConvertCommentDoc(comment string, sign string) string {
	comment = strings.TrimPrefix(comment, "/*")
	comment = strings.TrimSuffix(comment, "*/")

	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = sign + " " + strings.TrimLeft(line, " *\t\n\r")
	}

	return strings.Join(lines, "\n")
}

func ConvertCommentTriple(comment string, sign string) string {
	comment = strings.Trim(comment, " \t\n\r")
	if strings.HasPrefix(comment, "///") {
		return sign + " " + strings.TrimLeft(comment, " /\t\n\r")
	}

	comment = strings.TrimPrefix(comment, "//")
	return sign + " " + strings.Trim(comment, " \t\n\r")
}

func ConvertCommentTripleSQL(comment string, sign string) string {
	comment = strings.Trim(comment, " \t\n\r")
	if strings.HasPrefix(comment, "---") {
		return sign + " " + strings.TrimLeft(comment, " -\t\n\r")
	}

	comment = strings.TrimPrefix(comment, "--")
	return sign + " " + strings.Trim(comment, " \t\n\r")
}

func ConvertCommentDocToPascal(comment string) string {
	content := strings.TrimPrefix(comment, "/*")
	content = strings.TrimSuffix(content, "*/")

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = " " + strings.TrimLeft(line, " *\t\n\r")
	}

	return "(*" + strings.Join(lines, "\n") + "*)"
}

func ConvertCommentTripeToPascal(comment string) string {
	comment = strings.Trim(comment, " \t\n\r")
	if strings.HasPrefix(comment, "///") {
		return "(* " + strings.TrimLeft(comment, " /\t\n\r") + " *)"
	}

	comment = strings.TrimPrefix(comment, "//")
	return "(* " + strings.Trim(comment, " \t\n\r") + " *)"
}
