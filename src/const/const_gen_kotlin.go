// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenKotlin(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenKotlin(option, builder)

	path := option.GetOutputFile("kt")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func kotlinTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "Boolean"

	case src.Typo_I8:
		return "Byte"

	case src.Typo_U8:
		return "UByte"

	case src.Typo_I16:
		return "Short"

	case src.Typo_U16:
		return "UShort"

	case src.Typo_I32:
		return "Int"

	case src.Typo_U32:
		return "UInt"

	case src.Typo_I64:
		return "Long"

	case src.Typo_U64:
		return "ULong"

	case src.Typo_F32:
		return "Float"

	case src.Typo_F64:
		return "Double"

	case src.Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *ConstObject) GenKotlin(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := adjustFloat(constObj.Value, constObj.Typo)

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + "(" + value + ")"
				if idx < len(obj.Children)-1 {
					str += ","
				} else {
					str += ";"
				}
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj)
				if idx < len(obj.Children)-1 {
					str += ","
				}
			}
		} else {
			str = "\n" + option.GetTabWidth() + "const val " + genIdentName(option, prefix, constObj) + ": " + kotlinTypo(constObj.Typo) + " = " + value
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genKotlinStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "const val " + genIdentName(option, prefix, constObj) + ": String = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genKotlinId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "const val " + genIdentName(option, prefix, constObj) + ": " + kotlinTypo(constObj.Typo) + " = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *NameSpace) GenKotlin(option *src.Options, builder *strings.Builder) {
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

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "object " + getIdentName(option, rootName, &obj.ConstObject) + " {")
				option.Level++
			}
		}
	} else {
		option.PushPrefixes(namespace)

		if !option.NestedDefine {
			namespace = option.GetPrefixes("")
		}

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "object " + getIdentName(option, namespace, &obj.ConstObject) + " {")
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenKotlin(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genKotlinId(prefix, option, builder)
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

			enumObj.GenKotlin(enumName, option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "enum class " + getIdentName(option, enumName, enumObj) + "(val value: " + kotlinTypo(enumObj.Typo) + ") {")
			option.Level++

			enumObj.GenKotlin("", option, builder)
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genKotlinStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	if option.NestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenKotlin(option, builder)
		}

		if obj.parent == nil {
			if option.RootNamespace {
				option.PopPrefixes()
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "}")
				}
			}
		} else {
			option.PopPrefixes()
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")
			}
		}
	} else {
		if obj.parent == nil {
			if option.RootNamespace {
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "}")
				}
			}
		} else {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "}")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenKotlin(option, builder)
		}

		if obj.parent == nil {
			if option.RootNamespace {
				option.PopPrefixes()
			}
		} else {
			option.PopPrefixes()
		}
	}
}
