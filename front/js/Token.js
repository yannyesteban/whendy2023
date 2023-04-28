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
    tkType[tkType["ADD"] = 15] = "ADD";
    tkType[tkType["SUB"] = 16] = "SUB";
    tkType[tkType["MUL"] = 17] = "MUL";
    tkType[tkType["QUO"] = 18] = "QUO";
    tkType[tkType["REM"] = 19] = "REM";
    tkType[tkType["LPAREN"] = 20] = "LPAREN";
    tkType[tkType["LBRACK"] = 21] = "LBRACK";
    tkType[tkType["LBRACE"] = 22] = "LBRACE";
    tkType[tkType["COMMA"] = 23] = "COMMA";
    tkType[tkType["PERIOD"] = 24] = "PERIOD";
    tkType[tkType["RPAREN"] = 25] = "RPAREN";
    tkType[tkType["RBRACK"] = 26] = "RBRACK";
    tkType[tkType["RBRACE"] = 27] = "RBRACE";
    tkType[tkType["SEMICOLON"] = 28] = "SEMICOLON";
    tkType[tkType["COLON"] = 29] = "COLON";
    tkType[tkType["ADD_ASSIGN"] = 30] = "ADD_ASSIGN";
    tkType[tkType["SUB_ASSIGN"] = 31] = "SUB_ASSIGN";
    tkType[tkType["MUL_ASSIGN"] = 32] = "MUL_ASSIGN";
    tkType[tkType["QUO_ASSIGN"] = 33] = "QUO_ASSIGN";
    tkType[tkType["REM_ASSIGN"] = 34] = "REM_ASSIGN";
    tkType[tkType["AND_ASSIGN"] = 35] = "AND_ASSIGN";
    tkType[tkType["OR_ASSIGN"] = 36] = "OR_ASSIGN";
    tkType[tkType["XOR_ASSIGN"] = 37] = "XOR_ASSIGN";
    tkType[tkType["SHL_ASSIGN"] = 38] = "SHL_ASSIGN";
    tkType[tkType["SHR_ASSIGN"] = 39] = "SHR_ASSIGN";
    tkType[tkType["AND_NOT_ASSIGN"] = 40] = "AND_NOT_ASSIGN";
    tkType[tkType["LAND"] = 41] = "LAND";
    tkType[tkType["LOR"] = 42] = "LOR";
    tkType[tkType["ARROW"] = 43] = "ARROW";
    tkType[tkType["INC"] = 44] = "INC";
    tkType[tkType["DEC"] = 45] = "DEC";
    tkType[tkType["EQL"] = 46] = "EQL";
    tkType[tkType["LSS"] = 47] = "LSS";
    tkType[tkType["GTR"] = 48] = "GTR";
    tkType[tkType["ASSIGN"] = 49] = "ASSIGN";
    tkType[tkType["NOT"] = 50] = "NOT";
    tkType[tkType["NEQ"] = 51] = "NEQ";
    tkType[tkType["LEQ"] = 52] = "LEQ";
    tkType[tkType["GEQ"] = 53] = "GEQ";
    tkType[tkType["DEFINE"] = 54] = "DEFINE";
    tkType[tkType["ELLIPSIS"] = 55] = "ELLIPSIS";
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
    ADD_ASSIGN: "+=",
    SUB_ASSIGN: "-=",
    MUL_ASSIGN: "*=",
    QUO_ASSIGN: "/=",
    REM_ASSIGN: "%=",
    AND_ASSIGN: "&=",
    OR_ASSIGN: "|=",
    XOR_ASSIGN: "^=",
    SHL_ASSIGN: "<<=",
    SHR_ASSIGN: ">>=",
    AND_NOT_ASSIGN: "&^=",
    LAND: "&&",
    LOR: "||",
    ARROW: "<-",
    INC: "++",
    DEC: "--",
    EQL: "==",
    LSS: "<",
    GTR: ">",
    ASSIGN: "=",
    NOT: "!",
    NEQ: "!=",
    LEQ: "<=",
    GEQ: ">=",
    DEFINE: ":=",
    ELLIPSIS: "...",
    LPAREN: "(",
    LBRACK: "[",
    LBRACE: "{",
    COMMA: ",",
    PERIOD: ".",
    RPAREN: ")",
    RBRACK: "]",
    RBRACE: "}",
    SEMICOLON: ";",
    COLON: ":",
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