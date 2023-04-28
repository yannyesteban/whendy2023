enum Token {
	IDENT = 1, // main
	EOF,
	COMMENT,
	INT, // 12345
	FLOAT, // 123.45
	//IMAG,   // 123.45i
	//CHAR,   // 'a'
	STRING, // "abc"
	
	

	ADD, // +
	SUB, // -
	MUL, // *
	DIV, // /
	MOD, // %

	BIT_AND,     // &
	BIT_OR,      // |
	BIT_XOR,     // ^
	BIT_SHL,     // <<
	BIT_SHR,     // >>
	BIT_AND_NOT, // &^

	LPAREN, // (
	LBRACK, // [
	LBRACE, // {
	COMMA, // ,
	DOT, // .

	RPAREN, // )
	RBRACK, // ]
	RBRACE, // }
	SEMICOLON, // ;
	COLON, // :

	ADD_ASSIGN, // +=
	SUB_ASSIGN, // -=
	MUL_ASSIGN, // *=
	DIV_ASSIGN, // /=
	MOD_ASSIGN, // %=

	AND, // &&
	OR, // ||

	INCR, // ++
	DECR, // --
	POW, // **

	EQL, // ==
	LSS, // <
	GTR, // >
	ASSIGN, // =
	NOT, // !

	NEQ, // !=
	LEQ, // <=
	GEQ, // >=
	LET, // :=

	SYMBOL,


	IF, // "if"
	ELSE, // "else"
	CASE, // "case"
	WHEN, // "when"
	WHILE, // "while"
	DEFAULT, // "default"
	FOR, // "for"
	EACH, // "each"
}

const keywords = {
	if: Token.IF,
	ELSE: Token.ELSE,
	case: Token.CASE,
	when: Token.WHEN,
	while: Token.WHILE,
	default: Token.DEFAULT,
	for: Token.FOR,
	each: Token.EACH,

	
};

class Keyword {
	constructor() { }

	isKeyword(key) {
		return keywords[key] ?? Token.IDENT;
	}
	
}

console.log(keywords["f"]);
