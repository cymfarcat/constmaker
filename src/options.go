// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync/atomic"

	"github.com/spf13/pflag"
)

// global id
var globalId uint32 = 0

func GlobalId() uint32 {
	return atomic.AddUint32(&globalId, 1)
}

func GlobalIdString() string {
	return fmt.Sprintf("%04d", GlobalId())
}

// generated random string
const defaultCharset = "abcdefghijklmnopqrstuvwxyz0123456789" // + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterCharset = "abcdefghijklmnopqrstuvwxyz"            // + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var generatedRandomString = make(map[string]bool)

type Options struct {
	InputFile  string
	OutputPath string
	FileName   string
	fileExt    string
	CheckValue bool
	Verbose    bool

	Prefix            string
	prefixes          []string
	Suffix            string
	tab               int8
	Level             int8
	rootName          string
	RootNamespace     bool
	NestedDefine      bool
	MacroDefine       bool
	freeze            bool
	UpperIdent        bool
	LowerIdent        bool
	UpperIdentCamel   bool
	LowerIdentCamel   bool
	EnumToPrefix      bool
	NamespaceToPrefix bool
	BitFlagNoneAll    bool
	BitFlagNone       string
	BitFlagAll        string

	// for sqlite
	SQLite             bool
	GenSql             bool
	UsePrefix          bool
	NameMin            uint8
	NameMax            uint8
	TablePrefix        string
	IndexPrefix        string
	ViewPrefix         string
	TriggerPrefix      string
	VirtualTablePrefix string
	fieldWhiteList     string
	fieldWhiteMap      map[string]bool

	// all support language
	GenAll bool

	// c/c++
	GenCpp   bool
	StdCpp11 bool

	// c#
	GenCShape bool

	// dart
	GenDart bool

	// go
	GenGo     bool
	GoPackage string

	// java
	GenJava     bool
	JavaPackage string

	// javascript
	GenJavaScript bool

	// json
	GenJson bool

	// kotlin
	GenKotlin bool

	// markdown
	GenMarkdown bool

	// ObjC
	GenObjC bool

	// pascal
	GenPascal bool

	// perl
	GenPerl bool

	// php
	GenPhp bool

	// python
	GenPython bool

	// qml
	GenQml       bool
	QmlSingleton bool

	// ruby
	GenRuby bool

	// rust
	GenRust bool

	// swift
	GenSwift bool

	// text
	GenText bool

	// typescript
	GenTypeScript bool

	// xml
	GenXml bool
}

