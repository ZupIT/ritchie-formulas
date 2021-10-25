#!/usr/bin/python3
import os

from formula import formula


number_letters = os.environ.get("RIT_INPUT_LETTERS")
number_symbols = os.environ.get("RIT_INPUT_SYMBOLS")
number_numbers = os.environ.get("RIT_INPUT_NUMBERS")
formula.Run(number_letters, number_symbols, number_numbers)
