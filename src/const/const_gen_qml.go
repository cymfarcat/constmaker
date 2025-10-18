// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenQml(option *src.Options) {
	builder := new(strings.Builder)

	// add pragma Singleton
	if option.QmlSingleton {
		builder.WriteString("\npragma Singleton\n")
	}

	// generate root
	option.Level = 0
	obj.NameSpace.GenQml(option, builder)

	path := option.GetOutputFile("qml")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func qmlTypo(typo string) string {
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
		return "real"

	case src.Typo_F64:
		return "real"

	case src.Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *ConstObject) GenQml(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj)
			}
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else {
			str = "\n" + option.GetTabWidth() + option.GetQmlReadonly(constObj.freeze) + "property " + qmlTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + ": " + constObj.Value
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genQmlStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + option.GetQmlReadonly(constObj.freeze) + "property string " + genIdentName(option, prefix, constObj) + ": " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genQmlId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + option.GetQmlReadonly(constObj.freeze) + "property " + qmlTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + ": " + constObj.Value

		builder.WriteString(str)
	}
}

/*
 * In Qml, constants must be declared within a QtObject.
 */
func (obj *NameSpace) GenQml(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		rootName := option.GetIdentName(option.GetRootName())

		// option.rootAsNode=true force do this
		builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "QtObject {")
		option.Level++

		// write id
		builder.WriteString("\n" + option.GetTabWidth() + "id: " + strings.ToLower(rootName) + "\n")
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.PushPrefixes(namespace)
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "QtObject {")
			option.Level++

			// write id
			builder.WriteString("\n" + option.GetTabWidth() + "id: " + strings.ToLower(namespace) + "\n")
		}
	}

	// generate consts
	obj.ConstObject.GenQml(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genQmlId(option.GetPrefixes(""), option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.IndentComment(enumObj.CommentDoc))
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PushPrefixes(enumName)
			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "// enum " + option.GetPrefixes(""))

			enumObj.GenQml(option.GetPrefixes(""), option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenQml("", option, builder)
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumName = option.JoinPrefix(enumObj.Prefix, enumObj.Ident)

			option.PushPrefixes(enumName)
			enumObj.genQmlStr(option.GetPrefixes(""), option, builder)
			option.PopPrefixes()
		}
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenQml(option, builder)
	}

	if obj.parent == nil {
		option.Level--
		builder.WriteString("\n" + option.GetTabWidth() + "}")
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}
	}
}
