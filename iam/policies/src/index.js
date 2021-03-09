const run = require("./formula/formula");

const ritFormat = process.env.RIT_FORMAT;
const ritOutput = process.env.RIT_OUTPUT;

run(ritFormat, ritOutput);
