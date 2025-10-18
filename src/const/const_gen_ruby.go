// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strconv"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenRuby(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenRuby(option, builder)

	path := option.GetOutputFile("rb")
	data := builder.String()
	obj.writeFile(path, data, true, "# ", "")

	option.ClearPrefixes()
}

/*
 * Ruby no typo
 */
func rubyTypo(typo string) string {
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
		return "float"

	case src.Typo_F64:
		return "double"

	case src.Typo_Str:
		return "str"
	}

	return "str"
}

func (obj *ConstObject) GenRuby(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := constObj.Value

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.IndentComment(src.ConvertCommentDoc(constObj.CommentDoc, "#")))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + value
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + strconv.FormatInt(int64(idx), 10)
			}
		} else {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + value
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "#")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genRubyStr(prefix string, isEnum bool, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if isEnum {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genRubyId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *NameSpace) GenRuby(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(src.ConvertCommentDoc(obj.CommentDoc, "#")))
	}

	if obj.parent == nil {
		if option.RootNamespace {
			rootName := option.GetIdentName(option.GetRootName())

			option.Level = 0
			option.PushPrefixes(rootName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "module " + getIdentName(option, rootName, &obj.ConstObject))
				option.Level++
			}
		}
	} else {
		option.PushPrefixes(namespace)

		if !option.NestedDefine {
			namespace = option.GetPrefixes("")
		}

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "# namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "module " + getIdentName(option, namespace, &obj.ConstObject))
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenRuby(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genRubyId(prefix, option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.IndentComment(src.ConvertCommentDoc(enumObj.CommentDoc, "#")))
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PushPrefixes(enumName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes("")
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "# enum " + getIdentName(option, enumName, enumObj))

			enumObj.GenRuby(enumName, option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genRubyStr(enumName, false, option, builder)
			}
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "module " + getIdentName(option, enumName, enumObj))
			option.Level++

			enumObj.GenRuby("", option, builder)

			if len(enumObj.genObjs) > 0 {
				enumObj.genRubyStr(enumName, true, option, builder)
			}
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "end")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	if option.NestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenRuby(option, builder)
		}

		if obj.parent == nil {
			if option.RootNamespace {
				option.PopPrefixes()
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "end")
				}
			}
		} else {
			option.PopPrefixes()
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "end")
			}
		}
	} else {
		if obj.parent == nil {
			if option.RootNamespace {
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "end")
				}
			}
		} else {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "end")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenRuby(option, builder)
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
