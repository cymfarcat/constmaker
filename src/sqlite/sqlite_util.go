// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package sqlite

import (
	"regexp"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func typeStr(typoStr string, name string) string {
	switch typoStr {
	case src.CPP:
		return "const char* " + name

	case src.Dart:
		return "const String " + name

	case src.Go:
		return name + " string"

	case src.Rust:
		return "pub const " + name + ": &str"
	}
	return ""
}

func quoteStr(typoStr string, str string) string {
	quoteChar := `"`

	if typoStr == src.Go {
		if strings.Count(str, "\n") >= 1 {
			quoteChar = "`"
		}
	}

	count := 1
	if typoStr == src.Dart {
		// if str no \n
		if strings.Count(str, "\n") >= 1 {
			count = 3
		}
	}

	ret := strings.Repeat(quoteChar, count) + strings.Replace(str, `"`, `\"`, -1) + strings.Repeat(quoteChar, count)
	if typoStr != src.Go {
		ret += ";"
	}
	return ret
}

func getSQLiteName(option *src.Options, objType uint8, prefix string, name string) string {
	if option.UsePrefix {
		switch objType {
		case TypeTable:
			name = option.TablePrefix + src.UnderScore + name

		case TypeIndex:
			name = option.IndexPrefix + src.UnderScore + name

		case TypeView:
			name = option.ViewPrefix + src.UnderScore + name

		case TypeTrigger:
			name = option.TriggerPrefix + src.UnderScore + name

		case TypeVirtualTable:
			name = option.VirtualTablePrefix + src.UnderScore + name

		case TypeField:
			name = prefix + src.UnderScore + name
		}
	}

	return option.GetIdentName(name)
}

var stmtShortRegex = regexp.MustCompile(`#\s*\[\s*(?i:stmt)\s*\(\s*(?i:short)\s*=\s*"([^"]+)"\s*\)\s*\]`)
var stmtFullRegex = regexp.MustCompile(`#\s*\[\s*(?i:stmt)\s*\(\s*(?i:full)\s*=\s*"([^"]+)"\s*\)\s*\]`)

// #[stmt(short = "Answer")]
// #[stmt(full = "alter_user_bio")]
func parseStmtName(input string) (bool, string) {
	input = strings.Trim(input, " -\t\n\r")
	if len(input) == 0 {
		return false, ""
	}
	if stmtShortRegex.MatchString(input) {
		var ret = stmtShortRegex.FindStringSubmatch(input)
		return false, ret[1]
	}
	if stmtFullRegex.MatchString(input) {
		var ret = stmtFullRegex.FindStringSubmatch(input)
		return true, ret[1]
	}
	return false, ""
}

func getStmtName(option *src.Options, stmt string, name string, stmtFull bool, stmtName string) string {
	name = src.UnwrapIdent(name)

	if len(stmtName) == 0 {
		return option.GetIdentName(stmt + src.UnderScore + name + src.UnderScore + src.GlobalIdString())
	}
	if stmtFull {
		return stmtName
	}
	return option.GetIdentName(stmt + src.UnderScore + name + src.UnderScore + stmtName)
}
