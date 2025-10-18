// Copyright © 2025 cymlifework (cymlifework@163.com). All rights reserved.
// Licensed under the Apache License, Version 2.0 (see LICENSE file).

package sqlite

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/cymfarcat/constmaker/parser/sqlite"
	src "github.com/cymfarcat/constmaker/src"
)

type SQLWriterAny struct {
	parser.BaseSQLiteParserListener
	parser   *SQLiteParser
	option   *src.Options
	builder  *strings.Builder
	stmtGen  bool
	stmtFull bool
	stmtName string
	inCreate bool

	typoStr string
}

func NewSQLWriterAny(parser *SQLiteParser, typoStr string) *SQLWriterAny {
	writer := &SQLWriterAny{}
	writer.parser = parser
	writer.option = parser.option
	writer.builder = new(strings.Builder)
	writer.typoStr = typoStr
	return writer
}

func (obj *SQLWriterAny) GenConsts() {
	// 1. gather keys
	keys := make([]string, 0, len(obj.parser.consts))
	for key := range obj.parser.consts {
		keys = append(keys, key)
	}

	// 2. sort key
	sort.Strings(keys)

	// 3. output
	obj.builder.WriteString(obj.option.GetTabWidth() + "// Consts\n")

	for _, key := range keys {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, key) + " = \"" + obj.parser.consts[key] + "\";\n")
	}
}

func (obj *SQLWriterAny) GenSql(tree parser.IParseContext) {
	if obj.option.GenSql {
		obj.stmtGen = false
		obj.option.Level = 0
		obj.builder.Reset()

		//tree.Accept(obj) //obj.Visit(tree)
		antlr.ParseTreeWalkerDefault.Walk(obj, tree)

		path := obj.option.GetOutputFile("sql")
		data := obj.builder.String()
		obj.parser.writeFile(path, data, true, "-- ", "")
	}
}

func (obj *SQLWriterAny) GenCpp(tree parser.IParseContext) {
	obj.builder.WriteString("\n#ifndef CONSTMAKER_GENERATED_" + strings.ToUpper(obj.option.FileName) + "_H\n")
	obj.builder.WriteString("#define CONSTMAKER_GENERATED_" + strings.ToUpper(obj.option.FileName) + "_H\n")

	// first write .h
	obj.stmtGen = true
	obj.option.Level = 0

	// write consts
	obj.builder.WriteString("\n")
	obj.GenConsts()

	//tree.Accept(obj) //obj.Visit(tree)
	antlr.ParseTreeWalkerDefault.Walk(obj, tree)

	obj.builder.WriteString("\n\n#endif // CONSTMAKER_GENERATED_" + strings.ToUpper(obj.option.FileName) + "_H\n")

	path := obj.option.GetOutputFile("h")
	data := obj.builder.String()
	obj.parser.writeFile(path, data, true, "", "")

	// then write .sql
	obj.GenSql(tree)
}

func (obj *SQLWriterAny) GenDart(tree parser.IParseContext) {
	// first write .dart
	obj.stmtGen = true
	obj.option.Level = 0

	// write consts
	obj.builder.WriteString("\n")
	obj.GenConsts()

	//tree.Accept(obj) //obj.Visit(tree)
	antlr.ParseTreeWalkerDefault.Walk(obj, tree)

	path := obj.option.GetOutputFile("dart")
	data := obj.builder.String()
	obj.parser.writeFile(path, data, true, "", "")

	// then write .sql
	obj.GenSql(tree)
}

func (obj *SQLWriterAny) GenGo(tree parser.IParseContext) {
	// first write .go
	obj.stmtGen = true
	obj.option.Level = 0

	if len(obj.option.GoPackage) > 0 {
		obj.builder.WriteString("\npackage " + obj.option.GoPackage + "\n")
	}

	// write consts
	obj.builder.WriteString("\n")
	obj.builder.WriteString("const (\n")
	obj.option.Level++
	obj.GenConsts()
	obj.option.Level--
	obj.builder.WriteString(")\n\n")

	obj.builder.WriteString("const (")
	obj.option.Level++
	//tree.Accept(obj) //obj.Visit(tree)
	antlr.ParseTreeWalkerDefault.Walk(obj, tree)
	obj.option.Level--
	obj.builder.WriteString(")\n")

	path := obj.option.GetOutputFile("go")
	data := obj.builder.String()
	obj.parser.writeFile(path, data, true, "", "")

	// then write .sql
	obj.GenSql(tree)
}

