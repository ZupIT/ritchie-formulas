const { existsSync, readFileSync, unlinkSync, writeFileSync } = require("fs")
const { execSync } = require("child_process")
const { homedir } = require('os')

const outputFile = homedir + '/.rit/output.json'

function runFormula(cmd) {
    try {
        execSync(cmd)
    } catch (error) {
        throw error
    }
}

function writeOutput(data) {
    try {
        const dataJSON = JSON.stringify(data)

        writeFileSync(outputFile, dataJSON)
    } catch (error) {
        throw error
    }
}

function readOutput() {
    try {
        if(existsSync(outputFile)) {
            const rawData = readFileSync(outputFile)
            const parseData = JSON.parse(rawData)

            return parseData
        }
    } catch (error) {
        throw error
    }
}

function removeFile() {
    try {
        unlinkSync(outputFile)
    } catch (error) {
        throw error
    }
}

const api = { readOutput, removeFile, runFormula, writeOutput }
module.exports = api
