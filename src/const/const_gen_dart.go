// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenDart(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenDart(option, builder)

	path := option.GetOutputFile("dart")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func dartTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "bool"

	case src.Typo_I8:
		return "int"

	case src.Typo_U8:
		return "int"

	case src.Typo_I16:
		return "int"

	case src.Typo_U16:
		return "int"

	case src.Typo_I32:
		return "int"

	case src.Typo_U32:
		return "int"

	case src.Typo_I64:
		return "int"

	case src.Typo_U64:
		return "int"

	case src.Typo_F32:
		return "double"

	case src.Typo_F64:
		return "double"

	case src.Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *ConstObject) GenDart(prefix string, option *src.Options, builder *strings.Builder) {

	for idx, constObj := range obj.Children {
		value := constObj.Value
		// PATCH HERE
		if constObj.Value == "\"$\"" {
			value = "r" + constObj.Value
		}

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + "(" + value + ")"
				if idx < len(obj.Children)-1 {
					str += ","
				} else {
					str += ";"
				}
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + ","
			}
		} else if obj.parent == nil && !option.RootNamespace {
			str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		} else if obj.objType == TypeEnum {
			if obj.EnumToPrefix || option.EnumToPrefix {
				if obj.IdentMapStr || obj.StrMapIdent {
					str = "\n" + option.GetTabWidth() + "const " + src.UpperCamelCase(prefix+"ID", false) /*dartTypo(constObj.Typo)*/ + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
				} else {
					str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
				}
			} else {
				str = "\n" + option.GetTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
			}
		} else if !(obj.NamespaceToPrefix || option.NamespaceToPrefix) {
			str = "\n" + option.GetTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genDartStr(prefix string, static string, option *src.Options, builder *strings.Builder) {
	builder.WriteString("\n")
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + static + "const String " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genDartId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if obj.parent != nil && !(obj.NamespaceToPrefix || option.NamespaceToPrefix) {
			str = "\n" + option.GetTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genDartIdentMapStr(identName string, option *src.Options, builder *strings.Builder) {
	identName2 := getIdentName(option, identName, obj)
	mapValue := identName2
	if obj.EnumToPrefix {
		mapValue = src.UpperCamelCase(identName+"ID", false)
	}
	mapName := src.LowerCamelCase("g_str_Map_"+identName, false)

	//
	builder.WriteString("\n" + option.GetTabWidth() + "//*********************************************")
	builder.WriteString("\n" + option.GetTabWidth() + "final Map<String, " + mapValue + "> " + mapName + " = {")
	option.Level++

	index := 0
	for _, constObj := range obj.Children {

		str := ""
		if obj.NamespaceToPrefix || obj.EnumToPrefix {
			str = "\n" + option.GetTabWidth() + genIdentName(option, identName, obj.genObjs[index]) + ": " + genIdentName(option, identName, constObj) + ","
		} else if obj.enumSetValue {
			str = "\n" + option.GetTabWidth() + identName2 + "." + genIdentName(option, "", obj.genObjs[index]) + ": " + identName2 + "." + genIdentName(option, "", constObj) + ","
		} else {
			str = "\n" + option.GetTabWidth() + genIdentName(option, identName, obj.genObjs[index]) + ": " + identName2 + "." + genIdentName(option, "", constObj) + ","
		}

		builder.WriteString(str)
		index++
	}

	option.Level--
	builder.WriteString("\n" + option.GetTabWidth() + "};\n")

	//
	str := mapValue + "? " + identName2 + "FromString(String str) {"
	builder.WriteString("\n" + option.GetTabWidth() + str)

	option.Level++
	builder.WriteString("\n" + option.GetTabWidth() + "return " + mapName + "[str];")

	option.Level--
	builder.WriteString("\n" + option.GetTabWidth() + "}")
}

func (obj *ConstObject) genDartStrMapIdent(identName string, option *src.Options, builder *strings.Builder) {
	identName2 := getIdentName(option, identName, obj)
	mapValue := identName2
	if obj.EnumToPrefix {
		mapValue = src.UpperCamelCase(identName+"ID", false)
	}
	mapName := src.LowerCamelCase("g_"+identName+"_Map_Str", false)

	//
	builder.WriteString("\n" + option.GetTabWidth() + "//*********************************************")
	builder.WriteString("\n" + option.GetTabWidth() + "final Map<" + mapValue + ", String> " + mapName + " = {")
	option.Level++

	index := 0
	for _, constObj := range obj.Children {

		str := ""

		if obj.NamespaceToPrefix || obj.EnumToPrefix {
			str = "\n" + option.GetTabWidth() + genIdentName(option, identName, constObj) + ": " + genIdentName(option, identName, obj.genObjs[index]) + ","
		} else if obj.enumSetValue {
			str = "\n" + option.GetTabWidth() + identName2 + "." + genIdentName(option, "", constObj) + ": " + identName2 + "." + genIdentName(option, "", obj.genObjs[index]) + ","
		} else {
			str = "\n" + option.GetTabWidth() + identName2 + "." + genIdentName(option, "", constObj) + ": " + genIdentName(option, identName, obj.genObjs[index]) + ","
		}

		builder.WriteString(str)
		index++
	}

	option.Level--
	builder.WriteString("\n" + option.GetTabWidth() + "};\n")

	//
	str := "String? " + src.LowerCamelCase("stringFrom_"+identName, false) + "(" + mapValue + " value) {"
	builder.WriteString("\n" + option.GetTabWidth() + str)

	option.Level++
	builder.WriteString("\n" + option.GetTabWidth() + "return " + mapName + "[value];")

	option.Level--
	builder.WriteString("\n" + option.GetTabWidth() + "}")
}

// HACK for qualified_name
func (obj *ConstObject) genDartQName(namespace *NameSpace, prefix string, option *src.Options, builder *strings.Builder) {
	mod := ""
	mod_namespace := ""
	mod_namespace_URI := ""

	for _, constObj := range namespace.Children {
		if constObj.Ident == "mod" {
			mod = src.UnquoteStr(constObj.Value)
			mod_namespace = src.LowerCamelCase(mod+"_namespace", true)
			mod_namespace_URI = src.LowerCamelCase(mod+"_namespace_URI", true)
		}
	}

	qualified_name := "QualifiedName"
	is_node := strings.Contains(strings.ToLower(obj.Ident), "_node")
	if is_node {
		builder.WriteString("\n" + option.GetTabWidth() + "class " + strings.ToUpper(mod) + qualified_name + " extends " + qualified_name + " {")
		option.Level += 1
		builder.WriteString("\n" + option.GetTabWidth() + strings.ToUpper(mod) + qualified_name + "(super.prefix, super.localName, super.namespaceUri, [super.flag = 0, super.namespace = 0, super.nodeId = 0]);")
		option.Level -= 1
		builder.WriteString("\n" + option.GetTabWidth() + "}\n")

		qualified_name = strings.ToUpper(mod) + qualified_name
	}

	for _, constObj := range obj.Children {

		arg := ""
		if is_node {
			arg = "gEmptyStr, " + genIdentName(option, prefix, constObj) + "Str, " + mod_namespace_URI + ", 1, " + mod_namespace + ", " + genIdentName(option, prefix, constObj)
		} else {
			arg = "gEmptyStr, " + genIdentName(option, prefix, constObj) + "Str, gEmptyStr" + ", 1, " + mod_namespace + ", " + genIdentName(option, prefix, constObj)
		}

		str := ""
		if obj.parent != nil && !(obj.NamespaceToPrefix || option.NamespaceToPrefix) {
			str = "\n" + option.GetTabWidth() + "final " + qualified_name + " " + genIdentName(option, prefix, constObj) + "Tag = " + qualified_name + "(" + arg + ");"
		} else {
			str = "\n" + option.GetTabWidth() + "final " + qualified_name + " " + genIdentName(option, prefix, constObj) + "Tag = " + qualified_name + "(" + arg + ");"
		}

		builder.WriteString(str)
	}
}

/*
 * Dart does not allow nested definitions of classes or enums.
 */
func (obj *NameSpace) GenDart(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
	}

	// HACK for qualified_name
	for _, constObj := range obj.Children {
		if constObj.Ident == "mod" {
			builder.WriteString("\nimport 'package:lib_appbase/lib_appbase.dart';")
			builder.WriteString("\nimport '../dom/qualified_name.dart';\n")
		}
	}

	if obj.parent == nil {
		if option.RootNamespace {
			rootName := option.GetIdentName(option.GetRootName())

			option.Level = 0
			option.PushPrefixes(rootName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "class " + getIdentName(option, rootName, &obj.ConstObject) + " {")
				option.Level++
			}
		}
	} else {
		option.PushPrefixes(namespace)

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
			if obj.IdentMapStr || obj.StrMapIdent {
				builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "typedef " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject) + " = String;")
			}
		} else {
			option.Level = 0
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "class " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject) + " {")
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenDart(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genDartId(prefix, option, builder)
	}

	// gen IdentMapStr/StrMapIdent
	if obj.IdentMapStr {
		builder.WriteString("\n")
		obj.genDartIdentMapStr(prefix, option, builder)
	}

	if obj.StrMapIdent {
		builder.WriteString("\n")
		obj.genDartStrMapIdent(prefix, option, builder)
	}

	if obj.StrMapQName {
		builder.WriteString("\n")
		obj.genDartQName(obj, prefix, option, builder)
	}

	if obj.parent == nil {
		if option.RootNamespace {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")
				option.Level = 0
			}
		}
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
			option.Level = 0
		}
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.IndentComment(enumObj.CommentDoc))
		}

		option.Level = 0
		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PushPrefixes(enumName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes("")
			}

			if obj.IdentMapStr || obj.StrMapIdent {
				builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "typedef " + src.UpperCamelCase(enumName+"ID", false) + " = " + dartTypo(enumObj.Typo) + ";")
			}

			enumObj.GenDart(option.GetPrefixes(""), option, builder)
		} else {
			enumName = option.GetPrefixes(enumName)

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenDart("", option, builder)

			if len(enumObj.genObjs) > 0 && enumObj.enumSetValue {
				builder.WriteString("\n")
				enumObj.genDartStr("", "static ", option, builder)
			}
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			if enumObj.enumSetValue {
				builder.WriteString("\n\n" + option.GetTabWidth() + "final int value;")
				builder.WriteString("\n" + option.GetTabWidth() + "const " + getIdentName(option, enumName, enumObj) + "(this.value);")

				// generate fromValue
				builder.WriteString("\n\n" + option.GetTabWidth() + "factory " + getIdentName(option, enumName, enumObj) + ".fromValue(int value) {")
				option.Level++
				builder.WriteString("\n" + option.GetTabWidth() + "switch (value) {")

				option.Level++
				for _, constObj := range enumObj.Children {
					builder.WriteString("\n" + option.GetTabWidth() + "case " + src.CalcShift(constObj.Value) + ": return " + getIdentName(option, enumName, enumObj) + "." + genIdentName(option, "", constObj) + ";")
				}
				builder.WriteString("\n" + option.GetTabWidth() + "default: throw StateError('" + getIdentName(option, enumName, enumObj) + ".fromValue: invalid value=$value');")

				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")

				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")
			}
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}

		// enum need gen str
		if (enumObj.EnumToPrefix || option.EnumToPrefix || !enumObj.enumSetValue) && len(enumObj.genObjs) > 0 {
			enumName = option.JoinPrefix(enumObj.Prefix, enumObj.Ident)

			option.PushPrefixes(enumName)
			enumObj.genDartStr(option.GetPrefixes(""), "", option, builder)
			option.PopPrefixes()
		}

		// gen IdentMapStr/StrMapIdent
		if enumObj.IdentMapStr {
			enumName = option.GetPrefixes(enumName)
			builder.WriteString("\n")
			enumObj.genDartIdentMapStr(enumName, option, builder)
		}

		if enumObj.StrMapIdent {
			enumName = option.GetPrefixes(enumName)
			builder.WriteString("\n")
			enumObj.genDartStrMapIdent(enumName, option, builder)
		}

		if enumObj.StrMapQName {
			enumName = option.GetPrefixes(enumName)
			builder.WriteString("\n")
			enumObj.genDartQName(obj, enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenDart(option, builder)
	}

	if obj.parent == nil {
		if option.RootNamespace {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				option.PopPrefixes()
			}
		}
	} else {
		option.PopPrefixes()
	}
}
