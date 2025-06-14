// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package src

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cymfarcat/constmaker/parser"
)

// ConstParser default is NameSpace
type ConstParser struct {
	NameSpace
	parser.BaseConstMakerListener

	XMLName xml.Name `xml:"" json:"-"`

	option        *ConstOption
	objectStack   []*Object
	currentObject *Object
}

func NewConstParser(option *ConstOption) *ConstParser {
	parser := &ConstParser{}
	parser.option = option

	parser.initialize(nil)
	parser.objectStack = make([]*Object, 0)

	// push this to stack
	parser.pushStack(&parser.Object)

	return parser
}

func (obj *ConstParser) ParseFile() int {
	content, err := os.ReadFile(obj.option.inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file: %s, error: %s\n", obj.option.inputFile, err)
		return -1
	}

	is := antlr.NewInputStream(string(content))
	lexer := parser.NewConstMakerLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewConstMakerParser(stream)
	tree := parser.File()

	if obj.option.verbose {
		str := tree.ToStringTree([]string{}, parser)
		str = strings.ReplaceAll(str, "const ", "\nconst ")
		str = strings.ReplaceAll(str, "enum ", "\nenum ")
		str = strings.ReplaceAll(str, "namespace ", "\nnamespace ")
		fmt.Println(str)
	}

	antlr.ParseTreeWalkerDefault.Walk(obj, tree)

	return 0
}

func (obj *ConstParser) GenFile() {
	if obj.option.genAll || obj.option.genCpp {
		obj.genCpp(obj.option)
	}

	if obj.option.genAll || obj.option.genCShape {
		obj.genCShape(obj.option)
	}

	if obj.option.genAll || obj.option.genDart {
		obj.genDart(obj.option)
	}

	if obj.option.genAll || obj.option.genGo {
		obj.genGo(obj.option)
	}

	if obj.option.genAll || obj.option.genJava {
		obj.genJava(obj.option)
	}

	if obj.option.genAll || obj.option.genJavaScript {
		obj.genJavaScript(obj.option)
	}

	if obj.option.genAll || obj.option.genJson {
		obj.genJson(obj.option)
	}

	if obj.option.genAll || obj.option.genKotlin {
		obj.genKotlin(obj.option)
	}

	if obj.option.genAll || obj.option.genMarkdown {
		obj.genMarkdown(obj.option)
	}

	if obj.option.genAll || obj.option.genObjC {
		obj.genObjC(obj.option)
	}

	if obj.option.genAll || obj.option.genPascal {
		obj.genPascal(obj.option)
	}

	if obj.option.genAll || obj.option.genPerl {
		obj.genPerl(obj.option)
	}

	if obj.option.genAll || obj.option.genPhp {
		obj.genPhp(obj.option)
	}

	if obj.option.genAll || obj.option.genPython {
		obj.genPython(obj.option)
	}

	if obj.option.genAll || obj.option.genQml {
		obj.genQml(obj.option)
	}

	if obj.option.genAll || obj.option.genRuby {
		obj.genRuby(obj.option)
	}

	if obj.option.genAll || obj.option.genRust {
		obj.genRust(obj.option)
	}

	if obj.option.genAll || obj.option.genSwift {
		obj.genSwift(obj.option)
	}

	if obj.option.genAll || obj.option.genText {
		obj.genText(obj.option)
	}

	if obj.option.genAll || obj.option.genTypeScript {
		obj.genTypeScript(obj.option)
	}

	if obj.option.genAll || obj.option.genXml {
		obj.genXml(obj.option)
	}
}

func (obj *ConstParser) writeFile(path string, data string, headComment bool, startComment string, endComment string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create file: %s, error: %s\n", path, err)
		return
	}
	defer file.Close()

	if headComment {
		if len(startComment) == 0 && len(endComment) == 0 {
			startComment = "// "
		}

		_, err = file.WriteString(startComment + "This file is generated by ConstMaker, DON'T MODIFY." + endComment + "\n")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't write file: %s, error: %s\n", path, err)
		}
	}

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't write file: %s, error: %s\n", path, err)
	}
}

func (obj *ConstParser) pushStack(object *Object) {
	obj.objectStack = append(obj.objectStack, object)
}

