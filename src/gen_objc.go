// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genObjC(option *ConstOption) {
	builder := new(strings.Builder)
	//builderM := new(strings.Builder)

	builder.WriteString("\n#ifndef CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")
	builder.WriteString("#define CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")

	// add stdint.h
	builder.WriteString("\n#include <stdint.h>\n")

	// generate root
	oldNamespaceToPrefix := option.namespaceToPrefix
	option.namespaceToPrefix = true // ObjC no namespace

	option.level = 0
	obj.NameSpace.genObjC(option, builder)

	builder.WriteString("\n\n#endif // CONSTMAKER_GENERATED_" + strings.ToUpper(option.fileName) + "_H\n")

	path := option.getOutputFile("h")
	path = strings.Replace(path, ".h", "_objc.h", -1)
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()

	option.namespaceToPrefix = oldNamespaceToPrefix
}

func objCTypo(typo string) string {
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
		return "NSString*"
	}

	return "NSString*"
}

func (obj *Object) genObjC(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj)
			}
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			if constObj.Typo == Typo_Str {
				str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " @" + constObj.Value
			} else {
				str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " " + constObj.Value
			}
		} else if constObj.Typo == Typo_Str {
			str = "\n" + option.getTabWidth() + "NSString *const " + option.genIdentNameObj(prefix, constObj) + " = @" + constObj.Value + ";"
		} else {
			str = "\n" + option.getTabWidth() + "const " + objCTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genObjCStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " @" + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "NSString *const " + option.genIdentNameObj(prefix, constObj) + " = @" + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genObjCId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.macroDefine || obj.macroDefine || constObj.macroDefine {
			str = "\n" + option.getTabWidth() + "#define " + option.genIdentNameObj(prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "const " + objCTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + "_ID = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

/*
 * ObjC no namespace
 */
func (obj *NameSpace) genObjC(option *ConstOption, builder *strings.Builder) {
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
	obj.Object.genObjC(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genObjCId(option.getPrefixes(""), option, builder)
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

			enumObj.genObjC(enumName, option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "typedef NS_ENUM(" + objCTypo(enumObj.Typo) + ", " + option.getIdentNameObj(enumName, enumObj) + ") {")
			option.level++

			enumObj.genObjC("", option, builder)
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "};")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genObjCStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genObjC(option, builder)
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
