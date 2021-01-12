package com.ritchie.formula

import com.github.tomaslanger.chalk.Chalk

class Formula(private val inputText: String, private val inputBoolean: Boolean, private val inputList: String, private val inputPassword: String) {
    fun Run() {
        System.out.println("Hello World!")
        System.out.println(Chalk.on(String.format("My name is %s.", inputText)).green())
        if (inputBoolean) {
            System.out.println(Chalk.on("I've already created formulas using Ritchie.").cyan())
        } else {
            System.out.println(Chalk.on("I'm excited in creating new formulas using Ritchie.").red())
        }
        System.out.println(Chalk.on(String.format("Today, I want to automate %s.", inputList)).yellow())
        System.out.println(Chalk.on(String.format("My secret is %s.", inputPassword)).magenta())
    }
}