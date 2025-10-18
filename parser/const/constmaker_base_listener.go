// Code generated from ConstMaker.g4 by ANTLR 4.13.2. DO NOT EDIT.

package consts // ConstMaker
import "github.com/antlr4-go/antlr/v4"

// BaseConstMakerListener is a complete listener for a parse tree produced by ConstMakerParser.
type BaseConstMakerListener struct{}

var _ ConstMakerListener = &BaseConstMakerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConstMakerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConstMakerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConstMakerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConstMakerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseConstMakerListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseConstMakerListener) ExitFile(ctx *FileContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseConstMakerListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseConstMakerListener) ExitRules(ctx *RulesContext) {}

// EnterConst is called when production const is entered.
func (s *BaseConstMakerListener) EnterConst(ctx *ConstContext) {}

// ExitConst is called when production const is exited.
func (s *BaseConstMakerListener) ExitConst(ctx *ConstContext) {}

// EnterEnum is called when production enum is entered.
func (s *BaseConstMakerListener) EnterEnum(ctx *EnumContext) {}

// ExitEnum is called when production enum is exited.
func (s *BaseConstMakerListener) ExitEnum(ctx *EnumContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseConstMakerListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseConstMakerListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterType is called when production type is entered.
func (s *BaseConstMakerListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseConstMakerListener) ExitType(ctx *TypeContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseConstMakerListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseConstMakerListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterIdentAssignValue is called when production identAssignValue is entered.
func (s *BaseConstMakerListener) EnterIdentAssignValue(ctx *IdentAssignValueContext) {}

// ExitIdentAssignValue is called when production identAssignValue is exited.
func (s *BaseConstMakerListener) ExitIdentAssignValue(ctx *IdentAssignValueContext) {}

// EnterIdentAssignValueLoop is called when production identAssignValueLoop is entered.
func (s *BaseConstMakerListener) EnterIdentAssignValueLoop(ctx *IdentAssignValueLoopContext) {}

// ExitIdentAssignValueLoop is called when production identAssignValueLoop is exited.
func (s *BaseConstMakerListener) ExitIdentAssignValueLoop(ctx *IdentAssignValueLoopContext) {}

// EnterOption is called when production option is entered.
func (s *BaseConstMakerListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseConstMakerListener) ExitOption(ctx *OptionContext) {}

// EnterValue is called when production value is entered.
func (s *BaseConstMakerListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseConstMakerListener) ExitValue(ctx *ValueContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseConstMakerListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseConstMakerListener) ExitInteger(ctx *IntegerContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseConstMakerListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseConstMakerListener) ExitIdent(ctx *IdentContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseConstMakerListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseConstMakerListener) ExitKeywords(ctx *KeywordsContext) {}

// EnterDocComment is called when production docComment is entered.
func (s *BaseConstMakerListener) EnterDocComment(ctx *DocCommentContext) {}

// ExitDocComment is called when production docComment is exited.
func (s *BaseConstMakerListener) ExitDocComment(ctx *DocCommentContext) {}

// EnterTripleComment is called when production tripleComment is entered.
func (s *BaseConstMakerListener) EnterTripleComment(ctx *TripleCommentContext) {}

// ExitTripleComment is called when production tripleComment is exited.
func (s *BaseConstMakerListener) ExitTripleComment(ctx *TripleCommentContext) {}