func (obj *ConstParser) popStack() *Object {
	if len(obj.objectStack) == 0 {
		fmt.Println("popStack internal error: stack.isEmpty")
		return nil
	}

	topStack := obj.objectStack[len(obj.objectStack)-1]
	obj.objectStack = obj.objectStack[:len(obj.objectStack)-1]
	return topStack
}

func (obj *ConstParser) topStack() *Object {
	if len(obj.objectStack) == 0 {
		fmt.Println("topStack internal error: stack.isEmpty")
		return nil
	}

	return obj.objectStack[len(obj.objectStack)-1]
}

// EnterFile is called when production file is entered.
func (obj *ConstParser) EnterFile(ctx *parser.FileContext) {
}

// ExitFile is called when production file is exited.
func (obj *ConstParser) ExitFile(ctx *parser.FileContext) {
}

// EnterRules is called when production rules is entered.
func (obj *ConstParser) EnterRules(ctx *parser.RulesContext) {}

// ExitRules is called when production rules is exited.
func (obj *ConstParser) ExitRules(ctx *parser.RulesContext) {}

// EnterConst is called when production const is entered.
func (obj *ConstParser) EnterConst(ctx *parser.ConstContext) {
	topStack := obj.topStack()
	constObj := NewObject(topStack, TypeObj)

	obj.currentObject = constObj
}

// ExitConst is called when production const is exited.
func (obj *ConstParser) ExitConst(ctx *parser.ConstContext) {
	obj.currentObject.genValue(obj.option)

	topStack := obj.topStack()
	if topStack != nil {
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		topStack.addObject(obj.currentObject, line, column)
	}

	obj.currentObject = nil
}

// EnterEnum is called when production enum is entered.
func (obj *ConstParser) EnterEnum(ctx *parser.EnumContext) {
	topStack := obj.topStack()
	enum := NewObject(topStack, TypeEnum)

	obj.pushStack(enum)
	obj.currentObject = enum

	// add enum to namespace
	//if topStack != nil && topStack.objType == TypeNamespace {
	//	line := ctx.GetStart().GetLine()
	//	column := ctx.GetStart().GetColumn()
	//	(*NameSpace)(unsafe.Pointer(topStack)).addEnum(enum, line, column)
	//}
}

// ExitEnum is called when production enum is exited.
func (obj *ConstParser) ExitEnum(ctx *parser.EnumContext) {
	topStack := obj.topStack()
	if topStack != nil {
		topStack.genValue(obj.option)

		// add enum to namespace
		parent := topStack.parent
		if parent != nil && parent.objType == TypeNamespace && topStack.objType == TypeEnum {
			line := ctx.GetStart().GetLine()
			column := ctx.GetStart().GetColumn()

			(*NameSpace)(unsafe.Pointer(parent)).addEnum(topStack, line, column)
		}
	}

	obj.popStack()
	obj.currentObject = nil
}

// EnterNamespace is called when production namespace is entered.
func (obj *ConstParser) EnterNamespace(ctx *parser.NamespaceContext) {
	topStack := obj.topStack()
	namespace := NewNameSpace(topStack)

	// add namespace
	//if topStack != nil && topStack.objType == TypeNamespace {
	//	line := ctx.GetStart().GetLine()
	//	column := ctx.GetStart().GetColumn()
	//	(*NameSpace)(unsafe.Pointer(topStack)).addNamespace(namespace, line, column)
	//}

	obj.pushStack(&namespace.Object)
	obj.currentObject = &namespace.Object
}

// ExitNamespace is called when production namespace is exited.
func (obj *ConstParser) ExitNamespace(ctx *parser.NamespaceContext) {
	topStack := obj.topStack()
	if topStack != nil {
		topStack.genValue(obj.option)

		parent := topStack.parent
		if parent != nil && parent.objType == TypeNamespace && topStack.objType == TypeNamespace {
			line := ctx.GetStart().GetLine()
			column := ctx.GetStart().GetColumn()
			(*NameSpace)(unsafe.Pointer(parent)).addNamespace((*NameSpace)(unsafe.Pointer(topStack)), line, column)
		}
	}

	obj.popStack()
	obj.currentObject = nil
}

