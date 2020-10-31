const axios = require("axios");
const clc = require("cli-color");

async function Run(provider, token, user, project) {
    axios.post(
        `https://circleci.com/api/v1.1/project/${provider}/${user}/${project}/follow`,
        {},
        {headers: {Accept: 'application/json', 'Circle-Token': token}}
        ).then(() => {
        console.log(clc.green("CircleCI successfully followed your project"));
    }).catch(error => {
        console.log(clc.red("Error while following project"), error.response);
    });
}

const formula = Run;
module.exports = formula;
