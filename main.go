// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cymfarcat/constmaker/src"
)

func main() {
	fmt.Fprintf(os.Stdout, "ConstMaker %s\n", src.VersionString())

	// for test (F5 to run)
	if len(os.Args) <= 1 {
		exePath, _ := os.Executable()
		exeDir := filepath.Dir(exePath)

		os.Args = append(os.Args, "-i"+exeDir+"/tests/test1.cmt", "-o"+exeDir+"/tests",
			"--verbose",
			// "--root-name=",
			// "--root-namespace",
			"--nested-define=false",
			// "--macro-define",
			// "--enum-to-prefix",
			// "--namespace-to-prefix",
			"--upper-ident",
			// "--lower-ident",
			"--upper-ident-camel",
			// "--lower-ident-camel",
			"--cpp",
			"--c#",
			"--dart",
			"--go", "--go-package=tests",
			"--java", "--java-package=tests",
			"--javascript",
			"--json",
			"--kotlin",
			"--markdown",
			"--objc",
			"--pascal",
			"--perl",
			"--php",
			"--python",
			"--qml",
			"--ruby",
			"--rust",
			"--swift",
			"--text",
			"--typescript",
			"--xml")
	}

	src.RunParser()
}
