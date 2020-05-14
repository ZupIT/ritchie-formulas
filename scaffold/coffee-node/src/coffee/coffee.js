function Run(name, coffeeType, delivery) {
    console.log("Preparing your coffee "+name+" .....")
    sleep(1)
    console.log("......")
    sleep(1)
    console.log("......")
    sleep(1)
    console.log("......")
    sleep(1)

    if (delivery == 'true') {
        console.log("Your "+coffeeType+" coffee is ready, enjoy your trip")
    } else {
        console.log("Your "+coffeeType+" coffee is ready, have a seat and enjoy your drink")
    }
}

function msleep(n) {
    Atomics.wait(new Int32Array(new SharedArrayBuffer(4)), 0, 0, n);
}
function sleep(n) {
    msleep(n*1000);
}

const coffee = Run
module.exports = coffee