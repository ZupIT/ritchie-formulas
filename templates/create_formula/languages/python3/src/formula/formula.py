#!/usr/bin/python3
from colored import fg, attr


def Run(input1, input2, input3):
    print("Hello World!")
    print(f"{fg(2)}You receive {input1} in text.{attr(0)}")
    print(f"{fg(1)}You receive {input2} in list.{attr(0)}")
    print(f"{fg(3)}You receive {input3} in boolean.{attr(0)}")