func (obj *SQLWriterAny) GenRust(tree parser.IParseContext) {
	// first write .dart
	obj.stmtGen = true
	obj.option.Level = 0

	// write consts
	obj.builder.WriteString("\n")
	obj.GenConsts()

	//tree.Accept(obj) //obj.Visit(tree)
	antlr.ParseTreeWalkerDefault.Walk(obj, tree)

	path := obj.option.GetOutputFile("rs")
	data := obj.builder.String()
	obj.parser.writeFile(path, data, true, "", "")

	// then write .sql
	obj.GenSql(tree)
}

func (obj *SQLWriterAny) GenText(name string, prc antlr.BaseParserRuleContext) string {
	if prc.GetChildCount() == 0 {
		return ""
	}

	obj.option.Level++

	var str string
	for _, child := range prc.GetChildren() {
		value := reflect.ValueOf(child)
		typo := value.Type().String()

		item := ""
		switch typo {
		case "*sqlite.Column_defContext":
			item = obj.option.GetEnterCount(1) + obj.option.GetTabWidth() + obj.GenText(name, value.Interface().(*parser.Column_defContext).BaseParserRuleContext)
		case "*sqlite.Column_constraintContext":
			item = obj.GenText(name, value.Interface().(*parser.Column_constraintContext).BaseParserRuleContext)
		case "*sqlite.Foreign_key_clauseContext":
			item = obj.GenText(name, value.Interface().(*parser.Foreign_key_clauseContext).BaseParserRuleContext)
		case "*sqlite.ExprContext":
			item = obj.GenText(name, value.Interface().(*parser.ExprContext).BaseParserRuleContext)
		case "*sqlite.Limit_stmtContext":
			item = obj.GenText(name, value.Interface().(*parser.Limit_stmtContext).BaseParserRuleContext)
		case "*sqlite.Literal_valueContext":
			if obj.inCreate {
				//eg: CREATE VIRTUAL TABLE: content='posts'
				item = obj.GenText(name, value.Interface().(*parser.Literal_valueContext).BaseParserRuleContext)
				item = strings.TrimSuffix(item, " ")
				item = strings.Trim(item, "'")
				item = obj.parser.AltName(name, item)
				item = "'" + item + "'"
			} else {
				item = child.(antlr.ParseTree).GetText()
			}
		case "*sqlite.Module_argumentContext":
			item = obj.GenText(name, value.Interface().(*parser.Module_argumentContext).BaseParserRuleContext)
		case "*sqlite.Join_clauseContext":
			item = obj.GenText(name, value.Interface().(*parser.Join_clauseContext).BaseParserRuleContext)
		case "*sqlite.Join_constraintContext":
			item = obj.GenText(name, value.Interface().(*parser.Join_constraintContext).BaseParserRuleContext)
		case "*sqlite.Join_operatorContext":
			item = obj.GenText(name, value.Interface().(*parser.Join_operatorContext).BaseParserRuleContext)
		case "*sqlite.Table_constraintContext":
			item = obj.option.GetEnterCount(1) + obj.option.GetTabWidth() + obj.GenText(name, value.Interface().(*parser.Table_constraintContext).BaseParserRuleContext)
		case "*sqlite.Table_or_subqueryContext":
			item = obj.GenText(name, value.Interface().(*parser.Table_or_subqueryContext).BaseParserRuleContext)
		case "*sqlite.Indexed_columnContext":
			item = obj.GenText(name, value.Interface().(*parser.Indexed_columnContext).BaseParserRuleContext)
		case "*sqlite.Result_columnContext":
			item = obj.GenText(name, value.Interface().(*parser.Result_columnContext).BaseParserRuleContext)
		case "*sqlite.Insert_stmtContext":
			item = obj.GenText(name, value.Interface().(*parser.Insert_stmtContext).BaseParserRuleContext)
		case "*sqlite.Select_stmtContext":
			item = obj.GenText(name, value.Interface().(*parser.Select_stmtContext).BaseParserRuleContext)
		case "*sqlite.Select_coreContext":
			item = obj.GenText(name, value.Interface().(*parser.Select_coreContext).BaseParserRuleContext)
		case "*sqlite.Update_stmtContext":
			item = obj.GenText(name, value.Interface().(*parser.Update_stmtContext).BaseParserRuleContext)
		default:
			item = child.(antlr.ParseTree).GetText()
		}

		fmt.Fprintf(os.Stderr, "type=%s, value=%s\n", typo, item)

		item = obj.parser.AltName(name, item)

		if item == "(" || item == ")" || item == "." || item == "," || item == ";" {
			str = strings.TrimSuffix(str, " ")
		}

		str += item

		if len(item) == 0 || item == "(" || item == ")" || item == "." || item == "," || item == ";" || strings.HasSuffix(str, " ") {
			continue
		}

		str += " "
	}

	obj.option.Level--

	return str
}

