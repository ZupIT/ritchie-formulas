var path = require('path'),
run = require("./formula/formula"),
config_data = require(path.resolve('../config.json'));

const INPUT1 = process.env.INPUT_TEXT||config_data.inputs[0].name
const INPUT2 = process.env.INPUT_BOOLEAN||config_data.inputs[1].name
const INPUT3 = process.env.INPUT_LIST||config_data.inputs[2].name
const INPUT4 = process.env.INPUT_PASSWORD||config_data.inputs[3].name

run(INPUT1, INPUT2, INPUT3, INPUT4)
