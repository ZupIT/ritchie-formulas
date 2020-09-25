import run from "./formula/formula"

const input1: string = process.env.SAMPLE_TEXT
const input2: string = process.env.SAMPLE_LIST
const input3: boolean = Boolean(process.env.SAMPLE_BOOL)

run(input1, input2, input3)