func (obj *Options) InitUsage() {
	pflag.StringVarP(&obj.InputFile, "input", "i", "", "input file to const define.")
	pflag.StringVarP(&obj.OutputPath, "output", "o", "", "output path to generated files.")
	pflag.StringVarP(&obj.FileName, "file-name", "", "", "output file name.")
	pflag.StringVarP(&obj.fileExt, "file-ext", "", "", "output file name extension.")
	pflag.BoolVarP(&obj.CheckValue, "check-value", "", false, "check for duplicate constant values within the current scope.")
	pflag.BoolVarP(&obj.Verbose, "verbose", "v", false, "print verbose info.")

	pflag.StringVarP(&obj.Prefix, Option_Prefix, "", "", "prefix for every const ident name.")
	pflag.StringVarP(&obj.Suffix, Option_Suffix, "", "", "suffix for every const ident name.")
	pflag.StringVarP(&obj.rootName, "root-name", "", "", "root name, if empty substitute with file name.")
	pflag.BoolVarP(&obj.RootNamespace, "root-namespace", "", false, "forece root node into namspace.")
	pflag.BoolVarP(&obj.RootNamespace, "root-ns", "", false, "forece root node into namspace.")
	pflag.BoolVarP(&obj.NestedDefine, "nested-define", "", true, "force namespace use nested define like tree, or else use flat define like list.")
	pflag.BoolVarP(&obj.MacroDefine, Option_MacroDefine, "", false, "force use macro define to declare constants.")
	pflag.BoolVarP(&obj.freeze, Option_Freeze, "", false, "force use readonly to declare constants.")
	pflag.BoolVarP(&obj.UpperIdent, Option_UpperIdent, "", false, "force ident to upper.")
	pflag.BoolVarP(&obj.LowerIdent, Option_LowerIdent, "", false, "force ident to lower.")
	pflag.BoolVarP(&obj.UpperIdentCamel, Option_UpperIdentCamel, "", false, "force ident to upper camel case.")
	pflag.BoolVarP(&obj.LowerIdentCamel, Option_LowerIdentCamel, "", false, "force ident to lower camel case.")
	pflag.Int8VarP(&obj.tab, "tab", "t", 4, "tab width.")
	pflag.BoolVarP(&obj.EnumToPrefix, Option_EnumToPrefix, "", false, "use enum name as prefix.")
	pflag.BoolVarP(&obj.NamespaceToPrefix, Option_NamespaceToPrefix, "", false, "use namespace as prefix.")
	pflag.BoolVarP(&obj.NamespaceToPrefix, Option_NSToPrefix, "", false, "use namespace as prefix.")

	pflag.BoolVarP(&obj.BitFlagNoneAll, "bitflag-none-all", "", true, "bitflag generate None and All option.")
	pflag.StringVarP(&obj.BitFlagNone, "bitflag-none", "", "NONE", "bitflag generate None string.")
	pflag.StringVarP(&obj.BitFlagAll, "bitflag-all", "", "ALL", "bitflag generate All string.")

	// for sqlite
	pflag.BoolVarP(&obj.SQLite, SQLITE, "", false, "use sqlite consts.")
	pflag.BoolVarP(&obj.GenSql, "sql", "", false, "generate sql files.")
	pflag.BoolVarP(&obj.UsePrefix, "use-prefix", "", false, "force name use prefix.")
	pflag.Uint8VarP(&obj.NameMin, "min-name", "", 4, "min name length.")
	pflag.Uint8VarP(&obj.NameMax, "max-name", "", 8, "max name length.")
	pflag.StringVarP(&obj.TablePrefix, "table", "", "tbl", "prefix for every table name.")
	pflag.StringVarP(&obj.IndexPrefix, "index", "", "idx", "prefix for every index name.")
	pflag.StringVarP(&obj.ViewPrefix, "view", "", "vw", "prefix for every trigger name.")
	pflag.StringVarP(&obj.TriggerPrefix, "trigger", "", "trg", "prefix for every trigger name.")
	pflag.StringVarP(&obj.VirtualTablePrefix, "vtable", "", "vt", "prefix for every virtual table name.")
	pflag.StringVarP(&obj.fieldWhiteList, "field-wl", "", "", "semicolon-separated whitelist field.")

	// all
	pflag.BoolVarP(&obj.GenAll, "all", "a", false, "generate all supported language files.")

	// c/c++
	pflag.BoolVarP(&obj.GenCpp, "cpp", "c", false, "generate C/CPP head files.")
	pflag.BoolVarP(&obj.StdCpp11, "std-cpp11", "", true, "std c++11.")

	// c#
	pflag.BoolVarP(&obj.GenCShape, "c#", "", false, "generate C# files.")

	// dart
	pflag.BoolVarP(&obj.GenDart, "dart", "d", false, "generate Dart files.")

	// go
	pflag.BoolVarP(&obj.GenGo, "go", "g", false, "generate Go files.")
	pflag.StringVarP(&obj.GoPackage, "go-package", "", "", "generate Go package.")

	// java
	pflag.BoolVarP(&obj.GenJava, "java", "j", false, "generate Java files.")
	pflag.StringVarP(&obj.JavaPackage, "java-package", "", "", "generate Java package.")

	// javascript
	pflag.BoolVarP(&obj.GenJavaScript, "javascript", "", false, "generate JavaScript files.")

	// json
	pflag.BoolVarP(&obj.GenJson, "json", "", false, "generate Json files.")

	// kotlin
	pflag.BoolVarP(&obj.GenKotlin, "kotlin", "", false, "generate Kotlin files.")

	// markdown
	pflag.BoolVarP(&obj.GenMarkdown, "markdown", "", false, "generate Markdown files.")

	// ObjC
	pflag.BoolVarP(&obj.GenObjC, "objc", "", false, "generate Objective-C files.")

	// pascal
	pflag.BoolVarP(&obj.GenPascal, "pascal", "", false, "generate Pascal files.")

	// perl
	pflag.BoolVarP(&obj.GenPerl, "perl", "", false, "generate Perl files.")

	// php
	pflag.BoolVarP(&obj.GenPhp, "php", "", false, "generate Php files.")

	// python
	pflag.BoolVarP(&obj.GenPython, "python", "p", false, "generate Python files.")

	// qml
	pflag.BoolVarP(&obj.GenQml, "qml", "", false, "generate QML files.")
	pflag.BoolVarP(&obj.QmlSingleton, "qml-singleton", "", false, "generate QML singleton.")

	// ruby
	pflag.BoolVarP(&obj.GenRuby, "ruby", "", false, "generate Ruby files.")

	// rust
	pflag.BoolVarP(&obj.GenRust, "rust", "r", false, "generate Rust files.")

	// swift
	pflag.BoolVarP(&obj.GenSwift, "swift", "s", false, "generate Swift files.")

	// text
	pflag.BoolVarP(&obj.GenText, "text", "", false, "generate Text files.")

	// typescript
	pflag.BoolVarP(&obj.GenTypeScript, "typescript", "", false, "generate TypeScript files.")

	// xml
	pflag.BoolVarP(&obj.GenXml, "xml", "", false, "generate XML files.")
}

