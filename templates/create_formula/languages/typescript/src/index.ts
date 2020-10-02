import run from './formula/Formula'

const inputText: string = process.env.INPUT_TEXT
const inputBoolean: boolean = JSON.parse(process.env.INPUT_BOOLEAN.toLowerCase())
const inputList: string = process.env.INPUT_LIST
const inputPassword: string = process.env.INPUT_PASSWORD

run(inputText, inputBoolean, inputList, inputPassword)
