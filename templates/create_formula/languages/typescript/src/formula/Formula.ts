import * as chalk from 'chalk'
import IFormula from './IFormula'

function run(iFormula: IFormula) {
    console.log('Hello World!')

    console.log(chalk.green(`My name is ${iFormula.text}.`))

    if (iFormula.boolean) {
        console.log(chalk.blue('I’ve already created formulas using Ritchie.'))
    } else {
        console.log(chalk.red('I’m excited in creating new formulas using Ritchie.'))
    }

    console.log(chalk.yellow(`Today, I want to automate ${iFormula.list}.`))

    console.log(chalk.cyan(`My secret is ${iFormula.password}.`))
}

const Formula = run
export default Formula
