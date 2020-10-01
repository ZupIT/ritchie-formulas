import run from './formula/Formula'

const text: string = process.env.INPUT_TEXT
const boolean: boolean = JSON.parse(process.env.INPUT_BOOLEAN.toLowerCase())
const list: string = process.env.INPUT_LIST
const password: string = process.env.INPUT_PASSWORD

run(text, boolean, list, password)