// EnterParse is called when entering the parse production.
// func (obj *SQLWriterCPP) EnterParse(ctx *parser.ParseContext) {}

// EnterSql_stmt_list is called when entering the sql_stmt_list production.
// func (obj *SQLWriterCPP) EnterSql_stmt_list(ctx *parser.Sql_stmt_listContext) {}

// EnterSql_stmt is called when entering the sql_stmt production.
// func (obj *SQLWriterCPP) EnterSql_stmt(ctx *parser.Sql_stmtContext) {}

func (obj *SQLWriterAny) ExitSql_stmt(ctx *parser.Sql_stmtContext) {
	//clear stmt name
	obj.stmtName = ""
}

// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
func (obj *SQLWriterAny) EnterAlter_table_stmt(ctx *parser.Alter_table_stmtContext) {
	name := ctx.Table_name(0).GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "alter", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterAnalyze_stmt is called when entering the analyze_stmt production.
func (obj *SQLWriterAny) EnterAnalyze_stmt(ctx *parser.Analyze_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "analyze", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterAttach_stmt is called when entering the attach_stmt production.
func (obj *SQLWriterAny) EnterAttach_stmt(ctx *parser.Attach_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "attach", ctx.Schema_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterBegin_stmt is called when entering the begin_stmt production.
func (obj *SQLWriterAny) EnterBegin_stmt(ctx *parser.Begin_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "begin", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterCommit_stmt is called when entering the commit_stmt production.
func (obj *SQLWriterAny) EnterCommit_stmt(ctx *parser.Commit_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "commit", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterRollback_stmt is called when entering the rollback_stmt production.
func (obj *SQLWriterAny) EnterRollback_stmt(ctx *parser.Rollback_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "rollback", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterSavepoint_stmt is called when entering the savepoint_stmt production.
func (obj *SQLWriterAny) EnterSavepoint_stmt(ctx *parser.Savepoint_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "savepoint", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterRelease_stmt is called when entering the release_stmt production.
func (obj *SQLWriterAny) EnterRelease_stmt(ctx *parser.Release_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "release", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterCreate_index_stmt is called when entering the create_index_stmt production.
func (obj *SQLWriterAny) EnterCreate_index_stmt(ctx *parser.Create_index_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(ctx.Table_name().GetText(), ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "create_index", ctx.Index_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterIndexed_column is called when entering the indexed_column production.
func (obj *SQLWriterAny) EnterIndexed_column(ctx *parser.Indexed_columnContext) {}

// EnterCreate_table_stmt is called when entering the create_table_stmt production.
func (obj *SQLWriterAny) EnterCreate_table_stmt(ctx *parser.Create_table_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(ctx.Table_name().GetText(), ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "create_table", ctx.Table_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterColumn_def is called when entering the column_def production.
func (obj *SQLWriterAny) EnterColumn_def(ctx *parser.Column_defContext) {}

// EnterType_name is called when entering the type_name production.
func (obj *SQLWriterAny) EnterType_name(ctx *parser.Type_nameContext) {}

// EnterColumn_constraint is called when entering the column_constraint production.
func (obj *SQLWriterAny) EnterColumn_constraint(ctx *parser.Column_constraintContext) {}

// EnterSigned_number is called when entering the signed_number production.
func (obj *SQLWriterAny) EnterSigned_number(ctx *parser.Signed_numberContext) {}

// EnterTable_constraint is called when entering the table_constraint production.
func (obj *SQLWriterAny) EnterTable_constraint(ctx *parser.Table_constraintContext) {}

// EnterForeign_key_clause is called when entering the foreign_key_clause production.
func (obj *SQLWriterAny) EnterForeign_key_clause(ctx *parser.Foreign_key_clauseContext) {}

// EnterConflict_clause is called when entering the conflict_clause production.
func (obj *SQLWriterAny) EnterConflict_clause(ctx *parser.Conflict_clauseContext) {}

// EnterCreate_trigger_stmt is called when entering the create_trigger_stmt production.
func (obj *SQLWriterAny) EnterCreate_trigger_stmt(ctx *parser.Create_trigger_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	obj.inCreate = true
	str := strings.TrimSuffix(obj.GenText(ctx.Table_name().GetText(), ctx.BaseParserRuleContext), " ") + ";"
	obj.inCreate = false

	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "create_trigger", ctx.Trigger_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterCreate_view_stmt is called when entering the create_view_stmt production.
func (obj *SQLWriterAny) EnterCreate_view_stmt(ctx *parser.Create_view_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "create_view", ctx.View_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterCreate_virtual_table_stmt is called when entering the create_virtual_table_stmt production.
func (obj *SQLWriterAny) EnterCreate_virtual_table_stmt(ctx *parser.Create_virtual_table_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	obj.inCreate = true
	str := strings.TrimSuffix(obj.GenText(ctx.Table_name().GetText(), ctx.BaseParserRuleContext), " ") + ";"
	obj.inCreate = false

	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "create_virtual_table", ctx.Table_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterWith_clause is called when entering the with_clause production.
func (obj *SQLWriterAny) EnterWith_clause(ctx *parser.With_clauseContext) {}

// EnterCte_table_name is called when entering the cte_table_name production.
func (obj *SQLWriterAny) EnterCte_table_name(ctx *parser.Cte_table_nameContext) {}

// EnterRecursive_cte is called when entering the recursive_cte production.
func (obj *SQLWriterAny) EnterRecursive_cte(ctx *parser.Recursive_cteContext) {}

// EnterCommon_table_expression is called when entering the common_table_expression production.
func (obj *SQLWriterAny) EnterCommon_table_expression(ctx *parser.Common_table_expressionContext) {}

// EnterDelete_stmt is called when entering the delete_stmt production.
func (obj *SQLWriterAny) EnterDelete_stmt(ctx *parser.Delete_stmtContext) {
	name := ctx.Qualified_table_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "delete", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterDelete_stmt_limited is called when entering the delete_stmt_limited production.
func (obj *SQLWriterAny) EnterDelete_stmt_limited(ctx *parser.Delete_stmt_limitedContext) {
	name := ctx.Qualified_table_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "delete_limited", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterDetach_stmt is called when entering the detach_stmt production.
func (obj *SQLWriterAny) EnterDetach_stmt(ctx *parser.Detach_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "detach", ctx.Schema_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterDrop_stmt is called when entering the drop_stmt production.
func (obj *SQLWriterAny) EnterDrop_stmt(ctx *parser.Drop_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "drop", ctx.Any_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterExpr is called when entering the expr production.
func (obj *SQLWriterAny) EnterExpr(ctx *parser.ExprContext) {}

// EnterRaise_function is called when entering the raise_function production.
func (obj *SQLWriterAny) EnterRaise_function(ctx *parser.Raise_functionContext) {}

// EnterLiteral_value is called when entering the literal_value production.
func (obj *SQLWriterAny) EnterLiteral_value(ctx *parser.Literal_valueContext) {}

// EnterValue_row is called when entering the value_row production.
func (obj *SQLWriterAny) EnterValue_row(ctx *parser.Value_rowContext) {}

// EnterValues_clause is called when entering the values_clause production.
func (obj *SQLWriterAny) EnterValues_clause(ctx *parser.Values_clauseContext) {}

// EnterInsert_stmt is called when entering the insert_stmt production.
func (obj *SQLWriterAny) EnterInsert_stmt(ctx *parser.Insert_stmtContext) {
	name := ctx.Table_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "insert", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterReturning_clause is called when entering the returning_clause production.
func (obj *SQLWriterAny) EnterReturning_clause(ctx *parser.Returning_clauseContext) {}

// EnterUpsert_clause is called when entering the upsert_clause production.
func (obj *SQLWriterAny) EnterUpsert_clause(ctx *parser.Upsert_clauseContext) {}

// EnterPragma_stmt is called when entering the pragma_stmt production.
func (obj *SQLWriterAny) EnterPragma_stmt(ctx *parser.Pragma_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "pragma", ctx.Pragma_name().GetText(), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterPragma_value is called when entering the pragma_value production.
func (obj *SQLWriterAny) EnterPragma_value(ctx *parser.Pragma_valueContext) {}

// EnterReindex_stmt is called when entering the reindex_stmt production.
func (obj *SQLWriterAny) EnterReindex_stmt(ctx *parser.Reindex_stmtContext) {
	name := ctx.Collation_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "reindex", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterSelect_stmt is called when entering the select_stmt production.
func (obj *SQLWriterAny) EnterSelect_stmt(ctx *parser.Select_stmtContext) {
	//fmt.Fprintf(os.Stdout, "select= %s\n", ctx.GetText())

	var names []string

	// first gather all table names
	all_select := ctx.AllSelect_core()
	for _, one_select := range all_select {
		all_table := one_select.AllTable_or_subquery()
		for _, one_table := range all_table {
			names = append(names, one_table.Table_name().GetText())
		}

		join := one_select.Join_clause()
		if join != nil {
			all_table := join.AllTable_or_subquery()
			for _, one_table := range all_table {
				names = append(names, one_table.Table_name().GetText())
			}
		}
	}

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(strings.Join(names, src.SeparatorChar), ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "select", strings.Join(names, src.UnderScore), obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterJoin_clause is called when entering the join_clause production.
func (obj *SQLWriterAny) EnterJoin_clause(ctx *parser.Join_clauseContext) {}

// EnterSelect_core is called when entering the select_core production.
func (obj *SQLWriterAny) EnterSelect_core(ctx *parser.Select_coreContext) {}

// EnterFactored_select_stmt is called when entering the factored_select_stmt production.
func (obj *SQLWriterAny) EnterFactored_select_stmt(ctx *parser.Factored_select_stmtContext) {}

// EnterSimple_select_stmt is called when entering the simple_select_stmt production.
func (obj *SQLWriterAny) EnterSimple_select_stmt(ctx *parser.Simple_select_stmtContext) {}

// EnterCompound_select_stmt is called when entering the compound_select_stmt production.
func (obj *SQLWriterAny) EnterCompound_select_stmt(ctx *parser.Compound_select_stmtContext) {}

// EnterTable_or_subquery is called when entering the table_or_subquery production.
func (obj *SQLWriterAny) EnterTable_or_subquery(ctx *parser.Table_or_subqueryContext) {}

// EnterResult_column is called when entering the result_column production.
func (obj *SQLWriterAny) EnterResult_column(ctx *parser.Result_columnContext) {}

// EnterJoin_operator is called when entering the join_operator production.
func (obj *SQLWriterAny) EnterJoin_operator(ctx *parser.Join_operatorContext) {}

// EnterJoin_constraint is called when entering the join_constraint production.
func (obj *SQLWriterAny) EnterJoin_constraint(ctx *parser.Join_constraintContext) {}

// EnterCompound_operator is called when entering the compound_operator production.
func (obj *SQLWriterAny) EnterCompound_operator(ctx *parser.Compound_operatorContext) {}

// EnterUpdate_stmt is called when entering the update_stmt production.
func (obj *SQLWriterAny) EnterUpdate_stmt(ctx *parser.Update_stmtContext) {
	name := ctx.Qualified_table_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "update", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterColumn_name_list is called when entering the column_name_list production.
func (obj *SQLWriterAny) EnterColumn_name_list(ctx *parser.Column_name_listContext) {}

// EnterUpdate_stmt_limited is called when entering the update_stmt_limited production.
func (obj *SQLWriterAny) EnterUpdate_stmt_limited(ctx *parser.Update_stmt_limitedContext) {
	name := ctx.Qualified_table_name().GetText()

	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText(name, ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "update_limited", name, obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterQualified_table_name is called when entering the qualified_table_name production.
func (obj *SQLWriterAny) EnterQualified_table_name(ctx *parser.Qualified_table_nameContext) {}

// EnterVacuum_stmt is called when entering the vacuum_stmt production.
func (obj *SQLWriterAny) EnterVacuum_stmt(ctx *parser.Vacuum_stmtContext) {
	//obj.option.Level = 0
	obj.builder.WriteString(obj.option.GetEnterCount(1))

	str := strings.TrimSuffix(obj.GenText("", ctx.BaseParserRuleContext), " ") + ";"
	if obj.stmtGen {
		obj.builder.WriteString(obj.option.GetTabWidth() + typeStr(obj.typoStr, getStmtName(obj.option, "vacuum", "", obj.stmtFull, obj.stmtName)) + " = ")
		obj.builder.WriteString(quoteStr(obj.typoStr, str))
	} else {
		obj.builder.WriteString(str)
	}

	obj.builder.WriteString(obj.option.GetEnterCount(1))
}

// EnterFilter_clause is called when entering the filter_clause production.
func (obj *SQLWriterAny) EnterFilter_clause(ctx *parser.Filter_clauseContext) {}

// EnterWindow_defn is called when entering the window_defn production.
func (obj *SQLWriterAny) EnterWindow_defn(ctx *parser.Window_defnContext) {}

// EnterOver_clause is called when entering the over_clause production.
func (obj *SQLWriterAny) EnterOver_clause(ctx *parser.Over_clauseContext) {}

// EnterFrame_spec is called when entering the frame_spec production.
func (obj *SQLWriterAny) EnterFrame_spec(ctx *parser.Frame_specContext) {}

// EnterFrame_clause is called when entering the frame_clause production.
func (obj *SQLWriterAny) EnterFrame_clause(ctx *parser.Frame_clauseContext) {}

// EnterSimple_function_invocation is called when entering the simple_function_invocation production.
func (obj *SQLWriterAny) EnterSimple_function_invocation(ctx *parser.Simple_function_invocationContext) {
}

// EnterAggregate_function_invocation is called when entering the aggregate_function_invocation production.
func (obj *SQLWriterAny) EnterAggregate_function_invocation(ctx *parser.Aggregate_function_invocationContext) {
}

// EnterWindow_function_invocation is called when entering the window_function_invocation production.
func (obj *SQLWriterAny) EnterWindow_function_invocation(ctx *parser.Window_function_invocationContext) {
}

// EnterCommon_table_stmt is called when entering the common_table_stmt production.
func (obj *SQLWriterAny) EnterCommon_table_stmt(ctx *parser.Common_table_stmtContext) {}

// EnterOrder_by_stmt is called when entering the order_by_stmt production.
func (obj *SQLWriterAny) EnterOrder_by_stmt(ctx *parser.Order_by_stmtContext) {}

// EnterLimit_stmt is called when entering the limit_stmt production.
func (obj *SQLWriterAny) EnterLimit_stmt(ctx *parser.Limit_stmtContext) {}

// EnterOrdering_term is called when entering the ordering_term production.
func (obj *SQLWriterAny) EnterOrdering_term(ctx *parser.Ordering_termContext) {}

// EnterAsc_desc is called when entering the asc_desc production.
func (obj *SQLWriterAny) EnterAsc_desc(ctx *parser.Asc_descContext) {}

// EnterFrame_left is called when entering the frame_left production.
func (obj *SQLWriterAny) EnterFrame_left(ctx *parser.Frame_leftContext) {}

// EnterFrame_right is called when entering the frame_right production.
func (obj *SQLWriterAny) EnterFrame_right(ctx *parser.Frame_rightContext) {}

// EnterFrame_single is called when entering the frame_single production.
func (obj *SQLWriterAny) EnterFrame_single(ctx *parser.Frame_singleContext) {}

// EnterWindow_function is called when entering the window_function production.
func (obj *SQLWriterAny) EnterWindow_function(ctx *parser.Window_functionContext) {}

// EnterOffset is called when entering the offset production.
func (obj *SQLWriterAny) EnterOffset(ctx *parser.OffsetContext) {}

// EnterDefault_value is called when entering the default_value production.
func (obj *SQLWriterAny) EnterDefault_value(ctx *parser.Default_valueContext) {}

// EnterPartition_by is called when entering the partition_by production.
func (obj *SQLWriterAny) EnterPartition_by(ctx *parser.Partition_byContext) {}

// EnterOrder_by_expr is called when entering the order_by_expr production.
func (obj *SQLWriterAny) EnterOrder_by_expr(ctx *parser.Order_by_exprContext) {}

// EnterOrder_by_expr_asc_desc is called when entering the order_by_expr_asc_desc production.
func (obj *SQLWriterAny) EnterOrder_by_expr_asc_desc(ctx *parser.Order_by_expr_asc_descContext) {}

// EnterExpr_asc_desc is called when entering the expr_asc_desc production.
func (obj *SQLWriterAny) EnterExpr_asc_desc(ctx *parser.Expr_asc_descContext) {}

// EnterInitial_select is called when entering the initial_select production.
func (obj *SQLWriterAny) EnterInitial_select(ctx *parser.Initial_selectContext) {}

// EnterRecursive_select is called when entering the recursive_select production.
func (obj *SQLWriterAny) EnterRecursive_select(ctx *parser.Recursive_selectContext) {}

// EnterUnary_operator is called when entering the unary_operator production.
func (obj *SQLWriterAny) EnterUnary_operator(ctx *parser.Unary_operatorContext) {}

// EnterError_message is called when entering the error_message production.
func (obj *SQLWriterAny) EnterError_message(ctx *parser.Error_messageContext) {}

// EnterModule_argument is called when entering the module_argument production.
func (obj *SQLWriterAny) EnterModule_argument(ctx *parser.Module_argumentContext) {}

// EnterColumn_alias is called when entering the column_alias production.
func (obj *SQLWriterAny) EnterColumn_alias(ctx *parser.Column_aliasContext) {}

// EnterKeyword is called when entering the keyword production.
func (obj *SQLWriterAny) EnterKeyword(ctx *parser.KeywordContext) {}

// EnterName is called when entering the name production.
func (obj *SQLWriterAny) EnterName(ctx *parser.NameContext) {}

// EnterFunction_name is called when entering the function_name production.
func (obj *SQLWriterAny) EnterFunction_name(ctx *parser.Function_nameContext) {}

// EnterSchema_name is called when entering the schema_name production.
func (obj *SQLWriterAny) EnterSchema_name(ctx *parser.Schema_nameContext) {}

// EnterTable_name is called when entering the table_name production.
func (obj *SQLWriterAny) EnterTable_name(ctx *parser.Table_nameContext) {}

// EnterTable_or_index_name is called when entering the table_or_index_name production.
func (obj *SQLWriterAny) EnterTable_or_index_name(ctx *parser.Table_or_index_nameContext) {}

// EnterColumn_name is called when entering the column_name production.
func (obj *SQLWriterAny) EnterColumn_name(ctx *parser.Column_nameContext) {}

// EnterCollation_name is called when entering the collation_name production.
func (obj *SQLWriterAny) EnterCollation_name(ctx *parser.Collation_nameContext) {}

// EnterForeign_table is called when entering the foreign_table production.
func (obj *SQLWriterAny) EnterForeign_table(ctx *parser.Foreign_tableContext) {}

// EnterIndex_name is called when entering the index_name production.
func (obj *SQLWriterAny) EnterIndex_name(ctx *parser.Index_nameContext) {}

// EnterTrigger_name is called when entering the trigger_name production.
func (obj *SQLWriterAny) EnterTrigger_name(ctx *parser.Trigger_nameContext) {}

// EnterView_name is called when entering the view_name production.
func (obj *SQLWriterAny) EnterView_name(ctx *parser.View_nameContext) {}

// EnterModule_name is called when entering the module_name production.
func (obj *SQLWriterAny) EnterModule_name(ctx *parser.Module_nameContext) {}

// EnterPragma_name is called when entering the pragma_name production.
func (obj *SQLWriterAny) EnterPragma_name(ctx *parser.Pragma_nameContext) {}

// EnterSavepoint_name is called when entering the savepoint_name production.
func (obj *SQLWriterAny) EnterSavepoint_name(ctx *parser.Savepoint_nameContext) {}

// EnterTable_alias is called when entering the table_alias production.
func (obj *SQLWriterAny) EnterTable_alias(ctx *parser.Table_aliasContext) {}

// EnterTransaction_name is called when entering the transaction_name production.
func (obj *SQLWriterAny) EnterTransaction_name(ctx *parser.Transaction_nameContext) {}

// EnterWindow_name is called when entering the window_name production.
func (obj *SQLWriterAny) EnterWindow_name(ctx *parser.Window_nameContext) {}

// EnterAlias is called when entering the alias production.
func (obj *SQLWriterAny) EnterAlias(ctx *parser.AliasContext) {}

// EnterFilename is called when entering the filename production.
func (obj *SQLWriterAny) EnterFilename(ctx *parser.FilenameContext) {}

// EnterBase_window_name is called when entering the base_window_name production.
func (obj *SQLWriterAny) EnterBase_window_name(ctx *parser.Base_window_nameContext) {}

// EnterSimple_func is called when entering the simple_func production.
func (obj *SQLWriterAny) EnterSimple_func(ctx *parser.Simple_funcContext) {}

// EnterAggregate_func is called when entering the aggregate_func production.
func (obj *SQLWriterAny) EnterAggregate_func(ctx *parser.Aggregate_funcContext) {}

// EnterTable_function_name is called when entering the table_function_name production.
func (obj *SQLWriterAny) EnterTable_function_name(ctx *parser.Table_function_nameContext) {}

// EnterAny_name is called when entering the any_name production.
func (obj *SQLWriterAny) EnterAny_name(ctx *parser.Any_nameContext) {
	// fmt.Fprintf(os.Stdout, "anyName= %s\n", ctx.GetText())
}

// EnterDocComment is called when production docComment is entered.
func (obj *SQLWriterAny) EnterDocComment(ctx *parser.DocCommentContext) {
	//fmt.Fprintf(os.Stdout, "docComment= %s\n", ctx.GetText())

	//clear stmt name
	obj.stmtName = ""

	obj.builder.WriteString("\n" + obj.option.IndentComment(ctx.GetText()))
}

// EnterTripleComment is called when production tripleComment is entered.
func (obj *SQLWriterAny) EnterTripleComment(ctx *parser.TripleCommentContext) {
	//fmt.Fprintf(os.Stdout, "tripleComment= %s\n", ctx.GetText())

	// parse this comment is stmt name
	obj.stmtFull, obj.stmtName = parseStmtName(ctx.GetText())

	if len(obj.stmtName) == 0 {
		if obj.stmtGen {
			obj.builder.WriteString("\n" + obj.option.GetTabWidth() + src.ConvertCommentTripleSQL(ctx.GetText(), "//"))
		} else {
			obj.builder.WriteString("\n" + obj.option.GetTabWidth() + src.ConvertCommentTripleSQL(ctx.GetText(), "--"))
		}
	}
}
