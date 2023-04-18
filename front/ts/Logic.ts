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
        if (this.pos < this.root.length) {

            this.value = this.root[this.pos].value;
            this.typ = this.root[this.pos].type;

            this.item = this.root[this.pos];
            this.pos++;
        } else {
            this.eof = true;
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

    evalExp(level: number, priority: number) {
        console.log("*****")
        //console.log("-----", [this.value, ...arguments])
        let result = null;
        let mode = null;
        let value;
        while (true) {

            console.log("Value", this.value)

            if (this.eof) {
                break;
            }


            let item = this.item;
            if (item.type == expType.number) {
                if (result === null) {
                    result = item.value;
                    
                    this.next()
                    continue;
                }
                console.log(this.item)
                const nextItem = this.peek();

                if(nextItem){

                    if(nextItem.toRight){
                        console.log("POWER", nextItem.priority , priority)
                        if(nextItem.priority >= priority){
                            value = this.evalExp(++level, nextItem.priority);
                            result = resolve(result, value, mode)

                        }else{
                            //result = resolve(result, item.value, mode)
                        }
                    }

                    if(nextItem.type == expType.opsum || nextItem.type == expType.opmul){
                        if (nextItem.priority > priority) {
                            console.log(777, nextItem.priority)
    
                            value = this.evalExp(++level, nextItem.priority)
    
                            console.log(` VALUE < ${value}>`)
                            result = resolve(result, value, mode)
                        } else {
                            console.log(888)
                            result = resolve(result, item.value, mode)
                        }
                    }

                }else {
                    console.log(999, result, item.value, mode)
                    result = resolve(result, item.value, mode)
                }
                /*
                if (nextItem !== null && (nextItem.type == expType.opsum || nextItem.type == expType.opmul)) {
                    if (nextItem.priority > priority) {
                        console.log(777, nextItem.priority)

                        value = this.evalExp(++level, nextItem.priority)

                        console.log(` VALUE < ${value}>`)
                        result = resolve(result, value, mode)
                    } else {
                        console.log(888)
                        result = resolve(result, item.value, mode)
                    }


                } else {
                    console.log(999, result, item.value, mode)
                    result = resolve(result, item.value, mode)
                }


                if (mode == 1) {
                    //result = sumPar(result, item.value);
                }
                */
            }

            if (item.type == expType.opsum || item.type == expType.opmul || item.type == expType.oppow) {

                mode = item.type;
                priority = item.priority;
                console.log("Priority ", priority)

            }
            this.next()
        }



        return result;

    }

    public decode() {
        return this.evalExp(0, 1)

    }
}





function resolve(a, b, op) {
    
    switch (op) {

        case expType.opsum:
            return sumPar(a, b);

        case expType.opmul:
            return mulPar(a, b);
        case expType.oppow:
            return powPar(a, b);
    }

}


function sumPar(a, b) {
    return a + +b

}

const delPar = (a, b) => {
    return a - b

}

const mulPar = (a, b) => {
    return a * b

}

const divPar = (a, b) => {
    return a / b

}

const powPar = (a, b) => {
    
    return a ** b

}

let data: item[] = [
    {
        type: expType.number,
        value: 4,
    },
    {
        type: expType.opsum,
        value: "+",
        priority: 1,
    },
    {
        type: expType.number,
        value: 1,
    },
    {
        type: expType.opsum,
        value: "+",
        priority: 1,
    },
    {
        type: expType.number,
        value: 2,
    },
    {
        type: expType.opmul,
        value: "*",
        priority : 2,
    },
    {
        type: expType.number,
        value: 3,
    },
    {
        type: expType.opmul,
        value: "*",
        priority : 2,
    },
    {
        type: expType.number,
        value: 5,
    },
    {
        type: expType.opsum,
        value: "+",
        priority : 1,
    },
    {
        type: expType.number,
        value: 6,
    }
];

data = [
    {
        type: expType.number,
        value: 2,
    },
    {
            type: expType.oppow,
            value: "^",
            priority: 3,
            toRight: true,
    },
    {
        type: expType.number,
        value: 2,
    },
    {
            type: expType.oppow,
            value: "^",
            priority: 3,
            toRight: true,
    },
    {
        type: expType.number,
        value: 3,
    },
    {
        type: expType.opsum,
        value: "+",
        priority : 1,
    },
    {
        type: expType.number,
        value: 4,
    }
];



const tree = new Tree(data);

tree.next()

console.log(` RESULT << ${tree.decode()} >>`)
//https://developer.mozilla.org/es/docs/Web/JavaScript/Reference/Operators/Operator_precedence
