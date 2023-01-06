
let sentence = '';

function sayIt(msg) {
    if (msg == '') {
        return;
    }
    console.log(msg);

    return sayIt;
}

sayIt('my')('name', 'is')('jonathan')()