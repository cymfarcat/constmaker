// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

/*
 *The first programming language I learned still dwells‌ in my memory.
 */
func (obj *ConstParser) genPascal(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	option.level = 0
	obj.NameSpace.genPascal(option, builder)

	path := option.getOutputFile("pas")
	data := builder.String()
	obj.writeFile(path, data, true, "(* ", " *)")

	option.clearPrefixes()
}

func pascalTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "Boolean"

	case Typo_I8:
		return "Shortint"

	case Typo_U8:
		return "Byte"

	case Typo_I16:
		return "Smallint"

	case Typo_U16:
		return "Word"

	case Typo_I32:
		return "Longint" //Integer

	case Typo_U32:
		return "LongWord"

	case Typo_I64:
		return "Int64"

	case Typo_U64:
		return "QWord"

	case Typo_F32:
		return "Real"

	case Typo_F64:
		return "Double"

	case Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *Object) genPascal(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.indentComment(convertCommentDocToPascal(constObj.CommentDoc)))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj)
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else {
			str = "\n" + option.getTabWidth() + "const " + option.genIdentNameObj(prefix, constObj) + ": " + pascalTypo(constObj.Typo) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripeToPascal(constObj.CommentTripe)
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genPascalStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "const " + option.genIdentNameObj(prefix, constObj) + ": String = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genPascalId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "const " + option.genIdentNameObj(prefix, constObj) + ": " + pascalTypo(constObj.Typo) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *NameSpace) genPascal(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(convertCommentDocToPascal(obj.CommentDoc)))
	}

	if obj.parent == nil {
		if option.rootNamespace {
			rootName := option.getIdentName(option.getRootName())

			option.level = 0
			option.pushPrefixes(rootName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "type")
				option.level++

				builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + option.getIdentNameObj(rootName, &obj.Object) + " = class")
				option.level++
			}
		}
	} else {
		option.pushPrefixes(namespace)

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + "(* namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object) + " *)")
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "type")
			option.level++

			builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + option.getIdentNameObj(namespace, &obj.Object) + " = class")
			option.level++
		}
	}

	prefix := ""
	if obj.namespaceToPrefix || option.namespaceToPrefix {
		prefix = option.getPrefixes("")
	}

	// generate consts
	obj.Object.genPascal(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genPascalId(prefix, option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.indentComment(convertCommentDocToPascal(enumObj.CommentDoc)))
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.pushPrefixes(enumName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes("")
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "(* enum " + option.getIdentNameObj(enumName, enumObj) + " *)")

			enumObj.genPascal(enumName, option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "type " + option.getIdentNameObj(enumName, enumObj) + " = (")
			option.level++

			enumObj.genPascal("", option, builder)
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + ");")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genPascalStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genPascal(option, builder)
	}

	if obj.parent == nil {
		if option.rootNamespace {
			option.popPrefixes()
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "end;")
			}
		}
	} else {
		option.popPrefixes()
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "end;")

			option.level--
		}
	}
}
