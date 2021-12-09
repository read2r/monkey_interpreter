package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

// TokenTypes and Values
const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    IDENT = "IDENT"
    INT = "INT"

    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    EQ = "=="
    NOT_EQ = "!="

    LT = "<"
    GT = ">"

    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "{"

    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)

// in keywords, string mean identifier
// and this is connecting identifier to TokenType
var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

// parameter is identifier
// this function would return a TokenType already made and connected to the identifier
// if there is no TokenType to be connected to given identifier(keyword),
// return 'IDENT' TokenType.
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
