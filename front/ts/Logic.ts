enum expType {
    number = 1,
    opsum,
    opdel,
    opmul,
    opdiv,
    oppow,
}

interface item {
    type: expType.number | expType.opsum | expType.opmul | expType.opdiv | expType.oppow,
    value: any,
    priority?: number,
    toRight?: boolean
}


function sep(str){
    const items:item[] = [];
    for (let chr of str) {
        if(chr == " "){
            continue;
        }
        if(chr >="0" && chr<="9"){
            items.push({
                type:expType.number,
                value:+chr
            })
        }
        if(chr =="+"){
            items.push({
                type:expType.opsum,
                priority:1,
                value:chr
            })
        }
        if(chr =="*"){
            items.push({
                type:expType.opmul,
                priority:2,
                value:chr,
               
            })
        }
        if(chr =="^"){
            items.push({
                type:expType.oppow,
                priority:3,
                value:chr,
                toRight:true,
                

            })
        }
        
    }
    return items

}



class Tree {
    public root: item[];
    public pos = 0;

    public value = null
    public typ = null

    public eof = false;

    public mode: number = null;
    public item: item = null;

    constructor(data) {
        this.root = data
    }

    next() {
        console.log('NEXT...')
        if (this.pos < this.root.length) {

            this.value = this.root[this.pos].value;
            this.typ = this.root[this.pos].type;

            this.item = this.root[this.pos];
            this.pos++;
        } else {
            this.eof = true;
            this.value = null;
            this.typ = null;
        }

    }

    back() {
        this.item = this.root[this.pos - 1];
        this.value = this.root[this.pos - 1].value;
        this.typ = this.root[this.pos - 1].type;
        this.pos--;
    }
    peek() {
        if ((this.pos) < this.root.length) {
            return this.root[this.pos]
        }
        return null
    }

    evalExp2(level: number, priority: number) {
        
        //console.log("-----", [this.value, ...arguments])
        let result = null;
        let mode = null;
        let value;
        while (true) {

            if (this.eof) {
                break;
            }
            console.log("Value", this.value)

            let item = this.item;
            if (item.type == expType.number) {
                //console.log("Number .--->", item.value)
                if (result === null) {
                    result = item.value;
                    
                    this.next()
                    continue;
                }
                //console.log(this.item)
                const nextItem = this.peek();

                if(nextItem){

                    if(nextItem.priority < priority){
                        console.log("Partial ", result)
                        return result
                    }

                    console.log("TYPE: ", item.type)
                    if(nextItem.toRight){
                        console.log("POWER", nextItem.priority , priority)
                        if(nextItem.priority >= priority){
                            value = this.evalExp(++level, nextItem.priority);
                            console.log(` VALUE < ${value}>`)
                            result = resolve(result, value, mode)

                        }else{

                            //result = resolve(result, item.value, mode)
                        }
                        
                    }

                    
                    

                    if(nextItem.type == expType.opsum || nextItem.type == expType.opmul){
                        if (nextItem.priority > priority) {
                            //console.log(777, nextItem.priority)
    
                            value = this.evalExp(++level, nextItem.priority)
    
                            console.log(` VALUE < ${value}>`)
                            result = resolve(result, value, mode)
                        } else {
                            //console.log("que ", result, item.value, mode)
                            result = resolve(result, item.value, mode)
                        }
                    }

                }else {
                    //console.log(999, result, item.value, mode)
                    result = resolve(result, item.value, mode)
                }
                
            }

            //console.log(item)
            if (item.type == expType.opsum || item.type == expType.opmul || item.type == expType.oppow) {

                mode = item.type;
                priority = item.priority;
                //console.log("Priority ", priority)

            }
            this.next()
        }



        return result;

    }

