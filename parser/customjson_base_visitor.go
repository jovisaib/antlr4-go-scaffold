// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CustomJson
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseCustomJsonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCustomJsonVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomJsonVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomJsonVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomJsonVisitor) VisitArr(ctx *ArrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCustomJsonVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
