#!/usr/bin/python3
import os

from formula import formula

datapoints_amount = os.environ.get("DATAPOINTS_AMOUNT")
category = os.environ.get("CATEGORY")
labeled_data = os.environ.get("LABELED_DATA")
formula.Run(datapoints_amount, category, labeled_data)
