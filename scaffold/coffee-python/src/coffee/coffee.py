#!/usr/bin/python3
import time
from distutils.util import strtobool


def Run(name, type, delivery):
    print("Preparing your coffee {} .....".format(name))
    time.sleep(1)
    print("......")
    time.sleep(1)
    print("......")
    time.sleep(1)
    print("......")
    time.sleep(1)
    if strtobool(delivery):
        print("Your {} coffee is ready, enjoy your trip".format(type))
    else:
        print("Your {} coffee is ready, have a seat and enjoy your drink".format(type))
