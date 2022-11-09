// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // CustomJson
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseCustomJsonListener is a complete listener for a parse tree produced by CustomJsonParser.
type BaseCustomJsonListener struct{}

var _ CustomJsonListener = &BaseCustomJsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCustomJsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCustomJsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCustomJsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCustomJsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseCustomJsonListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseCustomJsonListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseCustomJsonListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseCustomJsonListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseCustomJsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseCustomJsonListener) ExitPair(ctx *PairContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseCustomJsonListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseCustomJsonListener) ExitArr(ctx *ArrContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCustomJsonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCustomJsonListener) ExitValue(ctx *ValueContext) {}
