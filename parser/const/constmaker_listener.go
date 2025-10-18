// Code generated from ConstMaker.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consts // ConstMaker
import "github.com/antlr4-go/antlr/v4"

// ConstMakerListener is a complete listener for a parse tree produced by ConstMakerParser.
type ConstMakerListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterConst is called when entering the const production.
	EnterConst(c *ConstContext)

	// EnterEnum is called when entering the enum production.
	EnterEnum(c *EnumContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterIdentAssignValue is called when entering the identAssignValue production.
	EnterIdentAssignValue(c *IdentAssignValueContext)

	// EnterIdentAssignValueLoop is called when entering the identAssignValueLoop production.
	EnterIdentAssignValueLoop(c *IdentAssignValueLoopContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// EnterDocComment is called when entering the docComment production.
	EnterDocComment(c *DocCommentContext)

	// EnterTripleComment is called when entering the tripleComment production.
	EnterTripleComment(c *TripleCommentContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitConst is called when exiting the const production.
	ExitConst(c *ConstContext)

	// ExitEnum is called when exiting the enum production.
	ExitEnum(c *EnumContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitIdentAssignValue is called when exiting the identAssignValue production.
	ExitIdentAssignValue(c *IdentAssignValueContext)

	// ExitIdentAssignValueLoop is called when exiting the identAssignValueLoop production.
	ExitIdentAssignValueLoop(c *IdentAssignValueLoopContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)

	// ExitDocComment is called when exiting the docComment production.
	ExitDocComment(c *DocCommentContext)

	// ExitTripleComment is called when exiting the tripleComment production.
	ExitTripleComment(c *TripleCommentContext)
}
