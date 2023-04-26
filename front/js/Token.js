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
    tkType[tkType["DEFAULT"] = 11] = "DEFAULT";
    tkType[tkType["FOR"] = 12] = "FOR";
    tkType[tkType["EACH"] = 13] = "EACH";
})(tkType || (tkType = {}));
const keywords = {
    "if": tkType.IF,
    "ELSE": tkType.ELSE,
    "case": tkType.CASE,
    "when": tkType.WHEN,
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