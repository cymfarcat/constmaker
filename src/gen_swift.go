// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strings"
)

func (obj *ConstParser) genSwift(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	option.level = 0
	obj.NameSpace.genSwift(option, builder)

	path := option.getOutputFile("swift")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func swiftTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "Bool"

	case Typo_I8:
		return "Int8"

	case Typo_U8:
		return "UInt8"

	case Typo_I16:
		return "Int16"

	case Typo_U16:
		return "UInt16"

	case Typo_I32:
		return "Int32"

	case Typo_U32:
		return "UInt32"

	case Typo_I64:
		return "Int64"

	case Typo_U64:
		return "UInt64"

	case Typo_F32:
		return "Float"

	case Typo_F64:
		return "Double"

	case Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *Object) genSwift(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + "case " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value
			} else {
				str = "\n" + option.getTabWidth() + "case " + option.genIdentNameObj(prefix, constObj)
			}
		} else if obj.parent == nil {
			str = "\n" + option.getTabWidth() + "static let " + option.genIdentNameObj(prefix, constObj) + ": " + swiftTypo(constObj.Typo) + " = " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + "static let " + option.genIdentNameObj(prefix, constObj) + ": " + swiftTypo(constObj.Typo) + " = " + constObj.Value
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genSwiftStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "static let " + option.genIdentNameObj(prefix, constObj) + ": String = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *Object) genSwiftId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "static let " + option.genIdentNameObj(prefix, constObj) + ": " + swiftTypo(constObj.Typo) + " = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *NameSpace) genSwift(option *ConstOption, builder *strings.Builder) {
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
				builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "struct " + option.getIdentNameObj(rootName, &obj.Object) + " {")
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
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "struct " + option.getIdentNameObj(namespace, &obj.Object) + " {")
			option.level++
		}
	}

	prefix := ""
	if obj.namespaceToPrefix || option.namespaceToPrefix {
		prefix = option.getPrefixes("")
	}

	// generate consts
	obj.Object.genSwift(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genSwiftId(prefix, option, builder)
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

			enumObj.genSwift(enumName, option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "enum " + option.getIdentNameObj(enumName, enumObj) + ": " + swiftTypo(enumObj.Typo) + " {")
			option.level++

			enumObj.genSwift("", option, builder)
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		}

		// enum need gen str
		if len(enumObj.genObjs) > 0 {
			enumObj.genSwiftStr(enumName, option, builder)
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	if option.nestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genSwift(option, builder)
		}

		if obj.parent == nil {
			if option.rootNamespace {
				option.popPrefixes()
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "}")
				}
			}
		} else {
			option.popPrefixes()
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
			}
		}
	} else {
		if obj.parent == nil {
			if option.rootNamespace {
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "}")
				}
			}
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "}")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genSwift(option, builder)
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
