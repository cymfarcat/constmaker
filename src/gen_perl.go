// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"strconv"
	"strings"
)

func (obj *ConstParser) genPerl(option *ConstOption) {
	builder := new(strings.Builder)

	// generate root
	oldNamespaceToPrefix := option.namespaceToPrefix
	option.namespaceToPrefix = true // perl no namespace

	option.level = 0
	obj.NameSpace.genPerl(option, builder)

	path := option.getOutputFile("pl")
	data := builder.String()
	obj.writeFile(path, data, true, "# ", "")

	option.clearPrefixes()
	option.namespaceToPrefix = oldNamespaceToPrefix
}

/*
 * Perl no typo
 */
func perlTypo(typo string) string {
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

func (obj *Object) genPerl(prefix string, option *ConstOption, builder *strings.Builder) {
	for idx, constObj := range obj.Children {

		// doc comment
		if len(constObj.CommentDoc) > 0 {
			builder.WriteString("\n" + option.indentComment(convertCommentDoc(constObj.CommentDoc, "#")))
		}

		str := ""
		if constObj.objType == TypeEnumValue && !(obj.enumToPrefix || option.enumToPrefix) {
			if obj.enumSetValue {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " => " + constObj.Value + ","
			} else {
				str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " => " + strconv.FormatInt(int64(idx), 10) + ","
			}
		} else {
			str = "\n" + option.getTabWidth() + "use constant " + option.genIdentNameObj(prefix, constObj) + " => " + constObj.Value + ";"
		}

		// triple comment
		if len(constObj.CommentTripe) > 0 {
			str += " " + convertCommentTripe(constObj.CommentTripe, "#")
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genPerlStr(prefix string, isEnum bool, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := ""
		if isEnum {
			str = "\n" + option.getTabWidth() + option.genIdentNameObj(prefix, constObj) + " => " + constObj.Value + ";"
		} else {
			str = "\n" + option.getTabWidth() + "use constant " + option.genIdentNameObj(prefix, constObj) + " => " + constObj.Value + ";"
		}

		builder.WriteString(str)
	}
}

func (obj *Object) genPerlId(prefix string, option *ConstOption, builder *strings.Builder) {
	for _, constObj := range obj.genObjs {

		str := "\n" + option.getTabWidth() + "use constant " + option.genIdentNameObj(prefix, constObj) + " => " + constObj.Value + ";"

		builder.WriteString(str)
	}
}

/*
 * Perl no namespace
 */
func (obj *NameSpace) genPerl(option *ConstOption, builder *strings.Builder) {
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
		}
	} else {
		option.pushPrefixes(namespace)

		if obj.namespaceToPrefix || option.namespaceToPrefix {
			option.level = 0
			builder.WriteString(option.getEnterCount(len(obj.CommentDoc)) + "# namespace " + option.getIdentNameObj(option.getPrefixes(""), &obj.Object))
		} else {
		}
	}

	// generate consts
	obj.Object.genPerl(option.getPrefixes(""), option, builder)

	// namespace need gen id
	if len(obj.genObjs) > 0 {
		obj.genPerlId(option.getPrefixes(""), option, builder)
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

			enumObj.genPerl(option.getPrefixes(""), option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genPerlStr(option.getPrefixes(""), false, option, builder)
			}
		} else {
			option.pushPrefixes(enumName)

			builder.WriteString(option.getEnterCount(len(enumObj.CommentDoc)) + option.getTabWidth() + "use constant {")
			option.level++

			enumObj.genPerl(option.getPrefixes(""), option, builder)

			// enum need gen str
			if len(enumObj.genObjs) > 0 {
				enumObj.genPerlStr(option.getPrefixes(""), true, option, builder)
			}
		}

		if enumObj.enumToPrefix || option.enumToPrefix {
			option.popPrefixes()
		} else {
			option.popPrefixes()
			option.level--
			builder.WriteString("\n" + option.getTabWidth() + "};")
		}
	}

	// if option.unqualNamespace, pop rootNode
	if obj.parent == nil && option.rootNamespace {
		option.popPrefixes()
	}

	// recursive generate namespaces
	for _, namespace := range obj.Namespaces {
		namespace.genPerl(option, builder)
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
