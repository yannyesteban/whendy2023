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

    getTokens(){
        let tokens = [];
        while (!this.eof) {
            
            tokens.push(this.scan())

        }
        return tokens;
    }

}

const source = `if for while if`;
let lexer = new Lexer(source);

console.log(lexer.getTokens());