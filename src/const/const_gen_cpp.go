// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenCpp(option *src.Options) {
	builder := new(strings.Builder)

	builder.WriteString("\n#ifndef CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")
	builder.WriteString("#define CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")

	// add stdint.h
	builder.WriteString("\n#include <stdint.h>\n")

	// generate root
	option.Level = 0
	obj.NameSpace.GenCpp(option, builder)

	builder.WriteString("\n\n#endif // CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")

	path := option.GetOutputFile("h")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func cppTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "bool"

	case src.Typo_I8:
		return "int8_t"

	case src.Typo_U8:
		return "uint8_t"

	case src.Typo_I16:
		return "int16_t"

	case src.Typo_U16:
		return "uint16_t"

	case src.Typo_I32:
		return "int32_t"

	case src.Typo_U32:
		return "uint32_t"

	case src.Typo_I64:
		return "int64_t"

	case src.Typo_U64:
		return "uint64_t"

	case src.Typo_F32:
		return "float"

	case src.Typo_F64:
		return "double"

	case src.Typo_Str:
		return "char*"
	}

	return "char*"
}

func (obj *ConstObject) GenCpp(prefix string, option *src.Options, builder *strings.Builder) {
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
		} else if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + "const " + cppTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genCppStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + "const char* " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genCppId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + "const " + cppTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *NameSpace) GenCpp(option *src.Options, builder *strings.Builder) {
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
				builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "namespace " + getIdentName(option, rootName, &obj.ConstObject) + " {")
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
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "namespace " + getIdentName(option, namespace, &obj.ConstObject) + " {")
			option.Level++
		}
	}

	prefix := ""
	if obj.NamespaceToPrefix || option.NamespaceToPrefix {
		prefix = option.GetPrefixes("")
	}

	// generate consts
	obj.ConstObject.GenCpp(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genCppId(prefix, option, builder)
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

			enumObj.GenCpp(enumName, option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			enumName2 := getIdentName(option, enumName, enumObj)
			if option.StdCpp11 && len(enumObj.Typo) > 0 {
				enumName2 = "class " + enumName2 + " : " + cppTypo(enumObj.Typo)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "enum " + enumName2 + " {")
			option.Level++

			enumObj.GenCpp("", option, builder)
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "};")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genCppStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	if option.NestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenCpp(option, builder)
		}

		if obj.parent == nil {
			if option.RootNamespace {
				option.PopPrefixes()
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "};")
				}
			}
		} else {
			option.PopPrefixes()
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "};")
			}
		}
	} else {
		if obj.parent == nil {
			if option.RootNamespace {
				if obj.NamespaceToPrefix || option.NamespaceToPrefix {
				} else {
					option.Level--
					builder.WriteString("\n" + option.GetTabWidth() + "};")
				}
			}
		} else {
			if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			} else {
				option.Level--
				builder.WriteString("\n" + option.GetTabWidth() + "};")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.GenCpp(option, builder)
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
