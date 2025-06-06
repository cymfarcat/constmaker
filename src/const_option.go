// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

type ConstOption struct {
	inputFile  string
	outputPath string
	fileName   string
	fileExt    string
	checkValue bool
	verbose    bool

	prefix            string
	prefixes          []string
	suffix            string
	tab               int
	level             int
	rootName          string
	rootNamespace     bool
	nestedDefine      bool
	macroDefine       bool
	freeze            bool
	upperIdent        bool
	lowerIdent        bool
	upperIdentCamel   bool
	lowerIdentCamel   bool
	enumToPrefix      bool
	namespaceToPrefix bool
	bitFlagNoneAll    bool
	bitFlagNone       string
	bitFlagAll        string

	// all support language
	genAll bool

	// c/c++
	genCpp   bool
	stdCpp11 bool

	// c#
	genCShape bool

	// dart
	genDart bool

	// go
	genGo     bool
	goPackage string

	// java
	genJava     bool
	javaPackage string

	// javascript
	genJavaScript bool

	// json
	genJson bool

	// kotlin
	genKotlin bool

	// markdown
	genMarkdown bool

	// ObjC
	genObjC bool

	// pascal
	genPascal bool

	// perl
	genPerl bool

	// php
	genPhp bool

	// python
	genPython bool

	// qml
	genQml       bool
	qmlSingleton bool

	// ruby
	genRuby bool

	// rust
	genRust bool

	// swift
	genSwift bool

	// text
	genText bool

	// typescript
	genTypeScript bool

	// xml
	genXml bool
}

