// Code generated from ConstMaker.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ConstMaker
import "github.com/antlr4-go/antlr/v4"

type BaseConstMakerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseConstMakerVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitRules(ctx *RulesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitConst(ctx *ConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitEnum(ctx *EnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitIdentAssignValue(ctx *IdentAssignValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitIdentAssignValueLoop(ctx *IdentAssignValueLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitKeywords(ctx *KeywordsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitDocComment(ctx *DocCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseConstMakerVisitor) VisitTripleComment(ctx *TripleCommentContext) interface{} {
	return v.VisitChildren(ctx)
}
