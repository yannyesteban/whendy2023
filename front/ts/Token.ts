enum tkType {
	IDENT =1, // main
	INT,    // 12345
	FLOAT,  // 123.45
	IMAG,   // 123.45i
	CHAR,   // 'a'
	STRING, // "abc"
    IF, // "abc"
    ELSE, // "abc"
    CASE, // "abc"
    WHEN, // "abc"
    WHILE, // "abc"
    DEFAULT, // "abc"
    FOR, // "abc"
    EACH, 
    
    ADD, // +
	SUB, // -
	MUL, // *
	QUO, // /
	REM, // %
    
    LPAREN, // (
        LBRACK, // [
        LBRACE, // {
        COMMA,  // ,
        PERIOD, // .
    


        RPAREN,    // )
        RBRACK,    // ]
        RBRACE,    // }
        SEMICOLON, // ;
        COLON,     // :

        ADD_ASSIGN, // +=
        SUB_ASSIGN, // -=
        MUL_ASSIGN, // *=
        QUO_ASSIGN, // /=
        REM_ASSIGN, // %=
    
        AND_ASSIGN,     // &=
        OR_ASSIGN,      // |=
        XOR_ASSIGN,     // ^=
        SHL_ASSIGN,     // <<=
        SHR_ASSIGN,     // >>=
        AND_NOT_ASSIGN, // &^=




	LAND,  // &&
	LOR,   // ||
	ARROW, // <-
	INC,   // ++
	DEC,   // --

	EQL,    // ==
	LSS,    // <
	GTR,    // >
	ASSIGN, // =
	NOT,    // !

	NEQ,      // !=
	LEQ,      // <=
	GEQ,      // >=
	DEFINE,   // :=
	ELLIPSIS, // ...
}

const keywords = {
    "if":tkType.IF,
    "ELSE":tkType.ELSE,
    "case":tkType.CASE,
    "when":tkType.WHEN,
    "while":tkType.WHILE,
    "default":tkType.DEFAULT,
    "for":tkType.FOR,
    "each":tkType.EACH,

    ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

    LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",
}

class Token{

    constructor(){

    }

    isKeyword(key){
        return keywords[key] ?? false;
    }
}

console.log(keywords["f"])