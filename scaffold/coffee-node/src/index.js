const NAME = process.env.NAME
const COFFEE_TYPE = process.env.COFFEE_TYPE
const DELIVERY = process.env.DELIVERY

console.log(`Preparing your coffee ${NAME} .....`)
sleep(1)
console.log("......")
sleep(1)
console.log("......")
sleep(1)
console.log("......")
sleep(1)

if (DELIVERY == 'true') {
    console.log(`Your ${COFFEE_TYPE} coffee is ready, enjoy your trip`)
} else {
    console.log(`Your ${COFFEE_TYPE} coffee is ready, have a seat and enjoy your drink`)
}


function msleep(n) {
    Atomics.wait(new Int32Array(new SharedArrayBuffer(4)), 0, 0, n);
}
function sleep(n) {
    msleep(n*1000);
}