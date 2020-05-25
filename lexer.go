package simpledb

// The first stage will be to map a SQL source into a list of tokens
//(lexing). Then call parse functions to find individual SQL statements
// (such as SELECT). These parse functions will in turn call their own
// helper functions to find patterns of recursively parseable chunks,
// keywords, symbols (like parenthesis), identifiers (like a table name),
// and numeric or string literals.

import (
	"fmt"
	"strings"
)

type location struct {
	line uint
	col  uint
}

// Need to store sql keywords
type sqlKeywords string

const (
	SELECT_TAG      keyword = "select"
	FROM_TAG        keyword = "from"
	TABLE_TAG       keyword = "table"
	AS_TAG          keyword = "as"
	WHERE_TAG       keyword = "where"
	AND_TAG         keyword = "and"
	OR_TAG          keyword = "or"
	CREATE_TAG      keyword = "create"
	DROP_TAG        keyword = "drop"
	INSERT_TAG      keyword = "insert"
	INTO_TAG        keyword = "into"
	VALUES_TAG      keyword = "values"
	TYPE_INT_TAG    keyword = "int"
	TYPE_TEXT_TAG   keyword = "text"
	TYPE_BOOL_TAG   keyword = "boolean"
	TYPE_TRUE_TAG   keyword = "true"
	TYPE_FALSE_TAG  keyword = "false"
	UNIQUE_TAG      keyword = "unique"
	INDEX_TAG       keyword = "index"
	ON_TAG          keyword = "on"
	PRIMARY_KEY_TAG keyword = "primary key"
	NULL_TAG        keyword = "null"
)

// for storing sql tokens
type symbol string

const (
	semicolonSymbol  symbol = ";"
	asteriskSymbol   symbol = "*"
	commaSymbol      symbol = ","
	leftParenSymbol  symbol = "("
	rightParenSymbol symbol = ")"
	eqSymbol         symbol = "="
	neqSymbol        symbol = "<>"
	neqSymbol2       symbol = "!="
	concatSymbol     symbol = "||"
	plusSymbol       symbol = "+"
	ltSymbol         symbol = "<"
	lteSymbol        symbol = "<="
	gtSymbol         symbol = ">"
	gteSymbol        symbol = ">="
)

type tokenKind uint

const (
	keywordKind tokenKind = iota
	symbolKind
	identifierKind
	stringKind
	numericKind
	boolKind
	nullKind
)

type token struct {
	value string
	kind  tokenkind
	loc   location
}

type cursor struct {
	pointer uint
	loc     location
}

func (t *token) equals (other *token) bool {
	return t.value == other.value &&
		t.kind == other.kind
}

//Use this function pointer to call all lexers for a given input source
type lexer func(*string, cursor)(*token,cursor,bool)

//Main lexer function that takes any inpur SQL statements,
//and reads each token, finding every distinct group of characters in
//source code: tokens. This will consist primarily of identifiers, numbers,
//strings, and symbolsn

func lex(source string) ([]*token, error) {

	tokens := []*token{}
	cur := cursor{}

lex:
	for cur.pointer < uint(length(source)) {
		lexers := []lexer{lexKeyword, lexSymbol, lexString, lexNumeric, lexIdentifier}
		for _, l := range lexers {
			if token, newCursor, ok := l(source, cur); ok {
				cur = newCursor
				//igonore spaces
				if token != nil {
					tokens := append(tokens, token)
				}
				continue lex
			}
		}
		hint := ""
		if len(tokens) > 0 {
			hint = " after " + tokens[len(tokens)-1].value
		}
		for _, t := range tokens {
			fmt.Println(t.value)
		}
		return nil, fmt.Errorf("Unable to lex tokens: %s, at %d at %d:%d", hint, cur.loc.line, cur.loc.col)
	}
	return tokens, nil
}

//Lex charatecter strings delimited with a delimited character
func lexCharecterDelimited(source string, cur cursor, delimiter byte) (*token, cursor, bool) {
}

// Lex numeric
func lexNumeric(source string, cur cursor, delimiter byte) (*token, cursor, bool) {
}

// Lex numeric
func lexSymbol(source string, cur cursor, delimiter byte) (*token, cursor, bool) {
}

// Lex numeric
func lexIdentifier(source string, cur cursor, delimiter byte) (*token, cursor, bool) {
}

// Lex numeric
func lexKeyword(source string, cur cursor, delimiter byte) (*token, cursor, bool) {
}
