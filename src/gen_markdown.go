// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genMarkdown(option *ConstOption) {
	builder := new(strings.Builder)

	oldEnumToPrefix := option.enumToPrefix
	oldNamespaceToPrefix := option.namespaceToPrefix
	option.enumToPrefix = false     // Markdown no enum
	option.namespaceToPrefix = true // Markdown no namespace

	// generate root
	option.level = 0
	obj.NameSpace.genMarkdown(option, builder)

	path := option.getOutputFile("md")
	data := builder.String()
	obj.writeFile(path, data, true, "<!--", "-->")

	option.clearPrefixes()
	option.enumToPrefix = oldEnumToPrefix
	option.namespaceToPrefix = oldNamespaceToPrefix
}

func markdownTypo(typo string) string {
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

func (obj *Object) genMarkdown(prefix string, option *ConstOption, builder *strings.Builder) {
	if len(obj.Children) > 0 {
		builder.WriteString("\n>|ident|type|value|comment|")
		builder.WriteString("\n>|---|---|---|---|")
	}

	for _, constObj := range obj.Children {

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n>|" + option.genIdentNameObj(prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
			} else {
				str = "\n>|" + option.genIdentNameObj(prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
			}
		} else {
			str = "\n>|" + option.genIdentNameObj(prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
		}

		// triple comment
		comment := ""
		if len(constObj.CommentDoc) > 0 {
			comment = cleanCommentDoc(constObj.CommentDoc, ";")
		}
		if len(constObj.CommentTripe) > 0 {
			comment += convertCommentTripe(constObj.CommentTripe, ";")
		}

		str += "|" + comment + "|"

		builder.WriteString(str)
	}
}

func (obj *Object) genMarkdownStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n>|" + option.genIdentNameObj(prefix, constObj) + "| string | " + constObj.Value + "||"

		builder.WriteString(str)
	}
}

func (obj *Object) genMarkdownId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n>|" + option.genIdentNameObj(prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value + "||"

		builder.WriteString(str)
	}
}

/*
 * Markdown no namespace and enum
 */
func (obj *NameSpace) genMarkdown(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	if obj.parent == nil {
		if option.rootNamespace {
			rootName := option.getIdentName(option.getRootName())

			option.level = 0
			option.pushPrefixes(rootName)

			builder.WriteString("\n# " + rootName)
		}
	} else {
		option.pushPrefixes(namespace)

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + "## namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
		}
	}

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n" + cleanCommentDoc(obj.CommentDoc, "\n"))
	}

	// generate consts
	if len(obj.Object.Children) > 0 {
		//builder.WriteString("\n" + option.getTabWidth() + "const (")
		//option.level++
	}

	obj.Object.genMarkdown(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genMarkdownId(option.getPrefixes(""), option, builder)
	}

	if len(obj.Object.Children) > 0 {
		//option.level--
		//builder.WriteString(option.getTabWidth() + "\n)")
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		if /*enumObj.enumToPrefix ||*/ option.enumToPrefix {
		} else {
			if /*enumObj.enumToPrefix ||*/ option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "### enum " + option.getIdentNameObj(enumName, enumObj))

			//builder.WriteString("\n" + option.getTabWidth() + "const (")
			//option.level++

			// doc comment
			if len(enumObj.CommentDoc) > 0 {
				builder.WriteString("\n" + cleanCommentDoc(enumObj.CommentDoc, "\n"))
			}

			enumObj.genMarkdown(enumName, option, builder)
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumName = option.joinPrefix(enumObj.prefix, enumObj.Ident)

			option.pushPrefixes(enumName)
			enumObj.genMarkdownStr(option.getPrefixes(""), option, builder)
			option.popPrefixes()
		}

		if /*enumObj.enumToPrefix ||*/ option.enumToPrefix {
		} else {
			//option.level--
			//builder.WriteString(option.getTabWidth() + "\n)")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genMarkdown(option, builder)
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
