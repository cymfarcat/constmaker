// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genDart(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	option.level = 0
	obj.NameSpace.genDart(option, builder)

	path := option.getOutputFile("dart")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func dartTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "int"

	case Typo_U8:
		return "int"

	case Typo_I16:
		return "int"

	case Typo_U16:
		return "int"

	case Typo_I32:
		return "int"

	case Typo_U32:
		return "int"

	case Typo_I64:
		return "int"

	case Typo_U64:
		return "int"

	case Typo_F32:
		return "double"

	case Typo_F64:
		return "double"

	case Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *Object) genDart(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + "(" + constObj.Value + ")"
				if idx < len(obj.Children)-1 {
					str += ","
				} else {
					str += ";"
				}
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + ","
			}
		} else if obj.parent == nil && !option.rootNamespace {
			str = "\n" + option.getTabWidth() + "const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		} else if obj.objType == TypeEnum {
			if obj.enumToPrefix || option.enumToPrefix {
				str = "\n" + option.getTabWidth() + "const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
			} else {
				str = "\n" + option.getTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
			}
		} else if !(obj.namespaceToPrefix || option.namespaceToPrefix) {
			str = "\n" + option.getTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		} else {
			str = "\n" + option.getTabWidth() + "const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genDartStr(prefix string, static string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + static + "const String " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genDartId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if obj.parent != nil && !(obj.namespaceToPrefix || option.namespaceToPrefix) {
			str = "\n" + option.getTabWidth() + "static const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		} else {
			str = "\n" + option.getTabWidth() + "const " + dartTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

/*
 * Dart does not allow nested definitions of classes or enums.
 */
func (obj *NameSpace) genDart(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		if option.rootNamespace {
			rootName := option.getIdentName(option.getRootName())

			option.level = 0
			option.pushPrefixes(rootName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "class " + option.getIdentNameObj(rootName, &obj.Object) + " {")
				option.level++
			}
		}
	} else {
		option.pushPrefixes(namespace)

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "class " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object) + " {")
			option.level++
		}
	}

	prefix := ""
	if obj.namespaceToPrefix || option.namespaceToPrefix {
		prefix = option.getPrefixes("")
	}

	// generate consts
	obj.Object.genDart(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genDartId(prefix, option, builder)
	}

	if obj.parent == nil {
		if option.rootNamespace {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
				option.level = 0
			}
		}
	} else {
		if obj.namespaceToPrefix || option.namespaceToPrefix {
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
			option.level = 0
		}
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.indentComment(enumObj.CommentDoc))
		}

		option.level = 0
		if enumObj.enumToPrefix || option.enumToPrefix {
			option.pushPrefixes(enumName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes("")
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "// enum " + option.getIdentNameObj(enumName, enumObj))

			enumObj.genDart(option.getPrefixes(""), option, builder)
		} else {
			enumName = option.getPrefixes(enumName)

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "enum " + option.getIdentNameObj(enumName, enumObj) + " {")
			option.level++

			enumObj.genDart("", option, builder)

			if len(enumObj.genObjs) > 0 && enumObj.enumSetValue {
				builder.WriteString("\n")
				enumObj.genDartStr("", "static ", option, builder)
			}
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			if enumObj.enumSetValue {
				builder.WriteString("\n\n" + option.getTabWidth() + "final int value;")
				builder.WriteString("\n" + option.getTabWidth() + "const " + option.getIdentNameObj(enumName, enumObj) + "(this.value);")

				// generate fromValue
				builder.WriteString("\n\n" + option.getTabWidth() + "factory " + option.getIdentNameObj(enumName, enumObj) + ".fromValue(int value) {")
				option.level++
				builder.WriteString("\n" + option.getTabWidth() + "switch (value) {")

				option.level++
				for _, constObj := range enumObj.Children {
					builder.WriteString("\n" + option.getTabWidth() + "case " + calcShift(constObj.Value) + ": return " + option.getIdentNameObj(enumName, enumObj) + "." + option.genIdentNameObj("", constObj) + ";")
				}
				builder.WriteString("\n" + option.getTabWidth() + "default: throw StateError('" + option.getIdentNameObj(enumName, enumObj) + ".fromValue: invalid value=$value');")

				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")

				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
			}
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		}

		// enum need gen str
		if (enumObj.enumToPrefix || option.enumToPrefix || !enumObj.enumSetValue) && len(enumObj.genObjs) > 0 {
			enumName = option.joinPrefix(enumObj.prefix, enumObj.Ident)

			option.pushPrefixes(enumName)
			enumObj.genDartStr(option.getPrefixes(""), "", option, builder)
			option.popPrefixes()
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genDart(option, builder)
	}

	if obj.parent == nil {
		if option.rootNamespace {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
				option.popPrefixes()
			}
		}
	} else {
		option.popPrefixes()
	}
}
