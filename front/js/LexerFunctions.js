function isLetter(ch) {
    if ((ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z")) {
        return true;
    }
    return false;
}
function isKey() { }
function isAlphaNumeric(ch) {
    if ((ch >= "a" && ch <= "z") ||
        (ch >= "A" && ch <= "Z") ||
        (ch >= "0" && ch <= "9") ||
        ch == "_") {
        return true;
    }
    return false;
}
function isDecimal(ch) {
    return ch >= "0" && ch <= "9";
}
function isHex(ch) {
    return '0' <= ch && ch <= '9' || 'a' <= ch.toLowerCase() && ch.toLowerCase() <= 'f';
}
//# sourceMappingURL=LexerFunctions.js.map