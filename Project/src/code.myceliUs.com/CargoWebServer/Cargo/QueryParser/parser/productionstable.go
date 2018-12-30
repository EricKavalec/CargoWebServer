
// generated by gocc; DO NOT EDIT.

package parser

import "code.myceliUs.com/CargoWebServer/Cargo/QueryParser/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Query	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Query : "(" Query ")"	<< ast.NewQueryAst(X[1]) >>`,
		Id: "Query",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewQueryAst(X[1])
		},
	},
	ProdTabEntry{
		String: `Query : Query operator Query	<< ast.AppendQueryAst(X[0], X[1], X[2]) >>`,
		Id: "Query",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendQueryAst(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Query : Object comparator Value	<< ast.NewExpressionAst(X[0], X[1], X[2]) >>`,
		Id: "Query",
		NTType: 1,
		Index: 3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpressionAst(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Value : int64	<< ast.NewIntegerValue(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntegerValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : float64	<< ast.NewFloatValue(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFloatValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : string	<< ast.NewStringValue(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : boolean	<< ast.NewBoolValue(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : regex	<< ast.NewRegexExpr(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRegexExpr(X[0])
		},
	},
	ProdTabEntry{
		String: `Value : null	<< ast.NewNullValue(X[0]) >>`,
		Id: "Value",
		NTType: 2,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNullValue(X[0])
		},
	},
	ProdTabEntry{
		String: `Object : id "." id	<< ast.NewIdValue(X[0], X[2]) >>`,
		Id: "Object",
		NTType: 3,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIdValue(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Object : Object "." id	<< ast.AppendIdValue(X[0], X[2]) >>`,
		Id: "Object",
		NTType: 3,
		Index: 11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendIdValue(X[0], X[2])
		},
	},
	
}
