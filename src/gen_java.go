// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"fmt"
	"os"
	"strings"
)

func (obj *ConstParser) genJava(option *ConstOption) {
	builder := new(strings.Builder)

	if len(option.javaPackage) > 0 {
		builder.WriteString("\npackage " + option.javaPackage + ";\n")
	}

	// generate root
	option.level = 0
	obj.NameSpace.genJava(option, builder)

	path := option.getOutputFile("java")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.clearPrefixes()
}

func javaTypo(typo string) string {
	switch typo {
	case Typo_Bool:
		return "boolean"

	case Typo_I8:
		return "byte"

	case Typo_U8:
		return "byte"

	case Typo_I16:
		return "short"

	case Typo_U16:
		return "short"

	case Typo_I32:
		return "int"

	case Typo_U32:
		return "int"

	case Typo_I64:
		return "long"

	case Typo_U64:
		return "long"

	case Typo_F32:
		return "float"

	case Typo_F64:
		return "double"

	case Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *Object) genJava(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		value := adjustFloat(constObj.Value, constObj.Typo)

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString(option.getEnterIndex(idx) + option.indentComment(constObj.CommentDoc))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + "(" + value + ")"
				if idx < len(obj.Children)-1 {
					str += ","
				} else {
					str += ";"
				}
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj)
				if idx < len(obj.Children)-1 {
					str += ","
				}
			}
		} else {
			str = "\n" + option.getTabWidth() + "public static final " + javaTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genJavaStr(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "public static final String " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *Object) genJavaId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "public static final " + javaTypo(constObj.Typo) + " " + option.genIdentNameObj(prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * In Java, constants must be declared within a class.
 */
func (obj *NameSpace) genJava(option *ConstOption, builder *strings.Builder) {
	namespace := obj.Ident

	// Java same scope interface name can't same.
	parent := obj.Object.parent
	for {
		if parent == nil {
			break
		}

		if option.genIdentNameObj("", parent) == namespace {
			fmt.Fprintf(os.Stderr, "Java: interface names cannot be duplicated within the same scope, ignore it: %s.\n", obj.Ident)
			option.pushPrefixes(namespace)
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
			option.popPrefixes()
			return
		}

		parent = parent.parent
	}

	if obj.isEmpty() {
		option.pushPrefixes(namespace)
		builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		option.popPrefixes()
		return
	}

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.indentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		// option.rootAsNode=true force do this
		builder.WriteString(option.getEnterCount(1) + option.getTabWidth() + "interface " + option.fileName + " {")
		option.level++
	} else {
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.pushPrefixes(namespace)
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "// namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + option.getTabWidth() + "interface " + option.getIdentNameObj(namespace, &obj.Object) + " {")
			option.level++
		}
	}

	// generate consts
	obj.Object.genJava(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genJavaId(option.getPrefixes(""), option, builder)
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

			enumObj.genJava(option.getPrefixes(""), option, builder)
		} else {
			if enumObj.namespaceToPrefix || option.namespaceToPrefix {
				enumName = option.getPrefixes(enumName)
			}

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "enum " + option.getIdentNameObj(enumName, enumObj) + " {")
			option.level++

			enumObj.genJava("", option, builder)

			if len(enumObj.genObjs) > 0 && enumObj.enumSetValue {
				builder.WriteString("\n")
				enumObj.genJavaStr("", option, builder)
			}
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			if enumObj.enumSetValue {
				builder.WriteString("\n\n" + option.getTabWidth() + "final int value;")
				builder.WriteString("\n" + option.getTabWidth() + option.getIdentNameObj(enumName, enumObj) + "(int value) { this.value = value; }")
			}
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		}

		// enum need gen str
		if (enumObj.enumToPrefix || option.enumToPrefix || !enumObj.enumSetValue) && len(enumObj.genObjs) > 0 {
			enumName = option.joinPrefix(enumObj.prefix, enumObj.Ident)

			option.pushPrefixes(enumName)
			enumObj.genJavaStr(option.getPrefixes(""), option, builder)
			option.popPrefixes()
		}
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genJava(option, builder)
	}

	if obj.parent == nil {
		option.level--
		builder.WriteString("\n" + option.getTabWidth() + "}")
	} else {
		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.popPrefixes()
		} else {
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "}")
		}
	}
}