    evalExp(level){
        let partial = null;
        let op = null;
        let priority = null;

        let value = null
        while (true) {
            
            if (this.eof) {
                return partial;
            }

            const item = this.item;

            if(item.type >= expType.opsum){
                op = item.type;
                priority = item.priority;
                this.next();
                continue;

            }
            
           
            console.log("Value: ", item.value)

            if(item.type == expType.number){

                value = item.value;

                if(partial === null){
                    partial = value;
                    this.next();
                    continue;
                    
                }
                
                
                let peek = this.peek();

                if(!peek){
                    partial = resolve(partial, value, op);
                    this.next();
                    continue;
                }


                if(peek.priority == priority){
                    if(peek.toRight){

                    }else{
                        partial = resolve(partial, value, op);
                        this.next();
                        continue;
                    }

                }

                if(peek){

                    if(peek.toRight){
                        if(peek.priority >= priority){
                            value = this.evalExp(++level);
                            partial = resolve(partial, value, op);
                            this.next();
                            continue;
                        }
                        partial = resolve(partial, value, op);
                        this.next();
                        continue;
                    }
                    

                    if(peek.priority == priority){
                        partial = resolve(partial, value, op);
                        this.next();
                        continue;
                    }

                    if (peek.priority > priority) {
                        
                        console.log("Exp Parcial 1.", partial)
                        value = this.evalExp(++level);
                        console.log("Exp Parcial 2.", value);
                        partial = resolve(partial, value, op);
                        this.next();
                        continue;
                        
                    }

                    if (peek.priority < priority) {
                       
                        if(level>0){
                            console.log("A: ",partial, item.value, op)
                            
                            //this.next()
                            return resolve(partial, value, op);
                        }
                        //let value = item.value
                        //this.next()
                        //return resolve(partial, value, op);
                    }
                    

                    

                    
                    
                }
                
                partial = resolve(partial, value, op);
                
            }

            

            
            this.next();
        }
        

    }

    evalTodo(){
        interface elem {
            type: expType.number | expType.opsum | expType.opmul | expType.opdiv | expType.oppow,
            value: any,
            priority?: number,
            toRight?: boolean
        }

        const levels:elem[] = [];
        let level = 0;
        let value = null;
        let peek:item = null;
        let index = 0
        while(true){
            if (this.eof) {
                break;
            }

            if(index++>=100){
                break;
            }


            if(!levels[level]){
                levels[level] = {
                    type:null,
                    value:null,
                }
            }

            const item = this.item;

            if(item.type >= expType.opsum){
                levels[level].type = item.type;
                levels[level].priority = item.priority;
                this.next();
                continue;

            }
            
            value = item.value;
            console.log("Value: ", item.value)

            if(levels[level].value === null){
                levels[level].value = value;
                this.next();
                continue;
            }

            peek = this.peek();
           
            if(peek && peek.priority>levels[level].priority){
                level++;
                continue;
            }
            console.log("*****************")
            if(!peek || peek.priority<levels[level].priority){
                console.log("<---- priority")
                levels[level].value = resolve(levels[level].value, value, levels[level].type);
                if(levels[level-1]){
                    console.log("error")
                    levels[level-1].value = resolve(levels[level-1].value, levels[level].value, levels[level-1].type)
                    level--;
                    this.next()
                    
                }else{
                    console.log("BIEN")
                    this.next()
                }

                
                continue;
            }
            
            levels[level].value = resolve(levels[level].value, value, levels[level].type);


            

            this.next();



        }
        console.log(levels)
    }

    public decode() {
        return this.evalTodo()

    }
}





function resolve(a, b, op) {
    console.log("resolve ----> ", a, op, b)
    
    switch (op) {

        case expType.opsum:
            console.log(`[${a} + ${b} = ${a+b}]`)
            return a + b

        case expType.opmul:
            console.log(`[${a} * ${b} = ${a*b}]`)
            return a * b
        case expType.oppow:
            console.log(`[${a} ** ${b} = ${a**b}]`)
            return a ** b
    }

}


const calc = "1+2+3 +4*5+1";//6+25+3+6+2

const tree = new Tree(sep(calc));

tree.next()

console.log(calc,` RESULT << ${tree.decode()} >>`)



function start(str){
    const tree = new Tree(sep(str));

    tree.next();
    return str+": "+tree.decode();
}
//https://developer.mozilla.org/es/docs/Web/JavaScript/Reference/Operators/Operator_precedence