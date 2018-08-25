/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"function_name",
		"(",
		")",
		"()",
		",",
		"operator_charactor",
		"?",
		":",
		"doublequotes_string",
		"singlequote_string",
		"number",
		"argument",
		"true",
		"false",
		"float",
		"nil",
		"null",
	},

	idMap: map[string]Type{
		"INVALID":             0,
		"$":                   1,
		"function_name":       2,
		"(":                   3,
		")":                   4,
		"()":                  5,
		",":                   6,
		"operator_charactor":  7,
		"?":                   8,
		":":                   9,
		"doublequotes_string": 10,
		"singlequote_string":  11,
		"number":              12,
		"argument":            13,
		"true":                14,
		"false":               15,
		"float":               16,
		"nil":                 17,
		"null":                18,
	},
}
