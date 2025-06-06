// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genRust(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	oldNamespaceToPrefix := option.namespaceToPrefix
	option.namespaceToPrefix = true // rust no namespace

	option.level = 0
	obj.NameSpace.genRust(option, builder)

	path := option.getOutputFile("rs")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
	option.namespaceToPrefix = oldNamespaceToPrefix
}

func rustTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "i8"

	case Typo_U8:
		return "u8"

	case Typo_I16:
		return "i16"

	case Typo_U16:
		return "u16"

	case Typo_I32:
		return "i32"

	case Typo_U32:
		return "u32"

	case Typo_I64:
		return "i64"

	case Typo_U64:
		return "u64"

	case Typo_F32:
		return "f32"

	case Typo_F64:
		return "f64"

	case Typo_Str:
		return "&str"
	}

	return "&str"
}

func (obj *Object) genRust(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ","
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + ","
			}
		} else {
			str = "\n" + option.getTabWidth() + "pub const " + option.genIdentNameObj(prefix, constObj) + ": " + rustTypo(constObj.Typo) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genRustStr(prefix string, static string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + static + "pub const " + option.genIdentNameObj(prefix, constObj) + ": &str = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genRustId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "pub const " + option.genIdentNameObj(prefix, constObj) + " : " + rustTypo(constObj.Typo) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * Rust no namespace
 */
func (obj *NameSpace) genRust(option *ConstOption, builder *strings.Builder) {
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
	obj.Object.genRust(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genRustId(option.getPrefixes(""), option, builder)
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

			enumObj.genRust(option.getPrefixes(""), option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "#[derive(Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Debug, Hash)]\npub enum " + option.getIdentNameObj(enumName, enumObj) + " {")
			option.level++

			enumObj.genRust("", option, builder)
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
			enumObj.genRustStr(option.getPrefixes(""), "", option, builder)
			option.popPrefixes()
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genRust(option, builder)
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
