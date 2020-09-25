import clc from "cli-color"

function Run(input1: string, input2: string, input3: boolean) {
    console.log("Hello World!")
    console.log(clc.green(`You receive ${input1} in text.`));
    console.log(clc.red(`You receive ${input2} in list.`));
    console.log(clc.yellow(`You receive ${input3} in boolean.`));
}

const formula = Run
export default formula
