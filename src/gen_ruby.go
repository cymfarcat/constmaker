// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strconv"
	"strings"
)

func (obj *ConstParser) genRuby(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	option.level = 0
	obj.NameSpace.genRuby(option, builder)

	path := option.getOutputFile("rb")
	data := builder.String()
	obj.writeFile(path, data, true, "# ", "")

	option.clearPrefixes()
}

/*
 * Ruby no typo
 */
func rubyTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "bool"

	case Typo_I8:
		return "int"

	case Typo_U8:
		return "int"

	case Typo_I16:
		return "int"

	case Typo_U16:
		return "int"

	case Typo_I32:
		return "int"

	case Typo_U32:
		return "int"

	case Typo_I64:
		return "int"

	case Typo_U64:
		return "int"

	case Typo_F32:
		return "float"

	case Typo_F64:
		return "double"

	case Typo_Str:
		return "str"
	}

	return "str"
}

func (obj *Object) genRuby(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := constObj.Value

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.indentComment(convertCommentDoc(constObj.CommentDoc, "#")))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + value
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + strconv.FormatInt(int64(idx), 10)
			}
		} else {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + value
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "#")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genRubyStr(prefix string, isEnum bool, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if isEnum {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value
		} else {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genRubyId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value

		builder.WriteString(str)
	}
}

func (obj *NameSpace) genRuby(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(convertCommentDoc(obj.CommentDoc, "#")))
	}

	if obj.parent == nil {
		if option.rootNamespace {
			rootName := option.getIdentName(option.getRootName())

			option.level = 0
			option.pushPrefixes(rootName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "module " + option.getIdentNameObj(rootName, &obj.Object))
				option.level++
			}
		}
	} else {
		option.pushPrefixes(namespace)

		if !option.nestedDefine {
			namespace = option.getPrefixes("")
		}

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "# namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "module " + option.getIdentNameObj(namespace, &obj.Object))
			option.level++
		}
	}

	prefix := ""
	if obj.namespaceToPrefix || option.namespaceToPrefix {
		prefix = option.getPrefixes("")
	}

	// generate consts
	obj.Object.genRuby(prefix, option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genRubyId(prefix, option, builder)
	}

	// generate enums
	for _, enumObj := range obj.Enums {
		enumName := enumObj.Ident

		// doc comment
		if len(enumObj.CommentDoc) > 0 {
			builder.WriteString("\n\n" + option.indentComment(convertCommentDoc(enumObj.CommentDoc, "#")))
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.pushPrefixes(enumName)

			if obj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes("")
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "# enum " + option.getIdentNameObj(enumName, enumObj))

			enumObj.genRuby(enumName, option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genRubyStr(enumName, false, option, builder)
			}
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "module " + option.getIdentNameObj(enumName, enumObj))
			option.level++

			enumObj.genRuby("", option, builder)

			if len(enumObj.genObjs) > 0 {
				enumObj.genRubyStr(enumName, true, option, builder)
			}
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "end")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	if option.nestedDefine {
		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genRuby(option, builder)
		}

		if obj.parent == nil {
			if option.rootNamespace {
				option.popPrefixes()
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "end")
				}
			}
		} else {
			option.popPrefixes()
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "end")
			}
		}
	} else {
		if obj.parent == nil {
			if option.rootNamespace {
				if obj.namespaceToPrefix || option.namespaceToPrefix {
				} else {
					option.level--
					builder.WriteString("\n" + option.getTabWidth() + "end")
				}
			}
		} else {
			if obj.namespaceToPrefix || option.namespaceToPrefix {
			} else {
				option.level--
				builder.WriteString("\n" + option.getTabWidth() + "end")
			}
		}

		// recursive generate namespaces
		for _, namespace := range obj.Namespaces {
			namespace.genRuby(option, builder)
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
