// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenTypeScript(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	option.Level = 0
	obj.NameSpace.GenTypeScript(option, builder)

	path := option.GetOutputFile("ts")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

/*
 * TypeScript no typo
 */
func typeScriptTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "boolean"

	case src.Typo_I8:
		return "number"

	case src.Typo_U8:
		return "number"

	case src.Typo_I16:
		return "number"

	case src.Typo_U16:
		return "number"

	case src.Typo_I32:
		return "number"

	case src.Typo_U32:
		return "number"

	case src.Typo_I64:
		return "number"

	case src.Typo_U64:
		return "number"

	case src.Typo_F32:
		return "float"

	case src.Typo_F64:
		return "float"

	case src.Typo_Str:
		return "string"
	}

	return "string"
}

func (obj *ConstObject) GenTypeScript(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := constObj.Value

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
			str = "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genTypeScriptStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if obj.EnumToPrefix || option.EnumToPrefix {
			str = "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ","
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genTypeScriptId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "const " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *NameSpace) GenTypeScript(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		if option.RootNamespace {
			rootName := option.GetIdentName(option.GetRootName())

			option.PushPrefixes(rootName)

			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "export namespace " + getIdentName(option, rootName, &obj.ConstObject) + " {")
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
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "export namespace " + getIdentName(option, namespace, &obj.ConstObject) + " {")
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenTypeScript(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genTypeScriptId(prefix, option, builder)
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

			enumObj.GenTypeScript(enumName, option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genTypeScriptStr(enumName, option, builder)
			}
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "export enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenTypeScript("", option, builder)

			if len(enumObj.genObjs) > 0 {
				builder.WriteString("\n")
				enumObj.genTypeScriptStr("", option, builder)
			}
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	if option.NestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenTypeScript(option, builder)
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
			namespace.GenTypeScript(option, builder)
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
