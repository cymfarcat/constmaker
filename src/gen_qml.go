// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genQml(option *ConstOption) {
	builder := new(strings.Builder)

	// add pragma Singleton
	if option.qmlSingleton {
		builder.WriteString("\npragma Singleton\n")
	}

	// generate root
	option.level = 0
	obj.NameSpace.genQml(option, builder)

	path := option.getOutputFile("qml")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func qmlTypo(typo string) string {
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
		return "real"

	case Typo_F64:
		return "real"

	case Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *Object) genQml(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj)
			}
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else {
			str = "\n" + option.getTabWidth() + option.getQmlReadonly(constObj.freeze) + "property " + qmlTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + ": " + constObj.Value
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genQmlStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + option.getQmlReadonly(constObj.freeze) + "property string " + option.genIdentNameObj(prefix, constObj) + ": " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *Object) genQmlId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + option.getQmlReadonly(constObj.freeze) + "property " + qmlTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + ": " + constObj.Value

		builder.WriteString(str)
	}
}

/*
 * In Qml, constants must be declared within a QtObject.
 */
func (obj *NameSpace) genQml(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		rootName := option.getIdentName(option.getRootName())

		// option.rootAsNode=true force do this
		builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "QtObject {")
		option.level++

		// write id
		builder.WriteString("\n" + option.getTabWidth() + "id: " + strings.ToLower(rootName) + "\n")
	} else {
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.pushPrefixes(namespace)
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "QtObject {")
			option.level++

			// write id
			builder.WriteString("\n" + option.getTabWidth() + "id: " + strings.ToLower(namespace) + "\n")
		}
	}

	// generate consts
	obj.Object.genQml(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genQmlId(option.getPrefixes(""), option, builder)
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
			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "// enum " + option.getPrefixes(""))

			enumObj.genQml(option.getPrefixes(""), option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "enum " + option.getIdentNameObj(enumName, enumObj) + " {")
			option.level++

			enumObj.genQml("", option, builder)
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
			enumObj.genQmlStr(option.getPrefixes(""), option, builder)
			option.popPrefixes()
		}
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genQml(option, builder)
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
}
