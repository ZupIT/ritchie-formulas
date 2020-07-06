const run = require("./coffee/coffee")

const NAME = process.env.NAME
const COFFEE_TYPE = process.env.COFFEE_TYPE
const DELIVERY = process.env.DELIVERY

run(NAME, COFFEE_TYPE, DELIVERY)