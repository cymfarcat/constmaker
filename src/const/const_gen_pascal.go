// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

/*
 *The first programming language I learned still dwells‌ in my memory.
 */
func (obj *ConstParser) GenPascal(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenPascal(option, builder)

	path := option.GetOutputFile("pas")
	data := builder.String()
	obj.writeFile(path, data, true, "(* ", " *)")

	option.ClearPrefixes()
}

func pascalTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "Boolean"

	case src.Typo_I8:
		return "Shortint"

	case src.Typo_U8:
		return "Byte"

	case src.Typo_I16:
		return "Smallint"

	case src.Typo_U16:
		return "Word"

	case src.Typo_I32:
		return "Longint" //Integer

	case src.Typo_U32:
		return "LongWord"

	case src.Typo_I64:
		return "Int64"

	case src.Typo_U64:
		return "QWord"

	case src.Typo_F32:
		return "Real"

	case src.Typo_F64:
		return "Double"

	case src.Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *ConstObject) GenPascal(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.IndentComment(src.ConvertCommentDocToPascal(constObj.CommentDoc)))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj)
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else {
			str = "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + ": " + pascalTypo(constObj.Typo) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTripeToPascal(constObj.CommentTriple)
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genPascalStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + ": String = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genPascalId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + ": " + pascalTypo(constObj.Typo) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *NameSpace) GenPascal(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(src.ConvertCommentDocToPascal(obj.CommentDoc)))
	}

	if obj.parent == nil {
		if option.RootNamespace {
			rootName := option.GetIdentName(option.GetRootName())

			option.Level = 0
			option.PushPrefixes(rootName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "type")
				option.Level++

				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + getIdentName(option, rootName, &obj.ConstObject) + " = class")
				option.Level++
			}
		}
	} else {
		option.PushPrefixes(namespace)

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "(* namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject) + " *)")
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "type")
			option.Level++

			builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + getIdentName(option, namespace, &obj.ConstObject) + " = class")
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenPascal(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genPascalId(prefix, option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.IndentComment(src.ConvertCommentDocToPascal(enumObj.CommentDoc)))
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PushPrefixes(enumName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes("")
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "(* enum " + getIdentName(option, enumName, enumObj) + " *)")

			enumObj.GenPascal(enumName, option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "type " + getIdentName(option, enumName, enumObj) + " = (")
			option.Level++

			enumObj.GenPascal("", option, builder)
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + ");")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genPascalStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenPascal(option, builder)
	}

	if obj.parent == nil {
		if option.RootNamespace {
			option.PopPrefixes()
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "end;")
			}
		}
	} else {
		option.PopPrefixes()
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "end;")

			option.Level--
		}
	}
}
