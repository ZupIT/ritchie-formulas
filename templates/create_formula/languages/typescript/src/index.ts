import run from "./formula/formula"

const INPUT1: string = process.env.SAMPLE_TEXT
const INPUT2: string = process.env.SAMPLE_LIST
const INPUT3: boolean = Boolean(process.env.SAMPLE_BOOL)

run(INPUT1, INPUT2, INPUT3)
