// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"fmt"
	"os"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenJava(option *src.Options) {
	builder := new(strings.Builder)

	if len(option.JavaPackage) > 0 {
		builder.WriteString("\npackage " + option.JavaPackage + ";\n")
	}

	// generate root
	option.Level = 0
	obj.NameSpace.GenJava(option, builder)

	path := option.GetOutputFile("java")
	data := builder.String()
	obj.writeFile(path, data, true, "", "")

	option.ClearPrefixes()
}

func javaTypo(typo string) string {
	switch typo {
	case src.Typo_Bool:
		return "boolean"

	case src.Typo_I8:
		return "byte"

	case src.Typo_U8:
		return "byte"

	case src.Typo_I16:
		return "short"

	case src.Typo_U16:
		return "short"

	case src.Typo_I32:
		return "int"

	case src.Typo_U32:
		return "int"

	case src.Typo_I64:
		return "long"

	case src.Typo_U64:
		return "long"

	case src.Typo_F32:
		return "float"

	case src.Typo_F64:
		return "double"

	case src.Typo_Str:
		return "String"
	}

	return "String"
}

func (obj *ConstObject) GenJava(prefix string, option *src.Options, builder *strings.Builder) {
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
			str = "\n" + option.GetTabWidth() + "public static final " + javaTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "//")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genJavaStr(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "public static final String " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genJavaId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "public static final " + javaTypo(constObj.Typo) + " " + genIdentName(option, prefix, constObj) + " = " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * In Java, constants must be declared within a class.
 */
func (obj *NameSpace) GenJava(option *src.Options, builder *strings.Builder) {
	namespace := obj.Ident

	// Java same scope interface name can't same.
	parent := obj.ConstObject.parent
	for {
		if parent == nil {
			break
		}

		if genIdentName(option, "", parent) == namespace {
			fmt.Fprintf(os.Stderr, "Java: interface names cannot be duplicated within the same scope, ignore it: %s.\n", obj.Ident)
			option.PushPrefixes(namespace)
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
			option.PopPrefixes()
			return
		}

		parent = parent.parent
	}

	if obj.isEmpty() {
		option.PushPrefixes(namespace)
		builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		option.PopPrefixes()
		return
	}

	// doc comment
	if len(obj.CommentDoc) > 0 {
		builder.WriteString("\n\n" + option.IndentComment(obj.CommentDoc))
	}

	if obj.parent == nil {
		// option.rootAsNode=true force do this
		builder.WriteString(option.GetEnterCount(1) + option.GetTabWidth() + "interface " + option.FileName + " {")
		option.Level++
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.PushPrefixes(namespace)
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "// namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + option.GetTabWidth() + "interface " + getIdentName(option, namespace, &obj.ConstObject) + " {")
			option.Level++
		}
	}

	// generate consts
	obj.ConstObject.GenJava(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genJavaId(option.GetPrefixes(""), option, builder)
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

			enumObj.GenJava(option.GetPrefixes(""), option, builder)
		} else {
			if enumObj.NamespaceToPrefix || option.NamespaceToPrefix {
				enumName = option.GetPrefixes(enumName)
			}

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "enum " + getIdentName(option, enumName, enumObj) + " {")
			option.Level++

			enumObj.GenJava("", option, builder)

			if len(enumObj.genObjs) > 0 && enumObj.enumSetValue {
				builder.WriteString("\n")
				enumObj.genJavaStr("", option, builder)
			}
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			if enumObj.enumSetValue {
				builder.WriteString("\n\n" + option.GetTabWidth() + "final int value;")
				builder.WriteString("\n" + option.GetTabWidth() + getIdentName(option, enumName, enumObj) + "(int value) { this.value = value; }")
			}
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}

		// enum need gen str
		if (enumObj.EnumToPrefix || option.EnumToPrefix || !enumObj.enumSetValue) && len(enumObj.genObjs) > 0 {
			enumName = option.JoinPrefix(enumObj.Prefix, enumObj.Ident)

			option.PushPrefixes(enumName)
			enumObj.genJavaStr(option.GetPrefixes(""), option, builder)
			option.PopPrefixes()
		}
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenJava(option, builder)
	}

	if obj.parent == nil {
		option.Level--
		builder.WriteString("\n" + option.GetTabWidth() + "}")
	} else {
		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.PopPrefixes()
		} else {
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "}")
		}
	}
}
