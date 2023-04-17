enum expType {
    number = 1,
    opsum,
    opdel,
    opmul,
    opdiv,

}

interface item {
    type: expType.number | expType.opsum | expType.opmul | expType.opdiv,
    value: any,
    priority? : number
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

    peek() {
        if ((this.pos) < this.root.length) {
            return this.root[this.pos]
        }
        return null
    }

    evalExp(level: number) {
        let result = null;
        let mode = null;

        while (true) {
            this.next()

            if (this.eof) {
                break;
            }

            let item = this.item;
            if (item.type == expType.number) {
                if (result === null) {
                    result = item.value;
                    continue;
                }

                const nextItem = this.peek();

                if(nextItem && nextItem.type == expType.opsum || nextItem.type == expType.opmul){
                    if(nextItem.l=)
                }

                if (mode == 1) {
                    result = sumPar(result, item.value);
                }
            }

            if (item.type == expType.opsum || item.type == expType.opmul) {
                mode = 1
                continue;
            }
        }



        return result;

    }

    public decode() {
        return this.evalExp(0)

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

const data: item[] = [
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
        value: 9,
    },
    {
        type: expType.opmul,
        value: "*",
        priority : 2,
    },
    {
        type: expType.number,
        value: 3,
    }
]

const tree = new Tree(data);



console.log("logica", tree.decode())