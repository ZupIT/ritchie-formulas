const { start } = require('@mhlabs/iam-policies-cli/src/input-wizard');

function Run(ritFormat, ritOutput) {
  start("", ritFormat, ritOutput);
}

const formula = Run
module.exports = formula
