var token = new Token();
class Lexer {
    constructor(input) {
        this.input = "";
        this.pos = null;
        this.eof = false;
        this.ch = "";
        this.input = input;
        this.pos = 0;
        this.ch = " ";
        this.eof = false;
        this.next();
        console.log(input);
    }
    skipWhitespace() {
        while (this.ch == ' ' || this.ch == '\t' || this.ch == '\n' || this.ch == '\r') {
            console.log("[][][]");
            this.next();
        }
    }
    scanIdentifier() {
        let init = this.pos;
        let end = init;
        let lit = "";
        while (!this.eof) {
            let ch = this.ch;
            if (isAlphaNumeric(ch)) {
                end++;
                lit += this.ch;
                this.next();
                continue;
            }
            break;
        }
        return lit;
    }
    scanNumber() {
        let offs = this.pos;
        let tok = tkType.INT;
        let base = 10; // number base
        let prefix = "0"; // one of 0 (decimal), '0' (0-octal), 'x', 'o', or 'b'
        //digsep = 0       // bit 0: digit present, bit 1: '_' present
        //invalid = -1     // index of invalid digit in literal, or < 0
        let lit = "c";
        // integer part
        if (this.ch != '.') {
            tok = tkType.INT;
            if (this.ch == '0') {
                this.next();
                switch (this.ch.toLowerCase()) {
                    case 'x':
                        this.next();
                        [base, prefix] = [16, 'x'];
                        break;
                    case 'o':
                        this.next();
                        ([base, prefix] = [8, 'o']);
                        break;
                    case 'b':
                        this.next();
                        [base, prefix] = [2, 'b'];
                        break;
                    default:
                        [base, prefix] = [8, '0'];
                    //digsep = 1 // leading 0
                }
            }
            //digsep |= 
            this.digits(base);
        }
        // fractional part
        if (this.ch == '.') {
            tok = tkType.FLOAT;
            if (prefix == 'o' || prefix == 'b') {
                //s.error(s.offset, "invalid radix point in "+litname(prefix))
            }
            this.next();
            this.digits(base);
        }
        /*
        if digsep&1 == 0 {
            s.error(s.offset, litname(prefix)+" has no digits")
        }
        */
        // exponent
        const e = this.ch.toLowerCase();
        if (e == 'e' || e == 'p') {
            /*
            switch {
            case e == 'e' && prefix != 0 && prefix != '0':
                s.errorf(s.offset, "%q exponent requires decimal mantissa", s.ch)
            case e == 'p' && prefix != 'x':
                s.errorf(s.offset, "%q exponent requires hexadecimal mantissa", s.ch)
            }
            */
            this.next();
            tok = tkType.FLOAT;
            if (this.ch == '+' || this.ch == '-') {
                this.next();
            }
            let ds = this.digits(10);
            /*digsep |= ds
            if ds&1 == 0 {
                s.error(s.offset, "exponent has no digits")
            }
            */
        }
        /*
        else if prefix == 'x' && tok == token.FLOAT {
            s.error(s.offset, "hexadecimal mantissa requires a 'p' exponent")
        }
    
        // suffix 'i'
        if s.ch == 'i' {
            tok = token.IMAG
            s.next()
        }
    
        lit = string(s.src[offs:s.offset])
        if tok == token.INT && invalid >= 0 {
            s.errorf(invalid, "invalid digit %q in %s", lit[invalid-offs], litname(prefix))
        }
        if digsep&2 != 0 {
            if i = invalidSep(lit); i >= 0 {
                s.error(offs+i, "'_' must separate successive digits")
            }
        }
        */
        lit = this.input.substring(offs - 1, this.pos);
        console.log("<>", offs, this.pos, " = ", lit);
        //return lit
        return { lit, tok };
    }
    digits(base) {
        if (base <= 10) {
            //max := rune('0' + base)
            while (isDecimal(this.ch)) {
                this.next();
            }
        }
        else {
            while (isHex(this.ch)) {
                this.next();
            }
        }
        return;
    }
    scan() {
        let ch = null;
        let lit = "";
        let tok = null;
        while (!this.eof) {
            this.skipWhitespace();
            ch = this.ch;
            console.log(this.ch);
            if (isLetter(ch) || ch == "_") {
                lit = this.scanIdentifier();
                console.log(".....", lit);
                if (lit.length > 1) {
                    tok = token.isKeyword(lit);
                }
                break;
            }
            else if (isDecimal(ch) || ch == '.' && isDecimal(this.peek())) {
                ({ lit, tok } = this.scanNumber());
                //console.log("this.scanNumber()", this.scanNumber())
                //{lit, tok} = this.scanNumber()
                console.log(lit, tok);
                break;
            }
            else {
                this.next();
                switch (ch) {
                    case ",":
                        tok = tkType.COMMA;
                        lit = ",";
                        break;
                    case ";":
                        tok = tkType.SEMICOLON;
                        lit = ";";
                        break;
                    case "+":
                        if (this.ch == "=") {
                            tok = tkType.AND_ASSIGN;
                            lit = "+=";
                            this.next();
                            break;
                        }
                        if (this.ch == ch) {
                            tok = tkType.INC;
                            lit = "++";
                            this.next();
                            break;
                        }
                        tok = tkType.ADD;
                        lit = "+";
                        break;
                    case "=":
                        tok = tkType.ASSIGN;
                        lit = "=";
                        break;
                }
            }
            break;
        }
        console.log(lit, tok);
        return {
            pos: this.pos,
            value: lit,
            tok
        };
    }
    next() {
        if (this.pos < this.input.length) {
            this.ch = this.input[this.pos];
            this.pos++;
        }
        else {
            this.eof = true;
            this.ch = "\0";
        }
    }
    peek() {
        if (this.pos + 1 < this.input.length) {
            return this.input[this.pos + 1];
        }
        return null;
    }
    getTokens() {
        let tokens = [];
        while (!this.eof) {
            tokens.push(this.scan());
        }
        return tokens;
    }
}
let source = `while if for while  987.368  5e+3 -24 0xaf01 0b2 if yanny, esteban, hello; wait; test 4==5 6=3 2+2 k+=8`;
source = `3 5`;
let lexer = new Lexer(source);
console.log(lexer.getTokens());
`
if(a>1){"445"}else{"aaa"};a=45;case(a){when(1){ 456}when(2){100}default{3001}}

`;
//# sourceMappingURL=Lexer.js.map