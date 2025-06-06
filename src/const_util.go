// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// ////////////////////////////////////////////////////////////////////////////
type Version struct {
	Parts [3]int   // Major.Minor.Patch
	Meta  []string // stable/build-time
}

var versionList = []Version{
	{Parts: [3]int{1, 0, 0}, Meta: []string{"initial release", "build20250531"}},
}

func VersionString() string {
	if len(versionList) > 0 {
		version := versionList[len(versionList)-1]

		str := "version: " + fmt.Sprintf("%d.%d.%d", version.Parts[0], version.Parts[1], version.Parts[2])

		return str + ", " + strings.Join(version.Meta, ".")
	}
	return "UNKOWN VERSION"
}

// ////////////////////////////////////////////////////////////////////////////
func unquoteStr(str string) string {
	result, err := strconv.Unquote(str)
	if err != nil {
		result = str
	}
	return result
}

func correctIdent(object *Object) {
	// keep origin value
	object.value = object.Ident

	if strings.Contains(object.Ident, HyphenChar) {
		// replace with HyphenChar="-"
		result := strings.ReplaceAll(object.Ident, HyphenChar, UnderScore)
		object.Ident = result
	}
}

func upperCamelCase(ident string) string {
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

func lowerCamelCase(ident string) string {
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

func intType(intVal int64) string {
	switch {
	case intVal >= math.MinInt8 && intVal <= math.MaxInt8:
		return Typo_I8
	case intVal >= math.MinInt16 && intVal <= math.MaxInt16:
		return Typo_I16
	case intVal >= math.MinInt32 && intVal <= math.MaxInt32:
		return Typo_I32
	default:
		return Typo_I64
	}
}

func uintType(uintVal uint64) string {
	switch {
	case uintVal <= math.MaxUint8:
		return Typo_U8
	case uintVal <= math.MaxUint16:
		return Typo_U16
	case uintVal <= math.MaxUint32:
		return Typo_U32
	default:
		return Typo_U64
	}
}

func bitFlagCount(uintVal int) int {
	switch {
	case uintVal <= 8:
		return 2
	case uintVal <= 16:
		return 4
	case uintVal <= 32:
		return 8
	default:
		return 16
	}
}

func bitFlagType(uintVal uint64, bitflag bool) string {
	if bitflag {
		uintVal += 2 //None, All
	}
	if uintVal <= 8 {
		return Typo_U8
	}
	if uintVal <= 16 {
		return Typo_U16
	}
	if uintVal <= 32 {
		return Typo_U32
	}
	if uintVal <= 64 {
		return Typo_U64
	}
	return uintType(uintVal)
}

func adjustLen(uintVal int, bitflag bool) int {
	if bitflag {
		uintVal += 2 //None, All
	}
	return uintVal
}

func adjustFloat(value string, typo string) string {
	if typo == Typo_F32 {
		if len(value) > 0 && (value[len(value)-1] != 'f' || value[len(value)-1] != 'F') {
			value += "f"
		}
	}
	return value
}

// calc like (1<<2)
func calcShift(expr string) string {
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

func cleanCommentDoc(comment string, sign string) string {
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

func convertCommentDoc(comment string, sign string) string {
	comment = strings.TrimPrefix(comment, "/*")
	comment = strings.TrimSuffix(comment, "*/")

	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = sign + " " + strings.TrimLeft(line, " *\t\n\r")
	}

	return strings.Join(lines, "\n")
}

func convertCommentTripe(comment string, sign string) string {
	comment = strings.Trim(comment, " \t\n\r")
	if strings.HasPrefix(comment, "///") {
		return sign + " " + strings.TrimLeft(comment, " /\t\n\r")
	}

	comment = strings.TrimPrefix(comment, "//")
	return sign + " " + strings.Trim(comment, " \t\n\r")
}

func convertCommentDocToPascal(comment string) string {
	content := strings.TrimPrefix(comment, "/*")
	content = strings.TrimSuffix(content, "*/")

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = " " + strings.TrimLeft(line, " *\t\n\r")
	}

	return "(*" + strings.Join(lines, "\n") + "*)"
}

func convertCommentTripeToPascal(comment string) string {
	comment = strings.Trim(comment, " \t\n\r")
	if strings.HasPrefix(comment, "///") {
		return "(* " + strings.TrimLeft(comment, " /\t\n\r") + " *)"
	}

	comment = strings.TrimPrefix(comment, "//")
	return "(* " + strings.Trim(comment, " \t\n\r") + " *)"
}
