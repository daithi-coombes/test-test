function sayIt(msg) {
    if (msg == '') {
        console.log(msg);

        return;
    }

    return sayIt()
}
