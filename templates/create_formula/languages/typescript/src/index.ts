import run from './formula/Formula'
import IFormula from './formula/IFormula'

const input1: string = process.env.INPUT_TEXT
const input2: boolean = Boolean(process.env.INPUT_BOOLEAN)
const input3: string = process.env.INPUT_LIST
const input4: string = process.env.INPUT_PASSWORD

const iFormula: IFormula = {
  text: input1,
  boolean: input2,
  list: input3,
  password: input4
}

run(iFormula)