func (obj *ConstOption) initUsage() {
	pflag.StringVarP(&obj.inputFile, "input", "i", "", "input file to const define.")
	pflag.StringVarP(&obj.outputPath, "output", "o", "", "output path to generated files.")
	pflag.StringVarP(&obj.fileName, "file-name", "", "", "output file name.")
	pflag.StringVarP(&obj.fileExt, "file-ext", "", "", "output file name extension.")
	pflag.BoolVarP(&obj.checkValue, "check-value", "", false, "check for duplicate constant values within the current scope.")
	pflag.BoolVarP(&obj.verbose, "verbose", "v", false, "print verbose info.")

	pflag.StringVarP(&obj.prefix, Option_Prefix, "", "", "prefix for every const ident name.")
	pflag.StringVarP(&obj.suffix, Option_Suffix, "", "", "suffix for every const ident name.")
	pflag.StringVarP(&obj.rootName, "root-name", "", "", "root name, if empty substitute with file name.")
	pflag.BoolVarP(&obj.rootNamespace, "root-namespace", "", false, "forece root node into namspace.")
	pflag.BoolVarP(&obj.rootNamespace, "root-ns", "", false, "forece root node into namspace.")
	pflag.BoolVarP(&obj.nestedDefine, "nested-define", "", true, "force namespace use nested define like tree, or else use flat define like list.")
	pflag.BoolVarP(&obj.macroDefine, Option_MacroDefine, "", false, "force use macro define to declare constants.")
	pflag.BoolVarP(&obj.freeze, Option_Freeze, "", false, "force use readonly to declare constants.")
	pflag.BoolVarP(&obj.upperIdent, Option_UpperIdent, "", false, "force ident to upper.")
	pflag.BoolVarP(&obj.lowerIdent, Option_LowerIdent, "", false, "force ident to lower.")
	pflag.BoolVarP(&obj.upperIdentCamel, Option_UpperIdentCamel, "", false, "force ident to upper camel case.")
	pflag.BoolVarP(&obj.lowerIdentCamel, Option_LowerIdentCamel, "", false, "force ident to lower camel case.")
	pflag.IntVarP(&obj.tab, "tab", "t", 4, "tab width.")
	pflag.BoolVarP(&obj.enumToPrefix, Option_EnumToPrefix, "", false, "use enum name as prefix.")
	pflag.BoolVarP(&obj.namespaceToPrefix, Option_NamespaceToPrefix, "", false, "use namespace as prefix.")
	pflag.BoolVarP(&obj.namespaceToPrefix, Option_NSToPrefix, "", false, "use namespace as prefix.")

	pflag.BoolVarP(&obj.bitFlagNoneAll, "bitflag-none-all", "", true, "bitflag generate None and All option.")
	pflag.StringVarP(&obj.bitFlagNone, "bitflag-none", "", "NONE", "bitflag generate None string.")
	pflag.StringVarP(&obj.bitFlagAll, "bitflag-all", "", "ALL", "bitflag generate All string.")

	// all
	pflag.BoolVarP(&obj.genAll, "all", "a", false, "generate all supported language files.")

	// c/c++
	pflag.BoolVarP(&obj.genCpp, "cpp", "c", false, "generate C/CPP head files.")
	pflag.BoolVarP(&obj.stdCpp11, "std-cpp11", "", true, "std c++11.")

	// c#
	pflag.BoolVarP(&obj.genCShape, "c#", "", false, "generate C# files.")

	// dart
	pflag.BoolVarP(&obj.genDart, "dart", "d", false, "generate Dart files.")

	// go
	pflag.BoolVarP(&obj.genGo, "go", "g", false, "generate Go files.")
	pflag.StringVarP(&obj.goPackage, "go-package", "", "", "generate Go package.")

	// java
	pflag.BoolVarP(&obj.genJava, "java", "j", false, "generate Java files.")
	pflag.StringVarP(&obj.javaPackage, "java-package", "", "", "generate Java package.")

	// javascript
	pflag.BoolVarP(&obj.genJavaScript, "javascript", "", false, "generate JavaScript files.")

	// json
	pflag.BoolVarP(&obj.genJson, "json", "", false, "generate Json files.")

	// kotlin
	pflag.BoolVarP(&obj.genKotlin, "kotlin", "", false, "generate Kotlin files.")

	// markdown
	pflag.BoolVarP(&obj.genMarkdown, "markdown", "", false, "generate Markdown files.")

	// ObjC
	pflag.BoolVarP(&obj.genObjC, "objc", "", false, "generate Objective-C files.")

	// pascal
	pflag.BoolVarP(&obj.genPascal, "pascal", "", false, "generate Pascal files.")

	// perl
	pflag.BoolVarP(&obj.genPerl, "perl", "", false, "generate Perl files.")

	// php
	pflag.BoolVarP(&obj.genPhp, "php", "", false, "generate Php files.")

	// python
	pflag.BoolVarP(&obj.genPython, "python", "p", false, "generate Python files.")

	// qml
	pflag.BoolVarP(&obj.genQml, "qml", "", false, "generate QML files.")
	pflag.BoolVarP(&obj.qmlSingleton, "qml-singleton", "", false, "generate QML singleton.")

	// ruby
	pflag.BoolVarP(&obj.genRuby, "ruby", "", false, "generate Ruby files.")

	// rust
	pflag.BoolVarP(&obj.genRust, "rust", "r", false, "generate Rust files.")

	// swift
	pflag.BoolVarP(&obj.genSwift, "swift", "s", false, "generate Swift files.")

	// text
	pflag.BoolVarP(&obj.genText, "text", "", false, "generate Text files.")

	// typescript
	pflag.BoolVarP(&obj.genTypeScript, "typescript", "", false, "generate TypeScript files.")

	// xml
	pflag.BoolVarP(&obj.genXml, "xml", "", false, "generate XML files.")
}

func (obj *ConstOption) getRootName() string {
	if len(obj.rootName) > 0 {
		return obj.rootName
	}
	return obj.fileName
}

func (obj *ConstOption) getIdentName(ident string) string {
	if obj.upperIdent {
		ident = strings.ToUpper(ident)
	} else if obj.lowerIdent {
		ident = strings.ToLower(ident)
	}

	if obj.upperIdentCamel {
		return upperCamelCase(ident)
	} else if obj.lowerIdentCamel {
		return lowerCamelCase(ident)
	}

	return ident
}

func (obj *ConstOption) getIdentNameObj(ident string, object *Object) string {
	if obj.upperIdent || object.upperIdent {
		ident = strings.ToUpper(ident)
	} else if obj.lowerIdent || object.lowerIdent {
		ident = strings.ToLower(ident)
	}

	if obj.upperIdentCamel || object.upperIdentCamel {
		return upperCamelCase(ident)
	} else if obj.lowerIdentCamel || object.lowerIdentCamel {
		return lowerCamelCase(ident)
	}

	return ident
}

