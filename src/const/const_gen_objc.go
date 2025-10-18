// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenObjC(option *src.Options) {
	builder := new(strings.Builder)
	//builderM := new(strings.Builder)

	builder.WriteString("\n#ifndef CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")
	builder.WriteString("#define CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")

	// add stdint.h
	builder.WriteString("\n#include <stdint.h>\n")

	// generate root
	oldNamespaceToPrefix := option.NamespaceToPrefix
	option.NamespaceToPrefix = true // ObjC no namespace

	option.Level = 0
	obj.NameSpace.GenObjC(option, builder)

	builder.WriteString("\n\n#endif // CONSTMAKER_GENERATED_" + strings.ToUpper(option.FileName) + "_H\n")

	path := option.GetOutputFile("h")
	path = strings.Replace(path, ".h", "_objc.h", -1)
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()

	option.NamespaceToPrefix = oldNamespaceToPrefix
}

func objCTypo(typo string) string {
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
		return "NSString*"
	}

	return "NSString*"
}

func (obj *ConstObject) GenObjC(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.GetEnterIndex(idx) + option.IndentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " = " + constObj.Value
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj)
			}
			if idx < len(obj.Children)-1 {
				str += ","
			}
		} else if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			if constObj.Typo == src.Typo_Str {
				str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " @" + constObj.Value
			} else {
				str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " " + constObj.Value
			}
		} else if constObj.Typo == src.Typo_Str {
			str = "\n" + option.GetTabWidth() + "NSString *const " + genIdentName(option, prefix, constObj) + " = @" + constObj.Value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + "const " + objCTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genObjCStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " @" + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + "NSString *const " + genIdentName(option, prefix, constObj) + " = @" + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genObjCId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if option.MacroDefine || obj.MacroDefine || constObj.MacroDefine {
			str = "\n" + option.GetTabWidth() + "#define " + genIdentName(option, prefix, constObj) + " " + constObj.Value
		} else {
			str = "\n" + option.GetTabWidth() + "const " + objCTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + "_ID = " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

/*
 * ObjC no namespace
 */
func (obj *NameSpace) GenObjC(option *src.Options, builder *strings.Builder) {
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
	obj.ConstObject.GenObjC(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genObjCId(option.GetPrefixes(""), option, builder)
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

			enumObj.GenObjC(enumName, option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "typedef NS_ENUM(" + objCTypo(enumObj.Typo) + ", " + getIdentName(option, enumName, enumObj) + ") {")
			option.Level++

			enumObj.GenObjC("", option, builder)
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "};")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genObjCStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenObjC(option, builder)
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
