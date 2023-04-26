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
    DEFAULT, // "abc"
    FOR, // "abc"
    EACH, 
}

const keywords = {
    "if":tkType.IF,
    "ELSE":tkType.ELSE,
    "case":tkType.CASE,
    "when":tkType.WHEN,
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