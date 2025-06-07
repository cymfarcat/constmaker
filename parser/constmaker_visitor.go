// Code generated from ConstMaker.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ConstMaker
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ConstMakerParser.
type ConstMakerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ConstMakerParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#rules.
	VisitRules(ctx *RulesContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#const.
	VisitConst(ctx *ConstContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#enum.
	VisitEnum(ctx *EnumContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#enumValue.
	VisitEnumValue(ctx *EnumValueContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#identAssignValue.
	VisitIdentAssignValue(ctx *IdentAssignValueContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#identAssignValueLoop.
	VisitIdentAssignValueLoop(ctx *IdentAssignValueLoopContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#keywords.
	VisitKeywords(ctx *KeywordsContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#docComment.
	VisitDocComment(ctx *DocCommentContext) interface{}

	// Visit a parse tree produced by ConstMakerParser#tripleComment.
	VisitTripleComment(ctx *TripleCommentContext) interface{}
}
