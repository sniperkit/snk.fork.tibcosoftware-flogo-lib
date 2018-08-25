/*
Sniperkit-Bot
- Status: analyzed
*/

package data

type Expr interface {
	Eval(scope Scope) (interface{}, error)
}
