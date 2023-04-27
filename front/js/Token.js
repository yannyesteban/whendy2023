var tkType;
(function (tkType) {
    tkType[tkType["IDENT"] = 1] = "IDENT";
    tkType[tkType["INT"] = 2] = "INT";
    tkType[tkType["FLOAT"] = 3] = "FLOAT";
    tkType[tkType["IMAG"] = 4] = "IMAG";
    tkType[tkType["CHAR"] = 5] = "CHAR";
    tkType[tkType["STRING"] = 6] = "STRING";
    tkType[tkType["IF"] = 7] = "IF";
    tkType[tkType["ELSE"] = 8] = "ELSE";
    tkType[tkType["CASE"] = 9] = "CASE";
    tkType[tkType["WHEN"] = 10] = "WHEN";
    tkType[tkType["WHILE"] = 11] = "WHILE";
    tkType[tkType["DEFAULT"] = 12] = "DEFAULT";
    tkType[tkType["FOR"] = 13] = "FOR";
    tkType[tkType["EACH"] = 14] = "EACH";
    tkType[tkType["LPAREN"] = 15] = "LPAREN";
    tkType[tkType["LBRACK"] = 16] = "LBRACK";
    tkType[tkType["LBRACE"] = 17] = "LBRACE";
    tkType[tkType["COMMA"] = 18] = "COMMA";
    tkType[tkType["PERIOD"] = 19] = "PERIOD";
    tkType[tkType["RPAREN"] = 20] = "RPAREN";
    tkType[tkType["RBRACK"] = 21] = "RBRACK";
    tkType[tkType["RBRACE"] = 22] = "RBRACE";
    tkType[tkType["SEMICOLON"] = 23] = "SEMICOLON";
    tkType[tkType["COLON"] = 24] = "COLON";
})(tkType || (tkType = {}));
const keywords = {
    "if": tkType.IF,
    "ELSE": tkType.ELSE,
    "case": tkType.CASE,
    "when": tkType.WHEN,
    "while": tkType.WHILE,
    "default": tkType.DEFAULT,
    "for": tkType.FOR,
    "each": tkType.EACH,
};
class Token {
    constructor() {
    }
    isKeyword(key) {
        var _a;
        return (_a = keywords[key]) !== null && _a !== void 0 ? _a : false;
    }
}
console.log(keywords["f"]);
//# sourceMappingURL=Token.js.map