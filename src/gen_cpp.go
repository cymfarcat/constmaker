// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genCpp(option *ConstOption) {
	builder := new(strings.Builder)

	builder.WriteString("\n#ifndef CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")
	builder.WriteString("#define CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")

	// add stdint.h
	builder.WriteString("\n#include <stdint.h>\n")

	// generate root
	option.level = 0
	obj.NameSpace.genCpp(option, builder)

	builder.WriteString("\n\n#endif // CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")

	path := option.getOutputFile("h")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func cppTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "int8_t"

	case Typo_U8:
		return "uint8_t"

	case Typo_I16:
		return "int16_t"

	case Typo_U16:
		return "uint16_t"

	case Typo_I32:
		return "int32_t"

	case Typo_U32:
		return "uint32_t"

	case Typo_I64:
		return "int64_t"

	case Typo_U64:
		return "uint64_t"

	case Typo_F32:
		return "float"

	case Typo_F64:
		return "double"

	case Typo_Str:
		return "char*"
	}

	return "char*"
}

func (obj *Object) genCpp(prefix string, option *ConstOption, builder *strings.Builder) {
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
		} else if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "const " + cppTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genCppStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "const char* " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genCppId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "const " + cppTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *NameSpace) genCpp(option *ConstOption, builder *strings.Builder) {
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

			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "namespace " + option.getIdentNameObj(rootName, &obj.Object) + " {")
				option.level++
			}
		}
	} else {
		option.pushPrefixes(namespace)

		if !option.nestedDefine {
			namespace = option.getPrefixes("")
		}

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "namespace " + option.getIdentNameObj(namespace, &obj.Object) + " {")
			option.level++
		}
	}

	prefix := ""
	if obj.namespaceToPrefix || option.namespaceToPrefix {
		prefix = option.getPrefixes("")
	}

	// generate consts
	obj.Object.genCpp(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genCppId(prefix, option, builder)
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

			enumObj.genCpp(enumName, option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			enumName2 := option.getIdentNameObj(enumName, enumObj)
			if option.stdCpp11 && len(enumObj.Typo) > 0 {
				enumName2 = "class " + enumName2 + " : " + cppTypo(enumObj.Typo)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "enum " + enumName2 + " {")
			option.level++

			enumObj.genCpp("", option, builder)
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "};")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genCppStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	if option.nestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genCpp(option, builder)
		}

		if obj.parent == nil {
			if option.rootNamespace {
				option.popPrefixes()
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "};")
				}
			}
		} else {
			option.popPrefixes()
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "};")
			}
		}
	} else {
		if obj.parent == nil {
			if option.rootNamespace {
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "};")
				}
			}
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "};")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genCpp(option, builder)
		}

		if obj.parent == nil {
			if option.rootNamespace {
				option.popPrefixes()
			}
		} else {
			option.popPrefixes()
		}
	}
}
