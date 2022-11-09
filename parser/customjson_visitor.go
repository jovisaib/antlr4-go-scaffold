// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CustomJson
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CustomJsonParser.
type CustomJsonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CustomJsonParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by CustomJsonParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by CustomJsonParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by CustomJsonParser#arr.
	VisitArr(ctx *ArrContext) interface{}

	// Visit a parse tree produced by CustomJsonParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