func (obj *Options) InitFieldWhiteList() {
	obj.fieldWhiteMap = make(map[string]bool)

	newItems := strings.Split(obj.fieldWhiteList, SeparatorChar)
	for _, item := range newItems {
		if trimmed := strings.TrimSpace(item); trimmed != "" {
			obj.fieldWhiteMap[trimmed] = true
		}
	}
}

func (obj *Options) IsFieldWhiteList(field string) bool {
	_, ok := obj.fieldWhiteMap[field]
	return ok
}

func (obj *Options) GenerateRandomString() string {
	// ensure valid length range
	if obj.NameMin > obj.NameMax {
		obj.NameMin, obj.NameMax = obj.NameMax, obj.NameMin
	}

	var count = 0
	for {
		// generate random length within specified range
		lengthRange := obj.NameMax - obj.NameMin + 1
		randLengthBig, _ := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
		randLength := uint8(randLengthBig.Int64()) + obj.NameMin

		// ensure rand length
		if randLength < 4 {
			randLength = 4
		}

		b := make([]byte, randLength)

		// ensure the first character is a letter
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterCharset))))
		b[0] = letterCharset[n.Int64()]

		// generate remaining characters
		for i := 1; i < int(randLength); i++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(defaultCharset))))
			b[i] = defaultCharset[n.Int64()]
		}

		str := string(b)

		if !generatedRandomString[str] {
			generatedRandomString[str] = true
			return str
		}

		count++
		if count > 65535 {
			fmt.Fprintf(os.Stderr, "Can't generate random string. Please adjust argument: min-name/max-name.")
			os.Exit(-1)
		}
	}
}

func (obj *Options) GetRootName() string {
	if len(obj.rootName) > 0 {
		return obj.rootName
	}
	return obj.FileName
}

func (obj *Options) GetIdentName(ident string) string {
	if obj.UpperIdent {
		ident = strings.ToUpper(ident)
	} else if obj.LowerIdent {
		ident = strings.ToLower(ident)
	}

	if obj.UpperIdentCamel {
		return UpperCamelCase(ident)
	} else if obj.LowerIdentCamel {
		return LowerCamelCase(ident)
	}

	return ident
}

func (obj *Options) GetOutputFile(ext string) string {
	if len(obj.fileExt) != 0 {
		ext = obj.fileExt
	}
	return obj.OutputPath + "/" + obj.FileName + UnderScore + DerivedSign + "." + ext
}

func (obj *Options) GetEnterCount(count int) string {
	if count == 0 {
		count = 2
	} else {
		count = 1
	}
	return strings.Repeat("\n", count)
}

func (obj *Options) GetEnterIndex(index int) string {
	if index == 0 {
		index = 1
	} else {
		index = 2
	}
	return strings.Repeat("\n", index)
}

func (obj *Options) GetTabWidth() string {
	count := obj.Level * obj.tab
	if count < 0 {
		count = 0
	}
	return strings.Repeat(" ", int(count))
}

func (obj *Options) GetQmlReadonly(freeze bool) string {
	if obj.freeze || freeze {
		return "readonly "
	}
	return ""
}

func (obj *Options) IndentComment(comment string) string {
	if (obj.Level * obj.tab) <= 0 {
		return comment
	}

	space := obj.GetTabWidth()

	if !strings.Contains(comment, "\n") {
		return space + comment
	}

	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = space + line
	}

	return strings.Join(lines, "\n")
}

func (obj *Options) JoinPrefix(prefix1 string, prefix2 string) string {
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

func (obj *Options) ClearPrefixes() {
	obj.prefixes = make([]string, 0)
}

func (obj *Options) GetPrefixes(extra string) string {
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

func (obj *Options) PushPrefixes(prefix string) {
	obj.prefixes = append(obj.prefixes, prefix)
}

func (obj *Options) PopPrefixes() string {
	if len(obj.prefixes) == 0 {
		return ""
	}

	prefix := obj.prefixes[len(obj.prefixes)-1]
	obj.prefixes = obj.prefixes[:len(obj.prefixes)-1]
	return prefix
}
