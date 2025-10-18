// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenRust(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	oldNamespaceToPrefix := option.NamespaceToPrefix
	option.NamespaceToPrefix = true // rust no namespace

	option.Level = 0
	obj.NameSpace.GenRust(option, builder)

	path := option.GetOutputFile("rs")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
	option.NamespaceToPrefix = oldNamespaceToPrefix
}

func rustTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "bool"

	case src.Typo_I8:
		return "i8"

	case src.Typo_U8:
		return "u8"

	case src.Typo_I16:
		return "i16"

	case src.Typo_U16:
		return "u16"

	case src.Typo_I32:
		return "i32"

	case src.Typo_U32:
		return "u32"

	case src.Typo_I64:
		return "i64"

	case src.Typo_U64:
		return "u64"

	case src.Typo_F32:
		return "f32"

	case src.Typo_F64:
		return "f64"

	case src.Typo_Str:
		return "&str"
	}

	return "&str"
}

func (obj *ConstObject) GenRust(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ","
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + ","
			}
		} else {
			str = "\n" + option.GetTabWidth() + "pub const " + genIdentName(option, prefix, constObj) + ": " + rustTypo(constObj.Typo) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genRustStr(prefix string, static string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + static + "pub const " + genIdentName(option, prefix, constObj) + ": &str = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genRustId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "pub const " + genIdentName(option, prefix, constObj) + " : " + rustTypo(constObj.Typo) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * Rust no namespace
 */
func (obj *NameSpace) GenRust(option *src.Options, builder *strings.Builder) {
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
		}
	} else {
		option.PushPrefixes(namespace)

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
		}
	}

	// generate consts
	obj.ConstObject.GenRust(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genRustId(option.GetPrefixes(""), option, builder)
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

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes("")
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "// enum " + getIdentName(option, enumName, enumObj))

			enumObj.GenRust(option.GetPrefixes(""), option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "#[derive(Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Debug, Hash)]\npub enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenRust("", option, builder)
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
			enumObj.genRustStr(option.GetPrefixes(""), "", option, builder)
			option.PopPrefixes()
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenRust(option, builder)
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
