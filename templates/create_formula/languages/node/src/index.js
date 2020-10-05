var path = require('path'),
run = require("./formula/formula"),
config_data = require(path.resolve('../config.json'));

const INPUT1 = process.env.INPUT_TEXT
const INPUT2 = process.env.INPUT_BOOLEAN
const INPUT3 = process.env.INPUT_LIST
const INPUT4 = process.env.INPUT_PASSWORD

run(INPUT1, INPUT2, INPUT3, INPUT4)
