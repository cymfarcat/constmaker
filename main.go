// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
	consts "github.com/cymfarcat/constmaker/src/const"
	sqlite "github.com/cymfarcat/constmaker/src/sqlite"
	"github.com/spf13/pflag"
)

func main() {
	fmt.Fprintf(os.Stdout, "ConstMaker %s\n", src.VersionString())

	// for test (F5 to run)
	if len(os.Args) <= 1 {
		exePath, _ := os.Executable()
		exeDir := filepath.Dir(exePath)

		// for test sqlite, comment test constmaker
		test := "" //sqlite.SQLITE_ARG

		switch test {
		case sqlite.SQLITE_ARG:
			os.Args = append(os.Args, "-i"+exeDir+"/tests/sqlite/test1.sql", "-o"+exeDir+"/tests/sqlite",
				sqlite.SQLITE_ARG,
				"--verbose",
				"--lower-ident",
				"--field-wl=  row-id;  id  ",
				"--cpp",
				"--dart",
				"--go", "--go-package=sqlite",
				"--rust",
			)

		default:
			os.Args = append(os.Args, "-i"+exeDir+"/tests/const/test1.cmt", "-o"+exeDir+"/tests/const",
				"--verbose",
				// "--root-name=",
				// "--root-namespace",
				"--nested-define=false",
				// "--macro-define",
				// "--enum-to-prefix",
				// "--namespace-to-prefix",
				// "--upper-ident",
				// "--lower-ident",
				// "--upper-ident-camel",
				"--lower-ident-camel",
				"--cpp",
				"--c#",
				"--dart",
				"--go", "--go-package=consts",
				"--java", "--java-package=consts",
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
	}

	// first check args do parser
	for _, arg := range os.Args {
		switch arg {
		case sqlite.SQLITE_ARG:
			RunParser()
			return
		}
	}

	// default is const parser
	RunParser()
}

func RunParser() {
	option := new(src.Options)
	option.InitUsage()

	// parse command-line flags
	pflag.Parse()

	option.InitFieldWhiteList()

	// check file isEmpty
	if len(option.InputFile) == 0 {
		pflag.Usage()
		os.Exit(-1)
	}

	// check file existed
	_, err := os.Stat(option.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file: %s, error: %s\n", option.InputFile, err)
		os.Exit(-1)
	}

	// convert input path
	absPath, err := filepath.Abs(option.InputFile)
	if err == nil {
		option.InputFile = absPath
	}

	if option.Verbose {
		fmt.Fprintf(os.Stdout, "intput file: %s\n", option.InputFile)
	}

	// get fileName
	if len(option.FileName) == 0 {
		fileName := filepath.Base(option.InputFile)
		option.FileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}

	// check output path
	if len(option.OutputPath) == 0 {
		exePath, _ := os.Executable()
		option.OutputPath = filepath.Dir(exePath)
	} else {
		// convert output path
		absPath, err := filepath.Abs(option.OutputPath)
		if err == nil {
			option.OutputPath = absPath
		}

		// create output path
		err = os.MkdirAll(option.OutputPath, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create dir: %s, error: %s\n", option.OutputPath, err)
			os.Exit(-1)
		}
	}

	if option.Verbose {
		//pflag.Visit(func(f *pflag.Flag) {
		//	fmt.Fprintf(os.Stdout, "arg=%s: value=%v (default=%v)\n", f.Name, f.Value, f.DefValue)
		//})

		fmt.Fprintf(os.Stdout, "output path: %s, fileName: %s\n", option.OutputPath, option.FileName)
	}

	// start parser
	if option.SQLite {
		parser := sqlite.NewSQLiteParser(option)
		if parser.ParseFile() == 0 {
			//parser.GenFile()
		}
	} else {
		parser := consts.NewConstParser(option)
		if parser.ParseFile() == 0 {
			parser.GenFile()
		}

	}
}
