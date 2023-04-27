var token = new Token();

class Lexer {

    public input: string = "";
    public pos: number = null;
    public eof: boolean = false;
    public ch: string = "";

    constructor(input: string) {
        this.input = input;
        this.pos = 1;
        this.ch = this.input[0];
        this.eof = !(this.ch);

        console.log(input);
    }

    scanIdentifier(){
        let init = this.pos;
        let end = init;
        let lit = "";
        while(!this.eof){
            let ch = this.ch;

            if(isAlphaNumeric(ch)){
                
                end++;
                lit += this.ch
                
                this.next();
                continue;
                
            }

            break;
        }

        return lit;
    }

    scanNumber  () {
        let offs = this.pos
        let tok = tkType.INT
    
        let base = 10        // number base
        let prefix = "0" // one of 0 (decimal), '0' (0-octal), 'x', 'o', or 'b'
        //digsep = 0       // bit 0: digit present, bit 1: '_' present
        //invalid = -1     // index of invalid digit in literal, or < 0
    
        let lit = "c"
        // integer part
        if (this.ch != '.') {
            tok = tkType.INT
            if (this.ch == '0') {
                this.next();
                switch (this.ch.toLowerCase()) {
                case 'x':
                    this.next()
                    [base, prefix] = [16, 'x']
                case 'o':
                    this.next()
                    [base, prefix] = [8, 'o'];
                case 'b':
                    this.next()
                    [base, prefix] = [2, 'b']
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
            tok = tkType.FLOAT
            if (prefix == 'o' || prefix == 'b' ){
                //s.error(s.offset, "invalid radix point in "+litname(prefix))
            }
            this.next()
            //digsep |= s.digits(base, &invalid)
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
            tok = tkType.FLOAT
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
       console.log("<>",this.input.substring(offs-1, this.pos))
        return lit
        //return {tok, lit}
    }

    digits(base: number)  {
        if (base <= 10) {
            //max := rune('0' + base)
            while(isDecimal(this.ch)) {
                
                this.next()
            }
        } else {
            while(isHex(this.ch)) {
                
                this.next()
            }
        }
        return
    }

    

    scan() {
        
        let ch = null;
        let lit = "";
        let typ:number = null;
        while (!this.eof) {
            ch = this.ch;
            console.log(this.ch)

            
            if(isLetter(ch) || ch=="_"){
                lit = this.scanIdentifier();
                console.log(".....", lit)
                if(lit.length>1){
                    typ = token.isKeyword(lit);
                }
                break;
                
            }else if(isDecimal(ch) || ch == '.' && isDecimal(this.peek())){

                lit = this.scanNumber()
            }

            this.next();
        }

        return {
            pos: this.pos,
            value: lit,
            typ
        };


    }

    next() {

        if (this.pos < this.input.length) {
            
            this.ch = this.input[this.pos];
            this.pos++;
        } else {
            this.eof = true;
        }

    }

    peek(){
        if (this.pos+1 < this.input.length) {
            return this.input[this.pos+1];
        }
        return null;
    }

    getTokens(){
        let tokens = [];
        while (!this.eof) {
            
            tokens.push(this.scan())

        }
        return tokens;
    }

}

const source = `while if for while  987.368 if`;
let lexer = new Lexer(source);

console.log(lexer.getTokens());