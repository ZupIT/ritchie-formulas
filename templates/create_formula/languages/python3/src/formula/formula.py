#!/usr/bin/python3
from colored import fg, attr


def Run(input1, input2, input3, input4):
    print("Hello World!")
    print("%sMy name is {}.%s".format(input1) % (fg(2), attr(0)))
    if input2:
        print("%sI've already created formulas using Ritchie.%s" % (fg(3), attr(0)))
    else:
        print(
            "%sI'm excited in creating new formulas using Ritchie.%s" % (fg(3), attr(0))
        )
    print("%sToday, I want to automate {}.%s".format(input3) % (fg(1), attr(0)))
    print("%sMy secret is {}.%s".format(input4) % (fg(3), attr(0)))
