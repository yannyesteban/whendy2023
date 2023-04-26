function isLetter(ch){

    if(ch>="a" && ch<="z" || ch>="A" && ch<="Z"){
        return true;
    }
    return false;
}

function isKey(){


}

function isAlphaNumeric(ch){

    if(ch>="a" && ch<="z" || ch>="A" && ch<="Z" || ch>="0" && ch<="9" || ch=="_"){
        return true;
    }
    return false;

}