func (obj *ConstOption) genIdentNameObj(prefix string, object *Object) string {
	// option.prefix + prefix
	ident := obj.joinPrefix(obj.prefix, prefix)

	// ident + object.prefix
	ident = obj.joinPrefix(ident, object.prefix)

	// ident + object.Ident
	ident = obj.joinPrefix(ident, object.Ident)

	// ident + object.suffix
	ident = obj.joinPrefix(ident, object.suffix)

	// ident + option.suffix
	ident = obj.joinPrefix(ident, obj.suffix)

	return obj.getIdentNameObj(ident, object)
}

func (obj *ConstOption) getOutputFile(ext string) string {
	if len(obj.fileExt) != 0 {
		ext = obj.fileExt
	}
	return obj.outputPath + "/" + obj.fileName + UnderScore + DerivedSign + "." + ext
}

func (obj *ConstOption) getEnterCount(count int) string {
	if count == 0 {
		count = 2
	} else {
		count = 1
	}
	return strings.Repeat("\n", count)
}

func (obj *ConstOption) getEnterIndex(index int) string {
	if index == 0 {
		index = 1
	} else {
		index = 2
	}
	return strings.Repeat("\n", index)
}

func (obj *ConstOption) getTabWidth() string {
	count := obj.level * obj.tab
	if count < 0 {
		count = 0
	}
	return strings.Repeat(" ", count)
}

func (obj *ConstOption) getQmlReadonly(freeze bool) string {
	if obj.freeze || freeze {
		return "readonly "
	}
	return ""
}

func (obj *ConstOption) indentComment(comment string) string {
	if (obj.level * obj.tab) <= 0 {
		return comment
	}

	space := obj.getTabWidth()

	if !strings.Contains(comment, "\n") {
		return space + comment
	}

	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = space + line
	}

	return strings.Join(lines, "\n")
}

func (obj *ConstOption) joinPrefix(prefix1 string, prefix2 string) string {
	if len(prefix1) != 0 && len(prefix2) != 0 {
		return prefix1 + UnderScore + prefix2
	}

	if len(prefix1) != 0 {
		return prefix1
	}

	if len(prefix2) != 0 {
		return prefix2
	}

	return ""
}

func (obj *ConstOption) clearPrefixes() {
	obj.prefixes = make([]string, 0)
}

func (obj *ConstOption) getPrefixes(extra string) string {
	result := ""

	for _, item := range obj.prefixes {
		if len(item) == 0 {
			continue
		}

		if len(result) == 0 {
			result = item
		} else {
			result += UnderScore + item
		}
	}

	if len(extra) > 0 {
		if len(result) == 0 {
			result = extra
		} else {
			result += UnderScore + extra
		}
	}

	return result
}

func (obj *ConstOption) pushPrefixes(prefix string) {
	obj.prefixes = append(obj.prefixes, prefix)
}

func (obj *ConstOption) popPrefixes() string {
	if len(obj.prefixes) == 0 {
		return ""
	}

	prefix := obj.prefixes[len(obj.prefixes)-1]
	obj.prefixes = obj.prefixes[:len(obj.prefixes)-1]
	return prefix
}

func RunParser() {
	option := new(ConstOption)
	option.initUsage()

	// parse command-line flags
	pflag.Parse()

	// check file isEmpty
	if len(option.inputFile) == 0 {
		pflag.Usage()
		os.Exit(-1)
	}

	// check file existed
	_, err := os.Stat(option.inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open file: %s, error: %s\n", option.inputFile, err)
		os.Exit(-1)
	}

	// convert input path
	absPath, err := filepath.Abs(option.inputFile)
	if err == nil {
		option.inputFile = absPath
	}

	if option.verbose {
		fmt.Fprintf(os.Stdout, "intput file: %s\n", option.inputFile)
	}

	// get fileName
	if len(option.fileName) == 0 {
		fileName := filepath.Base(option.inputFile)
		option.fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}

	// check output path
	if len(option.outputPath) == 0 {
		exePath, _ := os.Executable()
		option.outputPath = filepath.Dir(exePath)
	} else {
		// convert output path
		absPath, err := filepath.Abs(option.outputPath)
		if err == nil {
			option.outputPath = absPath
		}

		// create output path
		err = os.MkdirAll(option.outputPath, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't create dir: %s, error: %s\n", option.outputPath, err)
			os.Exit(-1)
		}
	}

	if option.verbose {
		fmt.Fprintf(os.Stdout, "output path: %s, fileName: %s\n", option.outputPath, option.fileName)
	}

	// start parser
	parser := NewConstParser(option)
	if parser.ParseFile() == 0 {
		parser.GenFile()
	}
}
