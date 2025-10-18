// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenMarkdown(option *src.Options) {
	builder := new(strings.Builder)

	oldEnumToPrefix := option.EnumToPrefix
	oldNamespaceToPrefix := option.NamespaceToPrefix
	option.EnumToPrefix = false     // Markdown no enum
	option.NamespaceToPrefix = true // Markdown no namespace

	// generate root
	option.Level = 0
	obj.NameSpace.GenMarkdown(option, builder)

	path := option.GetOutputFile("md")
	data := builder.String()
	obj.writeFile(path, data, true, "<!--", "-->")

	option.ClearPrefixes()
	option.EnumToPrefix = oldEnumToPrefix
	option.NamespaceToPrefix = oldNamespaceToPrefix
}

func markdownTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "bool"

	case src.Typo_I8:
		return "int8"

	case src.Typo_U8:
		return "uint8"

	case src.Typo_I16:
		return "int16"

	case src.Typo_U16:
		return "uint16"

	case src.Typo_I32:
		return "int32"

	case src.Typo_U32:
		return "uint32"

	case src.Typo_I64:
		return "int64"

	case src.Typo_U64:
		return "uint64"

	case src.Typo_F32:
		return "float32"

	case src.Typo_F64:
		return "float64"

	case src.Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *ConstObject) GenMarkdown(prefix string, option *src.Options, builder *strings.Builder) {
	if len(obj.Children) > 0 {
		builder.WriteString("\n>|ident|type|value|comment|")
		builder.WriteString("\n>|---|---|---|---|")
	}

	for _, constObj := range obj.Children {

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n>|" + genIdentName(option, prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
			} else {
				str = "\n>|" + genIdentName(option, prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
			}
		} else {
			str = "\n>|" + genIdentName(option, prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value
		}

		// triple comment
		comment := ""
		if len(constObj.CommentDoc) > 0 {
			comment = src.CleanCommentDoc(constObj.CommentDoc, ";")
		}
		if len(constObj.CommentTriple) > 0 {
			comment += src.ConvertCommentTriple(constObj.CommentTriple, ";")
		}

		str += "|" + comment + "|"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genMarkdownStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n>|" + genIdentName(option, prefix, constObj) + "| string | " + constObj.Value + "||"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genMarkdownId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n>|" + genIdentName(option, prefix, constObj) + "|" + markdownTypo(constObj.Typo) + "|" + constObj.Value + "||"

		builder.WriteString(str)
	}
}

/*
 * Markdown no namespace and enum
 */
func (obj *NameSpace) GenMarkdown(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	if obj.parent == nil {
		if option.RootNamespace {
			rootName := option.GetIdentName(option.GetRootName())

			option.Level = 0
			option.PushPrefixes(rootName)

			builder.WriteString("\n# " + rootName)
		}
	} else {
		option.PushPrefixes(namespace)

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "## namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
		}
	}

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n" + src.CleanCommentDoc(obj.CommentDoc, "\n"))
	}

	// generate consts
	if len(obj.ConstObject.Children) > 0 {
		//builder.WriteString("\n" + option.GetTabWidth() + "const (")
		//option.Level++
	}

	obj.ConstObject.GenMarkdown(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genMarkdownId(option.GetPrefixes(""), option, builder)
	}

	if len(obj.ConstObject.Children) > 0 {
		//option.Level--
		//builder.WriteString(option.GetTabWidth() + "\n)")
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		if /*enumObj.EnumToPrefix ||*/ option.EnumToPrefix {
		} else {
			if /*enumObj.EnumToPrefix ||*/ option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "### enum " + getIdentName(option, enumName, enumObj))

			//builder.WriteString("\n" + option.GetTabWidth() + "const (")
			//option.Level++

			// doc comment
			if len(enumObj.CommentDoc) > 0 {
				builder.WriteString("\n" + src.CleanCommentDoc(enumObj.CommentDoc, "\n"))
			}

			enumObj.GenMarkdown(enumName, option, builder)
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumName = option.JoinPrefix(enumObj.Prefix, enumObj.Ident)

			option.PushPrefixes(enumName)
			enumObj.genMarkdownStr(option.GetPrefixes(""), option, builder)
			option.PopPrefixes()
		}

		if /*enumObj.EnumToPrefix ||*/ option.EnumToPrefix {
		} else {
			//option.Level--
			//builder.WriteString(option.GetTabWidth() + "\n)")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenMarkdown(option, builder)
	}

	if obj.parent == nil {
		if option.RootNamespace {
			option.PopPrefixes()
			option.Level = 0
		}
	} else {
		option.PopPrefixes()
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
		} else {
		}
	}
}
