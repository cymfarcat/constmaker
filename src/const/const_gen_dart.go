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

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + "(" + constObj.Value + ")"
				if idx < len(obj.Children)-1 {
					str += ","
				} else {
					str += ";"
				}
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + ","
			}
		} else if obj.parent == nil && !option.RootNamespace {
			str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		} else if obj.objType == TypeEnum {
			if obj.EnumToPrefix || option.EnumToPrefix {
				str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
			} else {
				str = "\n" + option.GetTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
			}
		} else if !(obj.NamespaceToPrefix || option.NamespaceToPrefix) {
			str = "\n" + option.GetTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + "const " + dartTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genDartStr(prefix string, static string, option *src.Options, builder *strings.Builder) {
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

/*
 * Dart does not allow nested definitions of classes or enums.
 */
func (obj *NameSpace) GenDart(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
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
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
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

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "// enum " + getIdentName(option, enumName, enumObj))

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
