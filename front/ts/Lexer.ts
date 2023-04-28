var keyword = new Keyword();

class Lexer {

    public input: string = "";
    public pos: number = null;
    public rd: number = null;
    public eof: boolean = false;
    public ch: string = "";

    constructor(input: string) {
        this.input = input;
        this.pos = 0;
        this.rd = 0;
        this.ch = " ";
        this.eof = false;

        this.next()
        console.log(input);
    }

    evalOp(ch, tokenDefault, tokenAssign, tokenX2, tokenX3) {
        if (this.ch == "=") {
            this.next()
            return tokenAssign
        }
        if (tokenX2 && this.ch == ch) {
            this.next()
            if (tokenX3 && this.ch == ch) {
                this.next()
                return tokenX3
            }
            return tokenX2
        }
        return tokenDefault
    }



    skipWhitespace() {
        while (this.ch == ' ' || this.ch == '\t' || this.ch == '\n' || this.ch == '\r') {
            console.log("[][][]")
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
                lit += this.ch

                this.next();
                continue;

            }

            break;
        }

        return lit;
    }

    scanNumber() {
        let offs = this.pos;
        let tok = Token.INT;

        let base = 10;        // number base
        let prefix = "0"; // one of 0 (decimal), '0' (0-octal), 'x', 'o', or 'b'
        //digsep = 0       // bit 0: digit present, bit 1: '_' present
        //invalid = -1     // index of invalid digit in literal, or < 0

        let lit = "c"
        // integer part
        if (this.ch != '.') {
            tok = Token.INT
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
                        [base, prefix] = [2, 'b']
                        break;
                    default:
                        [base, prefix] = [8, '0']
                    //digsep = 1 // leading 0
                }
            }
            //digsep |= 
            this.digits(base)
        }

        // fractional part
        if (this.ch == '.') {
            tok = Token.FLOAT
            if (prefix == 'o' || prefix == 'b') {
                //s.error(s.offset, "invalid radix point in "+litname(prefix))
            }
            this.next()
            this.digits(base)

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
            this.next()
            tok = Token.FLOAT
            if (this.ch == '+' || this.ch == '-') {
                this.next()
            }
            let ds = this.digits(10)
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
        lit = this.input.substring(offs, this.pos)
        console.log("<>", offs, this.pos, " = ", lit)
        //return lit
        return { lit, tok }
    }

    digits(base: number) {
        if (base <= 10) {
            //max := rune('0' + base)
            while (isDecimal(this.ch)) {

                this.next()
            }
        } else {
            while (isHex(this.ch)) {

                this.next()
            }
        }
        return
    }



    scan() {

        let ch = null;
        let lit = "*** ERROR ***";

        let tok: number = null;
        let offs = 0;
        while (!this.eof) {
            this.skipWhitespace();

            ch = this.ch;
            offs = this.pos;
            console.log(this.ch)


            if (isLetter(ch) || ch == "_") {
                lit = this.scanIdentifier();
                console.log(".....", lit)
                if (lit.length > 1) {
                    tok = keyword.isKeyword(lit);
                } else {
                    tok = Token.IDENT;
                }
                break;

            } else if (isDecimal(ch) || ch == '.' && isDecimal(this.peek())) {


                ({ lit, tok } = this.scanNumber());

                //console.log("this.scanNumber()", this.scanNumber())
                //{lit, tok} = this.scanNumber()
                console.log(lit, tok)
                break;
            } else {

                this.next();
                switch (ch) {
                    case ":":
                        tok = this.evalOp(ch, Token.COLON, Token.LET, null, null);
                        lit = this.input.substring(offs, this.pos)
                        break;
                    case ".":
                        tok = Token.DOT;
                        lit = ".";
                        break;
                    case "(":
                        tok = Token.LPAREN;
                        lit = "(";
                        break;
                    case ")":
                        tok = Token.RPAREN;
                        lit = ")";
                        break;
                    case "[":
                        tok = Token.LBRACK;
                        lit = "[";
                        break;

                    case "]":
                        tok = Token.RBRACK;
                        lit = "]";
                        break;
                    case "{":
                        tok = Token.LBRACE;
                        lit = "{";
                        break;
                    case "}":
                        tok = Token.RBRACE;
                        lit = "}";
                        break;
                    case ",":
                        tok = Token.COMMA;
                        lit = ","
                        break;
                    case ";":
                        tok = Token.SEMICOLON;
                        lit = ";"
                        break;
                    case "+":
                        tok = this.evalOp(ch, Token.ADD, Token.ADD_ASSIGN, Token.INCR, null);
                        lit = this.input.substring(offs, this.pos);
                        break;
                    case "-":
                        tok = this.evalOp(ch, Token.SUB, Token.SUB_ASSIGN, Token.DECR, null);
                        lit = this.input.substring(offs, this.pos);
                        break;
                    case "*":
                        tok = this.evalOp(ch, Token.MUL, Token.MUL_ASSIGN, Token.POW, null);
                        lit = this.input.substring(offs, this.pos);
                        break;
                    case "/":
                        tok = this.evalOp(ch, Token.DIV, Token.DIV_ASSIGN, null, null);
                        lit = this.input.substring(offs, this.pos);
                        break;
                    case "%":
                        tok = this.evalOp(ch, Token.MOD, Token.MOD_ASSIGN, null, null);
                        lit = this.input.substring(offs, this.pos);
                        break;                        
                    case "=":
                        
                        tok = this.evalOp(ch, Token.ASSIGN, Token.EQL, null, null);
                        lit = this.input.substring(offs, this.pos);
                        
                        break;
                    case "!":
                        
                        tok = this.evalOp(ch, Token.NOT, Token.NEQ, null, null);
                        lit = this.input.substring(offs, this.pos);
                        
                        break;                        

                }



            }


            break;
        }
        console.log(lit, tok)
        return {
            pos: this.pos,
            value: lit,
            tok
        };


    }

    next() {

        if (this.rd < this.input.length) {
            this.pos = this.rd;
            this.ch = this.input[this.pos];
            this.rd++;
        } else {
            this.pos = this.input.length;
            this.eof = true;
            this.ch = "\0";
        }

    }

    peek() {
        if (this.pos < this.input.length) {
            return this.input[this.pos];
        }
        return null;
    }

    getTokens() {
        let tokens = [];
        while (!this.eof) {

            tokens.push(this.scan())

        }
        return tokens;
    }

}

let source = `while if for while  987.368  5e+3 -24 0xaf01 0b2 if yanny, esteban, hello; wait; test 4==5 6=3 2+2 k+=8`;

source = `if(a!=1){g:=!s}{e=6%3}`;
let lexer = new Lexer(source);

console.log(source, "\n", lexer.getTokens());

`
if(a>1){"445"}else{"aaa"};a=45;case(a){when(1){ 456}when(2){100}default{3001}}

`