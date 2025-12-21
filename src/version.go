// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"fmt"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////
type Version struct {
	Parts [3]int   // Major.Minor.Patch
	Meta  []string // stable/build-time
}

var versionList = []Version{
	{Parts: [3]int{1, 6, 0}, Meta: []string{"dart strMapIdent/identMapStr", "build20251231"}},
	{Parts: [3]int{1, 5, 1}, Meta: []string{"bug fixed", "build20251221"}},
	{Parts: [3]int{1, 5, 0}, Meta: []string{"sqlite consts", "build20251031"}},
	{Parts: [3]int{1, 0, 0}, Meta: []string{"initial release", "build20250531"}},
}

func VersionString() string {
	if len(versionList) > 0 {
		version := versionList[0]

		str := "version: " + fmt.Sprintf("%d.%d.%d", version.Parts[0], version.Parts[1], version.Parts[2])

		return str + ", " + strings.Join(version.Meta, ".")
	}
	return "UNKOWN VERSION"
}
