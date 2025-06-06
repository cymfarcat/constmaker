// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genGo(option *ConstOption) {
	builder := new(strings.Builder)

	if len(option.goPackage) > 0 {
		builder.WriteString("\npackage " + option.goPackage + "\n")
	}

	oldEnumToPrefix := option.enumToPrefix
	oldNamespaceToPrefix := option.namespaceToPrefix
	option.enumToPrefix = false     // Go no enum
	option.namespaceToPrefix = true // Go no namespace

	// generate root
	option.level = 0
	obj.NameSpace.genGo(option, builder)

	path := option.getOutputFile("go")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
	option.enumToPrefix = oldEnumToPrefix
	option.namespaceToPrefix = oldNamespaceToPrefix
}

func goTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "int8"

	case Typo_U8:
		return "uint8"

	case Typo_I16:
		return "int16"

	case Typo_U16:
		return "uint16"

	case Typo_I32:
		return "int32"

	case Typo_U32:
		return "uint32"

	case Typo_I64:
		return "int64"

	case Typo_U64:
		return "uint64"

	case Typo_F32:
		return "float32"

	case Typo_F64:
		return "float64"

	case Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *Object) genGo(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " " + goTypo(constObj.Typo) + " = " + constObj.Value
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " " + goTypo(constObj.Typo) + " = " + constObj.Value
			}
		} else {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " " + goTypo(constObj.Typo) + " = " + constObj.Value
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genGoStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " string = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genGoId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " " + goTypo(constObj.Typo) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * Go no namespace and enum
 */
func (obj *NameSpace) genGo(option *ConstOption, builder *strings.Builder) {
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
		}
	} else {
		option.pushPrefixes(namespace)

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
		}
	}

	// generate consts
	if len(obj.Object.Children) > 0 {
		builder.WriteString("\n" + option.getTabWidth() + "const (")
		option.level++
	}

	obj.Object.genGo(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genGoId(option.getPrefixes(""), option, builder)
	}

	if len(obj.Object.Children) > 0 {
		option.level--
		builder.WriteString(option.getTabWidth() + "\n)")
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.indentComment(enumObj.CommentDoc))
		}

		if /*enumObj.enumToPrefix ||*/ option.enumToPrefix {
		} else {
			if /*enumObj.enumToPrefix ||*/ option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "// enum " + option.getIdentNameObj(enumName, enumObj))

			builder.WriteString("\n" + option.getTabWidth() + "const (")
			option.level++

			enumObj.genGo(enumName, option, builder)
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumName = option.joinPrefix(enumObj.prefix, enumObj.Ident)

			option.pushPrefixes(enumName)
			enumObj.genGoStr(option.getPrefixes(""), option, builder)
			option.popPrefixes()
		}

		if /*enumObj.enumToPrefix ||*/ option.enumToPrefix {
		} else {
			option.level--
			builder.WriteString(option.getTabWidth() + "\n)")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genGo(option, builder)
	}

	if obj.parent == nil {
		if option.rootNamespace {
			option.popPrefixes()
			option.level = 0
		}
	} else {
		option.popPrefixes()
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
		} else {
		}
	}
}
