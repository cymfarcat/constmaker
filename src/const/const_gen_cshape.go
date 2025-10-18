// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenCShape(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenCShape(option, builder)

	path := option.GetOutputFile("cs")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func cshapeTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "bool"

	case src.Typo_I8:
		return "sbyte"

	case src.Typo_U8:
		return "byte"

	case src.Typo_I16:
		return "short"

	case src.Typo_U16:
		return "ushort"

	case src.Typo_I32:
		return "int"

	case src.Typo_U32:
		return "uint"

	case src.Typo_I64:
		return "long"

	case src.Typo_U64:
		return "ulong"

	case src.Typo_F32:
		return "float"

	case src.Typo_F64:
		return "double"

	case src.Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *ConstObject) GenCShape(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := adjustFloat(constObj.Value, constObj.Typo)

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + value + ","
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + ","
			}
		} else {
			str = "\n" + option.GetTabWidth() + "public const " + cshapeTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genCShapeStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "public const string " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genCShapeId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "public const " + cshapeTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * In C#, constants must be declared within a class or struct.
 */
func (obj *NameSpace) GenCShape(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		rootName := option.GetIdentName(option.GetRootName())

		// option.rootAsNode=true force do this
		builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "public static class " + getIdentName(option, rootName, &obj.ConstObject) + " {")
		option.Level++
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.PushPrefixes(namespace)
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "public static class " + getIdentName(option, namespace, &obj.ConstObject) + " {")
			option.Level++
		}
	}

	// generate consts
	obj.ConstObject.GenCShape(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genCShapeId(option.GetPrefixes(""), option, builder)
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

			enumObj.GenCShape(option.GetPrefixes(""), option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "public enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenCShape("", option, builder)
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
			enumObj.genCShapeStr(option.GetPrefixes(""), option, builder)
			option.PopPrefixes()
		}
	}

	if option.NestedDefine || obj.NamespaceToPrefix || option.NamespaceToPrefix {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenCShape(option, builder)
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
	} else {
		if obj.parent == nil {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		} else {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				option.Level = 0
				//option.PopPrefixes()
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenCShape(option, builder)
		}

		if obj.parent == nil {
			if option.RootNamespace {
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
					option.PopPrefixes()
				}
			}
		} else {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				option.PopPrefixes()
			}
		}
	}
}