// EnterType is called when production type is entered.
func (obj *ConstParser) EnterType(ctx *parser.TypeContext) {
	obj.currentObject.Typo = ctx.GetText()
}

// ExitType is called when production type is exited.
func (obj *ConstParser) ExitType(ctx *parser.TypeContext) {
}

// EnterEnumValue is called when production enumValue is entered.
func (obj *ConstParser) EnterEnumValue(ctx *parser.EnumValueContext) {
	topStack := obj.topStack()
	enumValue := NewObject(topStack, TypeEnumValue)
	obj.currentObject = enumValue
}

// ExitEnumValue is called when production enumValue is exited.
func (obj *ConstParser) ExitEnumValue(ctx *parser.EnumValueContext) {
	obj.currentObject.genValue(obj.option)

	topStack := obj.topStack()
	if topStack != nil {
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		topStack.addObject(obj.currentObject, line, column)
	}

	obj.currentObject = nil
}

// EnterIdentAssignValue is called when production identAssignValue is entered.
func (obj *ConstParser) EnterIdentAssignValue(ctx *parser.IdentAssignValueContext) {
	obj.pushStack(obj.currentObject)
	obj.currentObject = NewObject(obj.currentObject, TypeProperty)
}

// ExitIdentAssignValue is called when production identAssignValue is exited.
func (obj *ConstParser) ExitIdentAssignValue(ctx *parser.IdentAssignValueContext) {
	topStack := obj.popStack()
	if topStack != nil {
		topStack.addOption(obj.currentObject)
	}

	obj.currentObject = topStack
}

// EnterIdentAssignValueLoop is called when production identAssignValueLoop is entered.
func (obj *ConstParser) EnterIdentAssignValueLoop(ctx *parser.IdentAssignValueLoopContext) {}

// ExitIdentAssignValueLoop is called when production identAssignValueLoop is exited.
func (obj *ConstParser) ExitIdentAssignValueLoop(ctx *parser.IdentAssignValueLoopContext) {}

// ExitOption is called when production option is entered.
func (obj *ConstParser) EnterOption(ctx *parser.OptionContext) {}

// ExitOption is called when production option is exited.
func (obj *ConstParser) ExitOption(ctx *parser.OptionContext) {}

// EnterValue is called when production value is entered.
func (obj *ConstParser) EnterValue(ctx *parser.ValueContext) {
	obj.currentObject.Value = ctx.GetText()
}

// ExitValue is called when production value is exited.
func (obj *ConstParser) ExitValue(ctx *parser.ValueContext) {}

// EnterInteger is called when production integer is entered.
func (obj *ConstParser) EnterInteger(ctx *parser.IntegerContext) {
	obj.currentObject.Value = ctx.GetText()
}

// ExitInteger is called when production integer is exited.
func (obj *ConstParser) ExitInteger(ctx *parser.IntegerContext) {}

// EnterIdent is called when production ident is entered.
func (obj *ConstParser) EnterIdent(ctx *parser.IdentContext) {
	if len(obj.currentObject.Ident) == 0 {
		obj.currentObject.Ident = ctx.GetText()
	} else {
		// NEXT: check is references
		obj.currentObject.Value = ctx.GetText()
	}
}

// ExitIdent is called when production ident is exited.
func (obj *ConstParser) ExitIdent(ctx *parser.IdentContext) {
}

// EnterKeywords is called when production keywords is entered.
func (obj *ConstParser) EnterKeywords(ctx *parser.KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (obj *ConstParser) ExitKeywords(ctx *parser.KeywordsContext) {}

// EnterDocComment is called when production docComment is entered.
func (obj *ConstParser) EnterDocComment(ctx *parser.DocCommentContext) {
	obj.currentObject.CommentDoc = ctx.GetText()
}

// ExitDocComment is called when production docComment is exited.
func (obj *ConstParser) ExitDocComment(ctx *parser.DocCommentContext) {}

// EnterTripleComment is called when production tripleComment is entered.
func (obj *ConstParser) EnterTripleComment(ctx *parser.TripleCommentContext) {
	obj.currentObject.CommentTripe = ctx.GetText()
}

// ExitTripleComment is called when production tripleComment is exited.
func (obj *ConstParser) ExitTripleComment(ctx *parser.TripleCommentContext) {}
