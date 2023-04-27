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
}

class Token{

    constructor(){

    }

    isKeyword(key){
        return keywords[key] ?? false;
    }
}

console.log(keywords["f"])