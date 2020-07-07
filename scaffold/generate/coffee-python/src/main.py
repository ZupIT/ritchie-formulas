#!/usr/bin/python3
import os

from coffee import coffee

name = os.environ.get('NAME')
type = os.environ.get('COFFEE_TYPE')
delivery = os.environ.get('DELIVERY')
coffee.Run(name, type, delivery)
