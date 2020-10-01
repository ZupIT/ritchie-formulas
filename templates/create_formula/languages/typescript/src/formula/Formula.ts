import * as chalk from 'chalk'

function run(text: string, boolean: boolean, list: string, password: string) {
    console.log('Hello World!')

    console.log(chalk.green(`My name is ${text}.`))

    if (boolean) {
        console.log(chalk.blue('I’ve already created formulas using Ritchie.'))
    } else {
        console.log(chalk.red('I’m excited in creating new formulas using Ritchie.'))
    }

    console.log(chalk.yellow(`Today, I want to automate ${list}.`))

    console.log(chalk.cyan(`My secret is ${password}.`))
}

const Formula = run
export default Formula
