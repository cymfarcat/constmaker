// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"math"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

// ////////////////////////////////////////////////////////////////////////////
func getIdentName(option *src.Options, ident string, object *ConstObject) string {
	if option.UpperIdent || object.UpperIdent {
		ident = strings.ToUpper(ident)
	} else if option.LowerIdent || object.LowerIdent {
		ident = strings.ToLower(ident)
	}

	if option.UpperIdentCamel || object.UpperIdentCamel {
		return src.UpperCamelCase(ident)
	} else if option.LowerIdentCamel || object.UpperIdentCamel {
		return src.LowerCamelCase(ident)
	}

	return ident
}

func genIdentName(option *src.Options, prefix string, object *ConstObject) string {
	// option.prefix + prefix
	ident := option.JoinPrefix(option.Prefix, prefix)

	// ident + object.prefix
	ident = option.JoinPrefix(ident, object.Prefix)

	// ident + object.Ident
	ident = option.JoinPrefix(ident, object.Ident)

	// ident + object.suffix
	ident = option.JoinPrefix(ident, object.Suffix)

	// ident + option.suffix
	ident = option.JoinPrefix(ident, option.Suffix)

	return getIdentName(option, ident, object)
}

func correctIdent(object *ConstObject) {
	// keep origin value
	object.value = object.Ident

	if strings.Contains(object.Ident, src.HyphenChar) {
		// replace with HyphenChar="-"
		result := strings.ReplaceAll(object.Ident, src.HyphenChar, src.UnderScore)
		object.Ident = result
	}
}

func intType(intVal int64) string {
	switch {
	case intVal >= math.MinInt8 && intVal <= math.MaxInt8:
		return src.Typo_I8
	case intVal >= math.MinInt16 && intVal <= math.MaxInt16:
		return src.Typo_I16
	case intVal >= math.MinInt32 && intVal <= math.MaxInt32:
		return src.Typo_I32
	default:
		return src.Typo_I64
	}
}

func uintType(uintVal uint64) string {
	switch {
	case uintVal <= math.MaxUint8:
		return src.Typo_U8
	case uintVal <= math.MaxUint16:
		return src.Typo_U16
	case uintVal <= math.MaxUint32:
		return src.Typo_U32
	default:
		return src.Typo_U64
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
		return src.Typo_U8
	}
	if uintVal <= 16 {
		return src.Typo_U16
	}
	if uintVal <= 32 {
		return src.Typo_U32
	}
	if uintVal <= 64 {
		return src.Typo_U64
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
	if typo == src.Typo_F32 {
		if len(value) > 0 && (value[len(value)-1] != 'f' || value[len(value)-1] != 'F') {
			value += "f"
		}
	}
	return value
}
