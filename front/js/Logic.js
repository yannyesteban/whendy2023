var expType;
(function (expType) {
    expType[expType["number"] = 1] = "number";
    expType[expType["opsum"] = 2] = "opsum";
    expType[expType["opdel"] = 3] = "opdel";
    expType[expType["opmul"] = 4] = "opmul";
    expType[expType["opdiv"] = 5] = "opdiv";
})(expType || (expType = {}));
class Tree {
    constructor(data) {
        this.pos = 0;
        this.value = null;
        this.typ = null;
        this.eof = false;
        this.mode = null;
        this.item = null;
        this.root = data;
    }
    next() {
        if (this.pos < this.root.length) {
            this.value = this.root[this.pos].value;
            this.typ = this.root[this.pos].type;
            this.item = this.root[this.pos];
            this.pos++;
        }
        else {
            this.eof = true;
        }
    }
    peek() {
        if ((this.pos) < this.root.length) {
            return this.root[this.pos];
        }
        return null;
    }
    evalExp(level) {
        let result = null;
        let mode = null;
        while (true) {
            this.next();
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
                if (nextItem && nextItem.type == expType.opsum || nextItem.type == expType.opmul) {
                    if (nextItem.l = )
                        ;
                }
                if (mode == 1) {
                    result = sumPar(result, item.value);
                }
            }
            if (item.type == expType.opsum || item.type == expType.opmul) {
                mode = 1;
                continue;
            }
        }
        return result;
    }
    decode() {
        return this.evalExp(0);
    }
}
function sumPar(a, b) {
    return a + +b;
}
const delPar = (a, b) => {
    return a - b;
};
const mulPar = (a, b) => {
    return a * b;
};
const divPar = (a, b) => {
    return a / b;
};
const data = [
    {
        type: expType.number,
        value: 5,
    },
    {
        type: expType.opsum,
        value: "+",
        priority: 1,
    },
    {
        type: expType.number,
        value: 9,
    },
    {
        type: expType.opmul,
        value: "*",
        priority: 2,
    },
    {
        type: expType.number,
        value: 3,
    }
];
const tree = new Tree(data);
console.log("logica", tree.decode());
//# sourceMappingURL=Logic.js.map