// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package consts

import (
	"strconv"
	"strings"

	src "github.com/cymfarcat/constmaker/src"
)

func (obj *ConstParser) GenPerl(option *src.Options) {
	builder := new(strings.Builder)

	// generate root
	oldNamespaceToPrefix := option.NamespaceToPrefix
	option.NamespaceToPrefix = true // perl no namespace

	option.Level = 0
	obj.NameSpace.GenPerl(option, builder)

	path := option.GetOutputFile("pl")
	data := builder.String()
	obj.writeFile(path, data, true, "# ", "")

	option.ClearPrefixes()
	option.NamespaceToPrefix = oldNamespaceToPrefix
}

/*
 * Perl no typo
 */
func perlTypo(typo string) string {
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

func (obj *ConstObject) GenPerl(prefix string, option *src.Options, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.IndentComment(src.ConvertCommentDoc(constObj.CommentDoc, "#")))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.EnumToPrefix || option.EnumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " => " + constObj.Value + ","
			} else {
				str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " => " + strconv.FormatInt(int64(idx), 10) + ","
			}
		} else {
			str = "\n" + option.GetTabWidth() + "use constant " + genIdentName(option, prefix, constObj) + " => " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTriple) > 0 {
			str += " " + src.ConvertCommentTriple(constObj.CommentTriple, "#")
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genPerlStr(prefix string, isEnum bool, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if isEnum {
			str = "\n" + option.GetTabWidth() + genIdentName(option, prefix, constObj) + " => " + constObj.Value + ";"
		} else {
			str = "\n" + option.GetTabWidth() + "use constant " + genIdentName(option, prefix, constObj) + " => " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *ConstObject) genPerlId(prefix string, option *src.Options, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.GetTabWidth() + "use constant " + genIdentName(option, prefix, constObj) + " => " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * Perl no namespace
 */
func (obj *NameSpace) GenPerl(option *src.Options, builder *strings.Builder) {
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
		}
	} else {
		option.PushPrefixes(namespace)

		if obj.NamespaceToPrefix || option.NamespaceToPrefix {
			option.Level = 0
			builder.WriteString(option.GetEnterCount(len(obj.CommentDoc)) + "# namespace " + getIdentName(option, option.GetPrefixes(""), &obj.ConstObject))
		} else {
		}
	}

	// generate consts
	obj.ConstObject.GenPerl(option.GetPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genPerlId(option.GetPrefixes(""), option, builder)
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

			enumObj.GenPerl(option.GetPrefixes(""), option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genPerlStr(option.GetPrefixes(""), false, option, builder)
			}
		} else {
			option.PushPrefixes(enumName)

			builder.WriteString(option.GetEnterCount(len(enumObj.CommentDoc)) + option.GetTabWidth() + "use constant {")
			option.Level++

			enumObj.GenPerl(option.GetPrefixes(""), option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genPerlStr(option.GetPrefixes(""), true, option, builder)
			}
		}

		if enumObj.EnumToPrefix || option.EnumToPrefix {
			option.PopPrefixes()
		} else {
			option.PopPrefixes()
			option.Level--
			builder.WriteString("\n" + option.GetTabWidth() + "};")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.RootNamespace {
		option.PopPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.GenPerl(option, builder)
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
