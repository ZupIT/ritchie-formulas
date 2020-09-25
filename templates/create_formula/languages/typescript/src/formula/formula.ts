import * as chalk from 'chalk'

function Run(input1: string, input2: string, input3: boolean) {
    console.log('Hello World!')
    console.log(chalk.green(`You receive ${input1} in text.`));
    console.log(chalk.red(`You receive ${input2} in list.`));
    console.log(chalk.yellow(`You receive ${input3} in boolean.`));
}

const formula = Run
export default formula
