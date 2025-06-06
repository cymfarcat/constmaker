// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genCShape(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	option.level = 0
	obj.NameSpace.genCShape(option, builder)

	path := option.getOutputFile("cs")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func cshapeTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "sbyte"

	case Typo_U8:
		return "byte"

	case Typo_I16:
		return "short"

	case Typo_U16:
		return "ushort"

	case Typo_I32:
		return "int"

	case Typo_U32:
		return "uint"

	case Typo_I64:
		return "long"

	case Typo_U64:
		return "ulong"

	case Typo_F32:
		return "float"

	case Typo_F64:
		return "double"

	case Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *Object) genCShape(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := adjustFloat(constObj.Value, constObj.Typo)

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + value + ","
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + ","
			}
		} else {
			str = "\n" + option.getTabWidth() + "public const " + cshapeTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genCShapeStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "public const string " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genCShapeId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "public const " + cshapeTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * In C#, constants must be declared within a class or struct.
 */
func (obj *NameSpace) genCShape(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		rootName := option.getIdentName(option.getRootName())

		// option.rootAsNode=true force do this
		builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "public static class " + option.getIdentNameObj(rootName, &obj.Object) + " {")
		option.level++
	} else {
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.pushPrefixes(namespace)
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "public static class " + option.getIdentNameObj(namespace, &obj.Object) + " {")
			option.level++
		}
	}

	// generate consts
	obj.Object.genCShape(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genCShapeId(option.getPrefixes(""), option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.indentComment(enumObj.CommentDoc))
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.pushPrefixes(enumName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes("")
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "// enum " + option.getIdentNameObj(enumName, enumObj))

			enumObj.genCShape(option.getPrefixes(""), option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "public enum " + option.getIdentNameObj(enumName, enumObj) + " {")
			option.level++

			enumObj.genCShape("", option, builder)
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumName = option.joinPrefix(enumObj.prefix, enumObj.Ident)

			option.pushPrefixes(enumName)
			enumObj.genCShapeStr(option.getPrefixes(""), option, builder)
			option.popPrefixes()
		}
	}

	if option.nestedDefine || obj.namespaceToPrefix || option.namespaceToPrefix {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genCShape(option, builder)
		}

		if obj.parent == nil {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
				option.popPrefixes()
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
			}
		}
	} else {
		if obj.parent == nil {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
				option.level = 0
				//option.popPrefixes()
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genCShape(option, builder)
		}

		if obj.parent == nil {
			if option.rootNamespace {
				if obj.namespaceToPrefix || option.namespaceToPrefix {
					option.popPrefixes()
				}
			}
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
				option.popPrefixes()
			}
		}
	}
}
