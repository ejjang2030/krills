package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL 	= "ILLEGAL"
	EOF      	= "EOF"
	EOL 		= "EOL"

	// identifier + literal
	IDENT  		= "IDENT" // add, foobar, x, y, ...
	INT    		= "INT"   // 12345
	STRING 		= "STRING"

	// operator
	ASSIGN   	= "="
	PLUS     	= "+"
	MINUS    	= "-"
	BANG     	= "!"
	ASTERISK 	= "*"
	SLASH    	= "/"
	POWER	 	= "^"

	LT 			= "<"
	GT 			= ">"
	LE 			= "<="
	GE 			= ">="

	EQ     		= "=="
	NOT_EQ 		= "!"

	// seperator
	COMMA     	= ","
	SEMICOLON 	= ";"
	COLON     	= ":"

	LPAREN   	= "("
	RPAREN   	= ")"
	LBRACE  	= "{"
	RBRACE   	= "}"
	LBRACKET 	= "["
	RBRACKET 	= "]"

	// keywords
	FUNCTION 	= "FUNCTION"
	LET     	= "LET"
	TRUE     	= "TRUE"
	FALSE    	= "FALSE"
	IF       	= "IF"
	ELSE     	= "ELSE"
	RETURN   	= "RETURN"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
